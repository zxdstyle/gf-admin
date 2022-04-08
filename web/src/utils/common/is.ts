export function is(val: unknown, type: string) {
    return toString.call(val) === `[object ${type}]`;
}

export function isBoolean(val: unknown): val is boolean {
    return is(val, 'Boolean');
}

export function isFunction(val: unknown): val is Function {
    return typeof val === 'function';
}

export function isNumber(val: unknown): val is number {
    return is(val, 'Number');
}

export function isObject(val: any): val is Record<any, any> {
    return val !== null && is(val, 'Object');
}

export function isArray(val: any): val is Array<any> {
    return val && Array.isArray(val);
}

export function isNull(val: unknown): val is null {
    return val === null;
}

export function isDef<T = unknown>(val?: T): val is T {
    return typeof val !== 'undefined';
}

export function isString(val: unknown): val is string {
    return is(val, 'String');
}

export function isUnDef<T = unknown>(val?: T): val is T {
    return !isDef(val);
}

export function isNullAndUnDef(val: unknown): val is null | undefined {
    return isUnDef(val) && isNull(val);
}

export function isNullOrUnDef(val: unknown): val is null | undefined {
    return isUnDef(val) || isNull(val);
}
