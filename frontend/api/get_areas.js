import { API_URL, HEADERS } from "./options";

const PATH = "api/areas"

export default async (tokens)=>{
    const headers = {
        ...HEADERS,
        'Authorization': `Bearer ${tokens["idToken"]}`, 
    }
    let response = await fetch(API_URL + PATH,{
        method:"GET",
        headers:headers
    })
    return response
}