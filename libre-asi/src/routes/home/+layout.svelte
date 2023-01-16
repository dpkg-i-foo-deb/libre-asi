<script lang="ts">
	import { onDestroy, onMount } from 'svelte';
	import { session } from '$lib/stores/userStore';
	import type { Unsubscriber } from 'svelte/store';
	import { goto } from '$app/navigation';

	let activeSession = false;
	let subscription: Unsubscriber;

	onMount(async () => {
		subscription = session.subscribe((value) => {
			if (value == 'true') {
				activeSession = true;
			} else {
				activeSession = false;
				session.set('false');
				goto('login');
			}
		});
	});

    onDestroy(async () => {
        subscription
    })
</script>

{#if activeSession}
	<slot />
{/if}
