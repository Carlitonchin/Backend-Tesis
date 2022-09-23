<script setup>
    import { Form, Field } from 'vee-validate';
import get_areas from '~~/api/get_areas';
import create_area from '~~/api/create_area'
import refresh_tokens from '~~/utils/refresh_tokens';

definePageMeta({
  middleware: ["index"]
})
    async function handle_submit(){
        let response = await create_area(tokens, area_name.value)
        response = await refresh_tokens(response, tokens, create_area, area_name)
        if(response.error){
            console.log(response.error)
            return
        }

        areas.value.push({name:area_name})
    }

    const tokens = useCookie("tokens").value
    const areas = ref([])
    const area_name = ref("")
    let response = await get_areas(tokens)
    response = await refresh_tokens(response, tokens, get_areas, null)
    if(response.error)
        console.log(response.error)
    else{
        areas.value = response.areas
    }
</script>
<template>
        <NuxtLayout name="logged">
            <div class="pl-2 pr-2">
                <form class="space-y-1 w-full" method="POST" @submit="handle_submit">
      <input type="hidden" name="remember" value="true">
      <div class="-space-y-px rounded-md shadow-sm w-full">
        <div>
          <label for="area" class="sr-only">Nueva Área</label>
          <input id="area" name="area" type="text" autocomplete="name" required 
          class="relative block w-full appearance-none rounded-md border border-gray-300 px-3 py-2
           text-gray-900 placeholder-gray-500 focus:z-10 focus:border-red-500 focus:outline-none
            focus:ring-red-500 sm:text-sm" placeholder="Nueva Área" v-model="area_name"/>
           
        </div>
       
       
      </div>
      <div>
        <button type="submit" class="group relative flex w-full justify-center rounded-md border border-transparent bg-red-600 py-2 px-4 text-sm font-medium text-white hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2">
          Crear
        </button>
      </div>
      
    </form>

    <h2 class="text-primary mt-4">Áreas</h2>
    <p v-for="area in areas">
    {{area.name}}
    </p>
            </div>
        </NuxtLayout>
</template>