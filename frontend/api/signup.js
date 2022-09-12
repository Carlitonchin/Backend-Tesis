import { API_URL, HEADERS } from "./options";

const SIGNUP_PATH = "api/account/signup"

export default async (body)=>{
    let response = await fetch(API_URL + SIGNUP_PATH,{
        method:"POST",
        headers:HEADERS,
        body:JSON.stringify(body)
    })

    response = await response.json()
    return response
}