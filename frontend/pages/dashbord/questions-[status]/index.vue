<script setup>
import get_question from '~~/api/get_question_by_status'
import my_fetch from '~~/utils/my_fetch';

definePageMeta({
  middleware: ["index"]
})

const status_code = useRoute().params.status
if(!status_code)
    console.log("status not especified")


const tokens = useCookie("tokens").value

const questions = ref([])

let response = await my_fetch(tokens, get_question, status_code)
if(response.error)
    console.log(response.error)
else
    questions.value = response.questions

</script>

<template>
    <div class="p-6 w-full">
    <h2 class="text-secondary max-w-full">Preguntas clasificadas</h2>
    <div class="space-y-8 mt-5">
        <div v-if="questions.length" v-for="question in questions" class="flex space-x-3">
        <div class="leading-6">
           {{question.text}}
        </div>
        <ThreePointOptions :options="areas" :handle_click="(area_id) => clasify(question.ID, area_id)"/>
    </div>
    <p v-else>No hay</p>
    </div>
</div>
</template>