import { dev } from "$app/environment";

let API_URL = 'https://libre-asi-api.onrender.com/'
if (dev){
API_URL = 'http://127.0.0.1:3000/'

}

const SET_UP = 'set-up';
const INTERVIEWER_LOGIN = 'login/interviewer/';
const ADMIN_LOGIN = 'login/admin';
const SIGN_OUT = 'sign-out';
const REFRESH = 'refresh/';
const GET_ADMINS = 'administrators/';
const REGISTER_ADMIN = 'administrators/';
const PASSWORD_RESET = 'password-reset/';

export {
	API_URL,
	INTERVIEWER_LOGIN,
	ADMIN_LOGIN,
	SIGN_OUT,
	REFRESH,
	GET_ADMINS,
	REGISTER_ADMIN,
	SET_UP,
	PASSWORD_RESET
};
