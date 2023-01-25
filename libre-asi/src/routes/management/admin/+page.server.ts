import { apiUrl, getAdmins } from '$lib/api/constants';
import type Administrator from '$lib/models/Administrator';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async function ({ cookies, fetch, request }) {
	let response: Response;

	try {
		response = await fetch(apiUrl + getAdmins, {
			credentials: 'include',
			headers: { Cookie: 'access-token=' + cookies.get('access-token') }
		});

		if (response.ok) {
			const administrators: Administrator[] = await response.json();

			return { administrators };
		}
	} catch (e) {
		console.log(e);
	}

	//cookies.delete('access-token', { path: '/' });
	//cookies.delete('refresh-token', { path: '/' });

	return { error: true };
};
