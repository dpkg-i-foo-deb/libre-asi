<script lang="ts">
	import { onDestroy, onMount } from 'svelte';
	import { session } from '$lib/stores/userStore';
	import { goto } from '$app/navigation';
	import type { Unsubscriber } from 'svelte/store';

	let activeSession = false;
	let subscription: Unsubscriber;

	onMount(async () => {
		subscription = session.subscribe((value) => {
			if (value == 'true') {
				activeSession = true;
				goto('/home');
			} else {
				activeSession = false;
			}
		});
	});

	onDestroy(async () => {
		subscription;
	});
</script>

{#if !activeSession}
	<slot />
{/if}
