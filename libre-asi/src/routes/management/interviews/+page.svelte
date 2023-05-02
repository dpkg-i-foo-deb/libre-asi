<script lang="ts">
	import { API_URL, GET_INTERVIEWERS, GET_INTERVIEWS, GET_PATIENT } from '$lib/api/constants';
	import type { Interview } from '$lib/models/Interview';
	import type Interviewer from '$lib/models/Interviewer';
	import type Patient from '$lib/models/Patient';
	import { fetchWithRefresh } from '$lib/util/fetchRefresh';
	import { handleResponse } from '$lib/util/handleResponse';
	import {
		Button,
		DataTableSkeleton,
		DataTable,
		OverflowMenu,
		Toolbar,
		ToolbarContent,
		ToolbarSearch
	} from 'carbon-components-svelte';
	import type { DataTableRow } from 'carbon-components-svelte/types/DataTable/DataTable.svelte';
	import { onMount } from 'svelte';

	let searchValue = '';

	let rows: ReadonlyArray<DataTableRow>;
	let filteredRows: ReadonlyArray<DataTableRow>;

	let newInterview: Interview;

	onMount(async function () {
		newInterview = {
			ID: 0,
			AsiFormID: 0,
			CurrentQuestion: ''
		};

		await loadInterviews();
	});

	async function loadInterviews() {
		const response = await fetchWithRefresh(API_URL + GET_INTERVIEWS, { method: 'GET' });

		if (response.ok) {
			const existingInterviews = (await response.json()) as Interview[];

			rows = await Promise.all(
				existingInterviews.map(async function (value: Interview) {
					const patientResponse = await fetchWithRefresh(API_URL + GET_PATIENT + value.PatientID);

					const interviewerResponse = await fetchWithRefresh(
						API_URL + GET_INTERVIEWERS + value.InterviewerID
					);

					let patient: Patient = {};
					let interviewer: Interviewer = {};

					if (patientResponse.ok && interviewerResponse.ok) {
						patient = (await patientResponse.json()) as Patient;
						interviewer = (await interviewerResponse.json()) as Interviewer;
					}

					handleResponse(patientResponse.status, false);
					handleResponse(interviewerResponse.status, false);

					return {
						id: value.ID,
						date: value.StartDate,
						interviewver: interviewer.firstName ?? '' + interviewer.firstSurname ?? '',
						patient: patient.firstName ?? '' + patient.firstSurname ?? ''
					};
				})
			);

			filteredRows = rows;
		}

		handleResponse(response.status, false);
	}
</script>

<main>
	{#if rows == undefined}
		<DataTableSkeleton
			headers={[
				{ key: 'date', value: 'Fecha' },
				{ key: 'interviewer', value: 'Entrevistador' },
				{ key: 'patient', value: 'Paciente' }
			]}
			rows={5}
		/>
	{:else}
		<DataTable
			title="Entrevistas"
			description="Entrevistas realizadas"
			headers={[
				{ key: 'date', value: 'Fecha' },
				{ key: 'interviewer', value: 'Entrevistador' },
				{ key: 'patient', value: 'Paciente' },
				{ key: 'overflow', empty: true }
			]}
			bind:rows={filteredRows}
		>
			<svelte:fragment slot="cell" let:cell let:row>
				{#if cell.key === 'overflow'}
					<OverflowMenu flipped />
				{:else}{cell.value}{/if}
			</svelte:fragment>
			<Toolbar>
				<ToolbarContent>
					<ToolbarSearch bind:value={searchValue} />
					<Button on:click={() => {}}>Realizar entrevista</Button>
				</ToolbarContent>
			</Toolbar>
		</DataTable>
	{/if}
</main>

<style>
</style>
