<script lang="ts">
	import { role } from '$lib/stores/userStore';
</script>

<h1>Home</h1>

{#if $role == 'admin'}
	<h4>
		Use the menu on the top left corner to access functions like the management of interviewer
		profiles or creating another administrator account
	</h4>
{/if}

<style>
	h4 {
		padding-top: 20px;
	}
</style>
