import { browser } from '$app/environment';
import type { Session } from '$lib/models/Session';
import { writable, type Writable } from 'svelte/store';
import { SessionRole } from '$lib/models/Session';

let storedSession: Session = { active: false, role: SessionRole.None };
let session: Writable<Session> = writable(storedSession);

if (browser) {
	storedSession = JSON.parse(localStorage.getItem('session') ?? JSON.stringify(storedSession));
	session = writable(storedSession);
}

session.subscribe(function (value: Session) {
	if (browser) {
		localStorage.setItem('session', JSON.stringify(value));
	}
});

export default session;
