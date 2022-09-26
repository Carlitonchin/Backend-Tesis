<script setup>
    const {messages} = defineProps(["messages"])
    const user = useCookie("user").value
    const color = ['b2', 'b3', 'b4']
    const bubble_color = {}
    bubble_color[user.name] = 'b1'
    for(let i = 0; i < messages.length; i++){
        let author = messages[i].user_name 

        if(color.length > 0 && !bubble_color[author])
            bubble_color[author] = color.pop()

        else if(!bubble_color[author])
            bubble_color[author] = 'bdefault'
        
    }
</script>

<template>
    <div v-for="m in messages"
    :class="[bubble_color[m.user_name], 'font-bold', 'mb-4', 'p-8', 'w-fit', 'rounded-md']">
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