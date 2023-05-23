<script lang="ts">
	import type { Answer } from '$lib/models/Answer';
	import type QuestionOption from '$lib/models/Option';
	import type { Question } from '$lib/models/Question';
	import { RadioButton, RadioButtonGroup, TextInput } from 'carbon-components-svelte';
	import { onMount } from 'svelte';

	export let question: Question;

	let statement = '';

	let comment = '';

	onMount(function () {
		statement = question.special_id ?? '';
		statement += ': ' + question.statement ?? '';
	});

	function update(option: QuestionOption) {
		question.answers = [
			{
				comment: comment,
				optionID: option.id,
				questionID: question.id,
				value: option.value
			}
		];

		question.valid = true;
	}

	function setComment() {
		question.answers?.forEach(function (value: Answer) {
			value.comment = comment;
		});
	}
</script>

<main>
	<div class="title">
		<h4>
			{statement}
		</h4>
	</div>

	<div class="bq-container">
		<RadioButtonGroup>
			<div class="radio-button-container">
				{#each question.options ?? [] as option}
					<div class="radio-button">
						<RadioButton
							labelText={option.description}
							value={option.value?.toString()}
							on:change={function () {
								update(option);
							}}
						/>
					</div>
				{/each}
			</div>
		</RadioButtonGroup>
	</div>

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

	.radio-button-container {
		display: flex;
		flex-direction: column;
		align-items: flex-start;
	}

	.radio-button {
		margin-top: 0.7rem;
		margin-bottom: 0.7rem;
	}

	.comment-container {
		margin-top: 2rem;
		margin-bottom: 2rem;
	}
</style>
