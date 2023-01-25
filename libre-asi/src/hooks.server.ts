import { redirect, type HandleFetch } from '@sveltejs/kit';
import type { Handle } from '@sveltejs/kit';
import { PROTECTED_ROUTES } from '$lib/protected/protected';
import { API_URL, REFRESH } from '$lib/api/constants';

export const handleFetch: HandleFetch = (async ({ request, fetch, event }) => {
	request.headers.set('content-type', 'application/json');

	const response = await fetch(request);
	const cookies = event.cookies;

	if (
		response.status == 401 &&
		!request.url.includes('login') &&
		!request.url.includes('refresh')
	) {
		const refresh = await fetch(API_URL + REFRESH, {
			credentials: 'include',
			headers: request.headers
		});

		if (refresh.ok) {
			const newResponse = await fetch(request);

			if (newResponse.ok) {
				return newResponse;
			}
		}

		cookies.set('access-token', '', { path: '/' });
		cookies.set('refresh-token', '', { path: '/' });
		throw redirect(302, '/login');
	}

	if (response.status == 403) {
		//TODO redirect to a 'Not enough privileges page'
	}

	return response;
}) satisfies HandleFetch;

export const handle: Handle = (async ({ event, resolve }) => {
	//TODO use this to protect routes
	const cookies = event.cookies;

	const response = await resolve(event);

	return response;
}) satisfies Handle;
