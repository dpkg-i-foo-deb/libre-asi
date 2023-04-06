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

	let newInterviewer: Interviewer;
	let editedInterviewer: Interviewer;
	let rows: ReadonlyArray<DataTableRow>;
	let filteredRows: ReadonlyArray<DataTableRow>;
	let isRegisterFormOpen = false;
	let isSuccessRegisterOpen = false;

	let email = '';
	let username = '';
	let firstName = '';
	let firstSurName = '';
	let editEmail = '';
	let editUsername = '';
	let editFirstName = '';
	let editFirstSurName = '';

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
		newInterviewer = {
			ID: 0,
			CreatedAt: new Date(),
			UpdatedAt: new Date(),
			email: '',
			username: '',
			password: '',
			firstName: '',
			lastName: '',
			firstSurName: '',
			lastSurName: '',
			birthdate: new Date(),
			age: 0,
			personalId: ''
		};

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
					username: value.username
					// ,
					// firstName: value.firstName,
					// firstSurName: value.firstSurName
				};
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

	async function register() {
		duplicateCredentials = false;
		isSuccessRegisterOpen = false;

		if (!validateEmail() || !validateUsername()) {
			return;
		}

		newInterviewer = {
			ID: 0,
			CreatedAt: new Date(),
			UpdatedAt: new Date(),
			email: email,
			username: username,
			password: '',
			firstName: '',
			lastName: '',
			firstSurName: '',
			lastSurName: '',
			birthdate: new Date(),
			age: 0,
			personalId: ''
		};

		const response = await fetchWithRefresh(API_URL + REGISTER_INTERVIEWER, {
			method: 'POST',
			body: JSON.stringify(newInterviewer)
		});

		if (response.ok) {
			newInterviewer = (await response.json()) as Interviewer;
			isRegisterFormOpen = false;
			isSuccessRegisterOpen = true;
			loadInterviewers();
		}

		if (handleResponse(response.status, false))
			if (response.status == 409) {
				duplicateCredentials = true;
			}

		email = '';
		username = '';
	}

	function handleSearch() {
		const query = searchValue;

		if (searchValue == '') {
			filteredRows = rows;
			return;
		}

		filteredRows = rows.filter((row) => {
			return (
				row.email.toLocaleLowerCase().includes(query) ||
				row.username.includes(query) ||
				row.firtsName.toLocaleLowerCase().includes(query) ||
				row.firtsSurName.toLocaleLowerCase().includes(query)
			);
		});
	}

	function toggleEditForm() {
		isEditionFormOpen = true;

		let row = rows.find((row) => row.id === editingId);

		if (row) {
			editEmail = row.email;
			editUsername = row.username;
		}
	}
	function handleCancel() {
		isModalOpen = false;
	}

	async function editAdmin() {
		duplicateCredentials = false;
		isSuccessRegisterOpen = false;

		if (!validateEmail() || !validateUsername()) {
			return;
		}

		editedInterviewer = {
			ID: editingId,
			CreatedAt: new Date(),
			UpdatedAt: new Date(),

			email: editEmail,
			username: editUsername,
			password: '',
			firstName: '',
			lastName: '',
			firstSurName: '',
			lastSurName: '',
			birthdate: new Date(),
			age: 0,
			personalId: ''
		};

		const response = await fetchWithRefresh(API_URL + EDIT_INTERVIEWER, {
			method: 'PATCH',
			body: JSON.stringify(editedInterviewer)
		});

		if (response.ok) {
			editedInterviewer = (await response.json()) as Interviewer;
			isEditionFormOpen = false;
			loadInterviewers();
			sendSuccess('Account modified successfully', 'You can now use the new credentials');
			return;
		}

		if (handleResponse(response.status, false)) {
			return;
		}

		if (response.status == 409) {
			duplicateCredentials = true;
		}
	}

	async function handleDelete() {
		const response = await fetchWithRefresh(API_URL + DELETE_INTERVIEWER + deletingId.toString(), {
			method: 'DELETE'
		});

		if (response.ok) {
			loadInterviewers();
			sendSuccess('Account deleted successfully', '');
			isModalOpen = false;
			return;
		}

		if (handleResponse(response.status, false)) {
			sendError('Error trying to remove Interviewer', '');
			isModalOpen = false;
			return;
		}
	}
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
		title="Interviewers"
		description="Current Registered Interviewers"
		headers={[
			{ key: 'email', value: 'Email' },
			{ key: 'username', value: 'Username' },
			// { key: 'firstName', value: 'First Name' },
			// { key: 'firstSurName', value: 'First Surname' },
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
							editingId = row.id;
							toggleEditForm();
						}}>Edit</OverflowMenuItem
					>
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
					on:click={() => {
						isRegisterFormOpen = true;
					}}>Register Interviewer</Button
				>
			</ToolbarContent>
		</Toolbar>
	</DataTable>

	<ComposedModal bind:open={isRegisterFormOpen} selectorPrimaryFocus="#email" on:submit={register}>
		<ModalHeader label="Transaction" title="New Interviewer Account" />
		<ModalBody hasForm>
			<form>
				<h7 class="paragraph" />
				<br />
				<h6 class="paragraph">
					The account password will be generated automatically and the new user will be prompted to
					change it
				</h6>

				{#if duplicateCredentials}
					<InlineNotification title="Error:" subtitle="Email or username already registered" />
				{/if}
				<div class="input-field">
					<TextInput
						id="email"
						labelText="Email"
						placeholder="Enter email..."
						on:blur={validateEmail}
						bind:invalid={invalidEmail}
						bind:value={email}
						bind:invalidText={invalidEmailCaption}
					/>
				</div>
				<div class="input-field">
					<TextInput
						id="username"
						labelText="User name"
						placeholder="Enter user name..."
						on:blur={validateUsername}
						bind:invalid={invalidUsername}
						bind:value={username}
						bind:invalidText={invalidUsernameCaption}
					/>
				</div>
			</form>
		</ModalBody>
		<ModalFooter
			primaryButtonText="Register"
			secondaryButtonText="Cancel"
			on:click:button--secondary={() => {
				isRegisterFormOpen = false;
			}}
		/>
	</ComposedModal>

	<ComposedModal bind:open={isEditionFormOpen} selectorPrimaryFocus="#email" on:submit={editAdmin}>
		<ModalHeader label="Transaction" title="Edit Administrator Account" />
		<ModalBody hasForm>
			<form>
				<h7 class="paragraph"> Interviewers can edit patients and interviews information. </h7>
				<br />

				{#if duplicateCredentials}
					<InlineNotification title="Error:" subtitle="Email or username already registered" />
				{/if}
				<div class="input-field">
					<TextInput
						id="email"
						labelText="Email"
						placeholder="Enter email..."
						on:blur={validateEmail}
						bind:invalid={invalidEmail}
						bind:value={editEmail}
						bind:invalidText={invalidEmailCaption}
					/>
				</div>
				<div class="input-field">
					<TextInput
						id="username"
						labelText="User name"
						placeholder="Enter user name..."
						on:blur={validateUsername}
						bind:invalid={invalidUsername}
						bind:value={editUsername}
						bind:invalidText={invalidUsernameCaption}
					/>
				</div>
			</form>
		</ModalBody>
		<ModalFooter
			primaryButtonText="Edit"
			secondaryButtonText="Cancel"
			on:click:button--secondary={() => {
				isRegisterFormOpen = false;
			}}
		/>
	</ComposedModal>

	<Modal
		passiveModal
		on:close={() => {
			isSuccessRegisterOpen = false;
		}}
		bind:open={isSuccessRegisterOpen}
		modalHeading="Interviewer Registered"
		primaryButtonText="Finish"
	>
		<p>The account has been created and a random password has been generated.</p>
		<br />
		<p class="bold">
			The new user will be prompted to change their password once they try to log in.
		</p>
		<br />
		<p>The generated password is:</p>
		<div class="password-container">
			<CodeSnippet code={newInterviewer.password} />
		</div>
	</Modal>

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
