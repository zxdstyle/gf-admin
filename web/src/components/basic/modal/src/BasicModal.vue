<template>
    <n-modal v-bind="getBindValue" @update:show="handleCancel">
        <slot></slot>
    </n-modal>
</template>

<script lang="ts">
import { computed, defineComponent, getCurrentInstance, nextTick, ref, unref } from 'vue';
import type { ModalProps } from 'naive-ui';
import { omit, merge } from 'lodash-es';
import type { ModalMethods } from './typing';

export default defineComponent({
    name: 'BasicModal',
    components: {},
    emits: ['cancel', 'register'],
    setup(props, { emit, attrs }) {
        const visibleRef = ref(false);
        const propsRef = ref<Partial<ModalProps> | null>(null);
        const modalWrapperRef = ref<any>(null);

        const extHeightRef = ref(0);
        const modalMethods: ModalMethods = {
            setModalProps,
            emitVisible: undefined,
            redoModalHeight: () => {
                nextTick(() => {
                    if (unref(modalWrapperRef)) {
                        (unref(modalWrapperRef) as any).setModalHeight();
                    }
                });
            },
        };

        const instance = getCurrentInstance();
        if (instance) {
            emit('register', modalMethods, instance.uid);
        }

        const getMergeProps = computed(() => {
            return {
                preset: 'card',
                style: 'width: 760px;min-height: 420px',
                trapFocus: true,
                ...props,
                ...(unref(propsRef) as any),
            };
        });

        const getBindValue = computed(() => {
            return {
                ...attrs,
                ...unref(getMergeProps),
                show: unref(visibleRef),
            };
        });

        function setModalProps(props: Partial<ModalProps>): void {
            // Keep the last setModalProps
            propsRef.value = merge(unref(propsRef) || ({} as any), props);
            if (Reflect.has(props, 'show')) {
                visibleRef.value = !!props.show;
            }
        }

        async function handleCancel(show: boolean) {
            visibleRef.value = show;
            emit('cancel');
        }

        return { handleCancel, getBindValue };
    },
});
</script>

<style lang="less" scoped></style>
