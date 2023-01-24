<script lang="ts">
	import { ToastNotification } from 'carbon-components-svelte';
	import { notifications } from '$lib/stores/notificationStore';

	export let title = $notifications.title;
	export let subtitle = $notifications.subtitle;
	export let caption = $notifications.caption;
	export let kind = $notifications.kind;
	export let timeout = $notifications.timeout;
	export function toggleVisibility() {
		$notifications.visible = !$notifications.visible;
	}

	function hide() {
		$notifications.visible = false;
	}
</script>

<main>
	{#if $notifications.visible}
		<div>
			<ToastNotification {title} {subtitle} {caption} {timeout} {kind} on:close={hide} />
		</div>
	{/if}
</main>
