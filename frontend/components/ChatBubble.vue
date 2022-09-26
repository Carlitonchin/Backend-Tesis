<script setup>
    const {messages} = defineProps(["messages"])
    const user = useCookie("user").value
</script>

<template>
    <div v-for="m in messages"
    :class="[m.color, 'font-bold', 'mb-4', 'p-8', 'w-fit', 'rounded-md']">
    <p class="text-2xl text-red-500 mb-2">{{m.user_name == user.name ? "tú" : m.user_name}}</p>
    <p class="text-gray-700">{{m.text}}</p>
    </div>
</template>

<style scoped>
    .b1{
        background: #b0f2c2;
    }

    .b2{
        background: #b0c2f2;
    }

    .b3{
        background: #fabff2;
    }

    .b4{
        background: #fdf9f2;
    }

    .bdefault{
        background: #ffdaf2;
    }
</style>