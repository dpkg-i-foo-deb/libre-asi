<script lang="ts">
	import session from '$lib/stores/userStore';
</script>

{#if $session.active}
	<slot />
{/if}
