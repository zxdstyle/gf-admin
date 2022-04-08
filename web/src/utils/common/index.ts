import type { App } from 'vue';

export * from './typeof';
export * from './console';
export * from './color';
export * from './number';
export * from './object';
export * from './icon';
export * from './design-pattern';
export * from './theme';

export const withInstall = <T>(component: T, alias?: string) => {
    const comp = component as any;
    comp.install = (app: App) => {
        app.component(comp.name || comp.displayName, component);
        if (alias) {
            app.config.globalProperties[alias] = component;
        }
    };
    return component as T & Plugin;
};
