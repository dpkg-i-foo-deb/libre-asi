<script lang="ts">
	import { fetchWithRefresh } from '$lib/util/fetchRefresh';
	import { API_URL, DELETE_PATIENT, GET_PATIENTS, REGISTER_PATIENT } from '$lib/api/constants';
	import type Patient from '$lib/models/Patient';
	import {
		Button,
		ComposedModal,
		DataTable,
		DataTableSkeleton,
		Modal,
		ModalBody,
		ModalFooter,
		ModalHeader,
		OverflowMenu,
		OverflowMenuItem,
		TextInput,
		Toolbar,
		ToolbarContent,
		ToolbarSearch
	} from 'carbon-components-svelte';
	import type { DataTableRow } from 'carbon-components-svelte/types/DataTable/DataTable.svelte';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { sendSuccess } from '$lib/util/notifications';
	import { handleResponse } from '$lib/util/handleResponse';

	let newPatient: Patient = {
		email: '',
		username: '',
		age: 0,
		firstName: '',
		firstSurname: '',
		lastName: '',
		lastSurname: '',
		password: '',
		personalID: '',
		ID: 0
	};

	let isDeleteModalOpen = false;
	let deletingId: number;

	let rows: ReadonlyArray<DataTableRow>;
	let filteredRows: ReadonlyArray<DataTableRow>;
	let searchValue: string;

	let isRegisterFormOpen = false;

	let invalidFirstName = false;
	let invalidFirstSurname = false;
	let invalidPersonalId = false;

	let invalidForm = true;

	onMount(async function () {
		await loadPatients();
	});

	function validateForm() {
		invalidForm = true;

		if (!invalidFirstName && !invalidFirstSurname && !invalidPersonalId) {
			invalidForm = false;
		}

		if (
			newPatient.firstName == '' ||
			newPatient.firstSurname == '' ||
			newPatient.personalID == ''
		) {
			invalidForm = true;
		}
	}

	async function loadPatients() {
		const response = await fetchWithRefresh(API_URL + GET_PATIENTS, { method: 'GET' });

		if (response.ok) {
			const existingPatients = (await response.json()) as Patient[];

			rows = existingPatients.map(function (value: Patient) {
				return {
					id: value.ID,
					firstName: value.firstName ?? '' + value.lastName ?? '',
					personalID: value.personalID
				};
			});

			filteredRows = rows;
		}

		handleResponse(response.status, false);
	}

	async function handleRegister() {
		const response = await fetchWithRefresh(API_URL + REGISTER_PATIENT, {
			method: 'POST',
			body: JSON.stringify(newPatient)
		});

		if (response.ok) {
			await loadPatients();
			isRegisterFormOpen = false;
			sendSuccess('Paciente registrado', 'Paciente registrado exitosamente');
		}

		handleResponse(response.status, false);
	}

	async function deletePatient() {
		const response = await fetchWithRefresh(API_URL + DELETE_PATIENT + deletingId, {
			method: 'DELETE'
		});

		if (response.ok) {
			sendSuccess('Paciente eliminado', 'Paciente eliminado exitosamente');
		}

		handleResponse(response.status, false);

		isDeleteModalOpen = false;

		loadPatients();
	}
	function handleSearch() {
		const query = searchValue;

		if (searchValue == '') {
			filteredRows = rows;
			return;
		}

		filteredRows = rows.filter((row) => {
			return row.firstName.toLocaleLowerCase().includes(query) || row.personalID.includes(query);
		});
	}
</script>

<main>
	{#if rows == undefined}
		<DataTableSkeleton
			headers={[
				{ key: 'firstName', value: 'Nombre' },
				{ key: 'personalID', value: 'Identificación' }
			]}
			rows={5}
		/>
	{:else}
		<DataTable
			title="Pacientes"
			description="Pacientes registrados actualmente"
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
						<OverflowMenuItem
							text="Editar"
							on:click={function () {
								goto('patients/' + row.id);
							}}>Editar</OverflowMenuItem
						>
						<OverflowMenuItem
							danger
							text="Eliminar"
							on:click={function () {
								isDeleteModalOpen = true;

								deletingId = row.id;
							}}
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
						}}>Registrar paciente</Button
					>
				</ToolbarContent>
			</Toolbar>
		</DataTable>
	{/if}

	<ComposedModal
		bind:open={isRegisterFormOpen}
		selectorPrimaryFocus="#firstName"
		on:submit={handleRegister}
	>
		<ModalHeader label="Transacción" title="Registro de paciente" />

		<ModalBody hasForm>
			<form>
				<h7 class="paragraph"> El registro provee únicamente los datos mínimos</h7>

				<br />

				<div class="input-field">
					<TextInput
						id="firstName"
						labelText="Primer nombre"
						placeholder="Ingrese el primer nombre"
						bind:invalid={invalidFirstName}
						invalidText="Ingrese un nombre válido"
						on:blur={function () {
							invalidFirstName = false;

							if (newPatient.firstName == '') {
								invalidFirstName = true;
							}
						}}
						bind:value={newPatient.firstName}
						on:input={function () {
							validateForm();
						}}
					/>
				</div>

				<div class="input-field">
					<TextInput
						id="firstSurname"
						labelText="Primer apellido"
						placeholder="Ingrese el primer apellido"
						bind:invalid={invalidFirstSurname}
						invalidText="Ingrese un apellido válido"
						on:blur={function () {
							invalidFirstSurname = false;

							if (newPatient.firstSurname == '') {
								invalidFirstSurname = true;
							}
						}}
						on:input={function () {
							validateForm();
						}}
						bind:value={newPatient.firstSurname}
					/>
				</div>

				<div class="input-field">
					<TextInput
						id="id"
						labelText="Número de identificación"
						placeholder="Ingrese el número de identificación"
						bind:invalid={invalidPersonalId}
						invalidText="Ingrese una identificación válida"
						on:blur={function () {
							invalidPersonalId = false;

							if (newPatient.personalID == '') {
								invalidPersonalId = true;
							}
						}}
						on:input={function () {
							validateForm();
						}}
						bind:value={newPatient.personalID}
					/>
				</div>
			</form>
		</ModalBody>

		<ModalFooter
			primaryButtonDisabled={invalidForm}
			primaryButtonText="Registrar paciente"
			secondaryButtonText="Cancelar"
			on:click:button--secondary={function () {
				isRegisterFormOpen = false;
			}}
		/>
	</ComposedModal>

	<Modal
		danger
		bind:open={isDeleteModalOpen}
		modalHeading="Eliminar Paciente"
		primaryButtonText="Eliminar"
		secondaryButtonText="Cancelar"
		on:click:button--secondary={() => (isDeleteModalOpen = false)}
		on:open
		on:close
		on:submit={deletePatient}
	>
		<p>Esta acción es permanente.</p>
	</Modal>
</main>

<style>
	.input-field {
		padding-top: 10px;
		padding-bottom: 10px;
	}

	.paragraph {
		padding-top: 5px;
		padding-bottom: 5px;
	}
</style>
