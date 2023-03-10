import type PasswordReset from '$lib/models/PasswordReset';
import type { RequestHandler } from './$types';
import { API_URL, PASSWORD_RESET } from '$lib/api/constants';

export const POST: RequestHandler = async function({ fetch, request, cookies }) {
  const credentials = (await request.json()) as PasswordReset;


  const response = await fetch(API_URL + PASSWORD_RESET, {
    method: 'POST',
    credentials: 'include',
    headers: request.headers,
    body: JSON.stringify(credentials)
  });

  if (response.ok) {
    cookies.delete("password-reset-token", { path: '/' })
  }

  return response;
};
