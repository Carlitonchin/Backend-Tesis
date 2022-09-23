<script setup>
    import {ADMIN_ROLE, LEVEL1_SPECIALIST, 
    LEVEL2_SPECIALIST, STUDENT_ROLE, WORKER_ROLE} from '~~/api/options'

    const user_role = useCookie("user").role.name
    let options = []
    switch(user_role){
        case STUDENT_ROLE:
            options = [[
            {text:"Mis dudas", href:"/dashbord/my-questions"}
            ]]
            break
        case ADMIN_ROLE:
            options = [[
                {text:"Usuarios", href:"/dashbord/users"},
                {text:"Areas", href:"/dashbord/areas"},
                {text:"Preguntas Clasificadas", href:"/dashbord/questions"},
                {text:"Mis Preguntas", href:"/dashbord/questions/taken"}
            ]]
            break
        default:
        options = [[
                {text:"Preguntas Clasificadas", href:"/dashbord/questions"},
                {text:"Mis Preguntas", href:"/dashbord/questions/taken"}
            ]]
            break
    }
</script>

<template>
    <NuxtLayout>
    <div v-if="user_role != WORKER_ROLE" class="w-full flex flex-col space-y-4">
        <Options :options="options"/>
        <slot/>
    </div>
</NuxtLayout>
</template>