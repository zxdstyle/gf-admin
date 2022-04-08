import type { ComputedRef, Ref } from 'vue';
import { computed, onMounted, reactive, ref, unref, watch, watchEffect } from 'vue';
import { merge } from 'lodash-es';
import type { DataTableColumn, DataTableSortState, DataTableFilterState } from 'naive-ui';
import type { Recordable } from '@/typings/global';
import { PAGE_SIZE } from '@/components/basic/table/src/const';
import { isBoolean } from '@/utils/common/is';
import type { PaginationProps } from '../types/pagination';
import type { BasicTableProps, FetchParams } from '../types/table';

interface ActionType {
    getPaginationInfo: ComputedRef<boolean | PaginationProps>;
    setPagination: (info: Partial<PaginationProps>) => void;
    setLoading: (loading: boolean) => void;
    // getFieldsValue: () => Recordable;
    // clearSelectedRowKeys: () => void;
    tableData: Ref<Recordable[]>;
}

interface SearchState {
    sortInfo: Recordable;
    filterInfo: Record<string, string[]>;
}

export function useDataSource(propsRef: ComputedRef<BasicTableProps>, { setLoading, tableData, setPagination, getPaginationInfo }: ActionType) {
    const dataSourceRef = ref<Recordable[]>([]);
    const searchState = reactive<SearchState>({
        sortInfo: {},
        filterInfo: {},
    });

    watchEffect(() => {
        tableData.value = unref(dataSourceRef);
    });

    watch(
        () => unref(propsRef).data,
        () => {
            const { data, api } = unref(propsRef);
            !api && data && (dataSourceRef.value = data);
        },
        {
            immediate: true,
        }
    );

    async function reload(opt?: FetchParams) {
        await fetch(opt);
    }

    async function fetch(opt?: FetchParams) {
        try {
            setLoading(true);

            let pageParams: Recordable = {};

            const { api, pagination } = unref(propsRef);
            if (!api) return;

            const { page = 1, pageSize = PAGE_SIZE } = unref(getPaginationInfo) as PaginationProps;
            if ((isBoolean(pagination) && !pagination) || isBoolean(getPaginationInfo)) {
                pageParams = {};
            } else {
                pageParams.page = (opt && opt.page) || page;
                pageParams.pageSize = pageSize;
            }

            const params: Recordable = merge(pageParams, {});
            const { data, total } = await api(params);

            setPagination({ itemCount: total });
            dataSourceRef.value = data;
        } finally {
            setLoading(false);
        }
    }

    const getDataSourceRef = computed(() => {
        return unref(dataSourceRef);
    });

    onMounted(() => {
        setTimeout(() => fetch(), 16);
    });

    function handleSorterChange(options: DataTableSortState | DataTableSortState[] | null) {
        console.log(options);
        // todo
    }

    async function handlePageSizeChange(pageSize: number) {
        setPagination({ pageSize });
        await reload();
    }

    async function handlePageChange(page: number) {
        setPagination({ page });
        await reload();
    }

    async function handleFilterChange(filters: DataTableFilterState, initiatorColumn: DataTableColumn) {
        console.log(filters, initiatorColumn);
    }

    function setTableData<T = Recordable>(values: T[]) {
        dataSourceRef.value = values;
    }

    function setRowData<T = Recordable>(row: number, values: T) {
        dataSourceRef.value[row] = values;
    }

    function setFieldData(row: number, field: string, value: any) {
        dataSourceRef.value[row][field] = value;
    }

    return { reload, setTableData, setRowData, setFieldData, getDataSourceRef, handleSorterChange, handlePageChange, handlePageSizeChange, handleFilterChange };
}
