<script setup>
const status_code = useRoute().params.id
if(!status_code)
    console.log("status not especified")
</script>

<template>
    <h1>Holita {{route.params.status}}</h1>
</template>