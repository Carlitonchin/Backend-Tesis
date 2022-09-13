import { API_URL, HEADERS } from "./options";

const PATH = "api/account/signout"

export default async (tokens)=>{
    const headers = {
        ...HEADERS,
        'Authorization': `Bearer ${tokens["idToken"]}`, 
    }
    let response = await fetch(API_URL + PATH,{
        method:"POST",
        headers:headers,
    })
    return response
}