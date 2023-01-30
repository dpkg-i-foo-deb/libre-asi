import { API_URL } from '$lib/api/constants';
import type Administrator from '$lib/models/Administrator';
import type { RequestHandler } from '@sveltejs/kit';

export const POST: RequestHandler = async function ({ fetch, request }) {
	const body = 'Success';
	const options = { status: 200, statusText: 'Administrator Registered' };

	const admin = (await request.json()) as Administrator;

	try {
		const response = await fetch(API_URL);
		return response;
	} catch (e) {
		throw e;
	}
};
