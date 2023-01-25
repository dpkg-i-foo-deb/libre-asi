import type { HandleFetch } from '@sveltejs/kit';
import type { Handle } from '@sveltejs/kit';

export const handleFetch: HandleFetch = (async ({ request, fetch, event }) => {
	request.headers.set('content-type', 'application/json');

	const response = await fetch(request);
	const cookies = event.cookies;

	if (response.status == 401) {
		cookies.delete('access-token', { path: '/' });
		cookies.delete('refresh-token', { path: '/' });
	}

	return response;
}) satisfies HandleFetch;

export const handle: Handle = (async ({ event, resolve }) => {
	//TODO use this to protect routes
	const cookies = event.cookies;

	const response = await resolve(event);

	return response;
}) satisfies Handle;
