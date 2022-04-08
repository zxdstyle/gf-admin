import type { ComputedRef } from 'vue';
import type { ModalProps } from 'naive-ui';

export interface ModalMethods {
    setModalProps: (props: Partial<ModalProps>) => void;
    emitVisible?: (visible: boolean, uid: number) => void;
    redoModalHeight?: () => void;
}

export type RegisterFn = (modalMethods: ModalMethods, uuid: string) => void;

export interface ReturnMethods extends ModalMethods {
    openModal: <T = any>(show?: boolean, data?: T, openOnSet?: boolean) => void;
    closeModal: () => void;
    getVisible?: ComputedRef<boolean>;
}

export type UseModalReturnType = [RegisterFn, ReturnMethods];
