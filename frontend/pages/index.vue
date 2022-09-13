<script setup>
    import jwtDecode from 'jwt-decode'
    const tokens_cookie = useCookie("tokens").value
    let user = null
    if(tokens_cookie){
        user = jwtDecode(tokens_cookie["idToken"]).user
        useCookie("user").value = user
        navigateTo("/dashbord")
    }
    
</script>

<template>
    <div v-if="!user">
    <h1>Bienvenidos</h1>
    </div>
</template>