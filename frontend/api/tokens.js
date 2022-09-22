import { API_URL, HEADERS } from "./options";
const PATH = "api/account/tokens"

export default async (tokens)=>{
    const rf = {
        refreshToken:tokens['refreshToken']
    }

    let response = await fetch(API_URL + PATH,{
        method:"POST",
        headers:HEADERS,
        body:JSON.stringify(rf)
    })
    response = await response.json()
    return response
}