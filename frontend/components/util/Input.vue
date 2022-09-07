<script setup>
    const {type, placeholder, id, model, fontSize} = defineProps(["type", "placeholder", "id", "model", "fontSize"])
</script>

<template>
    <div id="s-container-input" class="border flex-horizontal flex-center">
        <input
        :type="type"
        :placeholder="placeholder"
        :id="id"
        :style="{fontSize:`${fontSize}`}"
        @input="e=>model.value = e.target.value"
        class="s-real-input"
        />
    </div>
</template>

<style scoped>
    #s-container-input{
        background:white;
        padding: 0.15rem;
    }

    .s-real-input{
        width: 100%;
        padding-left: 0.5rem;
        outline: none !important;
        border: none !important;
    }

    #s-container-input:focus-within{
        -webkit-box-shadow: 0px 0px 13px 1px var(--primary-color);
  -moz-box-shadow: 0px 0px 13px 1px var(--primary-color);
  box-shadow: 0px 0px 13px 1px var(--primary-color);
    }
</style>
