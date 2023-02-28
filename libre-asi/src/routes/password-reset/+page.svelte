<script lang="ts">
	import { checkPassword, checkPasswordConfirm } from '$lib/util/formUtils';
	import { Button, Form, PasswordInput } from 'carbon-components-svelte';

	let password = '';
	let invalidPassword = false;
	let invalidPasswordCaption = '';

	let passwordConfirm = '';
	let invalidPasswordConfirm = false;
	let invalidPasswordCaptionConfirm = '';

	function validatePassword() {
		invalidPassword = false;
		invalidPasswordCaption = '';

		const result = checkPassword(password);

		invalidPassword = !result[1];
		invalidPasswordCaption = result[0];
	}

	function validatePasswordConfirm() {
		invalidPasswordConfirm = false;
		invalidPasswordCaptionConfirm = '';

		const result = checkPasswordConfirm(password, passwordConfirm);

		invalidPasswordConfirm = !result[1];
		invalidPasswordCaptionConfirm = result[0];
	}

	function resetPassword() {}
</script>

<main>
	<div class="container">
		<Form
			on:submit={(e) => {
				e.preventDefault();
				resetPassword();
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
					on:blur={validatePassword}
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
					on:blur={validatePasswordConfirm}
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
