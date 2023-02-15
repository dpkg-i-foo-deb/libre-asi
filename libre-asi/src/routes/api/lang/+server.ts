import type { RequestHandler } from '@sveltejs/kit';
import type { Locales } from '$lib/i18n/i18n-types';

export const POST: RequestHandler = async function ({ cookies, request }) {
	const locale = (await request.json()) as Locales;

	cookies.set('lang', locale);

	return new Response('Language set', { status: 200 });
};
