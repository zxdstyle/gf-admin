import type { ComputedRef, Ref } from 'vue';
import { unref } from 'vue';
import type { Recordable } from '@/typings/global';
import { isNullOrUnDef } from '@/utils/common/is';
import type { FormProps, FormSchema } from '../types/form';

interface UseFormValuesContext {
    defaultValueRef: Ref<any>;
    getSchema: ComputedRef<FormSchema[]>;
    getProps: ComputedRef<FormProps>;
    formModel: Recordable;
}

export default function useFormValues({ defaultValueRef, getSchema, formModel, getProps }: UseFormValuesContext) {
    function initDefault() {
        const schemas = unref(getSchema);
        const obj: Recordable = {};
        schemas.forEach((item) => {
            const { defaultValue } = item;

            if (!isNullOrUnDef(defaultValue)) {
                obj[item.field] = defaultValue;
                formModel[item.field] = defaultValue;
            }
        });
        defaultValueRef.value = obj;
    }

    return { initDefault };
}
