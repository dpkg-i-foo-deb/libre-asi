import { browser } from '$app/environment';
import { writable } from 'svelte/store';

export const session = writable((browser && localStorage.getItem('activeSession')) || 'false');
export const role = writable((browser && localStorage.getItem('role')) || 'none');

session.subscribe((value) => {
	if (browser) {
		return (localStorage.activeSession = value);
	}
});

role.subscribe((value) => {
	if (browser) {
		return (localStorage.role = value);
	}
});
