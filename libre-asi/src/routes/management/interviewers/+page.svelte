<script lang="ts">
	import {
		Button,
		DataTable,
		Toolbar,
		ToolbarContent,
		ToolbarSearch,
		DataTableSkeleton,
		ComposedModal,
		ModalHeader,
		ModalFooter,
		ModalBody,
		TextInput,
		InlineNotification,
		Modal,
		CodeSnippet,
		OverflowMenu,
		OverflowMenuItem
	} from 'carbon-components-svelte';
	import { onMount } from 'svelte';
	import type Interviewer from '$lib/models/Interviewer';
	import type { DataTableRow } from 'carbon-components-svelte/types/DataTable/DataTable.svelte';
	import { checkEmail, checkUsername } from '$lib/util/formUtils';
	import { handleResponse } from '$lib/util/handleResponse';
	import {
		API_URL,
		GET_INTERVIEWERS,
		REGISTER_INTERVIEWER,
		EDIT_INTERVIEWER,
		DELETE_INTERVIEWER
	} from '$lib/api/constants';
	import { fetchWithRefresh } from '$lib/util/fetchRefresh';
	import { sendError, sendSuccess } from '$lib/util/notifications';

	let rows: ReadonlyArray<DataTableRow>;
	let filteredRows: ReadonlyArray<DataTableRow>;

	let email = '';
	let username = '';

	let invalidEmail = false;
	let invalidUsername = false;
	let invalidEmailCaption = '';
	let invalidUsernameCaption = '';
	let duplicateCredentials = false;
	let isEditionFormOpen = false;
	let editingId = 0;
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
	}

	async function register() {}

	function handleSearch() {}

	function handleCancel() {
		isModalOpen = false;
	}

	function handleDelete() {}
</script>

{#if rows == undefined}
	<DataTableSkeleton
		headers={[
			{ key: 'email', value: 'Email' },
			{ key: 'username', value: 'Username' }
			// { key: 'firstName', value: 'First Name' },
			// { key: 'firstSurName', value: 'First Surname' }
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
				<Button on:click={function () {}}>Register Interviewer</Button>
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
