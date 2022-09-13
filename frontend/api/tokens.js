import { API_URL, HEADERS } from "./options";
const PATH = "api/account/tokens"

export default async (tokens)=>{
    const body = {
        "refreshToken":tokens['refreshToken']
    }
    let response = await fetch(API_URL + SIGNUP_PATH,{
        method:"POST",
        headers:HEADERS,
        body:JSON.stringify(body)
    })

    response = await response.json()
    return response
}