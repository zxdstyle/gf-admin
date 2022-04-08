<template>
    <BasicForm @register="registerForm"></BasicForm>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import ApiUser from '@/service/api/users';
import { BasicForm, useForm } from '@/components/basic/form';

export default defineComponent({
    name: 'UserForm',
    components: { BasicForm },
    props: {
        model: {
            type: Object,
            default: () => {},
        },
    },
    emits: ['submit'],
    setup(props, { emit }) {
        const [registerForm, { resetFields }] = useForm({
            model: props.model,
            schemas: [
                { field: 'username', component: 'Input', label: '账户', componentProps: { placeholder: '请输入用户账户' } },
                { field: 'nickname', component: 'Input', label: '昵称' },
                { field: 'email', component: 'Input', label: 'Email' },
                { field: 'is_active', component: 'Switch', label: '是否启用', defaultValue: true },
                { field: 'password', component: 'Input', label: '初始密码', componentProps: { type: 'password', showPasswordOn: 'click' } },
            ],
            rules: {
                username: [{ required: true }],
                email: [{ required: true, type: 'email' }],
            },
            onSubmit: async (e) => {
                let res;
                if (e.id && e.id > 0) {
                    res = await ApiUser.Update(e.id, e);
                } else {
                    res = await ApiUser.Create(e);
                }
                await resetFields();
                emit('submit', res);
            },
        });

        return { registerForm };
    },
});
</script>
