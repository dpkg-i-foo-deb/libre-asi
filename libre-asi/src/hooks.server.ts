import { redirect, type HandleFetch } from '@sveltejs/kit';
import type { Handle } from '@sveltejs/kit';
import { PROTECTED_ROUTES } from '$lib/protected/protectedRoutes';
import { API_URL, REFRESH } from '$lib/api/constants';
import type { JwtPair } from '$lib/models/JwtPair';
import type { Session } from '$lib/models/Session';

export const handleFetch: HandleFetch = async ({ request, fetch, event }) => {
	request.headers.set('content-type', 'application/json');

	const response = await fetch(request);
	const cookies = event.cookies;

	if (
		response.status == 401 &&
		!request.url.includes('login') &&
		!request.url.includes('refresh')
	) {
		const refresh = await fetch(API_URL + REFRESH, {
			method: 'POST',
			credentials: 'include',
			headers: request.headers
		});

		if (refresh.ok) {
			//Set the new access and refresh token as cookies
			const newCookies = (await refresh.json()) as JwtPair;

			console.log(newCookies.access_token);

			cookies.set('access-token', newCookies.access_token.value, {
				path: '/',
				httpOnly: true
			});
			cookies.set('refresh-token', newCookies.refresh_token.value, {
				path: '/',
				httpOnly: true
			});

			//Make a new request and pray for it to work

			let newRequest = request.clone();
			newRequest.headers.set('cookie', 'access-token=' + newCookies.access_token.value);

			const newResponse = await fetch(newRequest, {
				credentials: 'include'
			});

			if (newResponse.ok) {
				return newResponse;
			}
		}

		cookies.set('access-token', '', { path: '/' });
		cookies.set('refresh-token', '', { path: '/' });
	}

	if (response.status == 403) {
		//TODO redirect to a 'Not enough privileges page'
	}

	return response;
};

export const handle: Handle = async ({ event, resolve }) => {
	//TODO use this to protect routes
	const cookies = event.cookies;
	let session: Session;
	let protectedRoute = false;

	const accessToken = cookies.get('access-token') ?? '';
	const refreshToken = cookies.get('refresh-token') ?? '';

	PROTECTED_ROUTES.forEach(function (route) {
		if (event.url.pathname.includes(route)) {
			protectedRoute = true;
		}
	});

	if (protectedRoute) {
		if (accessToken == '' && refreshToken == '') {
			throw redirect(302, '/login');
		}

		if (accessToken != '' || refreshToken != '') {
			throw redirect(302, '/');
		}
	}

	const response = await resolve(event);

	return response;
};
