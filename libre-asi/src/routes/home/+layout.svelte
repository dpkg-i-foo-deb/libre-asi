<script lang="ts">
	import { onDestroy, onMount } from 'svelte';
	import session from '$lib/stores/userStore';
	import type { Unsubscriber } from 'svelte/store';
	import { goto } from '$app/navigation';
</script>

{#if $session.active}
	<slot />
{/if}
