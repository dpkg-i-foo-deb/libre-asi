import { apiUrl, signOut as signOutRoute } from '$lib/api/constants';
import type { RequestHandler } from '@sveltejs/kit';

export const POST: RequestHandler = async ({ fetch }) => {
	const body = 'fail';
	const failOptions = { status: 500, statusText: 'Something went wrong' };

	let response: Response;
	try {
		response = await fetch(apiUrl + signOutRoute, { method: 'POST' });
		return response;
	} catch (e) {}

	return new Response(body, failOptions);
};
