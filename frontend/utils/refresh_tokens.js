import {UNAUTH_STATUS} from '../api/options'
import post_token from '../api/tokens'
import remove_user_cookie from './remove_user_cookie'

export default async (response, tokens, fetcher, value)=>{
    if(response.status == UNAUTH_STATUS){
        response = await post_token(tokens)
        if(response.error){
            remove_user_cookie()
            window.location.href("/")
            return
        }

        useCookie("tokens").value = response.tokens
        response = await fetcher(response.tokens, value)
        return await response.json()
    }

    return await response.json()
}