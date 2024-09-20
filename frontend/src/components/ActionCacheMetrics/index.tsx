import React, { useCallback, useState } from "react";
import { PieChart, Pie, Sector, Cell, Legend, BarChart, Bar, LabelList } from 'recharts';
import { Table, Row, Col, Statistic, Tooltip, Space } from 'antd';
import type { StatisticProps, TableColumnsType } from "antd/lib";
import CountUp from 'react-countup';
import { ActionCacheStatistics, ActionSummary, MissDetail, ActionData } from "@/graphql/__generated__/graphql";
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
import internal from "stream";

interface MissDetailDisplayDataType {
    key: React.Key;
    name: string;
    value: number;
    rate: string;
}

interface ActionDataGraphDisplayType {
    key: React.Key;
    name: string;
    value: number;
    color: string;
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

const ad_columns: TableColumnsType<ActionData> = [
    {
        title: "Mnemonic",
        dataIndex: "mnemonic"
    },
    {
        title: "Actions Executed",
        dataIndex: "actionsExecuted",
        sorter: (a, b) => (a.actionsExecuted ?? 0) - (b.actionsExecuted ?? 0),
    },
    {
        title: "Actions Created",
        dataIndex: "actionsCreated",
        sorter: (a, b) => (a.actionsCreated ?? 0) - (b.actionsCreated ?? 0),
    },
    {
        title: "First Started(ms)",
        dataIndex: "firstStartedMs",
        sorter: (a, b) => (a.firstStartedMs ?? 0) - (b.firstStartedMs ?? 0),
    },
    {
        title: "Last Ended(ms)",
        dataIndex: "lastEndedMs",
        sorter: (a, b) => (a.lastEndedMs ?? 0) - (b.lastEndedMs ?? 0),
    },
    {
        title: "System Time(ms)",
        dataIndex: "systemTime",
        sorter: (a, b) => (a.systemTime ?? 0) - (b.systemTime ?? 0),
    },
    {
        title: "User Time(ms)",
        dataIndex: "userTime",
        sorter: (a, b) => (a.userTime ?? 0) - (b.userTime ?? 0),
    },
]

const ac_columns: TableColumnsType<MissDetailDisplayDataType> = [
    {
        title: "Miss Reason",
        dataIndex: "name",
    },
    {
        title: "Count",
        dataIndex: "value",
        sorter: (a, b) => a.value - b.value,
    },
    {
        title: "Rate (%)",
        dataIndex: "rate",
        sorter: (a, b) => parseFloat(a.rate) - parseFloat(b.rate),
    }
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

const AcMetrics: React.FC<{ acMetrics: ActionSummary | undefined; }> = ({ acMetrics }) => {

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

    const actions_data: ActionData[] = [];
    const actions_graph_data: ActionDataGraphDisplayType[] = [];
    acMetrics?.actionData?.map((ad: ActionData, idx) => {
        actions_data.push(ad)
        var agd: ActionDataGraphDisplayType = {
            key: "actiondatagraphdisplaytype-" + String(idx),
            name: ad.mnemonic ?? "",
            value: ad.userTime ?? 0,
            color: newColorFind(idx) ?? "#333333"
        }
        actions_graph_data.push(agd)
    });


    const acMetricsData: ActionCacheStatistics | undefined = acMetrics?.actionCacheStatistics?.at(0)

    var hitMissTotal: number = (acMetricsData?.misses ?? 0) + (acMetricsData?.hits ?? 0);

    const hits_data = [
        {
            key: "hitMissBarChart",
            Hit: acMetricsData?.hits,
            Miss: acMetricsData?.misses,
            hit_label: nullPercent(acMetricsData?.hits, hitMissTotal, 0),
            miss_label: nullPercent(acMetricsData?.misses, hitMissTotal, 0)
        },
    ]

    const ac_data: MissDetailDisplayDataType[] = [];
    var missTotal: number = acMetricsData?.misses ?? 0;

    acMetricsData?.missDetails?.map((item: MissDetail, index) => {
        var acd: MissDetailDisplayDataType = {
            key: index,
            name: item.reason ?? "",
            value: item.count ?? 0,
            rate: nullPercent(item.count, missTotal),

        }
        ac_data.push(acd)
    });


    const [activeIndexRunner, setActiveIndexRunner] = useState(0);
    const onRunnerPieEnter = useCallback(
        (_: any, runner_idx: any) => {
            setActiveIndexRunner(runner_idx);
        },
        [setActiveIndexRunner]
    );
    const acTitle: React.ReactNode[] = [<span key="label">Action Cache Statitics</span>];
    const actionsTitle: React.ReactNode[] = [<span key="label">Actions Data</span>];
    const userTimeTitle: React.ReactNode[] = [<span key="label">User Time</span>];

    return (
        <Space direction="vertical" size="middle" style={{ display: 'flex' }} >
            <PortalCard icon={<BuildOutlined />} titleBits={actionsTitle}>
                <Row justify="space-around" align="middle">
                    <Col span="18">
                        <Table
                            columns={ad_columns}
                            dataSource={actions_data}
                            showSorterTooltip={{ target: 'sorter-icon' }}
                        />
                    </Col>
                    <Col span="6">
                        <PortalCard icon={<PieChartOutlined />} titleBits={userTimeTitle}>

                            <PieChart width={600} height={556}>

                                <Pie
                                    activeIndex={activeIndexRunner}
                                    activeShape={renderActiveShape}

                                    data={actions_graph_data}
                                    dataKey="value"
                                    cx="50%"
                                    cy="50%"
                                    innerRadius={50}
                                    outerRadius={90}
                                    onMouseEnter={onRunnerPieEnter}
                                >
                                    {
                                        actions_graph_data.map((entry, actions_index) => (
                                            <Cell key={`cell-${actions_index}`} fill={entry.color} />
                                        ))
                                    }
                                </Pie>
                            </PieChart>

                        </PortalCard>
                    </Col>
                </Row>
            </PortalCard>

            <PortalCard icon={<PieChartOutlined />} titleBits={acTitle} >
                <Row justify="space-around" align="middle" >
                    <Col span="2">
                        <BarChart width={170} height={150} data={hits_data} margin={{ top: 0, left: 10, bottom: 10, right: 10 }}>
                            <Bar dataKey="Miss" fill={"#8884d8"} stackId="a">
                                <LabelList dataKey="miss_label" position="center" />
                            </Bar>
                            <Bar dataKey="Hit" fill={"#82ca9d"} stackId="a">
                                <LabelList dataKey="hit_label" position="center" />
                            </Bar>
                            <Tooltip />
                            <Legend />
                        </BarChart>
                        <Statistic title="Hits" value={acMetricsData?.hits ?? 0} formatter={formatter} />
                        <Statistic title="Misses" value={acMetricsData?.misses ?? 0} formatter={formatter} />
                        <Statistic title="Size (bytes)" value={acMetricsData?.sizeInBytes ?? 0} formatter={formatter} />
                        {/* <Statistic title="Load Time(ms)" value={acMetricsData.loadTimeInMs ?? 0} formatter={formatter} /> */}
                        <Statistic title="Save Time(ms)" value={acMetricsData?.saveTimeInMs ?? 0} formatter={formatter} />
                    </Col>
                    <Col span="8">

                        <PieChart width={500} height={500}>

                            <Pie
                                activeIndex={activeIndexRunner}
                                activeShape={renderActiveShape}
                                data={ac_data}
                                dataKey="value"
                                nameKey="name"
                                cx="50%"
                                cy="50%"
                                innerRadius={70}
                                outerRadius={90}
                                onMouseEnter={onRunnerPieEnter}>
                                {
                                    ac_data.map((entry, runner_index) => (
                                        <Cell key={`cell-${runner_index}`} fill={ac_colors[runner_index]} />
                                    ))
                                }
                            </Pie>
                            <Legend layout="vertical" />
                        </PieChart>
                    </Col>
                    <Col span="12">
                        <Table
                            columns={ac_columns}
                            dataSource={ac_data}
                            showSorterTooltip={{ target: 'sorter-icon' }}
                        />
                    </Col>
                    <Col span="2" />
                </Row >
            </PortalCard>
        </Space>



    )
}

export default AcMetrics;