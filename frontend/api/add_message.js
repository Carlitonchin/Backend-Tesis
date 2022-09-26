import { API_URL, HEADERS } from "./options";

const PATH = "api/chat"

export default async (tokens,body)=>{
    const headers = {
        ...HEADERS,
        'Authorization': `Bearer ${tokens["idToken"]}`, 
    }
    let response = await fetch(API_URL + PATH,{
        method:"POST",
        headers:headers,
        body:JSON.stringify(body)
    })
    return response
}