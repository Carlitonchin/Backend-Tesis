<script setup>
    import get_questions from '../../api/student_questions'
    import refresh_tokens from '~~/utils/refresh_tokens';
    const tokens = useCookie("tokens").value
    const questions = ref([])
    onMounted(async ()=>{
        let response = await get_questions(tokens)
        response = await refresh_tokens(response, tokens, get_questions, null)
        if(response.error){
            alert(response.error)
            return
        }

        questions.value = response.questions
    })
</script>

<template>
    <div class="p-6">
        <a href="/dashbord">
            <button class="secondary-button">
                Volver
            </button>
        </a>
        <h2 class="text-2xl font-semibold mb-4 mt-4">Mis dudas:</h2>
        <div class="space-y-3">
        <div v-for="question,i in questions">
            <div class="flex space-x-4 items-center">
            <p class="leading-8"><span class="font-semibold text-red-500">{{(i+1) + ". "}}</span> {{question.text}}</p>
            <a class="flex items-center relative hover:scale-125" :href="'/dashbord/chat/' + question.ID">
                <img class="h-8 w-6" src="/chat.svg" alt="chat icon" />
                <p v-if="question.unreaded_chats > 0" 
                    class="bg-red-500 rounded-full w-6 h-6 text-center 
                    m-0 z-10 absolute -top-2 -right-4">{{question.unreaded_chats}}</p>
            </a>
        </div>
            <div class="flex items-center">
                <p v-if="question.response"> <span class="text-red-500">{{"Respuesta: "}}</span>{{question.response}}</p>
            <p v-else><span class="text-red-500 m-0">Respuesta:</span> -</p>
            </div>
        </div>
    </div>
    </div>
</template>