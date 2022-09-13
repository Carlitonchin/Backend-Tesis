<script setup>
import Student from "../components/dashbord/Student.vue";
import {STUDENT_ROLE} from '../api/options'
    const user = useCookie("user").value
    if(!user){
        navigateTo("/")
    }

    const role = user.role.name
</script>

<template>
        <Student :user="user" :v-if="role === STUDENT_ROLE"/>
</template>