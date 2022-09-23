<script setup>
    import get_question from '~~/api/get_taken_question_by_id'
    import my_fetch from '~~/utils/my_fetch';
    import response_question_fetcher from '~~/api/response_question'

    definePageMeta({
  middleware: ["index"]
})

    let question_id = useRoute().params.id
    if(!question_id){
        window.location.href = "/"
    }
    question_id = parseInt(question_id)
    if(isNaN(question_id))
        window.location.href = "/"

    const tokens = useCookie("tokens").value
    let response = await my_fetch(tokens, get_question, {question_id})
    const question = ref("")
    const error = ref("")
    const response_question = ref("")
    if(response.error)
        error.value = response.error
    else
        question.value = response.question.text

    async function handle_submit(){
        let body ={
            question_id,
            response: response_question.value
        }

        let response = await my_fetch(tokens, response_question_fetcher, body)
        if(response.error){
            console.log(response.error)
            return
        }

        window.location.href="/dashbord/questions/taken"
    }

</script>

<template>
    <NuxtLayout>
        <div class="p-4">
            <div v-if="error">{{error}}</div>
            <div v-else class="flex flex-col items-center">
            <h2 class="text-secondary text-bg-500 mb-1">Pregunta:</h2>
            <p>{{question}}</p>
            <form class="w-full max-w-7xl mt-8" method="POST" @submit.prevent="handle_submit">
        <textarea v-model="response_question" id="question" name="text" autocomplete="name" required
           class="relative block w-full h-80 appearance-none rounded-md border
            border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10
             focus:border-red-500 focus:outline-none focus:ring-red-500 sm:text-sm" 
             placeholder="Responde la pregunta"/>

             <button type="submit" class="group relative flex w-full mt-4 justify-center rounded-md border border-transparent
              bg-red-600 py-2 px-4 text-sm font-medium text-white hover:bg-red-700 
              focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2">
          Responder
        </button>
        <div v-if="error" class="flex flex-col justify-center items-center mt-4 text-yellow-200">
            <li class="flex">
                <ul>
                    {{error}}
                </ul>
            </li>
        </div>
    </form>
        </div>
        </div>
    </NuxtLayout>
</template>