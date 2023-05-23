<script lang="ts">
	import type { Answer } from '$lib/models/Answer';
	import type QuestionOption from '$lib/models/Option';
	import type { Question } from '$lib/models/Question';
	import { NumberInput, TextInput } from 'carbon-components-svelte';
	import { onMount } from 'svelte';

	export let question: Question;

	let statement = '';

	let comment = '';

	onMount(function () {
		if (question.options == null) {
			question.options = [{}];
		}

		question.answers = [{ value: 0, questionID: question.id, optionID: question.options[0].id }];

		question.options?.forEach(function (value: QuestionOption) {
			value.value = 0;
		});

		question.valid = true;

		statement = question.special_id ?? '';
		statement += ': ' + question.statement ?? '';
	});

	function setComment() {
		question.answers.forEach(function (value: Answer) {
			value.comment = comment;
		});
	}

	function update(option: QuestionOption) {
		question.answers = [
			{
				comment: comment,
				optionID: option.id,
				questionID: question.id,
				value: option.value
			}
		];
	}
</script>

<main>
	<div class="title">
		<h4>
			{statement}
		</h4>
	</div>

	{#each question.options ?? [] as option}
		<div class="number-input-container">
			<NumberInput
				label={option.description?.toUpperCase()}
				bind:value={option.value}
				on:change={function () {
					update(option);
				}}
			/>
		</div>
	{/each}

	<div class="comment-container">
		<TextInput
			labelText="Comentarios (opcional)"
			placeholder="Ingrese comentarios"
			on:input={setComment}
			bind:value={comment}
		/>
	</div>
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

	.comment-container {
		margin-top: 2rem;
		margin-bottom: 2rem;
	}
</style>
