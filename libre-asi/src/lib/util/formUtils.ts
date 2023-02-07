import emailValidator from "./emailValidator";
import emptyValidator from "./emptyValidator";

export function checkEmail(email: string): [caption: string, valid: boolean] {
    if (!emptyValidator(email)) {
        return ['Email is required', false];
    }

    if (!emailValidator(email)) {
        return ['Enter a valid email', false]
    }

    return ['', true]
}

export function checkUsername(username: string): [caption: string, valid: boolean] {

    if (!emptyValidator(username)) {
        return ['Username is required', false];
    }

    return ['', true]
}

export function checkPassword(password: string): [caption: string, valid: boolean] {

    if (!emptyValidator(password)) {
        return ['Password is required', false]
    }

    return ['', true]
}