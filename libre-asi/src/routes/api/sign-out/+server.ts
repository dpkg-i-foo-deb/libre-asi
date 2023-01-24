import { apiUrl, signOut as signOutRoute } from '$lib/api/constants';
import type { RequestHandler } from '@sveltejs/kit';

export const POST: RequestHandler = async ({ fetch, cookies }) => {
	const body = 'fail';
	const failOptions = { status: 500, statusText: 'Something went wrong' };

	let response: Response;
	try {
		response = await fetch(apiUrl + signOutRoute, { method: 'POST' });

		cookies.delete('access-token', { path: '/' });
		cookies.delete('refresh-token', { path: '/' });

		return new Response();
	} catch (e) {
		return new Response(body, failOptions);
	}
};
