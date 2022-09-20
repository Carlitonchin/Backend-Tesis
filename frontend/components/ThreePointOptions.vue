<script setup>
    const {options, handle_click} = defineProps(['options', 'handle_click'])
    const focus = ref(false)
    
    const take_focus = ()=>{
        focus.value = !focus.value
    }

    function exit_focus(){
        setTimeout(()=>{
            focus.value = false
        }, 200)
    }
</script>

<template>
    <div>
    <button class="flex space-x-1 h-fit w-fit p-2 cursor-pointer red-animation" @click="take_focus" @blur="exit_focus">
        <div class="w-2 h-2 bg-gray-50 rounded-full child"></div>
        <div class="w-2 h-2 bg-gray-50 rounded-full child"></div>
        <div class="w-2 h-2 bg-gray-50 rounded-full child"></div>
    </button>
    <div :class="['absolute', 'text-white','z-10','w-fit','origin-top-right','divide-y','divide-gray-100',
'bg-gray-800', 'shadow-sm','ring-1','ring-black','ring-opacity-5','focus:outline-none',
!focus ? 'hidden' : 'block']" role="menu" aria-orientation="vertical" aria-labelledby="menu-button" tabindex="-1">
    <div >
      <!-- Active: "bg-gray-100 text-gray-900", Not Active: "text-gray-700" -->
      <a v-for="opt in options" @click="handle_click(opt.ID)" class="block cursor-pointer px-4 py-2 text-sm hover:bg-gray-500">{{opt.name}}</a>
    </div>
  </div>
</div>
</template>

<style scoped>
.red-animation:hover>div{
    background: rgb(239, 68, 68);
}
</style>