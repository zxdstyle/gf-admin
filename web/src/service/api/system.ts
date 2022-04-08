import { request } from '../request';

// 获取所有数据表
export function fetchTables() {
    return request.get<ApiCommon.Option[]>('/api/v1/system/tables');
}

// 转换为单数
export function fetchSingular(word: string) {
    return request.get<string>('/api/v1/system/singular', { params: { word } });
}

// 获取数据表信息
export function fetchColumns(table: string) {
    return request.get<ApiScaffold.column[]>('/api/v1/system/columns', { params: { table } });
}

export function previewCode(data: ApiScaffold.FormModel) {
    return request.post<ApiScaffold.Preview[]>(`/api/v1/system/preview`, data);
}
