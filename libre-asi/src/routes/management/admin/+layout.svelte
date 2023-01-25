<script lang="ts">
	import { onMount } from 'svelte';
	import session from '$lib/stores/userStore';
	import { goto } from '$app/navigation';
	import { SessionRole } from '$lib/models/Session';
</script>

{#if $session.active && $session.role == SessionRole.Admin}
	<slot />
{/if}
