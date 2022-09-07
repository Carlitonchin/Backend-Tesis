
<script setup>
    definePageMeta({
    layout: false,
    });

    const username = ref("")
    const email = ref("")
    const pass = ref("")
    const pass2 = ref("")
    const worker = ref(false)
    const inputs = [
      {
        label:"Nombre de Usuario",
        type:"text",
        placeholder:"Ejemplo: Luis40",
        id:"username",
        model: username
        },

     {
      label:"Correo",
      type:"email",
      placeholder:"ejemplo@gmail.com",
      id:"email",
      model:email
      },

     {
      label:"Contraseña",
      type:"password",
      placeholder:"-------------",
      id:"pass",
      model:pass
      },

     {
      label:"Repite la Contraseña",
      type:"password",
      placeholder:"-------------",
      id:"pass2",
      model:pass2
      }
      ]
</script>

<template>
<div>
<NuxtLayout name="login-signup"  title="Ask Matcom | Registro">
  <UtilForm :inputs="inputs" :workerModel="worker"/>
</NuxtLayout>
</div>
</template>
