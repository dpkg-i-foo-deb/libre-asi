<script lang="ts">
	import { onMount } from 'svelte';
	import { session, role } from '$lib/stores/userStore';
	import { goto } from '$app/navigation';

	let activeSession = false;

	onMount(async () => {
		if ($session == 'false' || $role != 'admin') {
			activeSession = false;
			goto('/login');
		} else {
			activeSession = true;
		}
	});
</script>

{#if activeSession}
	<slot />
{/if}
