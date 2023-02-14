import { setLocale } from '$lib/i18n/i18n-svelte';
import { loadLocaleAsync } from '$lib/i18n/i18n-util.async';
import { Locale } from '$lib/models/Locale';

export const load = async (event) => {
	await setUpLocale();
	return event.data;
};

async function setUpLocale() {
	const locale = Locale.EN;

	await loadLocaleAsync(locale);

	setLocale(locale);
}
