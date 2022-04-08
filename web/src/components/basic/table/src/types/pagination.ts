import type { VNodeChild } from 'vue';
import type { PaginationProps as NPaginationProps } from 'naive-ui';

interface PaginationRenderProps {
    page: number;
    type: 'page' | 'prev' | 'next';
    originalElement: any;
}

interface pageSizeOption {
    value: number;
    label: string;
}

export interface PaginationProps extends NPaginationProps {}
