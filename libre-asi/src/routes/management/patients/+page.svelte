<script lang="ts">
	import { fetchWithRefresh } from '$lib/util/fetchRefresh';
	import { API_URL, DELETE_PATIENT, GET_PATIENTS } from '$lib/api/constants';
	import type Patient from '$lib/models/Patient';
	import {
		Button,
		DataTable,
		DataTableSkeleton,
		Modal,
		OverflowMenu,
		OverflowMenuItem,
		Toolbar,
		ToolbarContent,
		ToolbarSearch
	} from 'carbon-components-svelte';
	import type { DataTableRow } from 'carbon-components-svelte/types/DataTable/DataTable.svelte';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { sendSuccess } from '$lib/util/notifications';
	import { handleResponse } from '$lib/util/handleResponse';

	let openDeleteModal = false;
	let deletingId: number;

	let rows: ReadonlyArray<DataTableRow>;
	let filteredRows: ReadonlyArray<DataTableRow>;
	let searchValue: string;

	onMount(async function () {
		await loadPatients();
	});

	async function loadPatients() {
		const response = await fetchWithRefresh(API_URL + GET_PATIENTS, { method: 'GET' });

		if (response.ok) {
			const existingPatients = (await response.json()) as Patient[];

			rows = existingPatients.map(function (value: Patient) {
				return {
					id: value.ID,
					email: value.email,
					username: value.username,
					firstName: value.firstName,
					personalID: value.personalID
				};
			});

			filteredRows = rows;
		}
	}

	async function deletePatient() {
		const response = await fetchWithRefresh(API_URL + DELETE_PATIENT + deletingId, {
			method: 'DELETE'
		});

		if (response.ok) {
			sendSuccess('Éxito', 'Paciente eliminado exitosamente');
		}

		handleResponse(response.status, false);

		openDeleteModal = false;

		loadPatients();
	}
</script>

<main>
	{#if rows == undefined}
		<DataTableSkeleton
			headers={[
				{ key: 'email', value: 'Email' },
				{ key: 'firstName', value: 'First Name' },
				{ key: 'personalID', value: 'Personal ID' }
			]}
			rows={5}
		/>
	{:else}
		<DataTable
			title="Patients"
			description="Current Registered Patients"
			headers={[
				{ key: 'email', value: 'Email' },
				{ key: 'firstName', value: 'First Name' },
				{ key: 'personalID', value: 'Personal ID' },
				{ key: 'overflow', empty: true }
			]}
			bind:rows={filteredRows}
		>
			<svelte:fragment slot="cell" let:cell let:row>
				{#if cell.key === 'overflow'}
					<OverflowMenu flipped>
						<OverflowMenuItem
							text="Edit"
							on:click={function () {
								goto('patients/' + row.id);
							}}>Edit</OverflowMenuItem
						>
						<OverflowMenuItem
							danger
							text="Delete"
							on:click={function () {
								openDeleteModal = true;

								deletingId = row.id;
							}}
						/>
					</OverflowMenu>
				{:else}{cell.value}{/if}
			</svelte:fragment>
			<Toolbar>
				<ToolbarContent>
					<ToolbarSearch bind:value={searchValue} />
					<Button
						on:click={() => {
							goto('patients/create/');
						}}>Register Patient</Button
					>
				</ToolbarContent>
			</Toolbar>
		</DataTable>
	{/if}

	<Modal
		danger
		bind:open={openDeleteModal}
		modalHeading="Eliminar Paciente"
		primaryButtonText="Eliminar"
		secondaryButtonText="Cancelar"
		on:click:button--secondary={() => (openDeleteModal = false)}
		on:open
		on:close
		on:submit={deletePatient}
	>
		<p>Esta acción es permanente.</p>
	</Modal>
</main>

<style>
</style>
