<script lang="ts">
	import { goto } from '$app/navigation';
	import { API_URL, GET_PATIENTS, START_INTERVIEW } from '$lib/api/constants';
	import type Interview from '$lib/models/Interview';
	import type Patient from '$lib/models/Patient';
	import { fetchWithRefresh } from '$lib/util/fetchRefresh';
	import { handleResponse } from '$lib/util/handleResponse';
	import { Button, DataTable, DataTableSkeleton, Tile } from 'carbon-components-svelte';
	import type { DataTableRow } from 'carbon-components-svelte/types/DataTable/DataTable.svelte';
	import { onMount } from 'svelte';

	let rows: ReadonlyArray<DataTableRow>;
	let filteredRows: ReadonlyArray<DataTableRow>;

	let selectedRowIds: number[] = [];

	let newInterview: Interview = {};

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

	async function startInterview() {
		newInterview.patientID = selectedRowIds[0];

		const response = await fetchWithRefresh(API_URL + START_INTERVIEW, {
			method: 'POST',
			body: JSON.stringify(newInterview)
		});

		if (response.ok) {
			newInterview = (await response.json()) as Interview;
			goto('/management/interviews/perform/' + newInterview.id);
		}

		handleResponse(response.status, false);
	}
</script>

<main>
	<Tile>
		<div class="content">
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
		</div>

		<div class="button-container">
			<Button
				size="default"
				on:click={function () {
					startInterview();
				}}>Realizar entrevista</Button
			>
		</div></Tile
	>
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
