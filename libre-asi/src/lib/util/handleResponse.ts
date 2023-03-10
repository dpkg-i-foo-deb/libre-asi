import { goto } from '$app/navigation';
import setup from '$lib/stores/setupStore';
import { sendError } from './notifications';
import { sendInfo } from './notifications';

export function handleResponse(code: number, allow401: boolean): boolean {
	let shouldNavigate = true;

	switch (code) {
		//TODO not enough privileges page
		case 403:
			goto('/');
			break;

		case 503:
			goto('/cannot-connect');
			setup.set(false);
			break;
		case 500:
			sendError('Something went wrong', 'Internal server error');
			break;
		case 412:
			setup.set(false);
			goto('/set-up');
			break;

		case 401:
			if (!allow401) {
				sendError('Your session has expired', 'Log In again');
				goto('/login');
			}
			break;
		case 428:
			sendInfo('You need to set a new password to continue', 'Set new password');
			goto('/password-reset');
			break;
		default:
			shouldNavigate = false;
	}

	return shouldNavigate;
}
