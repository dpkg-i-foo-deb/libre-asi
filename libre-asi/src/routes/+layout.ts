import { setLocale } from '$lib/i18n/i18n-svelte';
import { loadLocaleAsync } from '$lib/i18n/i18n-util.async';
import { Locale } from '$lib/models/Locale';

export const load = async (event) => {
	const locale = Locale.ES;

	await loadLocaleAsync(locale);

	setLocale(locale);

	return event.data;
};
