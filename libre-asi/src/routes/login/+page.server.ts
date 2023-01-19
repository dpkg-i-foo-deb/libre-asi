import { apiUrl, interviewerLogin } from '$lib/api/constants';
import type User from '$lib/models/User';
import { fail, redirect } from '@sveltejs/kit';
import type { Action, Actions, PageServerLoad } from './$types';

const login: Action = async function ({ cookies, request, fetch }) {
	const data = await request.formData();

	const email = data.get('email');
	const password = data.get('password');

	if (typeof email !== 'string' || typeof password !== 'string') {
		return fail(400, { badRequest: true });
	}

	const user: User = { email: email, password: password };

	console.log(user);

	try {
		const response = await fetch('http://127.0.0.1:3000/login/interviewer', {
			mode: 'cors',
			method: 'POST',
			credentials: 'include',
			body: JSON.stringify(user),
			headers: {
				'content-type': 'application/json'
			}
		});

		if (response.status == 401) {
			return fail(401, { invalidCredentials: true });
		}
	} catch (e) {
		console.log(e);
		return fail(500, { cannotConnect: true });
	}
};

export const actions: Actions = { login };
