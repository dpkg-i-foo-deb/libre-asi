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
		OverflowMenuItem,
		ComposedModal,
		ModalHeader,
		ModalBody,
		TextInput,
		ModalFooter
	} from 'carbon-components-svelte';
	import { onMount } from 'svelte';
	import type Interviewer from '$lib/models/Interviewer';
	import type { DataTableRow } from 'carbon-components-svelte/types/DataTable/DataTable.svelte';
	import { API_URL, GET_INTERVIEWERS } from '$lib/api/constants';
	import { fetchWithRefresh } from '$lib/util/fetchRefresh';
	import { handleResponse } from '$lib/util/handleResponse';

	let newInterviewer: Interviewer = {
		email: '',
		username: '',
		age: 0,
		birthdate: new Date(),
		firstName: '',
		lastName: '',
		firstSurname: '',
		lastSurname: '',
		password: '',
		personalID: '',
		ID: 0
	};

	let rows: ReadonlyArray<DataTableRow>;
	let filteredRows: ReadonlyArray<DataTableRow>;

	let deletingId = 0;
	let isModalOpen = false;
	let searchValue = '';

	let isRegisterFormOpen = false;

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
					username: value.username,
					firstName: value.firstName + value.lastName,
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

	function handleRegister() {}

	function handleDelete() {}
</script>

{#if rows == undefined}
	<DataTableSkeleton
		headers={[
			{ key: 'firstName', value: 'Primer Nombre' },
			{ key: 'personalID', value: 'Identificación' }
		]}
		rows={5}
	/>
{:else}
	<div class="data-table">
		<DataTable
			title="Entrevistadores"
			description="Registrados actualmente"
			headers={[
				{ key: 'firstName', value: 'Nombre' },
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
							isRegisterFormOpen = true;
						}}>Registrar Entrevistador</Button
					>
				</ToolbarContent>
			</Toolbar>
		</DataTable>
	</div>

	<ComposedModal
		bind:open={isRegisterFormOpen}
		selectorPrimaryFocus="#email"
		on:submit={handleRegister}
	>
		<ModalHeader label="Transacción" title="Registro de entrevistador" />

		<ModalBody hasForm>
			<form>
				<h7 class="paragraph"> El registro provee únicamente los datos mínimos </h7>

				<br />

				<h6 class="paragraph">
					La contraseña de la cuenta será generada automáticamente y el nuevo usuario tendrá que
					actualizarla
				</h6>

				<div class="input-field">
					<TextInput
						id="email"
						labelText="Correo electrónico"
						placeholder="Ingrese el correo electrónico"
						bind:value={newInterviewer.email}
					/>
				</div>

				<div class="input-field">
					<TextInput
						id="username"
						labelText="Nombre de usuario"
						placeholder="Ingrese el nombre de usuario"
						bind:value={newInterviewer.username}
					/>
				</div>

				<div class="input-field">
					<TextInput
						id="id"
						labelText="Número de documento de identidad"
						placeholder="Ingrese el número de documento de identidad"
						bind:value={newInterviewer.personalID}
					/>
				</div>

				<div class="input-field">
					<TextInput
						id="firstName"
						labelText="Primer nombre"
						placeholder="Ingrese el primer nombre"
						bind:value={newInterviewer.firstName}
					/>
				</div>
			</form>
		</ModalBody>

		<ModalFooter
			primaryButtonText="Registrar administrador"
			secondaryButtonText="Cancelar"
			on:click:button--secondary={function () {
				isRegisterFormOpen = false;
			}}
		/>
	</ComposedModal>

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
	.data-table {
		padding-right: 0;
	}

	.input-field {
		padding-top: 10px;
		padding-bottom: 10px;
	}

	.paragraph {
		padding-top: 5px;
		padding-bottom: 5px;
	}
</style>
