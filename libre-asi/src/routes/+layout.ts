import { setLocale } from '$lib/i18n/i18n-svelte';
import { loadLocaleAsync } from '$lib/i18n/i18n-util.async';
import type { LayoutLoad } from './$types';

export const load = (async (event) => {
	const locale = 'en';

	await loadLocaleAsync(locale);

	setLocale(locale);

	return event.data;
}) satisfies LayoutLoad;
