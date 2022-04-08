import type { FormRules } from 'naive-ui';

interface State {
    tables: ApiCommon.Option[];
    model: ApiScaffold.FormModel;
    columns: ApiScaffold.column[];
    preview: ApiScaffold.Preview[];
    rules: FormRules;
}
