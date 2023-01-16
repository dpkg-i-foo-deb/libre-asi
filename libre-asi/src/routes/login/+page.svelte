<script lang="ts">
	import {
		TextInput,
		PasswordInput,
		Button,
		ToastNotification,
		RadioButtonGroup,
		RadioButton,
		Tooltip
	} from 'carbon-components-svelte';
	import type user from '$lib/models/user';
	import { goto } from '$app/navigation';
	import { apiUrl, interviewerLogin } from '$lib/api/constants';
	let email: string;
	let password: string;
	let invalidEmail = false;
	let invalidCredentials = false;
	let loginError = false;

	const emailRegex = /^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}$/i;
	let result: Response;

	async function login(): Promise<void> {
		invalidEmail = false;
		invalidCredentials = false;
		loginError = false;

		if (!emailRegex.test(email)) {
			invalidEmail = true;
			return;
		}

		const user: user = { email: email, password: password };

		try {
			result = await fetch(apiUrl + interviewerLogin, {
				headers: { 'Content-Type': 'application/json' },
				method: 'POST',
				body: JSON.stringify(user),
				credentials: 'include',
				mode: 'cors'
			});
		} catch (e) {
			loginError = true;
			console.log(e);
			return;
		}

		if (!result.ok && result.status == 401) {
			invalidCredentials = true;
		}

		if (result.ok) {
			goto('/');
		}
	}
</script>

<main>
	<div class="container">
		<form on:submit|preventDefault={login}>
			<h3>Please Log In to your Account</h3>

			<div class="form-element">
				<TextInput
					labelText="Email"
					placeholder="Enter email..."
					required
					bind:invalid={invalidEmail}
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
					bind:value={password}
				/>
			</div>

			<div class="form-element">
				<RadioButtonGroup selected="interviewer">
					<div slot="legendText" style="display:flex">
						Account type
						<Tooltip>
							<p>Your account type is determined by your administrator</p>
						</Tooltip>
					</div>
					<RadioButton labelText="Interviewer" value="interviewer" />
					<RadioButton labelText="Administrator" value="administrator" />
				</RadioButtonGroup>
			</div>

			{#if invalidCredentials}
				<div class="invalid-credentials">
					<p>Check your credentials</p>
				</div>
			{/if}

			<div class="button-container">
				<Button type="submit">Log In</Button>
			</div>
		</form>
	</div>
	{#if loginError}
		<div class="error-notification">
			<ToastNotification
				title="Login Error"
				subtitle="Try again later"
				caption="If the error persists, contact your administrator"
			/>
		</div>
	{/if}
</main>

<style>
	.container {
		padding: 20px;
		display: flex;
		justify-content: center;
		align-items: center;
		text-align: center;
		min-height: 99vh;
	}

	.form-element {
		padding-top: 15px;
		padding-bottom: 15px;
	}

	.button-container {
		padding-top: 10px;
		float: left;
	}

	.error-notification {
		position: absolute;
		bottom: 0px;
		margin: 10px;
	}

	.invalid-credentials {
		padding-top: 5px;
		padding-bottom: 5px;
		color: red;
	}
</style>
