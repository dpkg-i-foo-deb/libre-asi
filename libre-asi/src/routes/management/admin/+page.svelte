<script lang="ts">
	import {
		Button,
		DataTable,
		Toolbar,
		ToolbarContent,
		ToolbarSearch,
		DataTableSkeleton,
		ComposedModal,
		ModalHeader,
		ModalFooter,
		ModalBody,
		TextInput
	} from 'carbon-components-svelte';
	import { onMount } from 'svelte';
	import type { PageData } from './$types';
	import session from '$lib/stores/userStore';
	import { notifications } from '$lib/stores/notificationStore';
	import { SessionRole } from '$lib/models/Session';
	import { goto } from '$app/navigation';
	import type Administrator from '$lib/models/Administrator';
	import type { DataTableRow } from 'carbon-components-svelte/types/DataTable/DataTable.svelte';
	import { enhance } from '$app/forms';
	export let data: PageData;

	let rows: ReadonlyArray<DataTableRow>;
	let isOpen = false;

	onMount(function () {
		if (data.error) {
			$session.active = false;
			$session.role = SessionRole.None;

			$notifications.kind = 'error';
			$notifications.title = 'Your session has expired';
			$notifications.subtitle = 'Log In Again';
			$notifications.caption = new Date().toLocaleString();
			$notifications.visible = true;

			goto('/');
		}

		const existingAdmins = data.administrators ?? [];

		rows = existingAdmins.map(function (value: Administrator) {
			return { id: value.ID, email: value.email, username: value.username };
		});
	});
</script>

{#if rows == undefined}
	<DataTableSkeleton
		headers={[
			{ key: 'email', value: 'Email' },
			{ key: 'username', value: 'Username' }
		]}
		rows={5}
	/>
{:else}
	<DataTable
		title="Administrators"
		description="Current Registered Administrators"
		headers={[
			{ key: 'email', value: 'Email' },
			{ key: 'username', value: 'Username' }
		]}
		{rows}
	>
		<Toolbar>
			<ToolbarContent>
				<ToolbarSearch />
				<Button
					on:click={() => {
						isOpen = true;
					}}>Register Administrator</Button
				>
			</ToolbarContent>
		</Toolbar>
	</DataTable>

	<ComposedModal bind:open={isOpen} selectorPrimaryFocus="#email">
		<ModalHeader label="Transaction" title="New Administrator Account" />
		<ModalBody hasForm>
			<form>
				<h7 class="paragraph">
					Administrators can register other administrators, they can also trigger password resets,
					access interviews and accounts information.
				</h7>
				<br />
				<h6 class="paragraph">
					The account password will be generated automatically and the new user will be prompted to
					change it
				</h6>
				<div class="input-field">
					<TextInput id="email" labelText="Email" placeholder="Enter email..." />
				</div>
				<div class="input-field">
					<TextInput labelText="User name" placeholder="Enter user name..." />
				</div>
			</form>
		</ModalBody>
		<ModalFooter
			primaryButtonText="Register"
			secondaryButtonText="Cancel"
			on:click:button--secondary={() => {
				isOpen = false;
			}}
		/>
	</ComposedModal>
{/if}

<style>
	.input-field {
		padding-top: 10px;
		padding-bottom: 10px;
	}

	.paragraph {
		padding-top: 5px;
		padding-bottom: 5px;
	}
</style>
