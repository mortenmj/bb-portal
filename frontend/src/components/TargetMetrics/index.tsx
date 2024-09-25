import React from "react";
import { Space, Table, Row, Col, Statistic, Tag } from 'antd';
import { DeploymentUnitOutlined, SearchOutlined } from '@ant-design/icons';
import type { StatisticProps, TableColumnsType } from "antd/lib";
import CountUp from 'react-countup';
import { TargetMetrics, TargetPair } from "@/graphql/__generated__/graphql";
import PortalCard from "../PortalCard";
import { SearchFilterIcon, SearchWidget } from '@/components/SearchWidgets';
import NullBooleanTag from "../NullableBooleanTag";
import TargetTypeTag from "./targetType";
import { TestSizeEnum } from "./targetType";

interface TargetDataType {
    key: React.Key;
    name: string;   //label
    success: boolean;
    value: number;  //duration
    target_kind: string;
    test_size: string;
    failure_reason: string
    test_found: boolean;
}

const formatter: StatisticProps['formatter'] = (value) => (
    <CountUp end={value as number} separator="," />
);


const TargetMetricsDisplay: React.FC<{
    targetMetrics: TargetMetrics | undefined | null,
    targetData: TargetPair[] | undefined | null,
    testLabels: Map<string, boolean>
}> = ({
    targetMetrics,
    targetData,
    testLabels,
}) => {

        var target_data: TargetDataType[] = []
        var count = 0;
        var all_types: string[] = []
        targetData?.map(x => {
            count++;
            var targetKind = x.targetKind ?? ""
            var failureReason = x.abortReason ?? ""

            var targetLabel: string = x.label ?? "EMPTY12345"

            var row: TargetDataType = {
                key: "target_data_type" + count.toString(),
                name: x.label ?? "",
                success: x.success ?? false,
                value: x.durationInMs ?? 0,
                target_kind: targetKind,
                test_size: x.testSize ?? "",
                failure_reason: failureReason,
                test_found: testLabels.has(targetLabel)
            }
            all_types.push(targetKind)
            target_data.push(row)

        })

        const targets_analyzed: number = targetData?.length ?? 0

        const type_filters: string[] = Array.from(new Set(all_types))

        const target_columns: TableColumnsType<TargetDataType> = [
            {
                //TODO: make this a link to the test in the github repo so you can jump straight to the code and look at it
                title: "Mnemonic",
                dataIndex: "name",
                //TODO working search w/autocomplete
                filterSearch: true,
                filterDropdown: filterProps => (
                    <SearchWidget placeholder="Target Pattern..." {...filterProps} />
                ),
                filterIcon: filtered => <SearchFilterIcon icon={<SearchOutlined />} filtered={filtered} />,
                onFilter: (value, record) => (record.name.includes(value.toString()) ? true : false)
            },
            {
                title: "Duration(ms)",
                dataIndex: "value",
                sorter: (a, b) => a.value - b.value,
            },
            {
                title: "Target Type",
                dataIndex: "target_kind",
                filters: type_filters.map(x => ({ text: x, value: x })),
                filterIcon: filtered => <SearchFilterIcon icon={<SearchOutlined />} filtered={filtered} />,
                onFilter: (value, record) => (record.target_kind.includes(value.toString()) ? true : false),
                sorter: (a, b) => a.target_kind.localeCompare(b.target_kind),

            },
            {
                title: "Failure Reason",
                dataIndex: "failure_reason",
                sorter: (a, b) => a.failure_reason.localeCompare(b.failure_reason),

            },
            {
                title: "Test Found",
                dataIndex: "test_found",
                render: (x) => <NullBooleanTag key="test_found" status={x as boolean | null} />,
                sorter: (a, b) => Number(a.test_found) - Number(b.test_found),
                filters: [
                    {
                        text: "Yes",
                        value: true
                    },
                    {
                        text: "No",
                        value: false
                    }
                ],
                filterIcon: filtered => <SearchFilterIcon icon={<SearchOutlined />} filtered={filtered} />,
                onFilter: (value, record) => record.test_found == value,
            },
            {
                title: "Test Size",
                dataIndex: "test_size",
                filters: [
                    {
                        text: "None",
                        value: "UNKNOWN"
                    },
                    {
                        text: "Small",
                        value: "SMALL"
                    },
                    {
                        text: "Medium",
                        value: "MEDIUM"
                    },
                    {
                        text: "Large",
                        value: "LARGE"
                    },
                    {
                        text: "Enormous",
                        value: "ENORMOUS"
                    },
                ],
                filterIcon: filtered => <SearchFilterIcon icon={<SearchOutlined />} filtered={filtered} />,
                onFilter: (value, record) => record.test_size == value,
                render: (x) => <TargetTypeTag key="test_size" size={x as TestSizeEnum} />,
                sorter: (a, b) => a.test_size.localeCompare(b.test_size),
            },
            {
                title: "Overall Status",
                dataIndex: "success",
                render: (x) => <NullBooleanTag key="success" status={x as boolean | null} />,
                sorter: (a, b) => Number(a.success) - Number(b.success),
                filters: [
                    {
                        text: "Yes",
                        value: true
                    },
                    {
                        text: "No",
                        value: false
                    }
                ],
                filterIcon: filtered => <SearchFilterIcon icon={<SearchOutlined />} filtered={filtered} />,
                onFilter: (value, record) => record.success == value,
            },
        ]


        return (
            <Space direction="vertical" size="middle" style={{ display: 'flex' }} >
                <PortalCard icon={<DeploymentUnitOutlined />} titleBits={["Targets"]}>
                    <Row>
                        <Space size="large">

                            <Statistic title="Targets Analyzed" value={targets_analyzed} formatter={formatter} />
                            <Statistic title="Targets Configured" value={targetMetrics?.targetsConfigured ?? 0} formatter={formatter} />
                            <Statistic title="Targets Configured Not Including Aspects" value={targetMetrics?.targetsConfiguredNotIncludingAspects ?? 0} formatter={formatter} />
                            {/* <Statistic title="Targets Loaded" value={targetMetrics?.targetsLoaded ?? 0} formatter={formatter} /> */}

                        </Space>
                    </Row>
                    <Row justify="space-around" align="middle">
                        <Col span="1" />
                        <Col span="22">
                            <Table
                                columns={target_columns}
                                dataSource={target_data}
                                showSorterTooltip={{ target: 'sorter-icon' }}
                            />
                        </Col>
                        <Col span="1" />
                    </Row>
                </PortalCard>

            </Space>
        )
    }

export default TargetMetricsDisplay;