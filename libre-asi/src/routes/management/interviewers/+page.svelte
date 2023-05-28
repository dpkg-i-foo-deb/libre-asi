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
		ModalFooter,
		CodeSnippet
	} from 'carbon-components-svelte';
	import { onMount } from 'svelte';
	import type Interviewer from '$lib/models/Interviewer';
	import type { DataTableRow } from 'carbon-components-svelte/types/DataTable/DataTable.svelte';
	import {
		API_URL,
		DELETE_INTERVIEWER,
		GET_INTERVIEWERS,
		REGISTER_INTERVIEWER
	} from '$lib/api/constants';
	import { fetchWithRefresh } from '$lib/util/fetchRefresh';
	import { handleResponse } from '$lib/util/handleResponse';
	import { sendInfo, sendSuccess } from '$lib/util/notifications';
	import { goto } from '$app/navigation';
	import { ValueVariable } from 'carbon-icons-svelte';

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

	let invalidEmail = false;
	let invalidUsername = false;
	let invalidPersonalId = false;
	let invalidFirstname = false;
	let invalidForm = true;

	let rows: ReadonlyArray<DataTableRow>;
	let filteredRows: ReadonlyArray<DataTableRow>;

	let deletingId = 0;
	let isDeleteModalOpen = false;
	let searchValue = '';

	let isRegisterFormOpen = false;

	let isSuccessRegisterOpen = false;

	onMount(async function () {
		await loadInterviewers();
	});

	function validateForm() {
		invalidForm = true;

		if (!invalidEmail && !invalidFirstname && !invalidPersonalId && !invalidUsername) {
			invalidForm = false;
		}

		if (
			newInterviewer.email == '' ||
			newInterviewer.username == '' ||
			newInterviewer.personalID == '' ||
			newInterviewer.firstName == ''
		) {
			invalidForm = true;
		}
	}

	async function loadInterviewers() {
		const response = await fetchWithRefresh(API_URL + GET_INTERVIEWERS, { method: 'GET' });
		if (response.ok) {
			const existingInterviewers = (await response.json()) as Interviewer[];

			rows = existingInterviewers.map(function (value: Interviewer) {
				let name = value.firstName + ' ' + value.lastName;

				if (name.length > 10) {
					name = name.substring(0, 10) + '...';
				}

				if (value.personalID?.length ?? 0 > 15) {
					value.personalID = value.personalID?.substring(0, 15) + '...';
				}

				return {
					id: value.ID,
					username: value.username,
					firstName: name,
					personalID: value.personalID
				};
			});

			filteredRows = rows;
		}

		handleResponse(response.status, false);
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

	function handleCancel() {
		isDeleteModalOpen = false;
	}

	async function handleRegister() {
		const response = await fetchWithRefresh(API_URL + REGISTER_INTERVIEWER, {
			method: 'POST',
			body: JSON.stringify(newInterviewer)
		});

		if (response.ok) {
			await loadInterviewers();
			isRegisterFormOpen = false;
			sendSuccess('Entrevistador registrado', 'Entrevistador registrado exitosamente');

			newInterviewer = (await response.json()) as Interviewer;

			isRegisterFormOpen = false;
			isSuccessRegisterOpen = true;
		}

		handleResponse(response.status, false);
	}

	async function handleDelete() {
		const response = await fetchWithRefresh(API_URL + DELETE_INTERVIEWER + deletingId, {
			method: 'DELETE'
		});

		if (response.ok) {
			await loadInterviewers();
			isDeleteModalOpen = false;
			sendInfo('Entrevistador eliminado', 'Entrevistador eliminado exitosamente');
		}

		handleResponse(response.status, false);
	}
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
						<OverflowMenuItem
							text="Editar"
							on:click={function () {
								goto('interviewers/' + row.id);
							}}>Editar</OverflowMenuItem
						>
						<OverflowMenuItem
							danger
							text="Eliminar"
							on:click={() => ((isDeleteModalOpen = true), (deletingId = row.id))}
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
						invalidText="Ingrese un correo electrónico válido"
						bind:invalid={invalidEmail}
						bind:value={newInterviewer.email}
						on:blur={function () {
							invalidEmail = false;
							if (newInterviewer.email == '') {
								invalidEmail = true;
							}
							validateForm();
						}}
					/>
				</div>

				<div class="input-field">
					<TextInput
						id="username"
						labelText="Nombre de usuario"
						placeholder="Ingrese el nombre de usuario"
						bind:invalid={invalidUsername}
						invalidText="Ingrese un nombre de usuario válido"
						on:blur={function () {
							invalidUsername = false;
							if (newInterviewer.username == '') {
								invalidUsername = true;
							}
							validateForm();
						}}
						bind:value={newInterviewer.username}
					/>
				</div>

				<div class="input-field">
					<TextInput
						id="id"
						labelText="Número de documento de identidad"
						placeholder="Ingrese el número de documento de identidad"
						bind:invalid={invalidPersonalId}
						invalidText="Ingrese un documento de identidad válido"
						on:blur={function () {
							invalidPersonalId = false;
							if (newInterviewer.personalID == '') {
								invalidPersonalId = true;
							}
							validateForm();
						}}
						bind:value={newInterviewer.personalID}
					/>
				</div>

				<div class="input-field">
					<TextInput
						id="firstName"
						labelText="Primer nombre"
						placeholder="Ingrese el primer nombre"
						bind:invalid={invalidFirstname}
						invalidText="Este campo es requerido"
						on:blur={function () {
							invalidFirstname = false;

							if (newInterviewer.firstName == '') {
								invalidFirstname = true;
							}
							validateForm();
						}}
						bind:value={newInterviewer.firstName}
					/>
				</div>
			</form>
		</ModalBody>

		<ModalFooter
			primaryButtonDisabled={invalidEmail ||
				invalidFirstname ||
				invalidPersonalId ||
				invalidUsername ||
				invalidForm}
			primaryButtonText="Registrar entrevistador"
			secondaryButtonText="Cancelar"
			on:click:button--secondary={function () {
				isRegisterFormOpen = false;
			}}
		/>
	</ComposedModal>

	<Modal
		danger
		bind:open={isDeleteModalOpen}
		modalHeading="Eliminación de entrevistador"
		primaryButtonText="Eliminar entrevistador"
		secondaryButtonText="Cancelar"
		on:submit={handleDelete}
		on:click:button--secondary={handleCancel}
	>
		<p>¿Está seguro que desea eliminar este entrevistador?</p>
	</Modal>

	<Modal
		passiveModal
		on:close={function () {
			isSuccessRegisterOpen = false;
		}}
		bind:open={isSuccessRegisterOpen}
		modalHeading="Entrevistador registrado"
		primaryButtonText="Finalizar"
	>
		<p>El entrevistador ha sido registrado y una contraseña aleatoria ha sido generada.</p>
		<br />
		<p class="bold">
			El nuevo entrevistador tendrá que establecer su contraseña cuando inicie sesión.
		</p>
		<br />
		<p>La contraseña generada es:</p>
		<div class="password-container">
			<CodeSnippet code={newInterviewer.password} />
		</div>
	</Modal>
{/if}

<style>
	.input-field {
		padding-top: 10px;
		padding-bottom: 10px;
	}

	.paragraph {
		padding-top: 5px;
		padding-bottom: 5px;
	}

	.password-container {
		margin-top: 10px;
		padding-bottom: 40px;
	}

	.bold {
		font-weight: bold;
	}
</style>
