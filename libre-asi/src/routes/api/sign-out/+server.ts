import { apiUrl, signOut as signOutRoute } from '$lib/api/constants';
import type { RequestHandler } from '@sveltejs/kit';

export const POST: RequestHandler = async ({ fetch, cookies, request }) => {
	let body = 'fail';
	const options = { status: 500, statusText: 'Something went wrong' };

	console.log(cookies.get('access-token'));

	let response: Response;
	try {
		response = await fetch(apiUrl + signOutRoute, {
			method: 'POST',
			credentials: 'include',
			headers: request.headers
		});

		options.status = response.status;

		cookies.delete('access-token', { path: '/' });
		cookies.delete('refresh-token', { path: '/' });

		return new Response(body, options);
	} catch (e) {
		return new Response(body, options);
	}
};
