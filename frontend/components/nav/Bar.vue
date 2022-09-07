<script setup>
    const props = defineProps(["options"])
    const width = ref(0)

    const min_pc_width = 900;

    function handle_resize()
    {
        width.value = window.document.body.clientWidth;
    }

    const isMobile = computed(()=> width.value < min_pc_width);

    onMounted(()=>
    {
        width.value = window.document.body.clientWidth;
        window.addEventListener('resize', handle_resize);
    });

</script>

<template>
    <div id="bar">

        <div id="container" class="flex-horizontal flex-center-align flex-space-between p3 p1-top-bottom">
        <slot name="logo"/>

        <NavOptionsMobileComponent v-if="isMobile" :options="props.options"/>
        <NavOptionsPC v-else :options="props.options"/>

    </div>
    </div>
</template>

<style scoped>
    #bar
    {
        position: fixed;
        width: 100vw;
    }

#container
{
    height: 8vh;
        min-height: 40px;
    background: rgba(0, 0, 0, 0.281);
     -webkit-box-shadow: -4px -26px 30px 21px rgba(0,0,0,0.75);
  -moz-box-shadow: -4px -26px 30px 21px rgba(0,0,0,0.75);
  box-shadow: -4px -26px 30px 21px rgba(0,0,0,0.75);

}

</style>
