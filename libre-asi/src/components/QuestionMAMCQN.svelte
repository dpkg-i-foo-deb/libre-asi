<script lang="ts">
	import type { Answer } from '$lib/models/Answer';
	import type QuestionOption from '$lib/models/Option';
	import type { Question } from '$lib/models/Question';
	import { Checkbox, TextInput } from 'carbon-components-svelte';
	import { onMount } from 'svelte';

	export let question: Question = { answers: [], valid: false };

	let statement = '';

	let comment = '';

	onMount(function () {
		statement = question.special_id ?? '';
		statement += ': ' + question.statement ?? '';

		question.answers = [];
	});

	function addAnswer(option: QuestionOption) {
		let isNew = true;

		const answer: Answer = {
			comment: comment,
			optionID: option.id,
			questionID: question.id,
			value: option.value
		};

		question.answers.forEach(function (value: Answer, index: number) {
			if (value.optionID == option.id) {
				isNew = false;
				question.answers[index] = answer;
				return;
			}
		});

		if (isNew) {
			question.answers.push(answer);
		}

		console.log(question.answers);

		question.valid = true;
	}

	function removeAnswer(option: QuestionOption) {
		question.answers.forEach(function (value: Answer, index: number) {
			if (value.optionID == option.id) {
				question.answers.splice(index, 1);
				return;
			}
		});
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

	<div class="checkbox-container">
		{#each question.options ?? [] as option}
			<div class="checkbox">
				<Checkbox
					labelText={option.description}
					value={option.value}
					on:check={function () {
						if (option.checked) {
							addAnswer(option);
						}
						if (!option.checked) {
							removeAnswer(option);
						}
					}}
					bind:checked={option.checked}
				/>
			</div>
		{/each}
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

	.checkbox-container {
		display: flex;
		flex-direction: column;
		align-items: flex-start;
	}

	.checkbox {
		margin-top: 0.7rem;
		margin-bottom: 0.7rem;
	}

	.comment-container {
		margin-top: 2rem;
		margin-bottom: 2rem;
	}
</style>
