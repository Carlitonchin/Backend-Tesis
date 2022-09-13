import { API_URL, HEADERS } from "./options";

const PATH = "api/questions/add"

export default async (body)=>{
    let response = await fetch(API_URL + PATH,{
        method:"POST",
        headers:HEADERS,
        body:JSON.stringify(body)
    })

    response = await response.json()
    return response
}