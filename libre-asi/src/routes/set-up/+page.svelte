<script lang="ts">
	import { checkEmail, checkPassword, checkUsername } from '$lib/util/formUtils';
	import {
		ButtonSet,
		Form,
		Button,
		TextInput,
		PasswordInput,
		ProgressIndicator,
		ProgressStep
	} from 'carbon-components-svelte';

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
		invalidUsername = usernameField[1];

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

	function handleNext() {
		switch (stepIndex) {
			case 0:
				if (!validateEmail()) {
					return;
				}
				break;
			case 1:
				if (!validatePassword()) {
					return;
				}
				break;
		}

		if (stepIndex < 1) {
			stepIndex++;
			return;
		}
	}
</script>

<h1>Set-Up Required</h1>
<br />
<h3>The application was just initialized, you need to create an administrator account</h3>

<div class="container">
	<Form
		on:submit={(e) => {
			e.preventDefault();
			handleNext();
		}}
	>
		<div class="stepper">
			<ProgressIndicator bind:currentIndex={stepIndex} spaceEqually preventChangeOnClick>
				<ProgressStep complete={stepIndex > 0} label="Identification" bind:invalid={invalidEmail} />
				<ProgressStep complete={stepIndex > 1} label="Password" />
			</ProgressIndicator>
		</div>

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
				<PasswordInput
					labelText="Password"
					placeholder="Enter password..."
					autofocus
					bind:value={password}
				/>
			</div>
		{/if}

		<ButtonSet>
			<div class="main-button-container">
				<Button type="submit">Next</Button>
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
		margin-left: auto;
		margin-right: 0;
	}

	.container {
		display: flex;
		justify-content: center;
		align-items: center;
		text-align: center;
		min-height: 60vh;
		width: 100%;
	}

	.form-element {
		padding-top: 15px;
		padding-bottom: 15px;
		width: 350px;
	}

	.stepper {
		margin-top: 25px;
		margin-bottom: 25px;
	}
</style>
