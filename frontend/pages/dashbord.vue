<script setup>
    const user = useCookie("user").value
    if(!user){
        navigateTo("/")
    }
    
</script>

<template>
    <div class="flex flex-col space-y-10 justify-center items-center">
    <h2 class="text-secondary">Hola <span class="text-red-500">{{user.name}}</span></h2>

</div>
</template>