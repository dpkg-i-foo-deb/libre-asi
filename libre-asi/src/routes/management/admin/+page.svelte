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

	export let data: PageData;

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
	});
</script>

<DataTable title="Administrators" description="Current Registered Administrators">
	<Toolbar>
		<ToolbarContent>
			<ToolbarSearch />
			<Button>Register Administrator</Button>
		</ToolbarContent>
	</Toolbar>
</DataTable>
