<script lang="ts">
	import LL from '$lib/i18n/i18n-svelte';
	import {
		TextInput,
		PasswordInput,
		Button,
		RadioButtonGroup,
		RadioButton,
		Tooltip,
		InlineNotification,
		Form,
		ProgressIndicator,
		ProgressStep,
		ButtonSet,
		InlineLoading
	} from 'carbon-components-svelte';
	import session from '$lib/stores/userStore';
	import { SessionRole } from '$lib/models/Session';
	import { onMount } from 'svelte';
	import type User from '$lib/models/User';
	import { goto } from '$app/navigation';
	import { sendSuccess } from '$lib/util/notifications';
	import { handleResponse } from '$lib/util/handleResponse';
	import { checkEmail, checkPassword } from '$lib/util/formUtils';
	import { ADMIN_LOGIN, API_URL, INTERVIEWER_LOGIN } from '$lib/api/constants';
	import { fetchNoRefresh } from '$lib/util/fetch';

	let loading = false;

	let stepIndex = 0;

	let wantsAdmin = false;
	let wantsInterviewer = false;
	let invalidEmail = false;
	let invalidPassword = false;
	let invalidEmailCaption = '';
	let invalidPasswordCaption = '';
	let email = '';
	let password = '';
	let invalidCredentials = false;

	//If the user lands here, it means they need a session
	onMount(function () {
		$session.active = false;
		$session.role = SessionRole.None;
	});

	function validateEmail(): boolean {
		invalidEmailCaption = '';
		invalidEmail = true;

		const result = checkEmail(email);

		invalidEmail = !result[1];
		invalidEmailCaption = result[0];

		return result[1];
	}

	function validatePassword(): boolean {
		invalidPasswordCaption = '';
		invalidPassword = true;

		const result = checkPassword(password);

		invalidPassword = !result[1];
		invalidPasswordCaption = result[0];

		return result[1];
	}

	function back() {
		if (stepIndex > 0) {
			stepIndex--;
		}
	}

	function hadnleNext() {
		switch (stepIndex) {
			case 1:
				if (!validateEmail()) {
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
			if (!validateEmail() || !validatePassword()) {
				return;
			}

			login();
		}
	}

	async function login() {
		loading = true;

		invalidCredentials = false;
		let url = '';

		if (wantsAdmin) {
			url += API_URL + ADMIN_LOGIN;
		}

		if (wantsInterviewer) {
			url += API_URL + INTERVIEWER_LOGIN;
		}

		if (!(validateEmail() || validatePassword())) {
			return;
		}

		const user: User = { email: email, password: password };

		const response = await fetchNoRefresh(url, {
			method: 'POST',
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
				hadnleNext();
			}}
		>
			<div class="title">
				<h3>{$LL.login.TITLE()}</h3>
			</div>

			{#if invalidCredentials}
				<div class="invalid-credentials">
					<InlineNotification
						title={$LL.login.ERROR()}
						subtitle={$LL.login.INVALID_CREDENTIALS()}
					/>
				</div>
			{/if}

			<div class="stepper">
				<ProgressIndicator bind:currentIndex={stepIndex} spaceEqually preventChangeOnClick>
					<ProgressStep complete={stepIndex > 0} label={$LL.login.ROLE()} />
					<ProgressStep
						complete={stepIndex > 1}
						label="Identification"
						bind:invalid={invalidEmail}
					/>
					<ProgressStep
						complete={stepIndex > 2}
						label="Credentials"
						bind:invalid={invalidPassword}
					/>
				</ProgressIndicator>
			</div>

			{#if loading}
				<InlineLoading description={$LL.login.SUBMITTING()} />
			{/if}

			{#if stepIndex == 0}
				<div class="form-element">
					<RadioButtonGroup selected="interviewer" required>
						<div slot="legendText" style="display:flex;margin-top:20px">
							{$LL.login.ACCOUNT_TYPE()}
							<Tooltip>
								<p>{$LL.login.ACCOUNT_TOOLTIP()}</p>
							</Tooltip>
						</div>
						<RadioButton
							labelText={$LL.login.INTERVIEWER()}
							value="interviewer"
							bind:checked={wantsInterviewer}
							id="interviewer"
							name="interviewer"
							autofocus
						/>
						<RadioButton
							labelText={$LL.login.ADMINISTRATOR()}
							value="admin"
							bind:checked={wantsAdmin}
							id="administrator"
							name="administrator"
						/>
					</RadioButtonGroup>
				</div>
			{/if}

			{#if stepIndex == 1}
				<div class="form-element">
					<TextInput
						labelText={$LL.login.EMAIL()}
						placeholder={$LL.login.EMAIL_PLACEHOLDER()}
						id="email"
						name="email"
						bind:value={email}
						bind:invalid={invalidEmail}
						bind:invalidText={invalidEmailCaption}
						on:blur={validateEmail}
						autofocus
					/>
				</div>
			{/if}

			{#if stepIndex == 2}
				<div class="form-element">
					<PasswordInput
						labelText={$LL.login.PASSWORD()}
						placeholder={$LL.login.PASSWORD_PLACEHOLDER()}
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
			{/if}

			<ButtonSet>
				<div class="button-container">
					<Button kind="secondary" disabled={stepIndex == 0 || loading} on:click={back}
						>{$LL.login.BACK()}</Button
					>
					<Button type="submit" disabled={loading}>{$LL.login.NEXT()}</Button>
				</div>
			</ButtonSet>
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
		margin-top: 15px;
		margin-bottom: 15px;
		height: 50px;
	}

	.button-container {
		padding-top: 10px;
		float: left;
	}

	.invalid-credentials {
		padding-bottom: 10px;
	}

	.stepper {
		margin-bottom: 30px;
	}

	.button-container {
		display: flex;
		margin-top: 25px;
		flex-direction: row;
	}
</style>
