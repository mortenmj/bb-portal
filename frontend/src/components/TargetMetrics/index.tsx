import React, { useCallback, useState, RefAttributes } from "react";
import { PieChart, Pie, Sector, Cell, Legend } from 'recharts';
import IntrinsicAttributes = JSX.IntrinsicAttributes;
import { CardProps, Table, Row, Col, Statistic, Tag } from 'antd';
import {
    CheckCircleFilled,
    CloseCircleFilled,
    ExclamationCircleFilled,
    InfoCircleFilled,
    QuestionCircleFilled,
    StopFilled,
} from '@ant-design/icons';
import { JSX } from 'react/jsx-runtime';
import type { StatisticProps, TableColumnsType } from "antd/lib";
import CountUp from 'react-countup';
import themeStyles from '@/theme/theme.module.css';
import BuildStepResultTag, { BuildStepResultEnum } from "../BuildStepResultTag";
import { TargetMetrics } from "@/graphql/__generated__/graphql";

interface Props {
    targetsLoaded: number | null | undefined;
    targetsConfigured: number | null | undefined;
    targetsConfiguredNotIncludingAspects: number | null | undefined;
    //TODO: Capture target data including timing and pass in here
}

interface TargetDataType {
    key: React.Key;
    name: string;
    value: number;
    target_type: string;
    target_status: string;
    cached: boolean;
}
const formatter: StatisticProps['formatter'] = (value) => (
    <CountUp end={value as number} separator="," />
);

var target_colors = ["#00cc66", "#0099ff", "#9900cc", "#ff9900", "#ff5050", "#ffff66", "#ff00ff"]

const target_columns: TableColumnsType<TargetDataType> = [
    {
        //TODO: make this a link to the test in the github repo so you can jump straight to the code and look at it
        title: "Mnemonic",
        dataIndex: "name",
    },
    {
        title: "Duration(ms)",
        dataIndex: "value",
        sorter: (a, b) => a.value - b.value,
    },
    {
        title: "Target Type",
        dataIndex: "target_type"
        //sorter: (a, b) => a.target_type < b.target_type,
    },
    {
        title: "Target Status",
        dataIndex: "target_status",
        render: (x) => <BuildStepResultTag key="result" result={x as BuildStepResultEnum} />,
        showSorterTooltip: { target: 'full-header' },
        filters: [
            {
                text: 'Suceeded',
                value: 'SUCCESS',
            },
            {
                text: 'Unstable',
                value: 'UNSTABLE',
            },
            {
                text: "Parsing Failure",
                value: "PARSING_FAILURE"
            },
            {
                text: "Build Failed",
                value: "BUILD_FAILURE"
            },
            {
                text: "Tests Failed",
                value: "TESTS_FAILED"
            },
            {
                text: "Not Built",
                value: "NOT_BUILT"
            },
            {
                text: "Aborted",
                value: "ABORTED"
            },
            {
                text: "Interrupted",
                value: "INTERRUPTED"
            },
            {
                text: "Unknown",
                value: "UNKNOWN"
            },
        ],
        // specify the condition of filtering result
        // here is that finding the name started with `value`
        onFilter: (value, record) => record.target_status == value,

    },
    {
        title: "Cached",
        dataIndex: "cached"
    }

]



function nullPercent(val: number | null | undefined, total: number | null | undefined) {
    return (((val ?? 0) / (total ?? 1)) * 100).toFixed(2);
}


