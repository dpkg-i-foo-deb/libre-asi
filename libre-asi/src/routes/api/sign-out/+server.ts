import type { RequestHandler } from '@sveltejs/kit';

export const POST: RequestHandler = async function ({ cookies }) {
	let body = 'Success';
	const options = { status: 200, statusText: 'Logged out' };

	cookies.delete('access-token', { path: '/' });
	cookies.delete('refresh-token', { path: '/' });

	return new Response(body, options);
};
