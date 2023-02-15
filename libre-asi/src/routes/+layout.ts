//https://blog.encodeart.dev/typesafe-i18n-with-sveltekit
import type { LayoutLoad } from './$types';
import { loadLocaleAsync } from '$lib/i18n/i18n-util.async';
import { setLocale } from '$lib/i18n/i18n-svelte';

export const load = (async (event) => {
	// Locale now comes from the server instead of the route
	const locale = event.data.locale;
	// But we load and set it as before
	await loadLocaleAsync(locale);
	setLocale(locale);

	return event.data;
}) satisfies LayoutLoad;
