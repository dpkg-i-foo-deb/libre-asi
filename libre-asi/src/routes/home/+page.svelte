<script lang="ts">
	import { SessionRole } from '$lib/models/Session';
	import session from '$lib/stores/userStore';
</script>

<h1>Home</h1>

{#if $session.role == SessionRole.Admin}
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
