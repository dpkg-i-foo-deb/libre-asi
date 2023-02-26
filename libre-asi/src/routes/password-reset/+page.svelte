<script lang="ts">
	import emptyValidator from '$lib/util/emptyValidator';
	import { Button, Form, PasswordInput } from 'carbon-components-svelte';

	let password = '';
	let invalidPassword = false;
	let invalidPasswordCaption = '';

	let passwordConfirm = '';
	let invalidPasswordConfirm = false;
	let invalidPasswordCaptionConfirm = '';

	function checkPassword() {
		invalidPassword = false;
		invalidPasswordCaption = '';
		if (!emptyValidator(password)) {
			invalidPassword = true;
			invalidPasswordCaption = 'Password is required';
		}
	}

	function checkPasswordConfirm() {
		invalidPasswordConfirm = false;
		invalidPasswordCaptionConfirm = '';
		if (!emptyValidator(passwordConfirm)) {
			invalidPasswordConfirm = true;
			invalidPasswordCaptionConfirm = 'Password confirmation is required';
		}
	}
</script>

<main>
	<div class="container">
		<Form
			on:submit={(e) => {
				e.preventDefault();
			}}
		>
			<div class="title">
				<h3>Set a new password</h3>
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
					autofocus
				/>
			</div>
			<div class="form-element">
				<PasswordInput
					labelText="Confirm password"
					placeholder="Enter the same password..."
					id="password-confirm"
					name="password-confirm"
					type="password"
					bind:value={passwordConfirm}
					bind:invalid={invalidPasswordConfirm}
					bind:invalidText={invalidPasswordCaptionConfirm}
					on:blur={checkPasswordConfirm}
				/>
			</div>

			<div class="button-container">
				<div class="main-button">
					<Button type="submit">Submit</Button>
				</div>
			</div>
		</Form>
	</div>
</main>

<style>
	.container {
		padding: 20px;
		display: flex;
		justify-content: center;
		align-items: center;
		text-align: center;
		min-height: 70vh;
	}

	.title {
		margin-bottom: 35px;
	}

	.form-element {
		margin-top: 45px;
		margin-bottom: 45px;
		height: 50px;
		width: 350px;
	}

	.button-container {
		float: right;
		margin-right: 0;
	}

	.main-button {
		margin-top: 15px;
		width: 100%;
	}
</style>
