import {defineStore} from 'pinia'

export const userStore = defineStore('userStore',()=>{
    const user = ref(null)
    function setUser(value){
        user.value = value
    }

    return {user, setUser}
})