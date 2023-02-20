//https://blog.encodeart.dev/typesafe-i18n-with-sveltekit
import type { LayoutServerLoad } from './$types';
import { detectLocale } from '$lib/i18n/i18n-util';

const langParam = 'lang';

export const load = (async (event) => {
	// Get the locale from the cookie
	const locale = detectLocale(() => [event.cookies.get(langParam) ?? '']);

	return { locale };
}) satisfies LayoutServerLoad;
