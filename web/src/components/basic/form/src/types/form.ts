import type { CSSProperties, VNode } from 'vue';
import type { ButtonProps, FormItemRule, RowProps, FormProps as NFormProps, GridProps, GridItemProps } from 'naive-ui';
import type { Recordable } from '@/typings/global';
import type { TableActionType } from '@/components/basic/table/src/types/table';
import type { ColEx, ComponentType } from '@/components/basic/form/src/types/index';
import type { FormItem, NamePath } from '@/components/basic/form/src/types/formItem';

export type FieldMapToTime = [string, [string, string], string?][];

export type Rule = FormItemRule & {
    trigger?: 'blur' | 'change' | ['change', 'blur'];
};

export type RegisterFn = (formInstance: FormActionType) => void;

export type UseFormReturnType = [RegisterFn, FormActionType];

export interface RenderCallbackParams {
    schema: FormSchema;
    values: Recordable;
    model: Recordable;
    field: string;
}

export interface HelpComponentProps {
    maxWidth: string;
    // Whether to display the serial number
    showIndex: boolean;
    // Text list
    text: any;
    // colour
    color: string;
    // font size
    fontSize: string;
    icon: string;
    absolute: boolean;
    // Positioning
    position: any;
}

export interface FormActionType {
    submit: () => Promise<void>;
    setFieldsValue: <T>(values: T) => Promise<void>;
    resetFields: () => Promise<void>;
    getFieldsValue: () => Recordable;
    clearValidate: (name?: string | string[]) => Promise<void>;
    updateSchema: (data: Partial<FormSchema> | Partial<FormSchema>[]) => Promise<void>;
    resetSchema: (data: Partial<FormSchema> | Partial<FormSchema>[]) => Promise<void>;
    setProps: (formProps: Partial<FormProps>) => Promise<void>;
    removeSchemaByFiled: (field: string | string[]) => Promise<void>;
    appendSchemaByField: (schema: FormSchema, prefixField: string | undefined, first?: boolean | undefined) => Promise<void>;
    validateFields: (nameList?: NamePath[]) => Promise<any>;
    validate: (nameList?: NamePath[]) => Promise<any>;
    scrollToField: (name: NamePath, options?: ScrollOptions) => Promise<void>;
}

export interface FormProps extends NFormProps {
    // Row configuration for the entire form
    gridProps?: GridProps;
    // General row style
    baseRowStyle?: CSSProperties;
    // General col configuration
    baseGridProps?: GridItemProps;
    // Form configuration rules
    schemas?: FormSchema[];

    onSubmit?: (data: Recordable) => void;

    // layout?: 'vertical' | 'inline' | 'horizontal';
    // // Form value
    // model?: Recordable;
    // // The width of all items in the entire form
    // labelWidth?: number | string;
    // // alignment
    // labelAlign?: 'left' | 'right';
    // // Submit form on reset
    // submitOnReset?: boolean;
    // // Submit form on form changing
    // submitOnChange?: boolean;
    // // Col configuration for the entire form
    // labelCol?: Partial<ColEx>;
    // // Col configuration for the entire form
    // wrapperCol?: Partial<ColEx>;
    //
    //

    //

    // // Function values used to merge into dynamic control form items
    // mergeDynamicData?: Recordable;
    // // Compact mode for search forms
    // compact?: boolean;
    // // Blank line span
    // emptySpan?: number | Partial<ColEx>;
    // // Internal component size of the form
    // size?: 'default' | 'small' | 'large';
    // // Whether to disable
    // disabled?: boolean;
    // // Time interval fields are mapped into multiple
    // fieldMapToTime?: FieldMapToTime;
    // // Placeholder is set automatically
    // autoSetPlaceHolder?: boolean;
    // // Auto submit on press enter on input
    // autoSubmitOnEnter?: boolean;
    // // Check whether the information is added to the label
    // rulesMessageJoinLabel?: boolean;
    // // Whether to show collapse and expand buttons
    // showAdvancedButton?: boolean;
    // // Whether to focus on the first input box, only works when the first form item is input
    // autoFocusFirstItem?: boolean;
    // // Automatically collapse over the specified number of rows
    // autoAdvancedLine?: number;
    // // Always show lines
    // alwaysShowLines?: number;
    // // Whether to show the operation button
    // showActionButtonGroup?: boolean;
    //
    // // Reset button configuration
    // resetButtonOptions?: Partial<ButtonProps>;
    //
    // // Confirm button configuration
    // submitButtonOptions?: Partial<ButtonProps>;
    //
    // // Operation column configuration
    // actionColOptions?: Partial<ColEx>;
    //
    // // Show reset button
    // showResetButton?: boolean;
    // // Show confirmation button
    // showSubmitButton?: boolean;
    //
    // resetFunc?: () => Promise<void>;
    // submitFunc?: () => Promise<void>;
    // transformDateFunc?: (date: any) => string;
    // colon?: boolean;
}

export interface FormSchema {
    // Field name
    field: string;
    // Event name triggered by internal value change, default change
    changeEvent?: string;
    // Variable name bound to v-model Default value
    valueField?: string;
    // Label name
    label: string | VNode;
    // Auxiliary text
    subLabel?: string;
    // Help text on the right side of the text
    helpMessage?: string | string[] | ((renderCallbackParams: RenderCallbackParams) => string | string[]);
    // BaseHelp component props
    helpComponentProps?: Partial<HelpComponentProps>;
    // Label width, if it is passed, the labelCol and WrapperCol configured by itemProps will be invalid
    labelWidth?: string | number;
    // Disable the adjustment of labelWidth with global settings of formModel, and manually set labelCol and wrapperCol by yourself
    disabledLabelWidth?: boolean;
    // render component
    component: ComponentType;
    // Component parameters
    componentProps?: ((opt: { schema: FormSchema; tableAction: TableActionType; formActionType: FormActionType; formModel: Recordable }) => Recordable) | object;
    // Required
    required?: boolean | ((renderCallbackParams: RenderCallbackParams) => boolean);

    suffix?: string | number | ((values: RenderCallbackParams) => string | number);

    // Validation rules
    rules?: Rule[];
    // Check whether the information is added to the label
    rulesMessageJoinLabel?: boolean;

    // Reference formModelItem
    itemProps?: Partial<FormItem>;

    // col configuration outside formModelItem
    gridProps?: GridItemProps;

    // 默认值
    defaultValue?: any;
    isAdvanced?: boolean;

    // Matching details components
    span?: number;

    ifShow?: boolean | ((renderCallbackParams: RenderCallbackParams) => boolean);

    show?: boolean | ((renderCallbackParams: RenderCallbackParams) => boolean);

    // Render the content in the form-item tag
    render?: (renderCallbackParams: RenderCallbackParams) => VNode | VNode[] | string;

    // Rendering col content requires outer wrapper form-item
    renderColContent?: (renderCallbackParams: RenderCallbackParams) => VNode | VNode[] | string;

    renderComponentContent?: ((renderCallbackParams: RenderCallbackParams) => any) | VNode | VNode[] | string;

    // Custom slot, in from-item
    slot?: string;

    // Custom slot, similar to renderColContent
    colSlot?: string;

    dynamicDisabled?: boolean | ((renderCallbackParams: RenderCallbackParams) => boolean);

    dynamicRules?: (renderCallbackParams: RenderCallbackParams) => Rule[];
}