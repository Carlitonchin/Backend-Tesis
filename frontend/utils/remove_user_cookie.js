import { userStore } from "~~/store"

export default ()=>{
    useCookie("tokens").value = null
    useCookie("user").value = null
    userStore().setUser(null)
}