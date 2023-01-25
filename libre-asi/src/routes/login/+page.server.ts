import { API_URL } from '$lib/api/constants';
import type { JwtPair } from '$lib/models/JwtPair';
import type User from '$lib/models/User';
import { fail, redirect } from '@sveltejs/kit';
import type { Action, Actions } from './$types';

const login: Action = async function ({ cookies, request, fetch }) {
	const data = await request.formData();

	let response: Response;

	const email = data.get('email');
	const password = data.get('password');
	const administrator = data.get('administrator');
	const interviewer = data.get('interviewer');

	const fullUrl = `${API_URL}login/${administrator ?? interviewer}`;

	if (typeof email !== 'string' || typeof password !== 'string') {
		return fail(400, { badRequest: true });
	}

	const user: User = { email: email, password: password };

	try {
		response = await fetch(fullUrl, {
			mode: 'cors',
			method: 'POST',
			credentials: 'include',
			body: JSON.stringify(user)
		});

		if (response.status == 401) {
			return fail(401, { invalidCredentials: true });
		}
	} catch (e) {
		console.log(e);
		return fail(500, { cannotConnect: true });
	}

	if (response.status == 200) {
		const pair: JwtPair = await response.json();

		cookies.set(pair.access_token.name, pair.access_token.value, {
			httpOnly: pair.access_token.http_only,
			expires: new Date(pair.access_token.expires),
			path: pair.access_token.path,
			secure: pair.access_token.secure
		});

		cookies.set(pair.refresh_token.name, pair.refresh_token.value, {
			httpOnly: pair.refresh_token.http_only,
			expires: new Date(pair.refresh_token.expires),
			path: pair.refresh_token.path,
			secure: pair.refresh_token.secure
		});

		throw redirect(302, 'home');
	}
};

export const actions: Actions = { login };
