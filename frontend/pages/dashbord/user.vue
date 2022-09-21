<script setup>
    import refresh_tokens from '~~/utils/refresh_tokens';
    import { STUDENT_ROLE } from '~~/api/options';
import get_users from '../../api/get_users'
import ThreePointOptions1 from '~~/components/ThreePointOptions.vue';
import get_areas from '~~/api/get_areas';
import get_roles from '~~/api/get_roles';

definePageMeta({
  middleware: ["index"]
})

const users = ref([])
const areas = ref([])
const roles = ref([])
const tab_worker = ref(true)

const tokens = useCookie("tokens").value
const current_user = useCookie("user").value
console.log(current_user)
const users_to_show = computed(()=>{
    if(tab_worker.value)
        return users.value.filter(u => u.role.name != STUDENT_ROLE && u.ID != current_user.ID)

    return users.value.filter(u=>u.role.name == STUDENT_ROLE && u.ID != current_user.ID)
})

let response = await get_users(tokens)
response = await refresh_tokens(response, tokens, get_users, null)
if(response.error)
    console.log(response.error)
else
    users.value = response.users

let response_areas = await get_areas(tokens)
response_areas = await refresh_tokens(response_areas, tokens, get_areas, null)
if(response_areas.error)
    console.log(response_areas.error)
else
    areas.value = response_areas.areas

let response_roles = await get_roles(tokens)
response_roles = await refresh_tokens(response_roles, tokens, get_roles, null)
if(response_roles.error)
    console.log(response_roles.error)
else
    roles.value = response_roles.roles

</script>

<template>
    <NuxtLayout>
        <NuxtLayout name="admin">
            <div class="pl-4 pr-4 flex flex-col items-center">
                <h2 class="text-primary">Users:</h2>
                <div class="w-full flex space-x-6 justify-center">
                    <div class="cursor-pointer p-3 rounded-md border-solid border-2 border-b-gray-500">Trabajadores</div>
                    <div class="cursor-pointer p-3">Estudiantes</div>
                </div>
                <div class="space-y-4">
                <div v-for="user in users_to_show">
                    <h2 class="text-secondary text-yellow-500">{{user.name}}</h2>
                    <a :href="'mailto:' + user.email" class="text-red-500">{{user.email}}</a>

                    <div class="flex">
                    <p><span class="text-red-500">Rol: </span>{{user.role.name}}</p>
                    <ThreePointOptions :options="roles.filter(r=>r.ID != user.role_id)"/>
                </div>

                <div class="flex">
                    <p><span class="text-red-500">Área: </span> {{user.area ? user.area.name : "-"}}</p>
                    <ThreePointOptions :options="areas.filter(a=>user.area == null || user.area_id != a.ID)"/>
                </div>
                </div>
            </div>
            </div>
        </NuxtLayout>
    </NuxtLayout>
    

</template>