//https://blog.logrocket.com/intercepting-javascript-fetch-api-requests-responses/
//https://stackoverflow.com/questions/65638438/interceptor-for-fetch-and-fetch-retry-javascript
//https://joyofcode.xyz/sveltekit-window-is-not-defined

import { browser } from '$app/environment';
import { apiUrl, refresh } from './constants';

//This custom fetch function provides some utility when it comes to handling token refreshing

if (browser) {
	const { fetch: originalFetch } = window;

	window.fetch = async function (...args) {
		const [url, config] = args;

		let response: Response;

		return originalFetch.apply(url, args).then(async function (data) {
			if (data.status == 401) {
				response = await originalFetch(apiUrl + refresh, {
					method: 'POST',
					credentials: 'include',
					mode: 'cors'
				});

				if (response.status == 401) {
					return data;
				}
				return fetch(url, config);
			} else {
				return data;
			}
		});
	};
}

export default fetch;
