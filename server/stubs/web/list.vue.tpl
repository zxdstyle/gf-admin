<template>
    <n-card class="h-full shadow-sm rounded-10px">
        <BasicTable @register="registerTable">
            <template #toolbar>
                <n-button type="primary" size="small" @click="openCreateModal">
                    <Icon icon="ic:outline-plus" class="text-xl"></Icon>
                    新增用户
                </n-button>
            </template>
        </BasicTable>

        <BasicModal @register="registerModal"><UserForm v-bind="getFormBind" /></BasicModal>
    </n-card>
</template>

<script lang="tsx">
import { defineComponent } from 'vue';
import { Icon } from '@iconify/vue';
import { BasicTable } from '@/components/basic/table';
import { BasicModal } from '@/components/basic/modal';
import { useData } from './useData';
import UserForm from './form.vue';

export default defineComponent({
    name: 'Index',
    components: { BasicTable, BasicModal, UserForm, Icon },
    setup() {
        const { registerModal, registerTable, openCreateModal, getFormBind } = useData();

        return { registerTable, registerModal, openCreateModal, getFormBind };
    },
});
</script>

<style lang="less" scoped></style>
