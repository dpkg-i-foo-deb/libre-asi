import { goto } from "$app/navigation"
import setup from "$lib/stores/setupStore";

export function handleResponse(code: number) : boolean {

   let  shouldNavigate = true

    switch (code) {
        //TODO not enough privileges page
        case 403: goto('/')
            break;

        case 503: goto('/cannot-connect')
            setup.set(false)
            break;

        case 412: goto('/set-up')
            break;

        default: shouldNavigate = false
    }

    return shouldNavigate
}