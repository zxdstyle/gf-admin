// @ts-nocheck

import { computed, onMounted, reactive, toRefs, unref } from 'vue';
import type { TreeOption, TreeProps, TreeSelectOption } from 'naive-ui';
import { NButton } from 'naive-ui';
import { Icon } from '@iconify/vue';
import { useLoading } from '@/hooks';
import apiMenu from '@/service/api/system/menu';
import { listToTree } from '@/utils/helper/treeHelper';

interface State {
    data: ApiMenu.Model[] | undefined;
    allKeys: number[];
    expandKeys: number[];
}

export default function useData() {
    const { startLoading, endLoading, loading } = useLoading();

    const state: State = reactive({
        data: [],
        allKeys: [],
        expandKeys: [],
    });

    const expandAll = () => (state.expandKeys = [0, ...state.allKeys]);

    const getMenuData = async () => {
        startLoading();
        state.data = await apiMenu.Index({ order: 'desc', orderBy: 'sort_num' });
        state.allKeys = [];
        state.data?.forEach((item) => {
            state.allKeys.push(item.id);
        });
        endLoading();
        expandAll();
    };
    onMounted(() => getMenuData());

    const treeData = computed<TreeSelectOption[]>(() => {
        if (!state.data) return [];
        return [
            {
                key: 0,
                label: '根目录',
                icon: 'ph:folder-notch-open-bold',
                children: listToTree(state.data, {
                    pid: 'parent_id',
                    format: (node) => {
                        const val = { key: node.id, label: node.title, ...node, children: undefined };

                        if (node.children && node.children.length > 0) {
                            val.children = node.children;
                        }
                        console.log(node, val);
                        return val;
                    },
                }),
            },
        ];
    });

    const getFormBind = computed(() => {
        return {
            onSubmit: async () => {
                await getMenuData();
            },
            treeData: unref(treeData),
        };
    });

    const getTreeBind = computed<TreeProps>(() => {
        return {
            data: unref(treeData),
            cascade: true,
            blockLine: true,
            blockNode: true,
            draggable: true,
            expandedKeys: state.expandKeys,
            renderPrefix: ({ option }) => {
                return <Icon icon={option.icon} class="text-lg" />;
            },
            renderSuffix: ({ option }) => {
                return (
                    <div class="flex gap-1">
                        <Icon icon="bx:edit" class="text-lg" />
                        <Icon icon="ion:trash-outline" class="text-lg text-red-700" />
                    </div>
                );
            },
            onDrop: async (ele) => {
                console.log(ele);
                switch (ele.dropPosition) {
                    case 'inside':
                        await apiMenu.Update(ele.dragNode.key, { parent_id: ele.node.key });
                        break;
                    case 'after':
                        await apiMenu.Update(ele.dragNode.key, { parent_id: ele.node.parent_id, sort_num: ele.node.sort_num + 1 });
                        break;
                    default:
                        console.log(ele);
                }
                await getMenuData();
            },
            'onUpdate:expandedKeys': (keys) => {
                state.expandKeys = keys;
            },
        };
    });

    return { ...toRefs(state), getFormBind, getTreeBind, getMenuData, loading, expandAll, collapseAll: () => (state.expandKeys = []) };
}
