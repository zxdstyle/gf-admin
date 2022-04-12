<template>
    <NTreeSelect v-bind="getAttrs"></NTreeSelect>
</template>

<script lang="ts">
import type { PropType } from 'vue';
import { computed, defineComponent, onMounted, ref, unref, watch } from 'vue';
import { NTreeSelect } from 'naive-ui';
import { get, isArray, isFunction } from 'lodash-es';
import { useLoading } from '@/hooks';
import type { Recordable } from '@/typings/global';
import { propTypes } from '@/utils/common/propTypes';

export default defineComponent({
    name: 'ApiSelect',
    components: { NTreeSelect },
    props: {
        api: { type: Function as PropType<(arg?: Recordable) => Promise<Recordable>>, required: true },
        params: { type: Object, default: () => {} },
        resultField: propTypes.string.def(''),
        afterFetch: { type: Function, default: () => {} },
    },
    emits: ['options-change'],
    setup(props, { emit, attrs }) {
        const { startLoading, endLoading, loading } = useLoading();
        const treeData = ref<Recordable[]>([]);
        const isFirstLoaded = ref<boolean>(false);
        const getAttrs = computed(() => {
            return {
                ...attrs,
            };
        });

        watch(
            () => props.params,
            () => {
                !unref(isFirstLoaded) && fetch();
            },
            { deep: true }
        );

        onMounted(() => fetch());

        async function fetch() {
            const { api } = props;
            if (!api || !isFunction(api)) return;

            startLoading();
            treeData.value = [];
            let result;
            try {
                result = await api(props.params);
            } catch (e) {
                console.error(e);
            }

            endLoading();
            if (!result) return;
            if (!isArray(result)) {
                result = get(result, props.resultField);
            }
            result = props.afterFetch(result);
            treeData.value = (result as Recordable[]) || [];
            isFirstLoaded.value = true;
            emit('options-change', treeData.value);
        }

        return { getAttrs, loading };
    },
});
</script>

<style lang="less" scoped></style>
