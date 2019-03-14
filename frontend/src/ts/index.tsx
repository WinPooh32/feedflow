export * from "./api"
export * from "./configuration"

function contains(src: string, search: string): boolean{
    src = src.toLocaleLowerCase()
    return src.includes(search)
}

function suffix(src: string, search: string): boolean{
    src = src.toLocaleLowerCase()
    return src.slice(-search.length) === search
}

async function Render(){
    const path = window.location.pathname

    switch(true){
        case suffix(path, '/signin') :{
            const signon = await import("./renders/signon")

            signon.Render()
            break
        }
    
        default: {
            const content = await import("./renders/content") 
            const menu = await import("./renders/menu") 
            
            content.Render()
            menu.Render()
        }
    }
}

Render()