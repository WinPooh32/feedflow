export * from "./api";
export * from "./configuration";

import * as React from "react"
import * as ReactDOM from "react-dom"
import { FeedApi, PagesApi, NewPageContent } from "./api";

const feedApi = new FeedApi()
const pagesApi = new PagesApi()


async function getChunk(){
    let chunkResp = await feedApi.requestChunk(0, 0)
    
    if(chunkResp.data.status == 200){
        console.log("requestChunk OK")
    }
}

let content = {
    title: "Hello",
    content: "Wowo its new content", 
    tags: ["tag1"],
} as NewPageContent

pagesApi.add(content).then((resp)=>{
    resp
})

console.log("Hello, World!")
getChunk()