<script setup>
    import refresh_tokens from '~~/utils/refresh_tokens';
    import { STUDENT_ROLE } from '~~/api/options';
import get_users from '~~/api/get_users'
import ThreePointOptions1 from '~~/components/ThreePointOptions.vue';
import get_areas from '~~/api/get_areas';
import get_roles from '~~/api/get_roles';
import update_user_role from '~~/api/update_user_role'
import my_fetch from '~~/utils/my_fetch'

definePageMeta({
  middleware: ["index"]
})
const tokens = useCookie("tokens").value

const users = ref([])
const areas = ref([])
const roles = ref([])
const tab_worker = ref(true)

function change_tab(value){
    tab_worker.value = value
}

async function update_role(user_id, role_id){
    let body = {user_id, role_id}
    let response = await my_fetch(tokens, update_user_role, body)
    if(response.error){
        console.log(response.error)
        return
    }

    let role_name = roles.value.filter(r=>r.ID == role_id)[0].name

    users.value = users.value.map(u=>
    {
        if(u.ID != user_id)
            return u;
        
        u.role_id = role_id
        u.role.ID = role_id
        u.role.name = role_name
        return u;
    })
}

function update_area(user_id, role_id){
    console.log(user_id, role_id)
}

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
                <div class="w-full flex justify-center mt-2 mb-2">
                    <div @click="change_tab(true)" :class="['cursor-pointer', 'p-3', 'pl-5', 'pr-5', 'rounded-l-md',
                    'border-solid', 'border', 'border-blue-500', (tab_worker?'bg-blue-300 text-gray-900':''),
                    (!tab_worker?'hover:bg-gray-500':'')]">Trabajadores</div>
                    
                    <div @click="change_tab(false)" :class="['cursor-pointer', 'p-3', 'pl-5', 'pr-5', 'rounded-r-md',
                    'border-solid', 'border', 'border-blue-500', (!tab_worker?'bg-blue-300 text-gray-900':''),
                    (tab_worker?'hover:bg-gray-500':'')]">Estudiantes</div>
                </div>
                <div class="space-y-4">
                <div v-for="user in users_to_show">
                    <h2 class="text-secondary text-yellow-500">{{user.name}}</h2>
                    <a :href="'mailto:' + user.email" class="text-red-500">{{user.email}}</a>

                    <div class="flex">
                    <p><span class="text-red-500">Rol: </span>{{user.role.name}}</p>
                    <ThreePointOptions :options="roles.filter(r=>r.ID != user.role_id)" 
                        :handle_click="(role_id)=>update_role(user.ID, role_id)"/>
                </div>

                <div class="flex">
                    <p><span class="text-red-500">Área: </span> {{user.area ? user.area.name : "-"}}</p>
                    <ThreePointOptions :options="areas.filter(a=>user.area == null || user.area_id != a.ID)"
                        :handle_click="(area_id)=>update_area(user.ID, area_id)"/>
                </div>
                </div>
            </div>
            </div>
        </NuxtLayout>
    </NuxtLayout>
    

</template>