import { dev } from '$app/environment';

let API_URL = 'https://libre-asi-api.onrender.com/';
if (dev) {
	API_URL = 'http://127.0.0.1:3000/';
}

const SET_UP = 'set-up';
const INTERVIEWER_LOGIN = 'login/interviewer/';
const ADMIN_LOGIN = 'login/admin';
const SIGN_OUT = 'sign-out';
const REFRESH = 'refresh/';
const GET_ADMINS = 'administrators/';
const REGISTER_ADMIN = 'administrators/';
const ADMIN_PASSWORD_RESET = 'password-reset/admin';
const INTERVIEWER_PASSWORD_RESET = 'password-reset/interviewer';
const PASSWORD_RESET = 'password-reset/';
const EDIT_ADMINS = 'administrators/';
const DELETE_ADMIN = 'administrators/';
const GET_PATIENTS = 'patients/';
const GET_PATIENT = 'patients/';
const GET_INTERVIEWERS = 'interviewers/';
const REGISTER_INTERVIEWER = 'interviewers/';
const EDIT_INTERVIEWER = 'interviewers/';
const DELETE_INTERVIEWER = 'interviewers/';
const REGISTER_PATIENT = 'patients/';
const EDIT_PATIENT = 'patients/';
const DELETE_PATIENT = 'patients/';

export {
	API_URL,
	INTERVIEWER_LOGIN,
	ADMIN_LOGIN,
	SIGN_OUT,
	REFRESH,
	GET_ADMINS,
	REGISTER_ADMIN,
	SET_UP,
	PASSWORD_RESET,
	EDIT_ADMINS,
	DELETE_ADMIN,
	ADMIN_PASSWORD_RESET,
	INTERVIEWER_PASSWORD_RESET,
	GET_PATIENTS,
	GET_INTERVIEWERS,
	REGISTER_INTERVIEWER,
	EDIT_INTERVIEWER,
	DELETE_INTERVIEWER,
	REGISTER_PATIENT,
	EDIT_PATIENT,
	DELETE_PATIENT,
	GET_PATIENT
};
