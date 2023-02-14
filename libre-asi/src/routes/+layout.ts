import { setLocale } from '$lib/i18n/i18n-svelte';
import { loadLocaleAsync } from '$lib/i18n/i18n-util.async';
import { Locale } from '$lib/models/Locale';
import locale from '$lib/stores/localeStore';
import { get } from 'svelte/store';

export const load = async (event) => {
	await setUpLocale();
	return event.data;
};

async function setUpLocale() {
	const storedLocale = get(locale);

	switch (storedLocale) {
		case Locale.EN:
			await loadLocaleAsync(Locale.EN);
			setLocale(Locale.EN);
			break;
		case Locale.ES:
			await loadLocaleAsync(Locale.ES);
			setLocale(Locale.ES);
			break;

		default:
			await loadLocaleAsync(Locale.EN);
			setLocale(Locale.EN);
	}
}
