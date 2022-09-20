import { API_URL, HEADERS } from "./options";

const PATH = "api/areas/add"

export default async (tokens,name)=>{
    const headers = {
        ...HEADERS,
        'Authorization': `Bearer ${tokens["idToken"]}`, 
    }
    let response = await fetch(API_URL + PATH,{
        method:"POST",
        headers:headers,
        body:JSON.stringify({name})
    })
    return response
}