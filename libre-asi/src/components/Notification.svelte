<script lang="ts">
	import { ToastNotification } from 'carbon-components-svelte';

	export let title = 'Generic Notification';
	export let subtitle = 'Subtitle';
	export let caption = 'Caption';
	export let kind: undefined | 'info' | 'warning' | 'error' | 'success';
	export let timeout = 5000;
</script>

<main>
	<div>
		<ToastNotification {title} {subtitle} {caption} {timeout} {kind} />
	</div>
</main>
