<template>
    <n-popover placement="bottom-end" trigger="click">
        <template #trigger>
            <Icon icon="nimbus:cog" class="text-xl cursor-pointer focus:outline-none hover:text-primary" />
        </template>

        <template #header>
            <div>
                <n-checkbox v-model:checked="flexHeight">自适应高度</n-checkbox>
            </div>
        </template>
    </n-popover>
</template>

<script lang="ts">
import { computed, defineComponent, reactive, toRefs, unref } from 'vue';
import { Icon } from '@iconify/vue';
import { useTableContext } from '../../hooks/useTableContext';

export default defineComponent({
    name: 'ColumnSetting',
    components: { Icon },
    setup() {
        const table = useTableContext();

        const state = reactive({});

        const flexHeight = computed({
            get: () => unref(table.getBindValues).flexHeight,
            set: (val) => table.setProps({ flexHeight: val }),
        });

        return { ...toRefs(state), flexHeight };
    },
});
</script>

<style lang="less" scoped></style>
