import type { Component } from 'vue';
import { NInput, NSelect, NRadio, NCheckbox, NAutoComplete, NCascader, NDatePicker, NInputNumber, NSwitch, NTimePicker, NTreeSelect, NSlider, NRate, NDivider } from 'naive-ui';
import IconPicker from '@/components/basic/icon/IconPicker.vue';
import type { ComponentType } from './types/index';
import ApiTreeSelect from './components/ApiTreeSelect.vue';
// import ApiRadioGroup from './components/ApiRadioGroup.vue';
// import RadioButtonGroup from './components/RadioButtonGroup.vue';
// import ApiSelect from './components/ApiTreeSelect.vue';
// import ApiTree from './components/ApiTree.vue';
// import ApiCascader from './components/ApiCascader.vue';
// import { BasicUpload } from '/@/components/Upload';
// import { StrengthMeter } from '/@/components/StrengthMeter';
// import { CountdownInput } from '/@/components/CountDown';

const componentMap = new Map<ComponentType, Component>();

componentMap.set('Input', NInput);
componentMap.set('InputNumber', NInputNumber);
componentMap.set('AutoComplete', NAutoComplete);

componentMap.set('Select', NSelect);
// componentMap.set('ApiSelect', ApiSelect);
// componentMap.set('ApiTree', ApiTree);
componentMap.set('TreeSelect', NTreeSelect);
componentMap.set('ApiTreeSelect', ApiTreeSelect);
// componentMap.set('ApiRadioGroup', ApiRadioGroup);
componentMap.set('Switch', NSwitch);
// componentMap.set('RadioButtonGroup', RadioButtonGroup);
componentMap.set('RadioGroup', NRadio.Group);
componentMap.set('Checkbox', NCheckbox);
componentMap.set('CheckboxGroup', NCheckbox.Group);
// componentMap.set('ApiCascader', ApiCascader);
componentMap.set('Cascader', NCascader);
componentMap.set('Slider', NSlider);
componentMap.set('Rate', NRate);

componentMap.set('DatePicker', NDatePicker);
componentMap.set('MonthPicker', NDatePicker.MonthPicker);
componentMap.set('RangePicker', NDatePicker.RangePicker);
componentMap.set('WeekPicker', NDatePicker.WeekPicker);
componentMap.set('TimePicker', NTimePicker);
// componentMap.set('StrengthMeter', StrengthMeter);
componentMap.set('IconPicker', IconPicker);
// componentMap.set('InputCountDown', CountdownInput);
//
// componentMap.set('Upload', BasicUpload);
componentMap.set('Divider', NDivider);

export function add(compName: ComponentType, component: Component) {
    componentMap.set(compName, component);
}

export function del(compName: ComponentType) {
    componentMap.delete(compName);
}

export { componentMap };
