<script lang="ts">
	import session from '$lib/stores/userStore';
	import { SessionRole } from '$lib/models/Session';
</script>

{#if $session.active && $session.role == SessionRole.Admin}
	<slot />
{/if}
