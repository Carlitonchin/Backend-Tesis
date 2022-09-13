import { API_URL, HEADERS } from "./options";

const PATH = "api/account/signin"

export default async (body)=>{
    let response = await fetch(API_URL + PATH,{
        method:"POST",
        headers:HEADERS,
        body:JSON.stringify(body)
    })

    response = await response.json()
    return response
}