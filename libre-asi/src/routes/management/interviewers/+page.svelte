<script lang="ts">
	import {
		Button,
		DataTable,
		Toolbar,
		ToolbarContent,
		ToolbarSearch,
		DataTableSkeleton,
		Modal,
		OverflowMenu,
		OverflowMenuItem
	} from 'carbon-components-svelte';
	import { onMount } from 'svelte';
	import type Interviewer from '$lib/models/Interviewer';
	import type { DataTableRow } from 'carbon-components-svelte/types/DataTable/DataTable.svelte';
	import { API_URL, GET_INTERVIEWERS } from '$lib/api/constants';
	import { fetchWithRefresh } from '$lib/util/fetchRefresh';
	import { handleResponse } from '$lib/util/handleResponse';
	import { goto } from '$app/navigation';

	let rows: ReadonlyArray<DataTableRow>;
	let filteredRows: ReadonlyArray<DataTableRow>;

	let deletingId = 0;
	let isModalOpen = false;
	let searchValue = '';

	onMount(async function () {
		await loadInterviewers();
	});

	async function loadInterviewers() {
		const response = await fetchWithRefresh(API_URL + GET_INTERVIEWERS, { method: 'GET' });
		if (response.ok) {
			const existingInterviewers = (await response.json()) as Interviewer[];

			rows = existingInterviewers.map(function (value: Interviewer) {
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

		handleResponse(response.status, false);
	}

	function handleSearch() {}

	function handleCancel() {
		isModalOpen = false;
	}

	function handleDelete() {}
</script>

{#if rows == undefined}
	<DataTableSkeleton
		headers={[
			{ key: 'email', value: 'Correo' },
			{ key: 'firstName', value: 'Primer Nombre' },
			{ key: 'personalID', value: 'Identificación' }
		]}
		rows={5}
	/>
{:else}
	<DataTable
		title="Entrevistadores"
		description="Registrados actualmente"
		headers={[
			{ key: 'email', value: 'Correo' },
			{ key: 'firstName', value: 'Primer Nombre' },
			{ key: 'personalID', value: 'Identificación' },
			{ key: 'overflow', empty: true }
		]}
		bind:rows={filteredRows}
	>
		<svelte:fragment slot="cell" let:cell let:row>
			{#if cell.key === 'overflow'}
				<OverflowMenu flipped>
					<OverflowMenuItem text="Edit" on:click={function () {}}>Edit</OverflowMenuItem>
					<OverflowMenuItem
						danger
						text="Delete"
						on:click={() => ((isModalOpen = true), (deletingId = row.id))}
					/>
				</OverflowMenu>
			{:else}{cell.value}{/if}
		</svelte:fragment>
		<Toolbar>
			<ToolbarContent>
				<ToolbarSearch bind:value={searchValue} on:input={handleSearch} />
				<Button
					on:click={function () {
						goto('/management/interviewers/create');
					}}>Registrar Entrevistador</Button
				>
			</ToolbarContent>
		</Toolbar>
	</DataTable>

	<Modal
		danger
		bind:open={isModalOpen}
		modalHeading="Delete Interviewer"
		primaryButtonText="Delete"
		secondaryButtonText="Cancel"
		on:submit={handleDelete}
		on:click:button--secondary={handleCancel}
	>
		<p>Are you sure you want to delete this Interviewer?</p>
	</Modal>
{/if}

<style>
</style>
