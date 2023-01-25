import { redirect, type HandleFetch } from '@sveltejs/kit';
import type { Handle } from '@sveltejs/kit';
import { generalRoutes } from '$lib/protected/general';

export const handleFetch: HandleFetch = (async ({ request, fetch, event }) => {
	request.headers.set('content-type', 'application/json');

	const response = await fetch(request);
	const cookies = event.cookies;

	if (response.status == 401 && !request.url.includes('login')) {
		//Try to refresh the access token
		cookies.delete('access-token', { path: '/' });
		cookies.delete('refresh-token', { path: '/' });
	}

	if (response.status == 403) {
		//TODO redirect to a 'Not enough privileges page'
	}

	return response;
}) satisfies HandleFetch;

export const handle: Handle = (async ({ event, resolve }) => {
	//TODO use this to protect routes
	const cookies = event.cookies;

	if (cookies.get('access-token') == undefined) {
		generalRoutes.forEach(function (value) {
			if (event.url.pathname.includes(value)) {
				throw redirect(302, '/login');
			}
		});
	}

	const response = await resolve(event);

	return response;
}) satisfies Handle;
