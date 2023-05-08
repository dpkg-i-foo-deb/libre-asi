<script lang="ts">
	import type { Question } from '$lib/models/Question';
	import { RadioButton, RadioButtonGroup } from 'carbon-components-svelte';
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

	<RadioButtonGroup>
		<div class="radio-button-container">
			{#each question.options ?? [] as option}
				<div class="radio-button">
					<RadioButton labelText={option.description} value={option.value?.toString()} />
				</div>
			{/each}
		</div>
	</RadioButtonGroup>
</main>

<style>
	.title {
		margin-top: 1rem;
		margin-bottom: 1.5rem;
	}

	.radio-button-container {
		display: flex;
		flex-direction: column;
		align-items: flex-start;
	}

	.radio-button {
		margin-top: 0.7rem;
		margin-bottom: 0.7rem;
	}
</style>
