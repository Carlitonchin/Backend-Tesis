<script setup>
    import { Field, Form } from 'vee-validate';
    const user = useCookie("user").value
    if(!user){
        navigateTo("/")
    }
    
</script>

<template>
    <div class="flex min-h-screen w-full flex-col space-y-20 items-center p-6 pt-24">
    <h2 class="text-secondary">Hola <span class="text-red-500">{{user.name}}</span></h2>
    <div class="h-fit w-full flex flex-col items-center">
    <Form class="w-full max-w-md">
        <Field id="question" name="question" as="textarea" autocomplete="name" required
           class="relative block w-full h-52 appearance-none rounded-md border
            border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10
             focus:border-red-500 focus:outline-none focus:ring-red-500 sm:text-sm" 
             placeholder="Escribe tu duda"/>

             <button type="submit" class="group relative flex w-full mt-4 justify-center rounded-md border border-transparent
              bg-red-600 py-2 px-4 text-sm font-medium text-white hover:bg-red-700 
              focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2">
          Enviar
        </button>
    </Form>
</div>
</div>
</template>