<script setup>
import my_fetch from '~~/utils/my_fetch';
import get_messages from '~~/api/get_messages';

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

</script>

<template>
    <NuxtLayout name="logged">
        <h2 class="text-secondary">{{question.text}}</h2>
        <div v-for="m in messages_ordered">
            <p>{{m.user_name}}</p>
            <p>{{m.text}}</p>
        </div>
    </NuxtLayout>
</template>