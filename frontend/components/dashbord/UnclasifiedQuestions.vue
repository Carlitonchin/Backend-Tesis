<script setup>
    import {API_URL, HEADERS} from '../../api/options'
    import get_questions from '../../api/get_unclasified_questions'
    import get_areas from '~~/api/get_areas';
    import refresh_tokens from '~~/utils/refresh_tokens';

    const tokens = useCookie("tokens").value
    let response_questions = await get_questions(tokens)
    response_questions = await refresh_tokens(response_questions, tokens, get_questions, null)
    if(response_questions.error)
        console.log(response_questions.error)
    
    let response_areas = await get_areas(tokens)
    response_areas = await refresh_tokens(response_areas, tokens, get_areas, null)
    if(response_areas.error)
        console.log(response_areas.error)
    
    console.log(response_areas.areas)
    const questions = ref(response_questions.questions)
    const areas = ref(response_areas.areas)
</script>

<template>
    <div class="p-6 w-full">
    <h2 class="text-secondary max-w-full">Preguntas sin clasificar</h2>
    <div class="space-y-8 mt-5">
        <div v-for="question in questions" class="flex space-x-3">
        <div class="leading-6">
           {{question.text}}
        </div>
        <ThreePointOptions/>
    </div>
    </div>
</div>
</template>