import type { HandleFetch } from '@sveltejs/kit';
export const handleFetch = (({ request, fetch }) => {
	request.headers.set('content-type', 'application/json');

	return fetch(request);
}) satisfies HandleFetch;
