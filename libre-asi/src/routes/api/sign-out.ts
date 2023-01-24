import { apiUrl, signOut as signOutRoute } from '$lib/api/constants';
import type { RequestHandler } from '@sveltejs/kit';

export const POST: RequestHandler = async function signOut({ cookies, fetch, request }) {
	const response = fetch(apiUrl + signOutRoute);

	return response;
};
