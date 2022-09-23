<script setup>
    import {STUDENT_ROLE} from '../../api/options'
import StudentQuestions from '../../components/dashbord/StudentQuestions.vue'
definePageMeta({
  middleware: ["index"]
})

const user = useCookie("user").value
const role = user.role.name
</script>

<template>
    <NuxtLayout name="logged">
        <StudentQuestions v-if="role == STUDENT_ROLE"/>
</NuxtLayout>
</template>