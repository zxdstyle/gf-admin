import type { ComputedRef, Ref } from 'vue';

declare interface Fn<T = any, R = T> {
    (...arg: T[]): R;
}

interface Window {
    $loadingBar?: import('naive-ui').LoadingBarProviderInst;
    $dialog?: import('naive-ui').DialogProviderInst;
    $message?: import('naive-ui').MessageProviderInst;
    $notification?: import('naive-ui').NotificationProviderInst;
}

/** 通用类型 */
declare namespace Common {
    /**
     * 策略模式
     * [状态, 为true时执行的回调函数]
     */
    type StrategyAction = [boolean, () => void];
}

/** 构建时间 */
declare const PROJECT_BUILD_TIME: string;

declare type Nullable<T> = T | null;
declare type Recordable<T = any> = Record<string, T>;

export type DynamicProps<T> = {
    [P in keyof T]: Ref<T[P]> | T[P] | ComputedRef<T[P]>;
};

declare type EmitType = (event: 'register', ...args: any[]) => void;
