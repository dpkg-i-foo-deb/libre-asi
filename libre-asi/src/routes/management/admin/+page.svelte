<script lang="ts">
	import {
		Button,
		DataTable,
		Toolbar,
		ToolbarContent,
		ToolbarSearch
	} from 'carbon-components-svelte';
	import { onMount } from 'svelte';
	import type { PageData } from './$types';
	import session from '$lib/stores/userStore';
	import { notifications } from '$lib/stores/notificationStore';
	import { SessionRole } from '$lib/models/Session';
	import { goto } from '$app/navigation';
	import type Administrator from '$lib/models/Administrator';
	import type { DataTableRow } from 'carbon-components-svelte/types/DataTable/DataTable.svelte';

	export let data: PageData;

	let rows: ReadonlyArray<DataTableRow>;

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
			<Button>Register Administrator</Button>
		</ToolbarContent>
	</Toolbar>
</DataTable>
