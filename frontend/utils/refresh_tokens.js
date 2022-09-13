import {UNAUTH_STATUS} from '../api/options'
import post_token from '../api/tokens'

export default async (response, tokens, fetcher, value)=>{
    if(response.status == UNAUTH_STATUS){
        response = await post_token(tokens)
        response = await response.json()
        if(response.error){
            alert(response.error)
            return
        }

        useCookie("tokens").value = response.tokens
        response = await fetcher(response.tokens, value)
        return await response.json()
    }

    return response
}