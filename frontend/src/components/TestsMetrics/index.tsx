import React, { useCallback, useState, RefAttributes } from "react";
import { PieChart, Pie, Sector, Cell, Legend } from 'recharts';
import { CardProps, Table, Row, Col, Statistic, Tag, Space } from 'antd';
import type { StatisticProps, TableColumnsType } from "antd/lib";
import CountUp from 'react-countup';
import TestStatusTag from "../TestStatusTag";
import { TestCollection } from "@/graphql/__generated__/graphql";
import { TestStatusEnum } from "../TestStatusTag";
import NullBooleanTag from "../NullableBooleanTag";
import PortalCard from "../PortalCard";
import { SearchFilterIcon, SearchWidget, TimeRangeSelector } from '@/components/SearchWidgets';
import {

    SearchOutlined,
    ExperimentOutlined,
} from "@ant-design/icons";
import { string } from "zod";

interface TestDataType {
    key: React.Key;
    status: string;
    name: string;
    value: number;
    strategy: string;
    cached_local: boolean | null;
    cached_remote: boolean | null;
    duration: number;
}
const formatter: StatisticProps['formatter'] = (value) => (
    <CountUp end={value as number} separator="," />
);

const test_columns: TableColumnsType<TestDataType> = [

    {
        title: "Status",
        dataIndex: "status",
        render: (x) => <TestStatusTag key="status" status={x as TestStatusEnum} />,
        showSorterTooltip: { target: 'full-header' },
        filters: [
            {
                text: 'No Status',
                value: 'NO_STATUS',
            },
            {
                text: 'Passed',
                value: 'PASSED',
            },
            {
                text: "Flaky",
                value: "FLAKY"
            },
            {
                text: "Timeout",
                value: "TIMEOUT"
            },
            {
                text: "Failed",
                value: "FAILED"
            },
            {
                text: "Incomplete",
                value: "INCOMPLETE"
            },
            {
                text: "Remote Failure",
                value: "REMOTE_FAILURE"
            },
            {
                text: "Failed to Build",
                value: "FAILED_TO_BUILD"
            },
            {
                text: "Tool Halted Before Testing",
                value: "TOOL_HALTED_BEFORE_TESTING"
            },
        ],
        // specify the condition of filtering result
        // here is that finding the name started with `value`
        onFilter: (value, record) => record.status == value,

    },
    {
        //TODO: make this a link to the test in the github repo so you can jump straight to the code and look at it
        title: "Mnemonic",
        dataIndex: "name",

        //TODO working search w/autocomplete

        filterIcon: filtered => <SearchFilterIcon icon={<SearchOutlined />} filtered={filtered} />,

        // onFilter: (val, record) =>  record.name.includes(val) ,
    },

    {
        title: "Strategy",
        dataIndex: "strategy"
        //sorter: (a, b) => a.target_type < b.target_type,
    },
    {
        title: "Cached Locally",
        dataIndex: "cached_local",
        render: (x) => <NullBooleanTag key="cached_local" status={x as boolean | null} />,
    },
    {
        title: "Cached Remotely",
        dataIndex: "cached_remote",
        render: (x) => <NullBooleanTag key="cached_remotely" status={x as boolean | null} />,
    },
    {
        title: "Duration(ms)",
        dataIndex: "value",
        sorter: (a, b) => a.value - b.value,
    },

]



function nullPercent(val: number | null | undefined, total: number | null | undefined) {
    return (((val ?? 0) / (total ?? 1)) * 100).toFixed(2);
}


const TestMetricsDisplay: React.FC<{ testMetrics: TestCollection[] | undefined | null }> = ({
    testMetrics
}) => {
    const target_data = [
        { key: 1, name /* mnuemnic */: '//av/tools:rbe1', value: 21 /* duration */, target_type: "nodes_binary", target_status: "SUCCESS", cached: true },

    ]
    const test_data: TestDataType[] = []
    testMetrics?.map((item: TestCollection, index) => {
        var ts = item.testSummary
        var tr0 = item.testResults?.at(0)
        var row: TestDataType = {
            key: "test-data-type-row-" + index,
            name: item.label ?? "",
            value: ts?.totalRunDuration ?? 0,
            strategy: tr0?.executionInfo?.strategy ?? "",
            cached_local: tr0?.cachedLocally ?? null,
            cached_remote: tr0?.executionInfo?.cachedRemotely ?? null,
            duration: ts?.totalRunDuration ?? 0,
            status: ts?.overallStatus ?? ""
        }
        test_data.push(row);
    })

    return (
        <Space direction="vertical" size="middle" style={{ display: 'flex' }} >
            <PortalCard icon={<ExperimentOutlined />} titleBits={["Tests"]}>
            </PortalCard>
            <Row justify="space-around" align="middle">
                <Col span="2">

                </Col>
                <Col span="12">
                    <Table
                        columns={test_columns}
                        dataSource={test_data}
                        showSorterTooltip={{ target: 'sorter-icon' }}
                    />
                </Col>

                <Col span="2" />
            </Row>
        </Space>
    )
}

export default TestMetricsDisplay;