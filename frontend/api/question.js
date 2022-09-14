import { API_URL, HEADERS } from "./options";

const PATH = "api/questions/add"

export default async (tokens,text)=>{
    const headers = {
        ...HEADERS,
        'Authorization': `Bearer ${tokens["idToken"]}`, 
    }
    let response = await fetch(API_URL + PATH,{
        method:"POST",
        headers:headers,
        body:JSON.stringify({text:text})
    })
    return response
}