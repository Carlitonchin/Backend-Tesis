<script setup>
    const {user, tokens} = defineProps(['user', 'tokens'])
</script>

<template>
        <h1>Hola {{user.name}}</h1>
</template>