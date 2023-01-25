<script lang="ts">
	import { onMount } from 'svelte';
	import session from '$lib/stores/userStore';
	import { goto } from '$app/navigation';

	onMount(async () => {
		if ($session.active == true) {
			goto('/');
		}
	});
</script>

{#if !$session.active}
	<slot />
{/if}
