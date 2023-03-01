<script lang="ts">
	import { checkPassword, checkPasswordConfirm } from '$lib/util/formUtils';
	import { Button, Form, PasswordInput } from 'carbon-components-svelte';

	let currentPassword = '';
	let invalidCurrentPassword = false;
	let invalidCurrentPasswordCaption = '';

	let newPassword = '';
	let newInvalidPassword = false;
	let newInvalidPasswordCaption = '';

	let passwordConfirm = '';
	let invalidPasswordConfirm = false;
	let invalidPasswordCaptionConfirm = '';

	function validateCurrentPassword() {
		invalidCurrentPassword = false;
		invalidCurrentPasswordCaption = '';

		const result = checkPassword(currentPassword);

		invalidCurrentPassword = !result[1];
		invalidCurrentPasswordCaption = result[0];
	}

	function validatePassword() {
		newInvalidPassword = false;
		newInvalidPasswordCaption = '';

		const result = checkPassword(newPassword);

		newInvalidPassword = !result[1];
		newInvalidPasswordCaption = result[0];
	}

	function validatePasswordConfirm() {
		invalidPasswordConfirm = false;
		invalidPasswordCaptionConfirm = '';

		const result = checkPasswordConfirm(newPassword, passwordConfirm);

		invalidPasswordConfirm = !result[1];
		invalidPasswordCaptionConfirm = result[0];
	}

	function resetPassword() {}
</script>

//TODO use stepper widget
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
					labelText="Current password"
					placeholder="Enter current password..."
					id="current-password"
					name="current-password"
					type="password"
					bind:value={currentPassword}
					bind:invalid={invalidCurrentPassword}
					bind:invalidText={invalidCurrentPasswordCaption}
					on:blur={validateCurrentPassword}
				/>
			</div>

			<div class="form-element">
				<PasswordInput
					labelText="New password"
					placeholder="Enter new password..."
					id="new-password"
					name="newPassword"
					type="password"
					bind:value={newPassword}
					bind:invalid={newInvalidPassword}
					bind:invalidText={newInvalidPasswordCaption}
					on:blur={validatePassword}
				/>
			</div>
			<div class="form-element">
				<PasswordInput
					labelText="Confirm new password"
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
