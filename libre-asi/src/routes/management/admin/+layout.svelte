<script lang="ts">
	import session from '$lib/stores/userStore';
	import { SessionRole } from '$lib/models/Session';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';

	onMount(function () {
		if (!$session.active) {
			//goto('/login');
		}
	});
</script>

	<slot />
