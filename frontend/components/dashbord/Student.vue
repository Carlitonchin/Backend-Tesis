<script setup>
    import { Field, Form } from 'vee-validate';
    import post_question from '../../api/question'
    import post_token from '../../api/tokens'
    import refresh_tokens from '../../utils/refresh_tokens'

    const {user, tokens} = defineProps(['user', 'tokens'])
    const fields = ['text']
    const question = ref("")
    const error = ref("")
    async function handle_submit(){
        if(!question.value){
            error.value = "Escriba una duda"
            return
        }
        let response = await post_question(tokens, question.value)
        response = await refresh_tokens(response, tokens, post_question, question.value)

        if(response.error){
            console.log(response.error)
            return
        }

        question.value = ""
    }

    function change_handler(){
        error.value = ""
    }

    const options = [[
        {href:"/dashbord/my-questions", text:"Mis dudas"}
    ]]
</script>

<template>
    <div class="w-full flex flex-col space-y-4">
    <Options :options="options"/>
    <div class="flex min-h-screen w-full flex-col space-y-10 items-center pr-6 pl-6">
    <h2 class="text-secondary">Hola <span class="text-red-500">{{user.name}}</span></h2>
    <div class="h-fit w-full flex flex-col items-center">
    <Form class="w-full max-w-7xl" method="POST" @submit="handle_submit">
        <Field v-model="question" @input="change_handler" id="question" name="text" as="textarea" autocomplete="name" required
           class="relative block w-full h-80 appearance-none rounded-md border
            border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10
             focus:border-red-500 focus:outline-none focus:ring-red-500 sm:text-sm" 
             placeholder="Escribe tu duda"/>

             <button type="submit" class="group relative flex w-full mt-4 justify-center rounded-md border border-transparent
              bg-red-600 py-2 px-4 text-sm font-medium text-white hover:bg-red-700 
              focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2">
          Enviar
        </button>
        <div v-if="error" class="flex flex-col justify-center items-center mt-4 text-yellow-200">
            <li class="flex">
                <ul>
                    {{error}}
                </ul>
            </li>
        </div>
    </Form>
</div>
</div>
</div>
</template>