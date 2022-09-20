<script setup>
    import {API_URL, HEADERS} from '../../api/options'
    import get_questions from '../../api/get_unclasified_questions'
    import get_areas from '~~/api/get_areas';
    import refresh_tokens from '~~/utils/refresh_tokens';
    import clasify_question from '../../api/clasify_question'

    const {user, tokens} = defineProps(["user", "tokens"])
    const questions = ref(null)
    const areas = ref(null)
    
    function clasify(question_id, area_id){
        let response = clasify_question(tokens, {question_id, area_id})
        response = refresh_tokens(response, tokens, clasify_question, {question_id, area_id})
        if(response.error){
            console.log(response.error)
            return
        }
        
        questions.value = questions.value.filter(q => q.ID != question_id)
    }

   onMounted(async ()=>{
    let response_questions = await get_questions(tokens)
    response_questions = await refresh_tokens(response_questions, tokens, get_questions, null)
    if(response_questions.error){
        console.log(response_questions.error)
        return
    }

    let response_areas = await get_areas(tokens)
    response_areas = await refresh_tokens(response_areas, tokens, get_areas, null)
    if(response_areas.error){
        console.log(response_areas.error)
        return
    }
    
    questions.value = response_questions.questions
    areas.value = response_areas.areas
   })
    
</script>

<template>
    <div class="p-6 w-full">
    <h2 class="text-secondary max-w-full">Preguntas sin clasificar</h2>
    <div class="space-y-8 mt-5" v-if="areas && questions">
        <div v-if="questions.length" v-for="question in questions" class="flex space-x-3">
        <div class="leading-6">
           {{question.text}}
        </div>
        <ThreePointOptions :options="areas" :handle_click="(area_id) => clasify(question.ID, area_id)"/>
    </div>
    <p v-else>No hay</p>
    </div>
    <div v-else>
        Cargando ...
    </div>
</div>
</template>