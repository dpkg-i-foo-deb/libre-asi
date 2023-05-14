<script lang="ts">
	import type { Question } from '$lib/models/Question';
	import { Checkbox } from 'carbon-components-svelte';
	import { onMount } from 'svelte';

	export let question: Question;

	let statement = '';

	onMount(function () {
		statement = question.special_id ?? '';
		statement += ': ' + question.statement ?? '';
	});
</script>

<main>
	<div class="title">
		<h4>
			{statement}
		</h4>
	</div>

	<div class="checkbox-container">
		{#each question.options ?? [] as option}
			<div class="checkbox">
				<Checkbox labelText={option.description} value={option.value} />
			</div>
		{/each}
	</div>
</main>

<style>
	.title {
		margin-top: 1rem;
		margin-bottom: 1.5rem;
	}

	.checkbox-container {
		display: flex;
		flex-direction: column;
		align-items: flex-start;
	}

	.checkbox {
		margin-top: 0.7rem;
		margin-bottom: 0.7rem;
	}
</style>
