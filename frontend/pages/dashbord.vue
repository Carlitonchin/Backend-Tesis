<script setup>
import Student from "../components/dashbord/Student.vue";
import {STUDENT_ROLE} from '../api/options'
const tokens = useCookie("tokens").value
    const user = useCookie("user").value
    if(!user || !tokens){
        navigateTo("/")
    }

    const role = user.role.name
</script>

<template>
    <NuxtLayout>
        <Student :user="user" :tokens="tokens" v-if="role === STUDENT_ROLE"/>
    </NuxtLayout>
</template>