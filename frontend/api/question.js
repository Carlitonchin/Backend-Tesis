import { API_URL, HEADERS } from "./options";

const PATH = "api/questions/add"

export default async (tokens,body)=>{
    const headers = {
        ...HEADERS,
        'Authorization': `Bearer ${tokens["idToken"]}`, 
    }
    console.log(headers)
    let response = await fetch(API_URL + PATH,{
        method:"POST",
        headers:headers,
        body:JSON.stringify(body)
    })
    response = await response.json()
    return response
}