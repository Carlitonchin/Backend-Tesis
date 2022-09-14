<script setup>
import Student from "../../components/dashbord/Student.vue";
import {STUDENT_ROLE} from '../../api/options'
    
definePageMeta({
  middleware: ["index"]
})

    const user = useCookie("user").value
    const role = user.role.name
</script>

<template>
    <NuxtLayout>
        <Student :user="user" :tokens="tokens" v-if="role === STUDENT_ROLE"/>
    </NuxtLayout>
</template>