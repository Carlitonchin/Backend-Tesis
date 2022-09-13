export default async (response, tokens)=>{
    if(response.status == UNAUTH_STATUS){
        response = await post_token(tokens)
        if(response.error){
            alert(response.error)
            return
        }

        useCookie("tokens").value = response.tokens
        response = await post_question(response.tokens, value)
    }
}