import { writable, type Writable } from 'svelte/store';
import { browser } from '$app/environment';
import { Locale } from '$lib/models/Locale';
let storedLocale = 'en';

let locale: Writable<string> = writable(storedLocale);

if (browser) {
	try {
		storedLocale = JSON.parse(localStorage.getItem('locale') ?? JSON.stringify(storedLocale));
	} catch (e) {
		console.warn('Local storage locale is wrong, fixing...');
	}

	locale = writable(storedLocale);
}

locale.subscribe(function (value: string) {
	if (browser) {
		localStorage.setItem('locale', JSON.stringify(value));
	}
});

export default locale;
