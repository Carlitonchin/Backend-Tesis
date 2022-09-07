<script setup>
  import nav_state from '~~/state/nav-state';
  const {page, title} = defineProps(['page', 'title'])
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

      <NavBar :options="nav_state(page)">
      <template #logo>
        <Logo/>
      </template>
  </NavBar>
  <slot/>
  
</template>