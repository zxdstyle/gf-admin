<template>
    <n-tooltip trigger="hover">
        <template #trigger>
            <Icon icon="carbon:help" class="inline ml-1 text-amber-300" />
        </template>
        {{ message }}
    </n-tooltip>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { Icon } from '@iconify/vue';

export default defineComponent({
    name: 'HelpMessage',
    components: { Icon },
    props: {
        message: { type: String, default: '' },
    },
    setup() {},
});
</script>
