<template>
    <n-popover width="trigger" placement="bottom" :show-arrow="false" :show="show">
        <template #trigger>
            <n-input v-model:value="val" :clearable="true" @focus="show = true" @focusout="show = false">
                <template #suffix>
                    <Icon :icon="val && val.length > 0 ? val : 'tabler:apps'" class="text-xl" />
                </template>
            </n-input>
        </template>

        <template #header>
            图标选择器
            <!--            <n-input placeholder="输入图标名称搜索"></n-input>-->
        </template>

        <div class="grid grid-cols-8 max-h-60 overflow-y-scroll">
            <div v-for="icon in icons" :key="icon" class="p-2 hover:cursor-pointer hover:text-primary" @click="handleClick(icon)">
                <Icon :icon="icon" class="text-2xl" />
            </div>
        </div>
    </n-popover>
</template>

<script lang="ts">
import { computed, defineComponent, ref } from 'vue';
import { Icon } from '@iconify/vue';
import { icons } from './icon';

export default defineComponent({
    name: 'IconPicker',
    components: { Icon },
    props: {
        value: { type: String, default: '' },
    },
    emits: ['update:value'],
    setup(props, { emit }) {
        const val = computed({
            get: () => props.value,
            set: (val) => emit('update:value', val),
        });

        const show = ref(false);
        const handleClick = (val: string) => {
            emit('update:value', val);
            show.value = false;
        };

        return { handleClick, icons, val, show };
    },
});
</script>

<style lang="less" scoped></style>
