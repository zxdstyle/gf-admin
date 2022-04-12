import { request } from '@/service/request';
import type { Recordable } from '@/typings/global';

interface IndexReq {
    order?: 'asc' | 'desc';
    orderBy?: string;
    page?: number;
    pageSize?: number;
}

export default class Restful {
    version = 'v1';

    resource = 'users';

    Index = (params: IndexReq = {}) => {
        return request.get(`/api/${this.version}/${this.resource}`, { params });
    };

    Show = (id: number, params = {}) => {
        return request.get(`/api/${this.version}/${this.resource}/${id}`, { params });
    };

    Create = (data: Recordable) => {
        return request.post(`/api/${this.version}/${this.resource}`, data);
    };

    Update = (id: number, data = {}) => {
        return request.put(`/api/${this.version}/${this.resource}/${id}`, data);
    };

    Destroy = (id: number) => {
        return request.delete(`/api/${this.version}/${this.resource}/${id}`, {});
    };
}
