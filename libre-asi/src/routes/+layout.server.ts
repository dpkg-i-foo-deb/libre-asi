import { apiUrl } from '$lib/api/constants';
import { fail, type Action, type Actions } from '@sveltejs/kit';

const signOut: Action = async function ({ cookies, fetch }) {
	let response: Response;

	try {
		response = await fetch(apiUrl + 'sign-out');
	} catch (e) {
		return fail(500, { cannotConnect: true });
	}
};

export const actions: Actions = { signOut };
