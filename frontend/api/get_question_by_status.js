import { API_URL, HEADERS } from "./options";

const PATH = "api/questions/"

export default async (tokens, status)=>{
    const headers = {
        ...HEADERS,
        'Authorization': `Bearer ${tokens["idToken"]}`, 
    }
    let response = await fetch(API_URL + PATH + status,{
        method:"GET",
        headers:headers
    })
    return response
}