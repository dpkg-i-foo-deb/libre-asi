<script lang="ts">
	import { ToastNotification } from 'carbon-components-svelte';
	import { notifications } from '$lib/stores/notificationStore';

	function hide() {
		$notifications.visible = false;
	}
</script>

<main>
	{#if $notifications.visible}
		<div>
			<ToastNotification
				bind:title={$notifications.title}
				bind:kind={$notifications.kind}
				bind:caption={$notifications.caption}
				bind:timeout={$notifications.timeout}
				bind:subtitle={$notifications.subtitle}
				on:close={hide}
			/>
		</div>
	{/if}
</main>
