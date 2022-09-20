<script setup>
    const {user, tokens} = defineProps(['user', 'tokens'])
    const options = [[{text:"Áreas", href:"/dashbord/areas"}]]
</script>

<template>
    <div class="w-full flex flex-col space-y-4">
        <Options :options="options"/>
        <h1 class="text-primary">Hola {{user.name}}</h1>
    </div>
    
</template>