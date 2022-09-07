<script setup>
    const props = defineProps(['options'])
</script>

<template>
    <div id="container-options" class="flex-horizontal flex-space-between">
            <NuxtLink class="clickable without-decoration s-link"
            v-for="item in props.options" :key="item.id" :to="item.link">
            <h2 :class="{'first-color':item.current}">
                {{item.text}}
              </h2> 
            </NuxtLink>
               
        </div>

        <UtilLoginRegister/>
</template>

<style scoped>
    #container-options
    {
        width: 30%;
    }

    .s-link
    {
        transition: 150ms ease-in-out;
    }
    .s-link:hover
    {
        color:var(--primary-color);
    }
</style>