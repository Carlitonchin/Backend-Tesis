<script setup>
import Student from "../../components/dashbord/Student.vue";
import {STUDENT_ROLE, WORKER_ROLE, ADMIN_ROLE} from '../../api/options'
import UnclasifiedQuestions from "~~/components/dashbord/UnclasifiedQuestions.vue";
import Admin from "~~/components/dashbord/Admin.vue";
    
definePageMeta({
  middleware: ["index"]
})
    const tokens = useCookie("tokens").value
    const user = useCookie("user").value
    const role = user.role.name
</script>

<template>
    <NuxtLayout>
        <Student :user="user" :tokens="tokens" v-if="role === STUDENT_ROLE"/>
        <UnclasifiedQuestions :user="user" :tokens="tokens" v-else-if="role === WORKER_ROLE"/>
        <Admin :user="user" :tokens="tokens" v-else-if="role === ADMIN_ROLE"/>
    </NuxtLayout>
</template>