import { INTERVIEWER_LOGIN, API_URL } from '$lib/api/constants';
import type { JwtPair } from '$lib/models/JwtPair';
import type User from '$lib/models/User';
import type { RequestHandler } from '@sveltejs/kit';
import type { PasswordResetCookie } from '$lib/models/PasswordResetCookie';

export const POST: RequestHandler = async function({ fetch, request, cookies }) {
  const body = 'Logged in';
  const options = { status: 200, statusText: 'Ok' };

  const user = (await request.json()) as User;

  const response = await fetch(API_URL + INTERVIEWER_LOGIN, {
    body: JSON.stringify(user),
    credentials: 'include',
    method: 'POST'
  });

  if (response.ok) {
    const pair = (await response.json()) as JwtPair;

    cookies.set(pair.access_token.name, pair.access_token.value, {
      httpOnly: pair.access_token.http_only,
      expires: new Date(pair.access_token.expires),
      path: pair.access_token.path,
      secure: pair.access_token.secure
    });

    cookies.set(pair.refresh_token.name, pair.refresh_token.value, {
      httpOnly: pair.refresh_token.http_only,
      expires: new Date(pair.refresh_token.expires),
      path: pair.refresh_token.path,
      secure: pair.refresh_token.secure
    });

    return new Response(body, options);
  } else {
    if (response.status == 428) {

      const passwordResetCookie = (await response.json()) as PasswordResetCookie;

      cookies.set(passwordResetCookie.name, passwordResetCookie.value, {
        httpOnly: passwordResetCookie.http_only,
        expires: new Date(passwordResetCookie.expires),
        path: passwordResetCookie.path,
        secure: passwordResetCookie.secure
      })

      return new Response('', {
        status: response.status,
        statusText: response.statusText,
        headers: response.headers
      })
    }
  }


  return response;
};
