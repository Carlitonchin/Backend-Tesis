<script setup>
    const {user, tokens} = defineProps(['user', 'tokens'])
</script>

<template>
    <NuxtLayout name="admin">
        <h1>Hola {{user.name}}</h1>
    </NuxtLayout>

    
</template>