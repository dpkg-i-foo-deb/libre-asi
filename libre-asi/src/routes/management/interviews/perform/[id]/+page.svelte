<script lang="ts">
	import {
		API_URL,
		GET_QUESTION,
		NEXT_QUESTION,
		GET_INTERVIEW,
		ANSWER_QUESTION
	} from '$lib/api/constants';
	import type Interview from '$lib/models/Interview';
	import type { Question } from '$lib/models/Question';
	import { fetchWithRefresh } from '$lib/util/fetchRefresh';
	import { handleResponse } from '$lib/util/handleResponse';
	import { Tile, Button } from 'carbon-components-svelte';
	import QuestionBq from '../../../../../components/QuestionBQ.svelte';
	import QuestionDnq from '../../../../../components/QuestionDNQ.svelte';
	import QuestionMamcn from '../../../../../components/QuestionMAMCN.svelte';
	import QuestionMamcqn from '../../../../../components/QuestionMAMCQN.svelte';
	import QuestionMnq from '../../../../../components/QuestionMNQ.svelte';
	import QuestionOeymq from '../../../../../components/QuestionOEYMQ.svelte';
	import QuestionSamcqn from '../../../../../components/QuestionSAMCQN.svelte';
	import QuestionSaq from '../../../../../components/QuestionSAQ.svelte';
	import QuestionSma from '../../../../../components/QuestionSMA.svelte';
	import QuestionSkeleton from '../../../../../components/QuestionSkeleton.svelte';
	import QuestionTnq from '../../../../../components/QuestionTNQ.svelte';
	import QuestionWnq from '../../../../../components/QuestionWNQ.svelte';
	import QuestionTq from '../../../../../components/QuestionTQ.svelte';
	import QuestionYnq from '../../../../../components/QuestionYNQ.svelte';
	import { onMount } from 'svelte';
	import { page } from '$app/stores';

	const interviewID = $page.params.id;

	let interview: Interview = {};

	let currentQuestion: Question = { valid: false, answers: [] };

	onMount(async function () {
		await loadInterview();
	});

	async function loadInterview() {
		const response = await fetchWithRefresh(API_URL + GET_INTERVIEW + interviewID, {
			method: 'GET'
		});

		if (response.ok) {
			interview = (await response.json()) as Interview;

			await getQuestion();
		}

		handleResponse(response.status, false);
	}

	async function handleNext() {
		answerQuestion();
	}

	async function answerQuestion() {
		const response = await fetchWithRefresh(API_URL + ANSWER_QUESTION + interviewID, {
			method: 'POST',
			body: JSON.stringify(currentQuestion.answers)
		});

		if (response.ok) {
			currentQuestion.answers = [];
			currentQuestion.valid = false;

			await nextQuestion();
		}

		handleResponse(response.status, false);
	}

	async function getQuestion() {
		currentQuestion.type = 'NIL';

		const questionResponse = await fetchWithRefresh(
			API_URL + GET_QUESTION + interview.currentQuestion,
			{
				method: 'GET'
			}
		);

		if (questionResponse.ok) {
			currentQuestion = (await questionResponse.json()) as Question;
		}

		handleResponse(questionResponse.status, false);
	}

	async function nextQuestion() {
		const nextQuestionResponse = await fetchWithRefresh(API_URL + NEXT_QUESTION, {
			method: 'POST',
			body: JSON.stringify(interview)
		});

		if (nextQuestionResponse.ok) {
			interview = (await nextQuestionResponse.json()) as Interview;

			await getQuestion();
		}

		handleResponse(nextQuestionResponse.status, false);
	}
</script>

<main>
	<Tile>
		<div>
			<div class="title">
				{#if currentQuestion?.category == 'INF'}
					<h2>Preguntas generales</h2>
				{/if}

				{#if currentQuestion?.category == 'AL' ?? false}
					<h2>Alojamiento</h2>
				{/if}

				{#if currentQuestion?.category == 'MED' ?? false}
					<h2>Estado médico</h2>
				{/if}

				{#if currentQuestion?.category == 'EMP' ?? false}
					<h2>Empleo/Sustento</h2>
				{/if}

				{#if currentQuestion?.category == 'DRU' ?? false}
					<h2>Drogas/Alcohol</h2>
				{/if}

				{#if currentQuestion?.category == 'LAW' ?? false}
					<h2>Estado legal</h2>
				{/if}

				{#if currentQuestion?.category == 'FAM' ?? false}
					<h2>Familia/Relaciones</h2>
				{/if}

				{#if currentQuestion?.category == 'PSY'}
					<h2>Estado Mental</h2>
				{/if}

				{#if currentQuestion?.category == 'VAL'}
					<h2>Valoración</h2>
				{/if}
			</div>
			{#if currentQuestion}
				{#if currentQuestion.type == 'SAMCQN' ?? ''}
					<QuestionSamcqn bind:question={currentQuestion} />
				{/if}

				{#if currentQuestion.type == 'OEYMQ' ?? ''}
					<QuestionOeymq bind:question={currentQuestion} />
				{/if}

				{#if currentQuestion.type == 'MAMCN' ?? ''}
					<QuestionMamcn bind:question={currentQuestion} />
				{/if}

				{#if currentQuestion.type == 'BQ' ?? ''}
					<QuestionBq bind:question={currentQuestion} />
				{/if}

				{#if currentQuestion.type == 'MAMCQN' ?? ''}
					<QuestionMamcqn bind:question={currentQuestion} />
				{/if}

				{#if currentQuestion.type == 'TQ' ?? ''}
					<QuestionTq bind:question={currentQuestion} />
				{/if}

				{#if currentQuestion.type == 'DNQ' ?? ''}
					<QuestionDnq bind:question={currentQuestion} />
				{/if}

				{#if currentQuestion.type == 'TNQ' ?? ''}
					<QuestionTnq bind:question={currentQuestion} />
				{/if}

				{#if currentQuestion.type == 'MNQ' ?? ''}
					<QuestionMnq bind:question={currentQuestion} />
				{/if}

				{#if currentQuestion.type == 'WNQ' ?? ''}
					<QuestionWnq bind:question={currentQuestion} />
				{/if}

				{#if currentQuestion.type == 'SMA' ?? ''}
					<QuestionSma bind:question={currentQuestion} />
				{/if}

				{#if currentQuestion.type == 'SAQ' ?? ''}
					<QuestionSaq bind:question={currentQuestion} />
				{/if}

				{#if currentQuestion.type == 'YNQ' ?? ''}
					<QuestionYnq bind:question={currentQuestion} />
				{/if}

				{#if currentQuestion.type == 'NIL' ?? ''}
					<QuestionSkeleton />
				{/if}
			{/if}
		</div>

		<div class="button-container">
			<Button size="default" kind="secondary">Volver</Button>

			<Button
				size="default"
				disabled={!currentQuestion.valid}
				on:click={function () {
					handleNext();
				}}>Siguiente</Button
			>
		</div>
	</Tile>
</main>

<style>
	.title {
		margin-bottom: 2rem;
	}

	.button-container {
		display: flex;
		flex-direction: row;
		margin-top: 2rem;
	}
</style>
