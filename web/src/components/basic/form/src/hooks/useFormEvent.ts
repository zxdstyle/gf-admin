import type { ComputedRef, Ref } from 'vue';
import { unref } from 'vue';
import dayjs from 'dayjs';
import type { EmitType, Recordable } from '@/typings/global';
import { dateItemType, defaultValueComponents, handleInputNumberValue } from '../helper';
import type { FormProps, FormSchema } from '../types/form';

interface UseFormActionContext {
    emit: EmitType;
    getProps: ComputedRef<FormProps>;
    getSchema: ComputedRef<FormSchema[]>;
    formModel: Recordable;
    defaultValueRef: Ref<Recordable>;
}

export default function useFormEvent({ emit, getProps, formModel, getSchema, defaultValueRef }: UseFormActionContext) {
    async function resetFields(): Promise<void> {
        Object.keys(formModel).forEach((key) => {
            const schema = unref(getSchema).find((item) => item.field === key);
            let val;

            const { model } = unref(getProps) || {};
            if (model && model[key]) {
                val = model[key];
            } else if (defaultValueRef.value[key]) {
                val = defaultValueRef.value[key];
            } else {
                const isInput = schema?.component && defaultValueComponents.includes(schema.component);
                val = isInput ? defaultValueRef.value[key] || '' : val;
            }

            formModel[key] = val;
        });
    }

    /**
     * @description: Set form value
     */
    async function setFieldsValue(values: Recordable): Promise<void> {
        Object.keys(values).forEach((key) => {
            const schema = unref(getSchema).find((item) => item.field === key);
            let value = values[key];

            const hasKey = Reflect.has(values, key);

            value = handleInputNumberValue(schema?.component, value);
            // 0| '' is allow
            if (hasKey) {
                // time type
                if (itemIsDateType(key)) {
                    if (Array.isArray(value)) {
                        const arr: any[] = [];
                        for (const ele of value) {
                            arr.push(ele ? dayjs(ele) : null);
                        }
                        formModel[key] = arr;
                    } else {
                        const { componentProps } = schema || {};
                        let _props = componentProps as any;
                        if (typeof componentProps === 'function') {
                            _props = _props({ formModel });
                        }
                        formModel[key] = value ? (_props?.valueFormat ? value : dayjs(value)) : null;
                    }
                } else {
                    formModel[key] = value;
                }
            }
        });
    }

    function itemIsDateType(key: string) {
        return unref(getSchema).some((item) => {
            return item.field === key ? dateItemType.includes(item.component) : false;
        });
    }

    return { resetFields, setFieldsValue };
}
