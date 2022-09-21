<script setup>
    const options = [[
    {text:"Áreas", href:"/dashbord/areas/"},
    {text:"Usuarios", href:"/dashbord/users/"}
    ]]
</script>

<template>
    <div class="w-full flex flex-col space-y-4">
        <Options :options="options"/>
        <slot/>
    </div>
</template>