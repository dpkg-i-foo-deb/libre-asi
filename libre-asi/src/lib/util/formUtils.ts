import emailValidator from "./emailValidator";
import emptyValidator from "./emptyValidator";

function checkEmail(email: string): [caption: string, valid: boolean] {
    if (!emptyValidator(email)) {
        return ['Email is required', false];
    }

    if (!emailValidator(email)) {
        return ['Enter a valid email', false]
    }

    return ['', true]
}

function checkUsername(username: string): [caption: string, valid: boolean] {

    if (!emptyValidator(username)) {
        return ['Username is required', false];
    }

    return ['', true]
}

function checkPassword(password: string): [caption: string, valid: boolean] {

    if (!emailValidator(password)) {
        return ['Password is required', false]
    }

    return ['', true]
}

export default [checkEmail, checkUsername, checkPassword]