<script lang="ts">
	import { fetchWithRefresh } from '$lib/util/fetchRefresh';
	import {
		API_URL,
		GET_PATIENTS,
		REGISTER_PATIENT,
		EDIT_PATIENTS,
		DELETE_PATIENTS
	} from '$lib/api/constants';
	import type Patient from '$lib/models/Patient';
	import {
		Button,
		DataTable,
		DataTableSkeleton,
		OverflowMenu,
		OverflowMenuItem,
		Toolbar,
		ToolbarContent,
		ToolbarSearch
	} from 'carbon-components-svelte';
	import type { DataTableRow } from 'carbon-components-svelte/types/DataTable/DataTable.svelte';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { checkEmail, checkUsername } from '$lib/util/formUtils';
	import { handleResponse } from '$lib/util/handleResponse';

	let rows: ReadonlyArray<DataTableRow>;
	let filteredRows: ReadonlyArray<DataTableRow>;
	let duplicateCredentials = false;
	let searchValue: string;
	let isRegisterFormOpen = false;
	let invalidEmailCaption = '';
	let invalidUsernameCaption = '';
	let invalidUsername = false;
	let invalidEmail = false;
	let newPatient: Patient;
	let email = '';
	let username = '';
	let isSuccessRegisterOpen = false;

	onMount(async function () {
		await loadPatients();
	});

	async function loadPatients() {
		const response = await fetchWithRefresh(API_URL + GET_PATIENTS, { method: 'GET' });

		if (response.ok) {
			const existingPatients = (await response.json()) as Patient[];

			rows = existingPatients.map(function (value: Patient) {
				return { id: value.ID, email: value.email, username: value.username };
			});

			filteredRows = rows;
		}
	}

	function validateUsername(): boolean {
		const usernameField = checkUsername(username);

		invalidUsernameCaption = usernameField[0];
		invalidUsername = !usernameField[1];

		return usernameField[1];
	}

	function validateEmail(): boolean {
		const emailField = checkEmail(email);

		invalidEmailCaption = emailField[0];
		invalidEmail = !emailField[1];

		return emailField[1];
	}
	async function registerPatient() {
		duplicateCredentials = false;

		if (!validateEmail() || !validateUsername()) {
			return;
		}

		newPatient = {
			ID: 0,
			CreatedAt: new Date(),
			UpdatedAt: new Date(),
			email: email,
			username: username,
			password: '',
			firstName: '',
			lastName: '',
			firstSurname: '',
			secondSurname: '',
			birthdate: new Date(),
			personalID: ''
		};

		const response = await fetchWithRefresh(API_URL + REGISTER_PATIENT, {
			method: 'POST',
			body: JSON.stringify(newPatient)
		});

		if (response.ok) {
			newPatient = (await response.json()) as Patient;
			isRegisterFormOpen = false;
			isSuccessRegisterOpen = true;
			loadPatients();
		}

		if (handleResponse(response.status, false))
			if (response.status == 409) {
				duplicateCredentials = true;
			}

		email = '';
		username = '';
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
						<OverflowMenuItem text="Edit" on:click={function () {}}>Edit</OverflowMenuItem>
						<OverflowMenuItem danger text="Delete" />
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
</main>

<style>
</style>
