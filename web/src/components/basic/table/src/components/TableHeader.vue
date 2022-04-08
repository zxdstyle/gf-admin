<template>
    <div class="flex items-center justify-between pb-2 relative">
        <div></div>
        <div class="flex items-center">
            <slot name="toolbar"></slot>
            <n-divider vertical />
            <Icon icon="mi:refresh" class="text-xl cursor-pointer hover:text-primary" @click="reload" />

            <n-tooltip>
                <template #trigger>
                    <Icon icon="ant-design:column-height-outlined" class="text-xl mx-1 cursor-pointer focus:outline-none hover:text-primary" @click="changeTableSize" />
                </template>
                表格尺寸
            </n-tooltip>

            <ColumnSetting />
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent, unref } from 'vue';
import { Icon } from '@iconify/vue';
import { useTableContext } from '../hooks/useTableContext';
import ColumnSetting from './settings/ColumnSetting.vue';

export default defineComponent({
    name: 'TableHeader',
    components: { Icon, ColumnSetting },
    setup() {
        const table = useTableContext();

        const reload = () => table.reload();

        const changeTableSize = () => {
            let { size } = unref(table.getBindValues);
            console.log(size);
            switch (size) {
                case 'medium':
                    size = 'large';
                    break;
                case 'large':
                    size = 'small';
                    break;
                case 'small':
                    size = 'medium';
                    break;
                default:
                    size = 'large';
                    break;
            }
            table.setProps({ size });
        };

        return { reload, changeTableSize };
    },
});
</script>
