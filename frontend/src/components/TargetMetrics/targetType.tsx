import React from 'react';
import {
    CheckCircleFilled,
    CloseCircleFilled,
    InfoCircleFilled,
    MinusCircleFilled,
    QuestionCircleFilled,
    StopOutlined,
} from '@ant-design/icons';
import { Tag } from 'antd';
import themeStyles from '@/theme/theme.module.css';

export const ALL_STATUS_VALUES = [
    'UNKNOWN',
    'SMALL',
    'MEDIUM',
    'LARGE',
    'ENORMOUS',
] as const;
export type SizeTuple = typeof ALL_STATUS_VALUES;
export type TestSizeEnum = SizeTuple[number];

interface Props {
    size: TestSizeEnum;
}

const STATUS_TAGS: { [key in TestSizeEnum]: React.ReactNode } = {
    UNKNOWN: (
        <Tag icon={<QuestionCircleFilled />} className={themeStyles.tag} >
            <span className={themeStyles.tagContent}>None</span>
        </Tag>
    ),
    SMALL: (
        <Tag icon={<CheckCircleFilled />} color="green" className={themeStyles.tag}>
            <span className={themeStyles.tagContent}>Small</span>
        </Tag>
    ),
    MEDIUM: (
        <Tag icon={<InfoCircleFilled />} color="orange" className={themeStyles.tag}>
            <span className={themeStyles.tagContent}>Medium</span>
        </Tag>
    ),
    LARGE: (
        <Tag icon={<CloseCircleFilled />} color="red" className={themeStyles.tag}>
            <span className={themeStyles.tagContent}>Large</span>
        </Tag>
    ),
    ENORMOUS: (
        <Tag icon={<MinusCircleFilled />} color="red" className={themeStyles.tag}>
            <span className={themeStyles.tagContent}>Enormous</span>
        </Tag>
    ),
};

const TargetTypeTag: React.FC<Props> = ({ size }) => {
    const resultTag = STATUS_TAGS[size] || STATUS_TAGS.UNKNOWN;
    return <>{resultTag}</>;
};

export default TargetTypeTag;
