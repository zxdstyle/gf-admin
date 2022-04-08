<script lang="tsx">
import { computed, defineComponent, unref } from 'vue';
import type { PropType } from 'vue';
import { isBoolean, isFunction, upperFirst } from 'lodash-es';
import type { Nullable, Recordable } from '@/typings/global';
import { componentMap } from '@/components/basic/form/src/componentMap';
import type { FormProps, FormSchema } from '../types/form';

export default defineComponent({
    name: 'FormItem',
    props: {
        schema: {
            type: Object as PropType<FormSchema>,
            default: () => ({}),
        },
        formProps: {
            type: Object as PropType<FormProps>,
            default: () => ({}),
        },
        allDefaultValues: {
            type: Object as PropType<Recordable>,
            default: () => ({}),
        },
        formModel: {
            type: Object as PropType<Recordable>,
            default: () => ({}),
        },
        setFormModel: {
            type: Function as PropType<(key: string, value: any) => void>,
            default: null,
        },
        componentProps: {
            type: Object as PropType<Recordable>,
            default: () => {},
        },
    },
    setup(props) {
        const getValues = computed(() => {
            const { allDefaultValues, formModel, schema } = props;
            // const { mergeDynamicData } = props.formProps;
            return {
                field: schema.field,
                model: formModel,
                values: {
                    // ...mergeDynamicData,
                    ...allDefaultValues,
                    ...formModel,
                } as Recordable,
                schema,
            };
        });

        function getShow(): { isShow: boolean; isIfShow: boolean } {
            const { show, ifShow } = props.schema;

            let isShow = true;
            let isIfShow = true;

            if (isBoolean(show)) {
                isShow = show;
            }
            if (isBoolean(ifShow)) {
                isIfShow = ifShow;
            }
            if (isFunction(show)) {
                isShow = show(unref(getValues));
            }
            if (isFunction(ifShow)) {
                isIfShow = ifShow(unref(getValues));
            }
            return { isShow, isIfShow };
        }

        function renderComponent() {
            const { isIfShow, isShow } = getShow();
            if (!isIfShow || !isShow) {
                return;
            }
            const { renderComponentContent, component, field, changeEvent = 'update:value', valueField } = props.schema;
            const eventKey = `on${upperFirst(changeEvent)}`;

            const on = {
                [eventKey]: (...args: Nullable<Recordable>[]) => {
                    const [e] = args;
                    const target = e ? e.target : null;
                    const value = target ? target.value : e;
                    props.setFormModel(field, value);
                },
            };

            const Comp = componentMap.get(component) as ReturnType<typeof defineComponent>;
            const bindValue: Recordable = {
                [valueField || 'value']: props.formModel[field],
            };

            const compAttr: Recordable = {
                ...props.componentProps,
                ...on,
                ...bindValue,
            };

            if (!renderComponentContent) {
                return <Comp {...compAttr} />;
            }

            const compSlot = isFunction(renderComponentContent)
                ? { ...renderComponentContent(unref(getValues)) }
                : {
                      default: () => renderComponentContent,
                  };
            return <Comp {...compAttr}>{compSlot}</Comp>;
        }

        return () => {
            const { gridProps = {} } = props.schema;
            const { baseGridProps = {} } = props.formProps;
            const realGridItemProps = { ...baseGridProps, ...gridProps };

            return renderComponent();
        };
    },
});
</script>
