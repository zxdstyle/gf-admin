import type { ComputedRef } from 'vue';
import { computed, ref, unref, watch } from 'vue';
import { isBoolean } from '@/utils/common/is';
import { PAGE_SIZE, PAGE_SIZE_OPTIONS } from '@/components/basic/table/src/const';
import type { PaginationProps } from '../types/pagination';
import type { BasicTableProps } from '../types/table';

export function usePagination(refProps: ComputedRef<BasicTableProps>) {
    const configRef = ref<PaginationProps>({});
    const show = ref(true);

    watch(
        () => unref(refProps).pagination,
        (pagination) => {
            if (!isBoolean(pagination) && pagination) {
                configRef.value = {
                    ...unref(configRef),
                    ...(pagination ?? {}),
                };
            }
        }
    );

    const getPaginationInfo = computed((): PaginationProps | boolean => {
        const { pagination } = unref(refProps);

        if (!unref(show) || (isBoolean(pagination) && !pagination)) {
            return false;
        }

        return {
            defaultPage: 1,
            pageSize: PAGE_SIZE,
            defaultPageSize: PAGE_SIZE,
            showSizePicker: true,
            pageSizes: PAGE_SIZE_OPTIONS,
            showQuickJumper: true,
            ...(isBoolean(pagination) ? {} : pagination),
            ...unref(configRef),
        };
    });

    function setPagination(info: Partial<PaginationProps>) {
        const paginationInfo = unref(getPaginationInfo);
        configRef.value = {
            ...(!isBoolean(paginationInfo) ? paginationInfo : {}),
            ...info,
        };
    }

    function getPagination() {
        return unref(getPaginationInfo);
    }

    function getShowPagination() {
        return unref(show);
    }

    async function setShowPagination(flag: boolean) {
        show.value = flag;
    }

    return { getPagination, getPaginationInfo, setShowPagination, getShowPagination, setPagination };
}
