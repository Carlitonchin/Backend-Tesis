import {UNAUTH_STATUS} from '../api/options'
import post_token from '../api/tokens'
import remove_user_cookie from './remove_user_cookie'

export default async (response, tokens, fetcher, value)=>{
    console.log("entre")
    if(response.status == UNAUTH_STATUS){
        console.log("no autorizado")
        response = await post_token(tokens)
        console.log("post token ok")
        if(response.error){
            console.log("errorsito de post token")
            remove_user_cookie()
            //window.location.href = "/"
            return
        }

        console.log("intentando refrescar cookie")
        useCookie("tokens").value = response.tokens
        console.log("tuto bene")
        response = await fetcher(response.tokens, value)
        console.log("fetch por 2da vez ok")
        return await response.json()
    }

    return await response.json()
}