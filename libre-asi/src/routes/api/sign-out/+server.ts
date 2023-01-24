import { apiUrl, signOut as signOutRoute } from '$lib/api/constants';
import type { RequestHandler } from '@sveltejs/kit';

export const POST: RequestHandler = async () => {
	const response = fetch(apiUrl + signOutRoute, { method: 'POST' });

	return response;
};
