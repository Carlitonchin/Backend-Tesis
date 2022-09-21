import refresh_tokens from "./refresh_tokens"

export default async (tokens, fetcher, body)=>{
    let response = await fetcher(tokens, body)
    response = refresh_tokens(response, tokens, fetcher, body)
    return response
}