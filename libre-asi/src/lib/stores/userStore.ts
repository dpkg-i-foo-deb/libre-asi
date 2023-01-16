import { browser } from '$app/environment';
import { writable } from 'svelte/store';

export const session = writable((browser && localStorage.getItem('activeSession')) || 'false');

session.subscribe((value) => {
	if (browser) {
		return (localStorage.activeSession = value);
	}
});
