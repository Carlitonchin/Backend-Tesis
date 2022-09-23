<script setup>
    import get_taken_questions from '~~/api/get_taken_questions'
    import my_fetch from '~~/utils/my_fetch'
    import up_level_fetcher from '~~/api/up_level_2'
    import up_to_admin_fetcher from '~~/api/up_admin'
    import {LEVEL1_SPECIALIST, LEVEL2_SPECIALIST} from '~~/api/options'
import consolaGlobalInstance from 'consola';

    definePageMeta({
  middleware: ["index"]
})

const questions = ref([])
let tokens = useCookie("tokens").value
let user = useCookie("user").value
let response = await my_fetch(tokens, get_taken_questions, null)
if(response.error){
    console.log(response.error)
}
else
    questions.value = response.questions

const options = [{name:"Responder", ID:"response"}]

if(user.role.name == LEVEL1_SPECIALIST)
    options.push({name:"Subir al nivel 2", ID:"upLevel"})

if(user.role.name == LEVEL2_SPECIALIST)
    options.push({name:"Subir a la administración", ID:"upAdmin"})

function handle_click(question_id, func){
    switch(func){
        case 'response':
            response_question(question_id)
            break
        case 'upLevel':
            upLevel(question_id)
            break
        case 'upAdmin':
            upAdmin(question_id)
            break
    }
}

function response_question(question_id){
    window.location.href = `${question_id}/response`
}

async function upLevel(question_id){
    let response = await my_fetch(tokens, up_level_fetcher, {question_id})
    if(response.error){
        console.log(response.error)
        return
    }

    questions.value = questions.value.filter(q => q.ID != question_id)
}

async function upAdmin(question_id){
    let response = await my_fetch(tokens, up_to_admin_fetcher, {question_id})
    if(response.error){
        console.log(response.error)
        return
    }

    questions.value = questions.value.filter(q=>q.ID != question_id)
}

</script>

<template>
    <NuxtLayout>
        <div class="p-4">
        <h1 class="text-primary mb-4">
            Tus preguntas
        </h1>
        <div v-if="questions.length" class="space-y-4">
            <div class="flex" v-for="question in questions">
                 <p>{{question.text}}</p>
                 <ThreePointOptions :options="options" :handle_click="(func)=>handle_click(question.ID, func)"/>
            </div>
        </div>
        <p v-else>No tienes</p>
    </div>
    </NuxtLayout>
</template>