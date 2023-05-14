<script lang="ts">
	import { TagSkeleton, TextInputSkeleton } from 'carbon-components-svelte';
</script>

<main>
	<TagSkeleton />

	<div class="skeleton">
		<TextInputSkeleton />
	</div>

	<div class="skeleton">
		<TextInputSkeleton />
	</div>

	<div class="skeleton">
		<TextInputSkeleton />
	</div>
</main>

<style>
	.skeleton {
		margin-top: 1rem;
		margin-bottom: 1rem;
	}
</style>
