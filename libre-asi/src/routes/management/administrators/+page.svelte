<script lang="ts">
	import {
		Button,
		DataTable,
		Toolbar,
		ToolbarContent,
		ToolbarMenu,
		ToolbarMenuItem,
		ToolbarSearch
	} from 'carbon-components-svelte';
	import { onMount } from 'svelte';
	import type administrator from '$lib/models/Administrator';
	import { apiUrl, getAdmins } from '$lib/api/constants';

	let administrators: administrator[];
	let response: Response;

	onMount(async function () {
		try {
			response = await fetch(apiUrl + getAdmins, {
				method: 'GET',
				mode: 'cors',
				credentials: 'include'
			});
		} catch (e) {
			console.log(e);
			return;
		}

		if (response.ok) {
			administrators = await response.json();
			console.log(administrators);
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
