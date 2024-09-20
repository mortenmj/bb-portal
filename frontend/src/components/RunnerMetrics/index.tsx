import { RunnerCount } from "@/graphql/__generated__/graphql";
import React, { useCallback, useState } from "react";
import { PieChart, Pie, Sector, Cell, Legend } from 'recharts';
import { Table, Row, Col } from 'antd';
import type { TableColumnsType } from "antd/lib";

interface RunnerDataType {
    key: React.Key;
    name: string;
    exec: string;
    value: number;
    rate: string;
    color: string;
}

const runner_columns: TableColumnsType<RunnerDataType> = [
    {
        title: 'Runner Type',
        dataIndex: 'name',
    },
    {
        title: 'Execution Type',
        dataIndex: 'exec',
        showSorterTooltip: { target: 'full-header' },
        filters: [
            {
                text: 'Remote',
                value: 'Remote',
            },
            {
                text: 'Local',
                value: 'Local',
            },
        ],

        onFilter: (value, record) => record.exec == value,
    },
    {
        title: 'Count',
        dataIndex: 'value',
        sorter: (a, b) => a.value - b.value,
    },
    {
        title: 'Rate (%)',
        dataIndex: 'rate',
        sorter: (a, b) => parseFloat(a.rate) - parseFloat(b.rate),
    },
];


function colorSwitchOnExwc(exec: string) {
    switch (exec) {
        case "Remote": return "#8884d8"
        case "Local": return "#82ca9d"
        default: return "#333333"
    }
}

function nullPercent(val: number | null | undefined, total: number | null | undefined) {
    return (((val ?? 0) / (total ?? 1)) * 100).toFixed(2);
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


const RunnerMetrics: React.FC<{ runnerMetrics: RunnerCount[]; }> = ({ runnerMetrics }) => {
    const runner_data: RunnerDataType[] = [];

    var totalCount: number = runnerMetrics.find(x => x.name == "total")?.actionsExecuted ?? 0

    runnerMetrics.map((item: RunnerCount, count: number) => {
        var rd: RunnerDataType = {
            key: count,
            name: item.name ?? "",
            value: item.actionsExecuted ?? 0,
            exec: item.execKind ?? "",
            rate: nullPercent(item.actionsExecuted, totalCount),
            color: colorSwitchOnExwc(item.execKind ?? "")

        }
        count++;
        if (rd.name != "total") {
            runner_data.push(rd);
        }
    });

    runnerMetrics.sort((x, y) => {
        var a = x.execKind ?? ""
        var b = y.execKind ?? ""
        if (a < b) {
            return -1;
        }
        if (a > b) {
            return 1;
        }
        return 0;

    })

    const [activeIndexRunner, setActiveIndexRunner] = useState(0);
    const onRunnerPieEnter = useCallback(
        (_: any, runner_idx: any) => {
            setActiveIndexRunner(runner_idx);
        },
        [setActiveIndexRunner]
    );
    return (

        <Row justify="space-around" align="middle">
            <Col span="2" />

            <Col span="8">


                <PieChart width={500} height={500}>
                    <Pie
                        activeIndex={activeIndexRunner}
                        activeShape={renderActiveShape}
                        data={runner_data}
                        dataKey="value"
                        nameKey="name"
                        cx="50%"
                        cy="50%"
                        innerRadius={70}
                        outerRadius={90}
                        onMouseEnter={onRunnerPieEnter}>
                        {
                            runner_data.map((entry, runner_index) => (
                                <Cell key={`cell-${runner_index}`} fill={entry.color} />
                            ))
                        }
                    </Pie>
                    <Legend layout="vertical" />
                </PieChart>
            </Col>
            <Col span="12">
                <Table
                    columns={runner_columns}
                    dataSource={runner_data}
                    showSorterTooltip={{ target: 'sorter-icon' }}
                />
            </Col>
            <Col span="2" />
        </Row>
    )
}

export default RunnerMetrics;