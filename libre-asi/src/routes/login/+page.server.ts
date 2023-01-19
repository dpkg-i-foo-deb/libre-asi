import { fail, redirect } from '@sveltejs/kit';
import type { Action, Actions, PageServerLoad } from './$types';

const login: Action = async ({ cookies, request, fetch }) => {
	// redirect the user
	throw redirect(302, '/');
};

export const actions: Actions = { login };
