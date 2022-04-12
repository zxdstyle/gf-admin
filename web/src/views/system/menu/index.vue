<template>
    <div class="grid grid-cols-12 gap-6 h-full">
        <n-card class="h-full shadow-sm rounded-10px col-span-8">
            <template #header>
                <div class="flex justify-end">
                    <n-button-group size="small">
                        <n-button type="default" @click="expandAll">
                            <template #icon>
                                <Icon icon="mdi:plus-box-multiple-outline" />
                            </template>
                            展开
                        </n-button>
                        <n-button type="default" @click="collapseAll">
                            <template #icon>
                                <Icon icon="mdi:minus-box-multiple-outline" />
                            </template>
                            收起
                        </n-button>
                    </n-button-group>
                    <n-button class="ml-3" type="default" size="small" @click="getMenuData">
                        <template #icon>
                            <Icon icon="mi:refresh" class="text-base" />
                        </template>
                        刷新
                    </n-button>
                </div>
            </template>
            {{ getTreeBind }}
            <n-spin :show="loading">
                <n-tree v-bind="getTreeBind"></n-tree>

                <template #description> 加载中... </template>
            </n-spin>
        </n-card>
        <n-card class="h-full shadow-sm rounded-10px col-span-4" title="添加菜单">
            <MenuForm v-bind="getFormBind" />
        </n-card>
    </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { Icon } from '@iconify/vue';
import useData from './useData';
import MenuForm from './form.vue';

export default defineComponent({
    name: 'Index',
    components: { Icon, MenuForm },
    setup() {
        const { getFormBind, getTreeBind, getMenuData, loading, expandAll, collapseAll } = useData();

        return { getFormBind, getTreeBind, getMenuData, loading, expandAll, collapseAll };
    },
});
</script>
