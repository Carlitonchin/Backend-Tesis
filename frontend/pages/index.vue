<script setup>
        definePageMeta({
  middleware: ["auth"]
  // or middleware: 'auth'
})
</script>
<template>
    <NuxtLayout>
    <div>
    <h1>Bienvenido</h1>
    </div>
</NuxtLayout>
</template>