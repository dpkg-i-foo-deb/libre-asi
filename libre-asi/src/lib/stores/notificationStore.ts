import type { Notification } from '$lib/models/Notification';
import { writable, type Writable } from 'svelte/store';

const defaultNotification: Notification = {
	caption: '',
	kind: undefined,
	subtitle: '',
	timeout: 0,
	title: '',
	visible: false
};

export const notifications: Writable<Notification> = writable(defaultNotification);
