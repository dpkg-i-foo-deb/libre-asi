import { API_URL, REFRESH } from '$lib/api/constants';

export async function fetchWithRefresh(url: string, options?: RequestInit): Promise<Response> {
	const headers = new Headers(options?.headers);

	headers.append('Accept', 'application/json');
	headers.append('Content-Type', 'application/json');

	options = {
		...options,
		headers,
		credentials: 'include',
		mode: 'cors'
	};

	try {
		const response = await fetch(url, options);

		if (response.status === 401) {
			const refreshResponse = await fetch(API_URL + REFRESH, {
				method: 'POST',
				credentials: 'include',
				mode: 'cors'
			});

			if (refreshResponse.ok) {
				const newResponse = await fetch(url, options);

				if (newResponse.ok) {
					return newResponse;
				}
			}

			return refreshResponse;
		}

		return response;
	} catch (e) {
		return new Response('Service unavailable', { status: 503 });
	}
}
