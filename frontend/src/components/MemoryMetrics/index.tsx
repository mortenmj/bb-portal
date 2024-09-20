import React, { useCallback, useState } from "react";
import { PieChart, Pie, Sector, Cell, Legend, BarChart, Bar, LabelList } from 'recharts';
import { Table, Row, Col, Statistic, Tooltip, Space } from 'antd';
import type { StatisticProps, TableColumnsType } from "antd/lib";
import CountUp from 'react-countup';
import { ActionCacheStatistics, ActionSummary, MissDetail, ActionData, MemoryMetrics, GarbageMetrics } from "@/graphql/__generated__/graphql";
import PortalCard from "../PortalCard";
import {
    BuildOutlined,
    FileSearchOutlined,
    PieChartOutlined,
    ExclamationCircleOutlined,
    NodeCollapseOutlined,
    DeploymentUnitOutlined,
    ExperimentOutlined,
} from "@ant-design/icons";

interface GarbageMetricDetailDisplayType {
    key: React.Key;
    name: string;
    value: number;
    color: string;
    //    rate: string;
}


const formatter: StatisticProps['formatter'] = (value) => (
    <CountUp end={value as number} separator="," />
);


const generateColor = () => {
    let randomColorString = "#";
    const arrayOfColorFunctions = "0123456789abcdef";
    for (let x = 0; x < 6; x++) {
        let index = Math.floor(Math.random() * 16);
        let value = arrayOfColorFunctions[index];

        randomColorString += value;
    }
    return randomColorString;
};


const colorMap: Map<number, string> = new Map();
const selectedColors: Map<string, boolean> = new Map();

var ac_colors = ["#00cc66", "#0099ff", "#9900cc", "#ff9900", "#ff5050", "#ffff66", "#ff00ff"]


const garbage_columns: TableColumnsType<GarbageMetricDetailDisplayType> = [
    {
        title: "Type",
        dataIndex: "name",
    },
    {
        title: "Garbage Collected",
        dataIndex: "value",
        sorter: (a, b) => a.value - b.value,
    },
]

function nullPercent(val: number | null | undefined, total: number | null | undefined, fixed: number = 2) {
    return String((((val ?? 0) / (total ?? 1)) * 100).toFixed(fixed)) + "%";
}

const renderActiveShape = (props: any) => {
    const RADIAN = Math.PI / 180;
    const {
        cx,
        cy,
        midAngle,
        innerRadius,
        outerRadius,
        startAngle,
        endAngle,
        fill,
        payload,
        percent,
        value
    } = props;
    const sin = Math.sin(-RADIAN * midAngle);
    const cos = Math.cos(-RADIAN * midAngle);
    const sx = cx + (outerRadius + 10) * cos;
    const sy = cy + (outerRadius + 10) * sin;
    const mx = cx + (outerRadius + 30) * cos;
    const my = cy + (outerRadius + 30) * sin;
    const ex = mx + (cos >= 0 ? 1 : -1) * 22;
    const ey = my;
    const textAnchor = cos >= 0 ? "start" : "end";

    return (
        <g>
            <text x={cx} y={cy} dy={8} textAnchor="middle" fill={fill}>
                {value}
            </text>
            <Sector
                cx={cx}
                cy={cy}
                innerRadius={innerRadius}
                outerRadius={outerRadius}
                startAngle={startAngle}
                endAngle={endAngle}
                fill={fill}
            />
            <Sector
                cx={cx}
                cy={cy}
                startAngle={startAngle}
                endAngle={endAngle}
                innerRadius={outerRadius + 6}
                outerRadius={outerRadius + 10}
                fill={fill}
            />
            <path
                d={`M${sx},${sy}L${mx},${my}L${ex},${ey}`}
                stroke={fill}
                fill="none"
            />
            <circle cx={ex} cy={ey} r={2} fill={fill} stroke="none" />
            <text
                x={ex + (cos >= 0 ? 1 : -1) * 12}
                y={ey}
                textAnchor={textAnchor}
                fill="#333"
            >{`${payload.name}`}</text>
            <text
                x={ex + (cos >= 0 ? 1 : -1) * 12}
                y={ey}
                dy={18}
                textAnchor={textAnchor}
                fill="#999"
            >
                {`(Rate ${(percent * 100).toFixed(2)}%)`}
            </text>
        </g>
    );
};

const newColorFind = (id: number) => {
    // If already generated and assigned, return
    if (colorMap.get(id)) return colorMap.get(id);

    // Generate new random color
    let newColor;

    do {
        newColor = generateColor();
    } while (selectedColors.get(newColor));

    // Found a new random, unassigned color
    colorMap.set(id, newColor);
    selectedColors.set(newColor, true);

    // Return next new color
    return newColor;
}
const MemoryMetricsDisplay: React.FC<{ memoryMetrics: MemoryMetrics | undefined; }> = ({ memoryMetrics }) => {





    const garbage_data: GarbageMetricDetailDisplayType[] = [];
    memoryMetrics?.garbageMetrics?.map((item: GarbageMetrics, index) => {
        var gm: GarbageMetricDetailDisplayType = {
            key: index,
            name: item.type ?? "",
            value: item.garbageCollected ?? 0,
            color: newColorFind(index) ?? "#333333"
        }
        garbage_data.push(gm)
    });


    const [activeIndexRunner, setActiveIndexRunner] = useState(0);
    const onRunnerPieEnter = useCallback(
        (_: any, runner_idx: any) => {
            setActiveIndexRunner(runner_idx);
        },
        [setActiveIndexRunner]
    );

    return (
        <Space direction="vertical" size="middle" style={{ display: 'flex' }} >


            <PortalCard icon={<PieChartOutlined />} titleBits={["Memory Metrics"]} >
                <Row justify="space-around" align="middle" >
                    <Col span="2">

                        <Statistic title="Peak Post GC Heap Size" value={memoryMetrics?.peakPostGcHeapSize ?? 0} formatter={formatter} />
                        <Statistic title="Peak Post TC Tenured Space Heap Size" value={memoryMetrics?.peakPostGcTenuredSpaceHeapSize ?? 0} formatter={formatter} />
                        <Statistic title="Used Heap Size Post Build" value={memoryMetrics?.usedHeapSizePostBuild ?? 0} formatter={formatter} />
                    </Col>
                    <Col span="8">

                        <PieChart width={500} height={500}>

                            <Pie
                                activeIndex={activeIndexRunner}
                                activeShape={renderActiveShape}
                                data={garbage_data}
                                dataKey="value"
                                nameKey="name"
                                cx="50%"
                                cy="50%"
                                innerRadius={70}
                                outerRadius={90}
                                onMouseEnter={onRunnerPieEnter}>
                                {
                                    garbage_data.map((entry, runner_index) => (
                                        <Cell key={`cell-${runner_index}`} fill={entry.color} />
                                    ))
                                }
                            </Pie>
                            <Legend layout="vertical" />
                        </PieChart>
                    </Col>
                    <Col span="12">
                        <Table
                            columns={garbage_columns}
                            dataSource={garbage_data}
                            showSorterTooltip={{ target: 'sorter-icon' }}
                        />
                    </Col>
                    <Col span="2" />
                </Row >
            </PortalCard>
        </Space>



    )
}

export default MemoryMetricsDisplay;