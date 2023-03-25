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
	import type Administrator from '$lib/models/Administrator';
	import type { DataTableRow } from 'carbon-components-svelte/types/DataTable/DataTable.svelte';
	import { checkEmail, checkUsername } from '$lib/util/formUtils';
	import { handleResponse } from '$lib/util/handleResponse';
	import {
		API_URL,
		GET_ADMINS,
		REGISTER_ADMIN,
		EDIT_ADMINS,
		DELETE_ADMIN
	} from '$lib/api/constants';
	import { fetchWithRefresh } from '$lib/util/fetchRefresh';
	import { sendError, sendSuccess } from '$lib/util/notifications';
	let newAdministrator: Administrator;
	let editedAdministrator: Administrator;
	let rows: ReadonlyArray<DataTableRow>;
	let filteredRows: ReadonlyArray<DataTableRow>;
	let isRegisterFormOpen = false;
	let isSuccessRegisterOpen = false;
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
		newAdministrator = {
			ID: 0,
			CreatedAt: new Date(),
			UpdatedAt: new Date(),
			email: '',
			username: '',
			password: ''
		};

		await loadAdmins();
	});

	async function loadAdmins() {
		const response = await fetchWithRefresh(API_URL + GET_ADMINS, { method: 'GET' });
		if (response.ok) {
			const existingAdmins = (await response.json()) as Administrator[];

			rows = existingAdmins.map(function (value: Administrator) {
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

	async function register() {
		duplicateCredentials = false;
		isSuccessRegisterOpen = false;

		if (!validateEmail() || !validateUsername()) {
			return;
		}

		newAdministrator = {
			ID: 0,
			CreatedAt: new Date(),
			UpdatedAt: new Date(),
			email: email,
			username: username,
			password: ''
		};

		const response = await fetchWithRefresh(API_URL + REGISTER_ADMIN, {
			method: 'POST',
			body: JSON.stringify(newAdministrator)
		});

		if (response.ok) {
			newAdministrator = (await response.json()) as Administrator;
			isRegisterFormOpen = false;
			isSuccessRegisterOpen = true;
			loadAdmins();
			return;
		}

		if (handleResponse(response.status, false)) {
			return;
		}

		if (response.status == 409) {
			duplicateCredentials = true;
		}
	}

	function handleSearch() {
		const query = searchValue;

		if (searchValue == '') {
			filteredRows = rows;
			return;
		}

		filteredRows = rows.filter((row) => {
			return row.email.toLocaleLowerCase().includes(query) || row.username.includes(query);
		});
	}

	function toggleEditForm() {
		isEditionFormOpen = true;

		let row = rows.find((row) => row.id === editingId);

		if (row) {
			email = row.email;
			username = row.username;
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

		editedAdministrator = {
			ID: editingId,
			CreatedAt: new Date(),
			UpdatedAt: new Date(),

			email: email,
			username: username,
			password: ''
		};

		const response = await fetchWithRefresh(API_URL + EDIT_ADMINS, {
			method: 'PATCH',
			body: JSON.stringify(editedAdministrator)
		});

		if (response.ok) {
			editedAdministrator = (await response.json()) as Administrator;
			isEditionFormOpen = false;
			loadAdmins();
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
		const response = await fetchWithRefresh(API_URL + DELETE_ADMIN + deletingId.toString(), {
			method: 'DELETE'
		});

		if (response.ok) {
			loadAdmins();
			sendSuccess('Account deleted successfully', '');
			return;
		}

		if (handleResponse(response.status, false)) {
			sendError('Error trying to remove administrator', '');
			return;
		}
	}
</script>

{#if rows == undefined}
	<DataTableSkeleton
		headers={[
			{ key: 'email', value: 'Email' },
			{ key: 'username', value: 'Username' }
		]}
		rows={5}
	/>
{:else}
	<DataTable
		title="Administrators"
		description="Current Registered Administrators"
		headers={[
			{ key: 'email', value: 'Email' },
			{ key: 'username', value: 'Username' },
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
					}}>Register Administrator</Button
				>
			</ToolbarContent>
		</Toolbar>
	</DataTable>

	<ComposedModal bind:open={isRegisterFormOpen} selectorPrimaryFocus="#email" on:submit={register}>
		<ModalHeader label="Transaction" title="New Administrator Account" />
		<ModalBody hasForm>
			<form>
				<h7 class="paragraph">
					Administrators can register other administrators, they can also trigger password resets,
					access interviews and accounts information.
				</h7>
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
				<h7 class="paragraph"> Administrators can edit other administrators' information. </h7>
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
		modalHeading="Administrator Registered"
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
			<CodeSnippet code={newAdministrator.password} />
		</div>
	</Modal>

	<Modal
		danger
		bind:open={isModalOpen}
		modalHeading="Delete Administrator"
		primaryButtonText="Delete"
		secondaryButtonText="Cancel"
		on:submit={handleDelete}
		on:click:button--secondary={handleCancel}
	>
		<p>Are you sure you want to delete this Administrator?</p>
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
