import React from "react";


import { Statistic } from 'antd';

import {
    FieldTimeOutlined,
} from "@ant-design/icons";
import type { StatisticProps } from "antd/lib";
import CountUp from 'react-countup';

import { TimingMetrics } from "@/graphql/__generated__/graphql";
import PortalCard from "../PortalCard";


const formatter: StatisticProps['formatter'] = (value) => (
    <CountUp end={value as number} separator="," />
);

const TimingMetricsDisplay: React.FC<{ timingMetrics: TimingMetrics | undefined }> = ({
    timingMetrics
}) => {
    return (
        <PortalCard titleBits={["Timing Metrics"]} icon={<FieldTimeOutlined />}>
            <Statistic title="Wall Time(ms)" value={timingMetrics?.wallTimeInMs ?? 0} formatter={formatter} />
            <Statistic title="Analysis(ms)" value={timingMetrics?.analysisPhaseTimeInMs ?? 0} formatter={formatter} />
            <Statistic title="CPU Time(ms)" value={timingMetrics?.cpuTimeInMs ?? 0} formatter={formatter} />
            <Statistic title="Execuction(ms)" value={timingMetrics?.executionPhaseTimeInMs ?? 0} formatter={formatter} />
            <Statistic title="Actions Execution Start" value={timingMetrics?.actionsExecutionStartInMs ?? 0} formatter={formatter} />
        </PortalCard>

    )
}

export default TimingMetricsDisplay