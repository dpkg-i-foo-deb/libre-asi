<script lang="ts">
	import { applyAction, enhance } from '$app/forms';
	import {
		TextInput,
		PasswordInput,
		Button,
		RadioButtonGroup,
		RadioButton,
		Tooltip,
		InlineNotification
	} from 'carbon-components-svelte';
	import type { ActionData } from './$types';
	import session from '$lib/stores/userStore';
	import type { ActionResult } from '@sveltejs/kit';
	import { SessionRole } from '$lib/models/Session';
	import { onMount } from 'svelte';
	import type User from '$lib/models/User';
	import { goto } from '$app/navigation';
	import { sendSuccess } from '$lib/util/notifications';

	export let form: ActionData;

	let wantsAdmin = false;
	let wantsInterviewer = false;
	let email = '';
	let password = '';

	//TODO make this form less ugly

	onMount(function () {
		//If the user lands here, it means they need a session
		$session.active = false;
		$session.role = SessionRole.None;
	});

	async function login() {
		const user: User = { email: email, password: password };

		const response = await fetch('/api/login', {
			method: 'POST',
			credentials: 'include',
			body: JSON.stringify(user)
		});

		if (response.ok) {
			$session.active = true;
			if (wantsAdmin) {
				$session.role = SessionRole.Admin;
			}
			if (wantsInterviewer) {
				$session.role = SessionRole.Interviewer;
			}

			sendSuccess('Logged In', 'Welcome back');

			goto('/home');

			return;
		}
	}
</script>

<main>
	<div class="container">
		<form on:submit|preventDefault={login}>
			<h3>Please Log In to your Account</h3>

			{#if form?.invalidCredentials}
				<div class="invalid-credentials">
					<InlineNotification title="Error" subtitle="Invalid username or password" />
				</div>
			{/if}

			<div class="form-element">
				<TextInput
					labelText="Email"
					placeholder="Enter email..."
					required
					invalidText="Check your email"
					id="email"
					name="email"
					bind:value={email}
				/>
			</div>

			<div class="form-element">
				<PasswordInput
					labelText="Password"
					placeholder="Enter password..."
					id="password"
					name="password"
					required
					invalidText="Check your password"
					type="password"
					bind:value={password}
				/>
			</div>

			<div class="form-element">
				<RadioButtonGroup selected="interviewer" required>
					<div slot="legendText" style="display:flex">
						Account type
						<Tooltip>
							<p>Your account type is determined by your administrator</p>
						</Tooltip>
					</div>
					<RadioButton
						labelText="Interviewer"
						value="interviewer"
						bind:checked={wantsInterviewer}
						id="interviewer"
						name="interviewer"
					/>
					<RadioButton
						labelText="Administrator"
						value="admin"
						bind:checked={wantsAdmin}
						id="administrator"
						name="administrator"
					/>
				</RadioButtonGroup>
			</div>

			<div class="button-container">
				<Button type="submit">Log In</Button>
			</div>
		</form>
	</div>
</main>

<style>
	.container {
		padding: 20px;
		display: flex;
		justify-content: center;
		align-items: center;
		text-align: center;
		min-height: 85vh;
	}

	.form-element {
		padding-top: 15px;
		padding-bottom: 15px;
	}

	.button-container {
		padding-top: 10px;
		float: left;
	}

	.invalid-credentials {
		padding-top: 20px;
	}
</style>
