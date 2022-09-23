import { API_URL, HEADERS } from "./options";

const PATH = "api/questions/taken/"

export default async ({tokens, question_id})=>{
    const headers = {
        ...HEADERS,
        'Authorization': `Bearer ${tokens["idToken"]}`, 
    }
    let response = await fetch(API_URL + PATH + question_id,{
        method:'GET',
        headers:headers
    })
    return response
}