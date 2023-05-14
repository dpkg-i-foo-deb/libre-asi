<script lang="ts">
	import {
		API_URL,
		GET_PATIENTS,
		GET_QUESTION,
		NEXT_QUESTION,
		START_INTERVIEW
	} from '$lib/api/constants';
	import type Interview from '$lib/models/Interview';
	import type Patient from '$lib/models/Patient';
	import type { Question } from '$lib/models/Question';
	import { fetchWithRefresh } from '$lib/util/fetchRefresh';
	import { handleResponse } from '$lib/util/handleResponse';
	import { Button, DataTable, DataTableSkeleton, Tile } from 'carbon-components-svelte';
	import type { DataTableRow } from 'carbon-components-svelte/types/DataTable/DataTable.svelte';
	import { onMount } from 'svelte';
	import QuestionSamcqn from '../../../../components/QuestionSAMCQN.svelte';
	import QuestionOeymq from '../../../../components/QuestionOEYMQ.svelte';
	import QuestionMamcn from '../../../../components/QuestionMAMCN.svelte';
	import QuestionBq from '../../../../components/QuestionBQ.svelte';
	import QuestionMamcqn from '../../../../components/QuestionMAMCQN.svelte';
	import QuestionSkeleton from '../../../../components/QuestionSkeleton.svelte';
	import QuestionTq from '../../../../components/QuestionTQ.svelte';
	import QuestionDnq from '../../../../components/QuestionDNQ.svelte';
	import QuestionTnq from '../../../../components/QuestionTNQ.svelte';

	let rows: ReadonlyArray<DataTableRow>;
	let filteredRows: ReadonlyArray<DataTableRow>;

	let selectedRowIds: number[] = [];

	let isSelectingPatient = true;

	let newInterview: Interview = {};

	let currentQuestion: Question = {};

	onMount(async function () {
		await loadPatients();
	});

	async function loadPatients() {
		const response = await fetchWithRefresh(API_URL + GET_PATIENTS, {
			method: 'GET'
		});

		if (response.ok) {
			const patients = (await response.json()) as Patient[];

			rows = patients.map(function (value) {
				return {
					id: value.ID,
					name: value.firstName ?? '' + value.firstSurname ?? '',
					personalId: value.personalID
				};
			});

			filteredRows = rows;
		}

		handleResponse(response.status, false);
	}

	async function handleNext() {
		if (isSelectingPatient) {
			newInterview.patientID = selectedRowIds[0];

			const response = await fetchWithRefresh(API_URL + START_INTERVIEW, {
				method: 'POST',
				body: JSON.stringify(newInterview)
			});

			if (response.ok) {
				newInterview = (await response.json()) as Interview;

				isSelectingPatient = false;

				await getQuestion();
			}

			handleResponse(response.status, false);
		} else {
			await nextQuestion();
		}
	}

	async function getQuestion() {
		currentQuestion.type = 'NIL';

		const questionResponse = await fetchWithRefresh(
			API_URL + GET_QUESTION + newInterview.currentQuestion,
			{
				method: 'GET'
			}
		);

		if (questionResponse.ok) {
			currentQuestion = (await questionResponse.json()) as Question;
		}

		console.log(currentQuestion);

		handleResponse(questionResponse.status, false);
	}

	async function nextQuestion() {
		const nextQuestionResponse = await fetchWithRefresh(API_URL + NEXT_QUESTION, {
			method: 'POST',
			body: JSON.stringify(newInterview)
		});

		if (nextQuestionResponse.ok) {
			newInterview = (await nextQuestionResponse.json()) as Interview;

			await getQuestion();
		}

		handleResponse(nextQuestionResponse.status, false);
	}
</script>

<main>
	<Tile>
		<div class="content">
			{#if isSelectingPatient}
				<h4 class="title">Realización de entrevista</h4>

				<div class="subtitle">
					<h5>Seleccione un paciente</h5>
				</div>
				<h7>El paciente seleccionado será guardado en el reporte de la entrevista</h7>

				<div class="table">
					{#if rows == undefined}
						<DataTableSkeleton
							headers={[
								{ key: 'name', value: 'Nombre' },
								{ key: 'personalId', value: 'Identificación' }
							]}
							rows={5}
						/>
					{:else}
						<DataTable
							radio
							headers={[
								{ key: 'name', value: 'Nombre' },
								{ key: 'personalId', value: 'Identificación' }
							]}
							bind:rows={filteredRows}
							bind:selectedRowIds
						/>
					{/if}
				</div>
			{/if}

			{#if currentQuestion?.category == 'INF'}
				<div class="title">
					<h2>Preguntas generales</h2>
				</div>
			{/if}

			{#if currentQuestion?.category == 'AL' ?? false}
				<div class="title">
					<h2>Alojamiento</h2>
				</div>
			{/if}

			{#if currentQuestion?.category == 'MED' ?? false}
				<div class="title">
					<h2>Estado médico</h2>
				</div>
			{/if}

			{#if currentQuestion?.category == 'EMP' ?? false}
				<div class="title">
					<h2>Empleo/Sustento</h2>
				</div>
			{/if}
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

				{#if currentQuestion.type == 'NIL' ?? ''}
					<QuestionSkeleton />
				{/if}
			{/if}
		</div>

		<div class="button-container">
			<Button size="default" kind="secondary">Volver</Button>

			<Button
				size="default"
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

	.subtitle {
		margin-bottom: 1rem;
	}

	.table {
		margin-top: 4rem;
		margin-bottom: 4rem;
	}

	.button-container {
		display: flex;
		flex-direction: row;
		margin-top: 2rem;
	}
</style>
