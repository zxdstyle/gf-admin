<template>
    <div ref="wrapRef" class="h-full">
        <TableHeader>
            <template #toolbar><slot name="toolbar"></slot></template>
        </TableHeader>
        <n-data-table ref="tableElRef" v-bind="getBindValues" :style="resizeStyle" />
    </div>
</template>

<script lang="ts">
import { computed, defineComponent, ref, toRaw, unref } from 'vue';
import type { Recordable } from '@/typings/global';
import { useDataSource } from '@/components/basic/table/src/hooks/useDataSource';
import { usePagination } from '@/components/basic/table/src/hooks/usePagination';
import TableHeader from './components/TableHeader.vue';
import { useLoading } from './hooks/useLoading';
import type { BasicTableProps, TableActionType } from './types/table';
import { createTableContext } from './hooks/useTableContext';

export default defineComponent({
    name: 'BasicTable',
    components: { TableHeader },
    emits: ['register'],
    setup(props, { attrs, emit }) {
        const wrapRef = ref(null);
        const tableElRef = ref(null);
        const tableData = ref<Recordable[]>([]);

        const innerPropsRef = ref<Partial<BasicTableProps>>();

        const getProps = computed(() => {
            return { ...props, ...unref(innerPropsRef) } as BasicTableProps;
        });

        function setProps(props: Partial<BasicTableProps>) {
            innerPropsRef.value = { ...unref(innerPropsRef), ...props };
        }

        const { setLoading, getLoading } = useLoading(getProps);

        const { getPaginationInfo, setPagination } = usePagination(getProps);

        const { reload, setRowData, setTableData, setFieldData, getDataSourceRef, handleSorterChange, handlePageChange, handlePageSizeChange, handleFilterChange } = useDataSource(getProps, {
            setLoading,
            tableData,
            setPagination,
            getPaginationInfo,
        });

        const getBindValues = computed(() => {
            const data = unref(getDataSourceRef);
            const propsData: Recordable = {
                pagination: toRaw(unref(getPaginationInfo)),
                loading: unref(getLoading),
                remote: true,
                tableLayout: 'fixed',
                striped: true,
                flexHeight: true,
                'onUpdate:page': handlePageChange,
                'onUpdate:pageSize': handlePageSizeChange,
                'onUpdate:sorter': handleSorterChange,
                'onUpdate:filters': handleFilterChange,
                ...attrs,
                ...unref(getProps),
                data,
            };
            return propsData;
        });

        const tableAction: TableActionType = {
            reload,
            setProps,
            setTableData,
            setRowData,
            setFieldData,
        };

        const resizeStyle = computed(() => {
            if (unref(getBindValues).flexHeight) {
                return { height: 'calc(100% - 20px)' };
            }
            return {};
        });

        createTableContext({ ...tableAction, wrapRef, getBindValues });

        emit('register', tableAction);

        return { wrapRef, tableElRef, getBindValues, resizeStyle };
    },
});
</script>
