import type { MockMethod } from 'vite-plugin-mock';

const routes: AuthRoute.Route[] = [
    {
        name: 'system_dashboard',
        path: '/system/dashboard',
        component: 'self',
        meta: {
            title: '工作台',
            singleLayout: 'basic',
            icon: 'icon-park-outline:workbench',
            order: 1,
        },
    },
    {
        name: 'system',
        path: '/system',
        component: 'basic',
        meta: {
            title: '系统',
            icon: 'bx:cog',
            order: 3,
        },
        children: [
            {
                name: 'system_user',
                path: '/system/user',
                component: 'self',
                meta: {
                    title: '管理员',
                    requiresAuth: true,
                    icon: 'fa-solid:users-cog',
                },
            },
            {
                name: 'system_scaffold',
                path: '/system/scaffold',
                component: 'self',
                meta: {
                    title: '脚手架',
                    requiresAuth: true,
                    icon: 'ic:baseline-laptop',
                },
            },
        ],
    },
    {
        name: 'component',
        path: '/component',
        component: 'basic',
        children: [
            {
                name: 'component_button',
                path: '/component/button',
                component: 'self',
                meta: {
                    title: '按钮',
                    requiresAuth: true,
                    icon: 'ic:baseline-radio-button-checked',
                },
            },
            {
                name: 'component_card',
                path: '/component/card',
                component: 'self',
                meta: {
                    title: '卡片',
                    requiresAuth: true,
                    icon: 'mdi:card-outline',
                },
            },
            {
                name: 'component_table',
                path: '/component/table',
                component: 'self',
                meta: {
                    title: '表格',
                    requiresAuth: true,
                    icon: 'mdi:table-large',
                },
            },
        ],
        meta: {
            title: '组件示例',
            icon: 'fluent:app-store-24-regular',
            order: 3,
        },
    },
    {
        name: 'plugin',
        path: '/plugin',
        component: 'basic',
        children: [
            {
                name: 'plugin_map',
                path: '/plugin/map',
                component: 'self',
                meta: {
                    title: '地图',
                    requiresAuth: true,
                    icon: 'mdi:map',
                },
            },
            {
                name: 'plugin_video',
                path: '/plugin/video',
                component: 'self',
                meta: {
                    title: '视频',
                    requiresAuth: true,
                    icon: 'mdi:video',
                },
            },
            {
                name: 'plugin_editor',
                path: '/plugin/editor',
                component: 'multi',
                children: [
                    {
                        name: 'plugin_editor_quill',
                        path: '/plugin/editor/quill',
                        component: 'self',
                        meta: {
                            title: '富文本编辑器',
                            requiresAuth: true,
                            icon: 'mdi:file-document-edit-outline',
                        },
                    },
                    {
                        name: 'plugin_editor_markdown',
                        path: '/plugin/editor/markdown',
                        component: 'self',
                        meta: {
                            title: 'markdown编辑器',
                            requiresAuth: true,
                            icon: 'ri:markdown-line',
                        },
                    },
                ],
                meta: {
                    title: '编辑器',
                    icon: 'icon-park-outline:editor',
                },
            },
            {
                name: 'plugin_swiper',
                path: '/plugin/swiper',
                component: 'self',
                meta: {
                    title: 'Swiper插件',
                    requiresAuth: true,
                    icon: 'simple-icons:swiper',
                },
            },
            {
                name: 'plugin_copy',
                path: '/plugin/copy',
                component: 'self',
                meta: {
                    title: '剪贴板',
                    requiresAuth: true,
                    icon: 'mdi:clipboard-outline',
                },
            },
            {
                name: 'plugin_icon',
                path: '/plugin/icon',
                component: 'self',
                meta: {
                    title: '图标',
                    requiresAuth: true,
                    icon: 'ic:baseline-insert-emoticon',
                },
            },
            {
                name: 'plugin_print',
                path: '/plugin/print',
                component: 'self',
                meta: {
                    title: '打印',
                    requiresAuth: true,
                    icon: 'ic:baseline-local-printshop',
                },
            },
        ],
        meta: {
            title: '插件示例',
            icon: 'clarity:plugin-line',
            order: 4,
        },
    },
];

function dataMiddleware(data: AuthRoute.Route[]): ApiRoute.Route {
    const routeHomeName: AuthRoute.RouteKey = 'system_dashboard';

    function sortRoutes(sorts: AuthRoute.Route[]) {
        return sorts.sort((next, pre) => Number(next.meta?.order) - Number(pre.meta?.order));
    }

    return {
        routes: sortRoutes(data),
        home: routeHomeName,
    };
}

const apis: MockMethod[] = [
    {
        url: '/mock/getUserRoutes',
        method: 'post',
        response: (): Service.MockServiceResult => {
            return {
                code: 200,
                message: 'ok',
                data: dataMiddleware(routes),
            };
        },
    },
];

export default apis;
