<script lang="ts">
	import {
		TextInput,
		PasswordInput,
		Button,
		RadioButtonGroup,
		RadioButton,
		Tooltip,
		InlineNotification
	} from 'carbon-components-svelte';
	import session from '$lib/stores/userStore';
	import { SessionRole } from '$lib/models/Session';
	import { onMount } from 'svelte';
	import type User from '$lib/models/User';
	import { goto } from '$app/navigation';
	import { sendError, sendSuccess } from '$lib/util/notifications';
	import emptyValidator from '$lib/util/emptyValidator';
	import emailValidator from '$lib/util/emailValidator';

	let wantsAdmin = false;
	let wantsInterviewer = false;
	let invalidEmail = false;
	let invalidPassword = false;
	let invalidEmailCaption = '';
	let invalidPasswordCaption = '';
	let email = '';
	let password = '';
	let invalidCredentials = false;

	//TODO make this form less ugly

	onMount(function () {
		//If the user lands here, it means they need a session
		$session.active = false;
		$session.role = SessionRole.None;
	});

	function checkEmail(): boolean {
		invalidEmailCaption = '';
		invalidEmail = true;
		if (!emptyValidator(email)) {
			invalidEmailCaption = 'This field is mandatory';
			return false;
		}

		if (!emailValidator(email)) {
			invalidEmailCaption = 'Enter a valid email address';
			return false;
		}

		invalidEmail = false;
		return true;
	}

	function checkPassword(): boolean {
		invalidPasswordCaption = '';
		invalidPassword = true;

		if (!emptyValidator(password)) {
			invalidPasswordCaption = 'This field is mandatory';
			return false;
		}

		invalidPassword = false;
		return true;
	}

	async function login() {
		invalidCredentials = false;
		let url = '/api/login/';

		if (wantsAdmin) {
			url += 'administrator';
		}

		if (wantsInterviewer) {
			url += 'interviewer';
		}

		if (!(checkEmail() || checkPassword())) {
			return;
		}

		const user: User = { email: email, password: password };

		const response = await fetch(url, {
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

		if (response.status == 401) {
			invalidCredentials = true;
			return;
		}

		if (response.status == 503) {
			sendError(
				'Something went wrong',
				'Try again later, if the problem persists, contact your administrator'
			);
		}
	}
</script>

<main>
	<div class="container">
		<form on:submit|preventDefault={login}>
			<h3>Please Log In to your Account</h3>

			{#if invalidCredentials}
				<div class="invalid-credentials">
					<InlineNotification title="Error" subtitle="Invalid username or password" />
				</div>
			{/if}

			<div class="form-element">
				<TextInput
					labelText="Email"
					placeholder="Enter email..."
					id="email"
					name="email"
					bind:value={email}
					bind:invalid={invalidEmail}
					bind:invalidText={invalidEmailCaption}
					on:blur={checkEmail}
					autofocus
				/>
			</div>

			<div class="form-element">
				<PasswordInput
					labelText="Password"
					placeholder="Enter password..."
					id="password"
					name="password"
					type="password"
					bind:value={password}
					bind:invalid={invalidPassword}
					bind:invalidText={invalidPasswordCaption}
					on:blur={checkPassword}
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
