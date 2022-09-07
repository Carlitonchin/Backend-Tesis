<script setup>
  const {title} = defineProps(['title'])
  useHead({
    title: title,
    // or, instead:
    // titleTemplate: (title) => `My App - ${title}`,
    viewport: 'width=device-width, initial-scale=1',
    htmlAttrs:{lang:'es'},
    charset: 'utf-8',
    meta: [
      { name: 'description', content: 'Sitio de dudas y sugerencias de la facultad de matemática y computación de la Universidad de La Habana' }
    ],
    link: [{ rel: 'icon', type: 'image/x-icon', href: '/favicon.png' }]
})
</script>

<template>
<div id="s-container" class="flex-vertical">
  <div class="flex-horizontal flex-center-align flex-space-between">
    <NuxtLink to="/" class="without-decoration secondary button">Volver</NuxtLink>
    <h1 id="s-title">Regístrate en <span id="s-ask">Ask</span>Matcom</h1>
    <a id="s-hide">
      <!--Para que quede el texto en el centro
      en el css le pongo el color igual al fondo
      -->
      Volver
    </a>

  </div>

  <div id="s-container-form" class="shadow p2 border">
    <slot/>
  </div>

</div>
</template>

<style scoped>

#s-container{
  width: 50%;
  height: 65vh;
  margin: auto;
  margin-top: calc(65vh / 2 / 2 / 2);
}

  #s-ask{
    color:var(--primary-color);
  }

#s-hide{
  color:var(--back-color);
}

#s-container-form{
  margin-top:1rem;
  height: 100%;
}

</style>
