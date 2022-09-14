<template>
    <div class="flex space-x-1 h-fit w-fit p-2 cursor-pointer red-animation">
        <div class="w-2 h-2 bg-gray-50 rounded-full child"></div>
        <div class="w-2 h-2 bg-gray-50 rounded-full child"></div>
        <div class="w-2 h-2 bg-gray-50 rounded-full child"></div>
    </div>
</template>

<style scoped>
.red-animation:hover>div{
    background: rgb(239, 68, 68);
}
</style>