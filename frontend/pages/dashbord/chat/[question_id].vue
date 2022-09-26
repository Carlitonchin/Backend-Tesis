<script setup>
import my_fetch from '~~/utils/my_fetch';
import get_messages from '~~/api/get_messages';
import add_message from '~~/api/add_message'

definePageMeta({
  middleware: ["index"]
})

const question_id = parseInt(useRoute().params.question_id)
if(isNaN(question_id))
    navigateTo("/")

const user = useCookie("user").value
const tokens = useCookie("tokens").value

const messages = ref([])
const question = ref(null)

async function get_data(){
    let response = await my_fetch(tokens, get_messages, {question_id})

    if(response.error)
    {
        console.log(response.error)
        return {error:response.error}
    }

    return {messages:response.messages, question:response.question}
}

let response = await get_data()
if(response.error){
    console.log(response.error)
}
else{
    messages.value = response.messages
    question.value = response.question
}

const messages_ordered = computed(()=>{
    return messages.value.sort((a,b)=>{
        if(a.CreatedAt <= b.CreatedAt)
            return -1
        
        return 1
    })
})

onMounted(()=>{
    const f = async ()=>{
        let response = await get_data()
        if(response.error){
            console.log(response.error)
            return
        }

        messages.value = response.messages
        setTimeout(f, 1000)
    }

    f()
})

async function send_message(){
    let text_input = document.getElementById("text-input")
    let text = text_input.value
    let body = {question_id, text}
    let response = await my_fetch(tokens, add_message, body)

    if(response.error){
        console.log(response.error)
        return
    }

    text_input.value = ""
    messages.value.push(response.message)
}

</script>

<template>
    <NuxtLayout name="logged">
        <div class="p-4 pt-0 space-y-4">
        <h2 class="text-secondary">{{question.text}}</h2>
        
        <div>
            <ChatBubble :messages="messages_ordered"/>
        </div>

        <form class="w-full" @submit.prevent="send_message">
            <div class="w-full flex">
            <input id="text-input" class="relative block w-full appearance-none rounded-md rounded-r-none border
           border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500
            focus:z-10 focus:border-red-500 focus:outline-none
             focus:ring-red-500 sm:text-sm" 
             type="text" required placeholder="Escribe un mensaje"/>
             <button type="submit" class="group relative flex w-fit justify-center rounded-md rounded-l-none border border-transparent bg-red-600 py-2 px-4 text-sm font-medium text-white hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2">
          Enviar
        </button>
        </div>
        </form>
    </div>
    </NuxtLayout>
</template>