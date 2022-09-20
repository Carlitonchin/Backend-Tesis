<script setup>
    import {API_URL, HEADERS} from '../../api/options'
    import get_data from '../../api/get_unclasified_questions'
    import refresh_tokens from '~~/utils/refresh_tokens';

    const tokens = useCookie("tokens").value
    let response = await get_data(tokens)
    response = await refresh_tokens(response, tokens, get_data, null)
    if(response.error)
        alert(response.error)
    
    const questions = ref(response.questions)
</script>

<template>
    <div class="p-6 w-full">
    <h2 class="text-secondary max-w-full">Preguntas sin clasificar</h2>
    <div class="space-y-4 mt-5">
        <div v-for="question in questions" class="flex space-x-3">
        <div class="leading-6">
           {{question.text}}
        </div>
        <ThreePointOptions/>
    </div>
    </div>
</div>
</template>