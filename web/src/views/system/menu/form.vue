<template>
    <BasicForm @register="registerForm"></BasicForm>
</template>

<script lang="ts">
import type { PropType } from 'vue';
import { computed, defineComponent } from 'vue';
import type { TreeSelectOption } from 'naive-ui';
import ApiMenu from '@/service/api/system/menu';
import { BasicForm, useForm } from '@/components/basic/form';

export default defineComponent({
    name: 'MenuForm',
    components: { BasicForm },
    props: {
        model: {
            type: Object,
            default: () => {},
        },
        treeData: {
            type: Array as PropType<TreeSelectOption[]>,
            required: true,
        },
    },
    emits: ['submit'],
    setup(props, { emit }) {
        const treeData = computed(() => props.treeData);

        const [registerForm, { resetFields }] = useForm({
            model: props.model,
            labelPlacement: 'left',
            labelAlign: 'right',
            labelWidth: 100,
            schemas: [
                { field: 'parent_id', component: 'TreeSelect', label: '父级菜单', componentProps: { placeholder: '请选择父级菜单', options: treeData, defaultExpandAll: true }, defaultValue: 0 },
                { field: 'title', component: 'Input', label: '标题', componentProps: { placeholder: '请输入标题' } },
                { field: 'icon', component: 'IconPicker', label: '图标', componentProps: { placeholder: '请选择图标' } },
                { field: 'uri', component: 'Input', label: 'URI', componentProps: { placeholder: '请输入URI' } },
                { field: 'order', component: 'Slider', label: '排序值', componentProps: {}, helpMessage: '值越大排序越靠前', defaultValue: 1 },
                { field: 'roles', component: 'Select', label: '角色', componentProps: {}, helpMessage: '可访问该菜单的角色', defaultValue: 1 },
                { field: 'permissions', component: 'Select', label: '权限', componentProps: {}, helpMessage: '可访问该菜单的权限', defaultValue: 1 },
            ],
            rules: {
                parent_id: [{ required: true }],
                title: [{ required: true }],
            },
            onSubmit: async (e) => {
                let res;
                if (e.id && e.id > 0) {
                    res = await ApiMenu.Update(e.id, e);
                } else {
                    res = await ApiMenu.Create(e);
                }
                await resetFields();
                emit('submit', res);
            },
        });

        return { registerForm };
    },
});
</script>
