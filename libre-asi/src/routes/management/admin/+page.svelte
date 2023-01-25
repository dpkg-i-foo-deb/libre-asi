<script lang="ts">
	import {
		Button,
		DataTable,
		Toolbar,
		ToolbarContent,
		ToolbarSearch
	} from 'carbon-components-svelte';
</script>

<DataTable title="Administrators" description="Current Registered Administrators">
	<Toolbar>
		<ToolbarContent>
			<ToolbarSearch />
			<Button>Register Administrator</Button>
		</ToolbarContent>
	</Toolbar>
</DataTable>
