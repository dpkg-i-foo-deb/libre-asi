const pingUrl = import.meta.env.VITE_API_URL+"/"

async function ping():Promise<Boolean>{

    try{
        let data = await fetch(pingUrl,{
            method:'GET',
            mode:'cors'
        })
        
        if(!data.ok){
            console.error(data.statusText+data.statusText)
            return false
        }

        return true

    } catch(e){
        console.error(e.message)
        return false;  
    }
}

export{ping}