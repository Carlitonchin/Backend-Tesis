<script setup>
    import Logo from '../components/Logo.vue';
    import logout_post from '../api/logout'
    import refresh_tokens from '../utils/refresh_tokens';
    import remove_user_cookie from '../utils/remove_user_cookie';
    
      const user = useCookie("user").value
    
      const tokens = useCookie("tokens").value

      async function logout_handler()
      {
        let response = await logout_post(tokens)
        response = await refresh_tokens(response, tokens, logout_post, null)
    
        if(response.error){
          alert(response.error)
          return
        }
    
        remove_user_cookie()
        window.location.href="/"
    
      }
    </script>

<template>
    <nav class="h-fit shadow-md flex flex-col space-y-2 p-4 justify-between items-center bg-gray-800">
        <div>
      <Logo/>
    </div>
    <div>
    </div>
    <div>
      <button v-if="user" @click="logout_handler" class="primary-button">Cerrar Sesion</button>
      <div v-else class="flex space-x-2">
        <a href="signin">
        <button class="secondary-button">
          Inciar sesión
        </button>
          
    </a>
        <a href="/signup">
          <button class="primary-button">
          Registrarse
        </button>
    </a>
      </div>
    </div>
    </nav>
</template>