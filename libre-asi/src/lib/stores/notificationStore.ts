import type { Notification } from '$lib/models/Notification';
import { writable, type Writable } from 'svelte/store';

const defaultNotification: Notification = {
	caption: '',
	kind: 'error',
	subtitle: '',
	timeout: 5,
	title: ''
};

export const notifications: Writable<Notification> = writable(defaultNotification);
