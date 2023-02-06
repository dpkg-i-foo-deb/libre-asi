import { redirect, type HandleFetch } from '@sveltejs/kit';
import type { Handle } from '@sveltejs/kit';
import { PROTECTED_ROUTES } from '$lib/protected/protectedRoutes';
import { API_URL, REFRESH } from '$lib/api/constants';
import type { JwtPair } from '$lib/models/JwtPair';

export const handleFetch: HandleFetch = async ({ request, fetch, event }) => {
	request.headers.set('content-type', 'application/json');
	let body = 'Cannot connect to the server';
	let options = { status: 503, statusText: 'Cannot connect' };
	let response: Response
	const cookies = event.cookies;
	try {
		//TODO idk if this automatically forwards all the cookies
		response = await fetch(request);
	} catch (e) {
		console.error(e);
		cookies.set('access-token', '', { path: '/' });
		cookies.set('refresh-token', '', { path: '/' });
		return new Response(body, options);
	}

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

			const newRequest = request.clone();
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

	return response;

};

export const handle: Handle = async ({ event, resolve }) => {
	//TODO use this to protect routes
	const cookies = event.cookies;
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
	}

	if ((accessToken != '' || refreshToken != '') && event.url.pathname == '/login') {
		throw redirect(302, '/');
	}

	const response = await resolve(event);

	return response;
};