const TargetMetricsDisplay: React.FC<{ targetMetrics: TargetMetrics | undefined }> = ({
    targetMetrics
}) => {
    const target_data = [
        { key: 1, name /* mnuemnic */: '//av/tools:rbe1', value: 21 /* duration */, target_type: "nodes_binary", target_status: "SUCCESS", cached: true },
        { key: 2, name /* mnuemnic */: '//av/tools:rbe2', value: 120 /* duration */, target_type: "nodes_binary", target_status: 'BUILD_FAILURE', cached: false },
        { key: 3, name /* mnuemnic */: '//av/tools:rbe3', value: 35 /* duration */, target_type: "nodes_binary", target_status: 'TESTS_FAILED', cached: true },
        { key: 4, name /* mnuemnic */: '//av/tools:rbe4', value: 70 /* duration */, target_type: "nodes_binary", target_status: 'NOT_BUILT', cached: true },
        { key: 5, name /* mnuemnic */: '//av/tools:rbe5', value: 2 /* duration */, target_type: "nodes_binary", target_status: 'INTERRUPTED', cached: true },
        { key: 6, name /* mnuemnic */: '//av/tools:rbe1', value: 21 /* duration */, target_type: "nodes_binary", target_status: "SUCCESS", cached: true },
        { key: 7, name /* mnuemnic */: '//av/tools:rbe2', value: 120 /* duration */, target_type: "nodes_binary", target_status: 'BUILD_FAILURE', cached: false },
        { key: 8, name /* mnuemnic */: '//av/tools:rbe3', value: 35 /* duration */, target_type: "nodes_binary", target_status: 'TESTS_FAILED', cached: true },
        { key: 9, name /* mnuemnic */: '//av/tools:rbe4', value: 70 /* duration */, target_type: "nodes_binary", target_status: 'NOT_BUILT', cached: true },
        { key: 10, name /* mnuemnic */: '//av/tools:rbe5', value: 2 /* duration */, target_type: "nodes_binary", target_status: 'INTERRUPTED', cached: true },
        { key: 11, name /* mnuemnic */: '//av/tools:rbe1', value: 21 /* duration */, target_type: "nodes_binary", target_status: "SUCCESS", cached: true },
        { key: 12, name /* mnuemnic */: '//av/tools:rbe2', value: 120 /* duration */, target_type: "nodes_binary", target_status: 'BUILD_FAILURE', cached: false },
        { key: 13, name /* mnuemnic */: '//av/tools:rbe3', value: 35 /* duration */, target_type: "nodes_binary", target_status: 'TESTS_FAILED', cached: true },
        { key: 14, name /* mnuemnic */: '//av/tools:rbe4', value: 70 /* duration */, target_type: "nodes_binary", target_status: 'NOT_BUILT', cached: true },
        { key: 15, name /* mnuemnic */: '//av/tools:rbe5', value: 2 /* duration */, target_type: "nodes_binary", target_status: 'INTERRUPTED', cached: true },
        { key: 16, name /* mnuemnic */: '//av/tools:rbe1', value: 21 /* duration */, target_type: "nodes_binary", target_status: "SUCCESS", cached: true },
        { key: 17, name /* mnuemnic */: '//av/tools:rbe2', value: 120 /* duration */, target_type: "nodes_binary", target_status: 'BUILD_FAILURE', cached: false },
        { key: 18, name /* mnuemnic */: '//av/tools:rbe3', value: 35 /* duration */, target_type: "nodes_binary", target_status: 'TESTS_FAILED', cached: true },
        { key: 19, name /* mnuemnic */: '//av/tools:rbe4', value: 70 /* duration */, target_type: "nodes_binary", target_status: 'NOT_BUILT', cached: true },
        { key: 20, name /* mnuemnic */: '//av/tools:rbe5', value: 2 /* duration */, target_type: "nodes_binary", target_status: 'INTERRUPTED', cached: true },
        { key: 21, name /* mnuemnic */: '//av/tools:rbe1', value: 21 /* duration */, target_type: "nodes_binary", target_status: "SUCCESS", cached: true },
        { key: 22, name /* mnuemnic */: '//av/tools:rbe2', value: 120 /* duration */, target_type: "nodes_binary", target_status: 'BUILD_FAILURE', cached: false },
        { key: 23, name /* mnuemnic */: '//av/tools:rbe3', value: 35 /* duration */, target_type: "nodes_binary", target_status: 'TESTS_FAILED', cached: true },
        { key: 24, name /* mnuemnic */: '//av/tools:rbe4', value: 70 /* duration */, target_type: "nodes_binary", target_status: 'NOT_BUILT', cached: true },
        { key: 25, name /* mnuemnic */: '//av/tools:rbe5', value: 2 /* duration */, target_type: "nodes_binary", target_status: 'INTERRUPTED', cached: true },
        { key: 26, name /* mnuemnic */: '//av/tools:rbe1', value: 21 /* duration */, target_type: "nodes_binary", target_status: "SUCCESS", cached: true },
        { key: 27, name /* mnuemnic */: '//av/tools:rbe2', value: 120 /* duration */, target_type: "nodes_binary", target_status: 'BUILD_FAILURE', cached: false },
        { key: 28, name /* mnuemnic */: '//av/tools:rbe3', value: 35 /* duration */, target_type: "nodes_binary", target_status: 'TESTS_FAILED', cached: true },
        { key: 29, name /* mnuemnic */: '//av/tools:rbe4', value: 70 /* duration */, target_type: "nodes_binary", target_status: 'NOT_BUILT', cached: true },
        { key: 30, name /* mnuemnic */: '//av/tools:rbe5', value: 2 /* duration */, target_type: "nodes_binary", target_status: 'INTERRUPTED', cached: true },
    ]

    return (
        <Row justify="space-around" align="middle">
            <Col span="2">
                <Statistic title="Targets Configured" value={targetMetrics?.targetsConfigured ?? 0} formatter={formatter} />
                <Statistic title="Targets Configured Not Including Aspects" value={targetMetrics?.targetsConfiguredNotIncludingAspects ?? 0} formatter={formatter} />
                <Statistic title="Targets Loaded" value={targetMetrics?.targetsLoaded ?? 0} formatter={formatter} />


            </Col>
            <Col span="8">



            </Col>
            <Col span="12">
                <Table
                    columns={target_columns}
                    dataSource={target_data}
                    showSorterTooltip={{ target: 'sorter-icon' }}
                />
            </Col>
            <Col span="2" />
        </Row>
    )
}

export default TargetMetricsDisplay;