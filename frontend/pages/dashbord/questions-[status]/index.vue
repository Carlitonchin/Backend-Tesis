<script setup>
import get_question from '~~/api/get_question_by_status'
import my_fetch from '~~/utils/my_fetch';
import {ADMIN_ROLE, QuestionsStatusDict} from '~~/api/options'
import ThreePointOptions from '~~/components/ThreePointOptions.vue';

definePageMeta({
  middleware: ["index"]
})

const status_code = useRoute().params.status
if(!status_code)
    console.log("status not especified")

if(!QuestionsStatusDict[status_code]){
    console.log(`status code ${status_code} no reconocido`)
    window.location.href="/"
}

const tokens = useCookie("tokens").value
const user = useCookie("user").value
const questions = ref([])

let response = await my_fetch(tokens, get_question, status_code)
if(response.error)
    console.log(response.error)
else
    questions.value = response.questions

const options = [{name:'Tomar'}]

function take_question(question_id){
    console.log(user.ID, question_id)
}

</script>

<template>
    <NuxtLayout>
        <NuxtLayout name="admin" v-if="user.role.name === ADMIN_ROLE">
    <div class="p-6 w-full">
    <h2 class="text-secondary max-w-full">{{QuestionsStatusDict[status_code]}}</h2>
    <h3 v-if="user.role.name !== ADMIN_ROLE" class="text-xl">Área: {{user.area.name}}</h3>
    <div class="space-y-8 mt-5">
        <div v-if="questions.length" v-for="question in questions" class="flex space-x-3">
        <div class="leading-6">
           {{question.text}}
        </div>
        <ThreePointOptions :options="options" :handle_click="take_question.bind(this, question.ID)"/>
    </div>
    <p v-else>No hay</p>
    </div>
</div>
</NuxtLayout>

<div v-else class="p-6 w-full">
    <h2 class="text-secondary max-w-full">{{QuestionsStatusDict[status_code]}}</h2>
    <h3 class="text-xl">Área: {{user.area.name}}</h3>
    <div class="space-y-8 mt-5">
        <div v-if="questions.length" v-for="question in questions" class="flex space-x-3">
        <div class="leading-6">
           {{question.text}}
        </div>
        <ThreePointOptions :options="options"/>
    </div>
    <p v-else>No hay</p>
    </div>
</div>

</NuxtLayout>
</template>