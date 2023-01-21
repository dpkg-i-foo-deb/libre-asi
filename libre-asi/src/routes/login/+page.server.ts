import { apiUrl, interviewerLogin } from '$lib/api/constants';
import { parseNodeFetchCookies } from '$lib/api/parseCookie';
import type { JwtPair } from '$lib/models/JwtPair';
import type User from '$lib/models/User';
import { fail, redirect } from '@sveltejs/kit';
import type { Action, Actions, PageServerLoad } from './$types';

const login: Action = async function ({ cookies, request, fetch, }) {
	const data = await request.formData();

	const email = data.get('email');
	const password = data.get('password');
	const administrator = data.get('administrator');
	const interviewer = data.get('interviewer');

	console.log(administrator);
	console.log(interviewer);

	const fullUrl = `${apiUrl}login/${administrator ?? interviewer}`;

	console.log(fullUrl);

	if (typeof email !== 'string' || typeof password !== 'string') {
		return fail(400, { badRequest: true });
	}

	const user: User = { email: email, password: password };

	console.log(user);

	try {
		const response = await fetch(fullUrl, {
			mode: 'cors',
			method: 'POST',
			credentials: 'include',
			body: JSON.stringify(user)
		});

		if (response.status == 401) {
			return fail(401, { invalidCredentials: true });
		}

		if (response.status == 200) {

			const pair : JwtPair = await response.json()

			console.log(pair)

			return {status:200, headers:'owo'}

		}
	} catch (e) {
		console.log(e);
		return fail(500, { cannotConnect: true });
	}
};

export const actions: Actions = { login };
