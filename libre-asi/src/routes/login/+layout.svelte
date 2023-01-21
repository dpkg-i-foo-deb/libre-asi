<script lang="ts">
	import { onMount } from 'svelte';
	import session  from '$lib/stores/userStore';
	import { goto } from '$app/navigation';

	let activeSession = false;

	onMount(async () => {
		if ($session.active== true) {
			activeSession = true;
			goto('/home');
		} else {
			activeSession = false;
		}
	});
</script>

{#if !activeSession}
	<slot />
{/if}
