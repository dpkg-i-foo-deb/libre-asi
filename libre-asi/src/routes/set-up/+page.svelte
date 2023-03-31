<script lang="ts">
	import type Administrator from '$lib/models/Administrator';
	import { checkEmail, checkPassword, checkUsername } from '$lib/util/formUtils';
	import { handleResponse } from '$lib/util/handleResponse';
	import { sendSuccess } from '$lib/util/notifications';
	import setup from '$lib/stores/setupStore';
	import {
		ButtonSet,
		Form,
		Button,
		TextInput,
		PasswordInput,
		ProgressIndicator,
		ProgressStep,
		InlineLoading
	} from 'carbon-components-svelte';
	import { goto } from '$app/navigation';
	import { API_URL, SET_UP } from '$lib/api/constants';
	import { fetchNoRefresh } from '$lib/util/fetch';

	let loading = false;

	let stepIndex = 0;
	let email = '';
	let password = '';
	let username = '';

	let invalidEmail = false;
	let invalidPassword = false;
	let invalidUsername = false;

	let invalidEmailCaption = '';
	let invalidPasswordCaption = '';
	let invalidUsernameCaption = '';

	function validateUsername(): boolean {
		const usernameField = checkUsername(username);

		invalidUsernameCaption = usernameField[0];
		invalidUsername = !usernameField[1];

		return usernameField[1];
	}

	function validateEmail(): boolean {
		const emailField = checkEmail(email);

		invalidEmailCaption = emailField[0];
		invalidEmail = !emailField[1];

		return emailField[1];
	}

	function validatePassword(): boolean {
		const passwordField = checkPassword(password);

		invalidPasswordCaption = passwordField[0];
		invalidPassword = !passwordField[1];

		return passwordField[1];
	}

	function back() {
		if (stepIndex > 0) {
			stepIndex--;
		}
	}

	function handleNext() {
		switch (stepIndex) {
			case 0:
				if (!validateEmail()) {
					return;
				}
				break;
			case 1:
				if (!validateUsername()) {
					return;
				}
				break;

			case 2:
				if (!validatePassword()) {
					return;
				}
				break;

			default:
				stepIndex = 0;
		}

		if (stepIndex < 2) {
			stepIndex++;
			return;
		}

		if (stepIndex == 2) {
			if (!validateEmail() || !validateUsername() || !validatePassword()) {
				return;
			}

			register();
		}
	}

	async function register() {
		loading = true;

		const newAdmin: Administrator = {
			email: email,
			username: username,
			password: password
		};

		const response = await fetchNoRefresh(API_URL+SET_UP,{
			method:'POST',
			body:JSON.stringify(newAdmin)
		})

		if (response.ok) {
			//Set the setup store to true and go to the welcome page
			$setup = true;

			sendSuccess('Success', 'Set-up completed, you can log in now');

			goto('/');
		}

		handleResponse(response.status, false);

		loading = false;
	}
</script>

<h1>Set-Up Required</h1>
<br />
<h3>Create an administrator account</h3>

<div class="container">
	<Form
		on:submit={(e) => {
			e.preventDefault();
			handleNext();
		}}
	>
		<div class="stepper">
			<ProgressIndicator bind:currentIndex={stepIndex} spaceEqually preventChangeOnClick>
				<ProgressStep complete={stepIndex > 0} label="Step 1" bind:invalid={invalidEmail} />
				<ProgressStep complete={stepIndex > 1} label="Step 2" bind:invalid={invalidUsername} />
				<ProgressStep complete={stepIndex > 2} label="Step 3" bind:invalid={invalidPassword} />
			</ProgressIndicator>
		</div>

		{#if loading}
			<InlineLoading description="Submitting" />
		{/if}

		{#if stepIndex == 0}
			<div class="form-element">
				<TextInput
					labelText="Email"
					placeholder="Enter email..."
					autofocus
					bind:value={email}
					on:blur={validateEmail}
					bind:invalid={invalidEmail}
					bind:invalidText={invalidEmailCaption}
				/>
			</div>
		{/if}

		{#if stepIndex == 1}
			<div class="form-element">
				<TextInput
					labelText="Username"
					placeholder="Enter username..."
					bind:value={username}
					bind:invalid={invalidUsername}
					bind:invalidText={invalidUsernameCaption}
					on:blur={validateUsername}
					autofocus
				/>
			</div>
		{/if}

		{#if stepIndex == 2}
			<div class="form-element">
				<PasswordInput
					labelText="Password"
					placeholder="Enter password..."
					autofocus
					bind:value={password}
					bind:invalid={invalidPassword}
					bind:invalidText={invalidPasswordCaption}
					on:blur={validatePassword}
				/>
			</div>
		{/if}

		<ButtonSet>
			<div class="main-button-container">
				<Button kind="secondary" disabled={loading || stepIndex == 0} on:click={back}>Back</Button>
				<Button type="submit" disabled={loading}>Next</Button>
			</div>
		</ButtonSet>
	</Form>
</div>

<style>
	h1 {
		margin-top: 10px;
		margin-bottom: 10px;
	}

	.main-button-container {
		display: flex;
		margin-top: 15px;
		flex-direction: row;
	}

	.container {
		display: flex;
		flex-flow: 0;
		justify-content: center;
		align-items: center;
		text-align: center;
		min-height: 50px;
		width: 100%;
	}

	.form-element {
		padding-top: 15px;
		padding-bottom: 15px;
		width: 100%;
		height: 100px;
	}

	.stepper {
		margin-top: 25px;
		margin-bottom: 25px;
	}
</style>
