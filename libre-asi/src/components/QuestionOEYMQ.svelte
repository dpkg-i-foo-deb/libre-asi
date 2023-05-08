<script lang="ts">
	import type { Question } from '$lib/models/Question';
	import { NumberInput } from 'carbon-components-svelte';
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

	{#each question.options ?? [] as option}
		<div class="number-input-container">
			<NumberInput label={option.description?.toUpperCase()} value={0} />
		</div>
	{/each}
</main>

<style>
	.title {
		margin-top: 1rem;
		margin-bottom: 1.5rem;
	}

	.number-input-container {
		margin-top: 2rem;
		margin-bottom: 2rem;
	}
</style>
