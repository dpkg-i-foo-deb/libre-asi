import { goto } from '$app/navigation';
import setup from '$lib/stores/setupStore';

export const checkSetup = checkSetupFn();

function checkSetupFn() {

    


	if (!setup) {
		goto('/set-up');
	}
}
