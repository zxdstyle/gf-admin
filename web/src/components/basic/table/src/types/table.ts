import type { DataTableBaseColumn, DataTableCreateSummary } from 'naive-ui';
import type { CreateRowClassName, CreateRowKey, CreateRowProps } from 'naive-ui/es/data-table/src/interface';
import type { Recordable } from '@/typings/global';
import type { PaginationProps } from '../types/pagination';

export interface FetchParams {
    searchInfo?: Recordable;
    page?: number;
    sortInfo?: Recordable;
    filterInfo?: Recordable;
}

export interface BasicColumn extends DataTableBaseColumn {
    // 业务控制是否显示
    ifShow?: boolean | ((column: BasicColumn) => boolean);
}

export interface BasicTableProps {
    api?: (...arg: any) => Promise<any>;
    data?: Recordable[];

    // 是否显示 border
    bordered?: boolean;

    // 是否显示 bottom border
    bottomBordered?: boolean;

    // 被选中列的 key
    checkedRowKeys?: string[] | number[];

    // 在进行树型数据选择的时候是否级联
    cascade?: boolean;

    // 树形数据下后代节点在数据中的 key, 默认 children
    childrenKey?: string;

    // 默认选中的 key 值
    defaultCheckedRowKeys?: string[] | number[];

    // 默认展开行的 key 值
    defaultExpandedRowKeys?: string[] | number[];

    // 展开行的 key 值
    expandedRowKeys?: string[] | number[];

    // 使用树形数据时行内容的缩进，默认 16
    indent?: number;

    // 是否让表格主体的高度自动适应整个表格区域的高度，打开这个选项会让 table-layout 始终为 'fixed'
    flexHeight?: boolean;

    // 表格内容的最大高度，可以是 CSS 属性值
    maxHeight?: string | number;

    // 表格内容的最低高度，可以是 CSS 属性值
    minHeight?: string | number;

    // 表格是否自动分页数据，在异步的状况下你可能需要把它设为 true
    remote?: boolean;

    // 每一行上的类名
    rowClassName?: string | CreateRowClassName;

    // 通过行数据创建行的 key（如果你不想给每一行加上 key）
    rowKey?: CreateRowKey;

    // 自定义行属性
    rowProps?: CreateRowProps;

    // 表格内容的横向宽度，如果列被水平固定了，则需要设定它
    scrollX?: number | string;

    // 是否不设定行的分割线，当参数为true时，则单元格没有下边线
    singleColumn?: boolean;

    // 是否不设定列的分割线，当参数值为 true 时，则单元格没有右边线
    singleLine?: boolean;

    // 表格的尺寸
    size?: 'small' | 'medium' | 'large';

    // 是否使用斑马线条纹
    striped?: boolean;

    // 总结栏的数据
    summary?: DataTableCreateSummary;

    // 表格的 table-layout 样式属性，在设定 ellipsis 或 max-height 的情况下固定为 'fixed'
    tableLayout?: 'auto' | 'fixed';

    // 是否开启虚拟滚动，应对大规模数据，开启前请设定好 max-height。当 virtual-scroll 为 true 时，rowSpan 将不生效
    virtualScroll?: boolean;

    // loading加载
    loading?: boolean;

    // 列配置
    columns: BasicColumn[];

    // 分页配置
    pagination?: PaginationProps | boolean;
}

export interface TableActionType {
    setProps: (props: Partial<BasicTableProps>) => void;
    reload: (opt?: FetchParams) => Promise<void>;
    setTableData: (values: Recordable[]) => void;
    setRowData: (row: number, value: Recordable) => void;
    setFieldData: (row: number, field: string, value: any) => void;
}
