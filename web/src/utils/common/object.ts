import { unref } from 'vue';
import type { Recordable } from '@/typings/global';

/** 设置对象数据 */
export function objectAssign<T extends Record<string, any>>(target: T, source: Partial<T>) {
    Object.assign(target, source);
}

export function getDynamicProps<T, U>(props: T): Partial<U> {
    const ret: Recordable = {};

    Object.keys(props).forEach((key) => {
        ret[key] = unref((props as Recordable)[key]);
    });

    return ret as Partial<U>;
}
