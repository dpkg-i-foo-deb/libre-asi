<script lang="ts">
	import { goto } from '$app/navigation';
	import { ADMIN_PASSWORD_RESET, API_URL, INTERVIEWER_PASSWORD_RESET } from '$lib/api/constants';
	import type PasswordReset from '$lib/models/PasswordReset';
	import { fetchNoRefresh } from '$lib/util/fetch';
	import { checkPassword, checkPasswordConfirm } from '$lib/util/formUtils';
	import { handleResponse } from '$lib/util/handleResponse';
	import { sendSuccess } from '$lib/util/notifications';
	import {
		Button,
		Form,
		InlineLoading,
		InlineNotification,
		PasswordInput,
		RadioButton,
		RadioButtonGroup,
		Tooltip
	} from 'carbon-components-svelte';

	//TODO use stepper widget

	let currentPassword = '';
	let invalidCurrentPassword = false;
	let invalidCurrentPasswordCaption = '';

	let newPassword = '';
	let newInvalidPassword = false;
	let newInvalidPasswordCaption = '';

	let passwordConfirm = '';
	let invalidPasswordConfirm = false;
	let invalidPasswordCaptionConfirm = '';

	let invalidCredentials = false;

	let loading = false;

	let wantsAdmin = false;
	let wantsInterviewer = false;

	function validateCurrentPassword(): boolean {
		invalidCurrentPassword = false;
		invalidCurrentPasswordCaption = '';

		const result = checkPassword(currentPassword);

		invalidCurrentPassword = !result[1];
		invalidCurrentPasswordCaption = result[0];

		return result[1];
	}

	function validateNewPassword(): boolean {
		newInvalidPassword = false;
		newInvalidPasswordCaption = '';

		const result = checkPassword(newPassword);

		newInvalidPassword = !result[1];
		newInvalidPasswordCaption = result[0];

		return result[1];
	}

	function validatePasswordConfirm(): boolean {
		invalidPasswordConfirm = false;
		invalidPasswordCaptionConfirm = '';

		const result = checkPasswordConfirm(newPassword, passwordConfirm);

		invalidPasswordConfirm = !result[1];
		invalidPasswordCaptionConfirm = result[0];

		return result[1];
	}

	async function resetPassword() {
		loading = true;
		invalidCredentials = false;
		let url = API_URL;

		if (!validateCurrentPassword()) {
			return;
		}

		if (!validateNewPassword()) {
			return;
		}

		if (!validatePasswordConfirm()) {
			return;
		}

		if (wantsAdmin) {
			url += ADMIN_PASSWORD_RESET;
		}
		if (wantsInterviewer) {
			url += INTERVIEWER_PASSWORD_RESET;
		}

		const credentials: PasswordReset = {
			currentPassword: currentPassword,
			newPassword: newPassword
		};

		const response = await fetchNoRefresh(url, {
			method: 'POST',
			body: JSON.stringify(credentials)
		});

		if (response.ok) {
			sendSuccess('Password updated', 'You can now log in using that password');
			goto('/login');
			return;
		}

		handleResponse(response.status, true);

		if (response.status == 401) {
			invalidCredentials = true;
		}

		loading = false;
	}
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
				<h3>Update your password</h3>
			</div>

			<br />

			<div class="subtitle">
				<p>
					You need to update your password because it is the first time you log in or you have
					requested a password reset to your administrator.
				</p>
			</div>

			{#if invalidCredentials}
				<InlineNotification title="Error" subtitle="Check your credentials" />
			{/if}

			{#if loading}
				<InlineLoading description="submitting..." />
			{/if}

			<RadioButtonGroup selected="interviewer" required>
				<div slot="legendText" style="display:flex;margin-top:20px">
					Account Type
					<Tooltip>
						<p>Account type is determined by your administrator</p>
					</Tooltip>
				</div>
				<RadioButton
					labelText="Interviewer"
					value="interviewer"
					bind:checked={wantsInterviewer}
					id="interviewer"
					name="interviewer"
					autofocus
				/>
				<RadioButton
					labelText="Administrator"
					value="admin"
					bind:checked={wantsAdmin}
					id="administrator"
					name="administrator"
				/>
			</RadioButtonGroup>

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
					on:blur={validateNewPassword}
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
					<Button type="submit" disabled={loading}>Submit</Button>
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
		min-height: 30vh;
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

	.subtitle {
		width: 350px;
		text-align: left;
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
