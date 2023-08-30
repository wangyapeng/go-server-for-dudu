<template>
    <Section>
        <router-view></router-view>
    </Section>
</template>
<script setup>
import { onMounted } from 'vue';
import Section from './struct/Section.vue'
onMounted(() => {
})
</script>
<style>
@import './style.css';

@import './theme.css'

</style>