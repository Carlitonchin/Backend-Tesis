<script setup>
    import { Form, Field, ErrorMessage } from 'vee-validate';
    import {validate_email, validate_username, validate_pass} from '../utils/form_validations'

    const fields = ["username", "email", "password"]
</script>

<template>
    <div class="flex min-h-full items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
  <div class="w-full max-w-md space-y-8">
    <div class="flex flex-col space-y-5 justify-center items-center">
        <Logo/>
        <h2 class="text-secondary">Regístrate</h2>
    </div>
    <Form class="mt-8 space-y-6 w-full" action="#" method="POST">
      <input type="hidden" name="remember" value="true">
      <div class="-space-y-px rounded-md shadow-sm w-full">
        <div>
          <label for="username" class="sr-only">Nombre de usuario</label>
          <Field id="username" name="username" type="text" autocomplete="name" required
           class="relative block w-full appearance-none rounded-none rounded-t-md border
            border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10
             focus:border-red-500 focus:outline-none focus:ring-red-500 sm:text-sm" 
             placeholder="Nombre de usuario" :rules="validate_username"/>
        </div>
        <div>
          <label for="email" class="sr-only">Correo eléctronico</label>
          <Field id="email" name="email" type="email" autocomplete="email" required 
          class="relative block w-full appearance-none rounded-none border border-gray-300 px-3 py-2
           text-gray-900 placeholder-gray-500 focus:z-10 focus:border-red-500 focus:outline-none
            focus:ring-red-500 sm:text-sm" placeholder="Correo" :rules="validate_email"/>
           
        </div>
        <div>
          <label for="password" class="sr-only">Contraseña</label>
          <Field id="password" name="password" type="password" autocomplete="current-password" required 
          class="relative block w-full appearance-none rounded-none border
           border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500
            focus:z-10 focus:border-red-500 focus:outline-none
             focus:ring-red-500 sm:text-sm" placeholder="Contraseña" :rules="validate_pass"/>
        </div>
        <div>
          <label for="passwordrepeat" class="sr-only">Repite la contraseña</label>
          <Field id="passwordrepeat" name="passwordrepeat" type="password" autocomplete="current-password" required 
          class="relative block w-full appearance-none rounded-none 
          rounded-b-md border border-gray-300 px-3 py-2 text-gray-900
           placeholder-gray-500 focus:z-10 focus:border-red-500 focus:outline-none
            focus:ring-red-500 sm:text-sm" placeholder="Repite la contraseña"/>
        </div>
      </div>

      <div class="flex items-center justify-between">
        <div class="flex items-center">
          <Field id="worker" name="worker" type="checkbox" class="h-4 w-4 rounded border-gray-300 text-red-600 focus:ring-red-500"/>
          <label for="worker" class="ml-2 block text-sm">Soy trabajador de Matcom</label>
        </div>

      </div>

      <div>
        <button type="submit" class="group relative flex w-full justify-center rounded-md border border-transparent bg-red-600 py-2 px-4 text-sm font-medium text-white hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2">
          <span class="absolute inset-y-0 left-0 flex items-center pl-3">
            <!-- Heroicon name: mini/lock-closed -->
            <svg class="h-5 w-5 text-red-500 group-hover:text-red-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
              <path fill-rule="evenodd" d="M10 1a4.5 4.5 0 00-4.5 4.5V9H5a2 2 0 00-2 2v6a2 2 0 002 2h10a2 2 0 002-2v-6a2 2 0 00-2-2h-.5V5.5A4.5 4.5 0 0010 1zm3 8V5.5a3 3 0 10-6 0V9h6z" clip-rule="evenodd" />
            </svg>
          </span>
          Registrarse
        </button>
      </div>
      <div class="text-sm flex justify-center">
          <a href="#" class="font-medium text-red-600 hover:text-red-500">¿Ya te registraste? Inicia sesión aquí</a>
        </div>

        <div class="flex flex-col justify-center items-center text-yellow-200">
            <ValidationError v-for="field in fields" :field="field" />
        </div>
    </Form>
  </div>
</div>
</template>