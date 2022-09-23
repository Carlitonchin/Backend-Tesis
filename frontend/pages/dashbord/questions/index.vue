<script setup>
import get_question from '~~/api/get_question_by_status'
import my_fetch from '~~/utils/my_fetch';
import {ADMIN_ROLE, QuestionsStatusDict, LEVEL1_SPECIALIST, LEVEL2_SPECIALIST, 
    clasified1_code, clasified2_code, admin_code} from '~~/api/options'
import ThreePointOptions from '~~/components/ThreePointOptions.vue';
import take_question_req from '~~/api/take_question'

definePageMeta({
  middleware: ["index"]
})

let status_code = null

const tokens = useCookie("tokens").value
const user = useCookie("user").value
const questions = ref([])

switch(user.role.name){
    case LEVEL1_SPECIALIST:
        status_code = clasified1_code
        break
    case LEVEL2_SPECIALIST:
        status_code = clasified2_code
        break
    case ADMIN_ROLE:
        status_code = admin_code
        break
    default:
        window.location.href = "/"
}

let response = await my_fetch(tokens, get_question, status_code)
if(response.error)
    console.log(response.error)
else
    questions.value = response.questions

const options = [{name:'Tomar'}]

function take_question(question_id){
    let response = my_fetch(tokens, take_question_req, {question_id})
    if(response.error)
    {
        console.log(response.error)
        return
    }

    questions.value = questions.value.filter(q=>q.ID != question_id)
}

</script>

<template>
    <NuxtLayout name="logged">
    <div class="p-6 w-full" v-if="user.role.name === ADMIN_ROLE">
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

<div v-else class="p-6 w-full">
    <h2 class="text-secondary max-w-full">{{QuestionsStatusDict[status_code]}}</h2>
    <h3 class="text-xl">Área: {{user.area.name}}</h3>
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
</template>