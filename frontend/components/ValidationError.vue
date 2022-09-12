<script setup>
    import { ErrorMessage } from 'vee-validate';
    const {field} = defineProps(["field"])
</script>

<template>
    <ErrorMessage as="p" v-slot="{message}" :name="field">
           {{". " + message}}
            </ErrorMessage>
</template>