import { API_URL, SET_UP } from '$lib/api/constants';
import type { RequestHandler } from '@sveltejs/kit';

export const GET: RequestHandler = async function ({ fetch }) {
	return await fetch(API_URL + SET_UP, { method: 'GET' });
};
