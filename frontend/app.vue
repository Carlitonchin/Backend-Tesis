<script setup>
import Logo from './components/Logo.vue';
  const user = useCookie("user").value
</script>
<template>
  <section class="bg-gray-900 text-white min-h-screen min-w-full">
    <nav class="bg-gray-800 shadow-md flex p-4 justify-between items-center">
      <div>
      <Logo/>
    </div>
    <div>

    </div>
    <div>
      <button v-if="user" class="primary-button">Cerrar Sesion</button>
      <div v-else>
        <NuxtLink>Inciar sesión</NuxtLink>
        <NuxtLink>Registrarse</NuxtLink>
      </div>
    </div>
    </nav>
    <NuxtPage />
  </section>
</template>
