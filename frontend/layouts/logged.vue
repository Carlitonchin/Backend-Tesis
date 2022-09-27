<script setup>
    import {ADMIN_ROLE, LEVEL1_SPECIALIST, 
    LEVEL2_SPECIALIST, STUDENT_ROLE, WORKER_ROLE} from '~~/api/options'

    const user_role = useCookie("user").value.role.name
    let options = []
    switch(user_role){
        case STUDENT_ROLE:
            options = [[
            {text:"Mis dudas", href:"/dashbord/my-questions"},
            {text:"Escribir duda", href:"/dashbord"}
            ]]
            break
        case ADMIN_ROLE:
            options = [[
                {text:"Inicio", href:"/dashbord"},
                {text:"Usuarios", href:"/dashbord/users"},
                {text:"Areas", href:"/dashbord/areas"},
                {text:"Preguntas Clasificadas", href:"/dashbord/questions"},
                {text:"Mis Preguntas", href:"/dashbord/questions/taken"}
            ]]
            break
        default:
        options = [[
            {text:"Inicio", href:"/dashbord"},
                {text:"Preguntas Clasificadas", href:"/dashbord/questions"},
                {text:"Mis Preguntas", href:"/dashbord/questions/taken"}
            ]]
            break
    }
</script>

<template>
    <NuxtLayout>
    <div class="w-full flex flex-col space-y-4">
        <Options v-if="user_role != WORKER_ROLE" :options="options"/>
        <slot/>
    </div>
</NuxtLayout>
</template>