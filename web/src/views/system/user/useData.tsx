import { computed, reactive } from 'vue';
import { NSwitch, NButton, NTag } from 'naive-ui';
import { useLoading } from '@/hooks';
import { useTable } from '@/components/basic/table';
import ApiUser from '@/service/api/users';
import { useModal } from '@/components/basic/modal';
import type { Recordable } from '@/typings/global';

export function useData() {
    const state = reactive({
        model: {},
    });

    const { startLoading, endLoading, loading } = useLoading();

    const [registerModal, { openModal, closeModal, setModalProps }] = useModal();

    const [registerTable, { reload, setFieldData }] = useTable({
        api: ApiUser.Index,
        flexHeight: true,
        columns: [
            { key: 'id', title: 'Id', sorter: { multiple: 1 } },
            { key: 'username', title: '账户', sorter: { multiple: 2 } },
            { key: 'email', title: '邮箱' },
            {
                key: 'password',
                title: '密码',
                render(row) {
                    if (row.password && (row.password as string).length > 0) {
                        return <NTag type="success">已修改</NTag>;
                    }
                    return <NTag type="warning">初始密码</NTag>;
                },
            },
            {
                key: 'is_active',
                title: '状态',
                filter: true,
                filterMultiple: false,
                width: 100,
                filterOptions: [
                    { label: '正常', value: 1 },
                    { label: '禁用', value: 0 },
                ],
                render(row, index) {
                    return <NSwitch loading={loading.value} value={row.is_active as boolean} onUpdateValue={(val) => toggleActive(row.id as number, index, val)} />;
                },
            },
            { key: 'created_at', title: '注册时间' },
            {
                key: 'action',
                title: '操作',
                align: 'center',
                render(row) {
                    return (
                        <div>
                            <NButton onClick={() => openEditModal(row)} size="small">
                                编辑
                            </NButton>
                        </div>
                    );
                },
            },
        ],
    });

    async function toggleActive(id: number, row: number, val: boolean) {
        startLoading();
        try {
            await ApiUser.Update(id, { is_active: val });
            setFieldData(row, 'is_active', val);
        } finally {
            endLoading();
        }
    }

    const getFormBind = computed(() => {
        return {
            onSubmit: async () => {
                closeModal();
                await reload();
            },
            model: state.model,
        };
    });

    async function openEditModal(model: Recordable) {
        setModalProps({ title: '编辑用户信息' });
        state.model = model;
        openModal();
    }

    const openCreateModal = () => {
        setModalProps({ title: '新增用户' });
        state.model = {};
        openModal();
    };

    return { registerTable, registerModal, openCreateModal, getFormBind };
}
