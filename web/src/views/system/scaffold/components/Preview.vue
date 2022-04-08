<template>
    <div class="h-full">
        <div class="grid grid-cols-12 overflow-hidden" style="height: 758px">
            <div class="border-r h-full col-span-3">
                <n-tree :data="files" :default-expand-all="true" :leaf-only="true" :render-prefix="renderPrefix" :selectable="true" :node-props="nodeProps"></n-tree>
            </div>
            <div class="col-span-9 overflow-hidden">
                <div class="px-2 h-full overflow-hidden">
                    <div class="flex items-center mb-3">
                        <Icon v-if="content.lang === 'go'" icon="logos:gopher" class="mr-2 text-lg"></Icon>
                        <Icon v-else-if="content.lang === 'vue'" icon="logos:vue" class="mr-2 text-base"></Icon>
                        <Icon v-else-if="content.lang === 'javascript'" icon="vscode-icons:file-type-typescriptdef" class="mr-2 text-lg"></Icon>
                        <Icon v-else icon="bi:file-earmark-code-fill" class="mr-2 text-lg"></Icon>
                        <h3 class="text-base">{{ content.path }}</h3>
                        <Icon icon="ic:baseline-content-copy" class="ml-2 text-base cursor-pointer hover:text-primary"></Icon>
                    </div>
                    <div class="parent">
                        <div class="content">
                            <highlight ref="editorRef" class="rounded-md h-full min-h-full" v-bind="options" :code="content.content" autodetect></highlight>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="tsx">
import { computed, defineComponent, reactive, ref, toRefs, h } from 'vue';
import type { PropType } from 'vue';
import type { TreeOption } from 'naive-ui';
import { Icon } from '@iconify/vue';
import 'highlight.js/lib/common';
import highlight from '@highlightjs/vue-plugin';
import 'highlight.js/styles/atom-one-dark.css';

export default defineComponent({
    name: 'Preview',
    components: { Icon, highlight: highlight.component },
    props: {
        data: {
            type: Array as PropType<ApiScaffold.Preview[]>,
            required: true,
        },
    },
    setup(props) {
        const editorRef = ref(null);
        const state = reactive({
            options: {
                language: 'js',
            },
            activePath: '',
        });

        const findContent = (data: ApiScaffold.Preview[], key = ''): ApiScaffold.Preview => {
            if (key.length === 0) {
                for (let i = 0; i < data.length; i++) {
                    const item = data[i];
                    if (item.content && item.content.length > 0) return item;

                    if (item.children && item.children.length > 0) {
                        return findContent(item.children);
                    }
                }
            }

            for (let i = 0; i < data.length; i++) {
                const item = data[i];
                if (item.path === key && item.content && item.content.length > 0) return item;

                if (item.children && item.children.length > 0) {
                    const val = findContent(item.children, key);
                    if (val.content && val.content.length > 0) {
                        return val;
                    }
                }
            }

            return {};
        };

        const content = computed(() => findContent(props.data, state.activePath));

        const files = computed(() => {
            const search = (items: ApiScaffold.Preview[]) => {
                const data: any = [];
                items.forEach((item) => {
                    let icon = 'mdi:folder-open';
                    if (!item.children || item.children?.length === 0) {
                        switch (item.lang) {
                            case 'go':
                                icon = 'logos:gopher';
                                break;
                            case 'javascript':
                                icon = 'vscode-icons:file-type-typescriptdef';
                                break;
                            case 'vue':
                                icon = 'logos:vue';
                                break;
                            default:
                                icon = 'bi:file-earmark-code-fill';
                        }
                    }
                    const val = {
                        key: item.path,
                        label: item.path,
                        prefix: icon,
                    };

                    if (item.children && item.children.length > 0) {
                        // @ts-ignore
                        val.children = search(item.children);
                    }
                    data.push(val);
                });
                return data;
            };
            return search(props.data);
        });

        const renderPrefix = (info: { option: TreeOption }) => {
            const icon = info.option.prefix;
            return h(Icon, { icon });
        };

        const changeFile = (key: any) => {
            state.activePath = key;
        };

        const nodeProps = ({ option }: { option: TreeOption }) => {
            return {
                onClick() {
                    changeFile(option.key);
                },
            };
        };

        return { ...toRefs(state), content, editorRef, files, renderPrefix, changeFile, nodeProps };
    },
});
</script>

<style scoped>
.parent {
    width: 100%;
    height: 100%;
    overflow: hidden;
}

.content {
    height: 100%;

    /* 隐藏y轴滚动条 */
    margin-right: -50px; /* Maximum width of scrollbar */
    padding-right: 50px; /* Maximum width of scrollbar */
    overflow-y: scroll;

    /* 隐藏x轴滚动条 */
    overflow-x: scroll;
}
</style>
