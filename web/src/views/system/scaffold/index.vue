<template>
    <n-card class="h-full shadow-sm rounded-16px">
        <n-form ref="formRef" :model="model" :rules="rules" label-placement="left" label-align="right" label-width="120" :show-require-mark="false">
            <n-form-item path="table" label="Table">
                <n-select v-model:value="model.table" :style="{ maxWidth: '320px' }" :options="tables" @update:value="changeTable" />
            </n-form-item>
            <n-form-item path="model" label="Model">
                <n-input v-model:value="model.model" :style="{ maxWidth: '480px' }" />
            </n-form-item>
            <n-form-item path="handler" label="Handler">
                <n-input v-model:value="model.handler" :style="{ maxWidth: '480px' }" />
            </n-form-item>
            <n-form-item path="repository" label="Repository">
                <n-input v-model:value="model.repository" :style="{ maxWidth: '480px' }" />
            </n-form-item>
            <n-form-item v-if="model.create_view" path="view_path" label="View">
                <n-input v-model:value="model.view_path" :style="{ maxWidth: '480px' }" />
            </n-form-item>
            <n-form-item class="ml-32">
                <n-checkbox v-model:checked="model.create_view">Create View</n-checkbox>
                <n-checkbox v-model:checked="model.create_model">Create Model</n-checkbox>
                <n-checkbox v-model:checked="model.create_repository">Create Repository</n-checkbox>
                <n-checkbox v-model:checked="model.create_handler">Create Handler</n-checkbox>
                <n-checkbox v-model:checked="model.create_lang">Create Lang</n-checkbox>
            </n-form-item>
        </n-form>

        <n-table>
            <n-thead>
                <n-tr>
                    <n-th>Struct Name</n-th>
                    <n-th>Json Name</n-th>
                    <n-th>Translation</n-th>
                    <n-th>Type</n-th>
                    <n-th>Nullable</n-th>
                    <n-th>Default</n-th>
                    <n-th>Comment</n-th>
                    <n-th>Hidden</n-th>
                </n-tr>
            </n-thead>

            <draggable v-if="model.fields.length > 0" :list="model.fields" item-key="index" tag="tbody">
                <template #item="{ element }">
                    <tr class="cursor-move">
                        <td><n-input v-model:value="element.name" /></td>
                        <td><n-input v-model:value="element.json_name" /></td>
                        <td><n-input v-model:value="element.translation" /></td>
                        <td><n-input v-model:value="element.type" /></td>
                        <td><n-checkbox v-model:checked="element.nullable" /></td>
                        <td><n-input v-model:value="element.default_value" /></td>
                        <td><n-input v-model:value="element.comment" /></td>
                        <td><n-checkbox v-model:checked="element.hidden" /></td>
                    </tr>
                </template>
            </draggable>
            <tbody v-else>
                <tr>
                    <td colspan="9" class="py-12">
                        <n-empty />
                    </td>
                </tr>
            </tbody>
        </n-table>

        <div class="flex justify-end mt-12 pr-12">
            <n-button type="primary" class="ml-2" :loading="loading" @click="submit">
                <Icon icon="ion:paper-plane" class="text-lg mr-2" />
                提交
            </n-button>
        </div>

        <BasicModal title="代码预览" style="width: 76%; min-height: 820px" @register="register"><Preview :data="preview" /></BasicModal>
    </n-card>
</template>

<script lang="ts">
import { defineComponent, reactive, toRefs, onMounted, ref } from 'vue';
import { Icon } from '@iconify/vue';
import draggable from 'vuedraggable';
import { camelCase, snakeCase, upperFirst } from 'lodash-es';
import type { FormInst } from 'naive-ui';
import { useLoading } from '@/hooks';
import { fetchTables, fetchSingular, fetchColumns, previewCode } from '@/service/api/system';
import { BasicModal, useModal } from '@/components/basic/modal';
import type { State } from './typing';
import Preview from './components/Preview.vue';

const defaultModelPath = `server\\app\\model\\`;
const defaultHandlerPath = `server\\app\\handler\\`;
const defaultRepositoryPath = `server\\app\\repository\\`;
const defaultViewPath = 'frontend\\src\\views\\';

export default defineComponent({
    name: 'Index',
    components: { draggable, BasicModal, Preview, Icon },
    setup() {
        const formRef = ref<FormInst | null>(null);
        const state: State = reactive({
            tables: [],
            model: {
                table: '',
                model: defaultModelPath,
                handler: defaultHandlerPath,
                repository: defaultRepositoryPath,
                view_path: defaultViewPath,
                create_view: true,
                create_handler: true,
                create_lang: true,
                create_model: true,
                create_repository: true,
                fields: [],
            },
            columns: [],
            preview: [],
            rules: {
                table: [{ required: true, message: '请选择数据库表', trigger: 'blur' }],
                model: [{ required: true, message: '请输入Model文件目录', trigger: 'blur' }],
                handler: [{ required: true, message: '请输入Handler文件目录', trigger: 'blur' }],
                repository: [{ required: true, message: '请输入Repository目录', trigger: 'blur' }],
            },
        });

        const getTables = async () => {
            const data = await fetchTables();
            if (data) {
                state.tables = data;
            }
        };

        const genFieldsByColumns = (columns: ApiScaffold.column[]) => {
            const fields: ApiScaffold.Field[] = [];
            columns.forEach((column, index) => {
                fields.push({
                    ...column,
                    name: upperFirst(camelCase(column.name)),
                    translation: column.comment,
                    json_name: snakeCase(column.name),
                    index,
                    label: column.label ? column.label : column.name,
                    hidden: ['id', 'created_at', 'updated_at'].includes(column.name),
                } as ApiScaffold.Field);
            });
            state.model.fields = fields;
        };

        const getColumns = async () => {
            const data = await fetchColumns(state.model.table);
            state.columns = data || [];
            genFieldsByColumns(state.columns);
        };

        const changeTable = async () => {
            const data = await fetchSingular(state.model.table);
            const res = snakeCase(data || '');
            state.model.model = `${defaultModelPath}${res}`;
            state.model.handler = `${defaultHandlerPath}${res}`;
            state.model.repository = `${defaultRepositoryPath}${res}`;
            state.model.view_path = `${defaultViewPath}${res}`;
            await getColumns();
        };

        onMounted(() => getTables());

        const [register, { openModal }] = useModal();

        const { loading, startLoading, endLoading } = useLoading();

        const submit = async () => {
            const errors = await formRef.value?.validate();
            if (errors) return;

            try {
                startLoading();
                state.preview = await previewCode(state.model);
                openModal();
            } finally {
                endLoading();
            }
        };

        return { ...toRefs(state), loading, changeTable, register, openModal, submit, formRef };
    },
});
</script>

<style lang="less" scoped></style>
