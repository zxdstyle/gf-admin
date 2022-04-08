import { onUnmounted, ref, unref, watch } from 'vue';
import type { WatchStopHandle } from 'vue';
import { values } from 'lodash-es';
import { getDynamicProps } from '@/utils';
import type { DynamicProps, Nullable } from '@/typings/global';
import type { BasicTableProps, FetchParams, TableActionType } from '../types/table';

type Props = Partial<DynamicProps<BasicTableProps>>;

type UseTableMethod = TableActionType;

export function useTable(tableProps?: Props): [(instance: TableActionType, formInstance: UseTableMethod) => void, TableActionType] {
    const tableRef = ref<Nullable<TableActionType>>(null);
    const loadedRef = ref<Nullable<boolean>>(false);

    let stopWatch: WatchStopHandle;

    function register(instance: TableActionType) {
        onUnmounted(() => {
            tableRef.value = null;
            loadedRef.value = null;
        });

        if (unref(loadedRef) && instance === unref(tableRef)) return;

        tableRef.value = instance;
        tableProps && instance.setProps(getDynamicProps(tableProps));
        loadedRef.value = true;

        stopWatch?.();

        stopWatch = watch(
            () => tableProps,
            () => {
                tableProps && instance.setProps(getDynamicProps(tableProps));
            },
            {
                immediate: true,
                deep: true,
            }
        );
    }

    function getTableInstance(): TableActionType {
        const table = unref(tableRef);
        if (!table) {
            console.error('The table instance has not been obtained yet, please make sure the table is presented when performing the table operation!');
        }
        return table as TableActionType;
    }

    const methods: TableActionType = {
        reload: async (opt?: FetchParams) => {
            await getTableInstance().reload(opt);
        },
        setProps: (props: Partial<BasicTableProps>) => getTableInstance().setProps(props),
        setTableData: (values) => getTableInstance().setTableData(values),
        setRowData: (row, value) => getTableInstance().setRowData(row, value),
        setFieldData: (row, field, value) => getTableInstance().setFieldData(row, field, value),
    };

    return [register, methods];
}
