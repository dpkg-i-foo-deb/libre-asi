import { API_URL, GET_ADMINS } from '$lib/api/constants';
import type Administrator from '$lib/models/Administrator';
import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async function ({ fetch, request }) {
	const response = await fetch(API_URL + GET_ADMINS, {
		credentials: 'include',
		headers: request.headers
	});

	if (response.ok) {
		const administrators: Administrator[] = await response.json();

		return { administrators };
	}

	if (response.status == 503) {
		throw redirect(301, '/cannot-connect');
	}

	return { error: true };
};
