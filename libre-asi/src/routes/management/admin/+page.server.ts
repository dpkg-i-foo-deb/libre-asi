import { API_URL, GET_ADMINS } from '$lib/api/constants';
import type Administrator from '$lib/models/Administrator';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async function ({ cookies, fetch, request }) {
	let response: Response;

	try {
		response = await fetch(API_URL + GET_ADMINS, {
			credentials: 'include',
			headers: request.headers
		});

		if (response.ok) {
			const administrators: Administrator[] = await response.json();

			return { administrators };
		}
	} catch (e) {
		console.log(e);
	}
	return { error: true };
};
