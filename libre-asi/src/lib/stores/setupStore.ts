import { browser } from '$app/environment';
import { writable, type Writable } from 'svelte/store';

let isSetup = false;

let setup: Writable<boolean> = writable(isSetup);

if (browser) {
	try {
		isSetup = JSON.parse(localStorage.getItem('setup') ?? JSON.stringify(isSetup));
	} catch (e) {
		console.warn('Local storage setup is wrong, fixing...');
		isSetup = false;
	}
	setup = writable(isSetup);
}

setup.subscribe(function (value: boolean) {
	if (browser) {
		localStorage.setItem('setup', JSON.stringify(value));
	}
});

export default setup;
