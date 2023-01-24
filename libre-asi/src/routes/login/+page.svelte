<script lang="ts">
	import { applyAction, enhance } from '$app/forms';
	import {
		TextInput,
		PasswordInput,
		Button,
		ToastNotification,
		RadioButtonGroup,
		RadioButton,
		Tooltip
	} from 'carbon-components-svelte';
	import type { ActionData } from './$types';
	import session from '$lib/stores/userStore';
	import type { ActionResult } from '@sveltejs/kit';
	import { notifications } from '$lib/stores/notificationStore';
	import Layout from '../+layout.svelte';

	export let form: ActionData;

	let wantsAdmin = false;
	let wantsInterviewer = false;

	//const emailRegex = /^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}$/i;

	function handleLogin(result: ActionResult<Record<string, any>, Record<string, any>>) {
		if (result.type == 'failure') {
			console.log('owo');
			$notifications.visible = true;
			if (form?.badRequest ?? false) {
			}
		}
	}
</script>

<main>
	<div class="container">
		<form
			use:enhance={() => {
				return async ({ result }) => {
					handleLogin(result);
					if (result.type === 'redirect') {
						$session.active = true;
					}
					await applyAction(result);
				};
			}}
			action="?/login"
			method="POST"
		>
			<h3>Please Log In to your Account</h3>

			<div class="form-element">
				<TextInput
					labelText="Email"
					placeholder="Enter email..."
					required
					invalidText="Check your email"
					id="email"
					name="email"
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

			{#if form?.invalidCredentials}
				<div class="invalid-credentials">
					<p>Check your credentials</p>
				</div>
			{/if}

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

	.error-notification {
		position: fixed;
		bottom: 0;
		right: 0;
		margin: 10px;
	}

	.invalid-credentials {
		padding-top: 5px;
		padding-bottom: 5px;
		color: red;
	}
</style>
