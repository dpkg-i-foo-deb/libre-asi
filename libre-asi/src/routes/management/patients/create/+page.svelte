<script lang="ts">
	import {
		DatePicker,
		DatePickerInput,
		TextInput,
		Tile,
		FileUploader,
		Button,
		ButtonSet,
		Tabs,
		Tab,
		TabContent,
		ProgressIndicator,
		ProgressStep,
		Toggle
	} from 'carbon-components-svelte';
	import { goto } from '$app/navigation';
	import type Patient from '$lib/models/Patient';
	import { API_URL, REGISTER_PATIENT } from '$lib/api/constants';
	import { fetchWithRefresh } from '$lib/util/fetchRefresh';
	import { sendSuccess } from '$lib/util/notifications';
	import { handleResponse } from '$lib/util/handleResponse';

	let file: File | null = null;
	let imageUrl: string | null = null;
	let files: File[] = [];

	let newPatient: Patient = {
		email: '',
		username: '',
		password: '',
		firstName: '',
		lastName: '',
		firstSurname: '',
		lastSurname: '',
		birthdate: new Date(),
		personalID: ''
	};

	function handleFileSelect(event: CustomEvent) {
		const selectedFiles = event.detail;
		if (selectedFiles && selectedFiles.length > 0) {
			file = selectedFiles[0];
			if (file) {
				imageUrl = URL.createObjectURL(file);
			}
		}
		files = [];
	}
	function removeImage() {
		imageUrl = null;
		files = [];
	}

	function handleFileRemove(event: CustomEvent<readonly File[]>) {
		const file = event.detail[0];
		if (file === files[0]) {
			imageUrl = null;
			files = [];
		}
	}

	async function handleRegisterPatient() {
		//TODO add validation

		const response = await fetchWithRefresh(API_URL + REGISTER_PATIENT, {
			method: 'POST',
			body: JSON.stringify(newPatient)
		});

		if (response.ok) {
			sendSuccess('Éxito', 'Paciente registrado exitosamente');
			goto('/management/patients/');
		}

		handleResponse(response.status, false);
	}
</script>

<main>
	<div class="form">
		<div class="title">
			<h1>Creando un nuevo paciente</h1>
		</div>

		<Tabs type="container">
			<Tab label="Datos básicos" />
			<Tab label="Datos extra" />

			<svelte:fragment slot="content">
				<TabContent>
					<div class="basic-data">
						<div class="subtitle">
							<p class="bold">Estos son los datos mínimos para registrar un paciente</p>
							<div class="image-container" style="display: flex; margin-top: 20px; padding: 70px;">
								<div
									class="file-uploader-container"
									style="width: 100px; height: 100px; margin-right: 50px; padding: 5px; position: relative;"
								>
									<FileUploader
										bind:files
										on:add={handleFileSelect}
										on:remove={handleFileRemove}
										multiple={false}
										labelTitle="Upload files"
										buttonLabel="Add Picture"
										labelDescription="Only JPEG files are accepted."
										accept={['.jpg', '.jpeg']}
										status="complete"
										style="font-size: 10px;"
									/>
								</div>

								<div
									class="image-preview-container"
									style="width: 100px; height: 100px; margin-left: 20px; padding: 10px; border: 1px solid black; position: relative;"
								>
									{#if imageUrl}
										<div style="position: relative;">
											<!-- svelte-ignore a11y-img-redundant-alt -->
											<img
												src={imageUrl}
												alt="Profile Picture"
												style="width: 100%; height: 100%; object-fit: cover; margin-bottom: 10px;"
											/>
											<div
												style="position: absolute; bottom: 0; left: 0; right: 0; text-align: center;"
											>
												Profile Picture
											</div>
											{#if imageUrl}
												<div
													style="position: absolute; bottom: -50px; left: 0; right: 0; text-align: center;"
												>
													<Button kind="danger" on:click={removeImage} style="margin: auto;"
														>Remove</Button
													>
												</div>
											{/if}
										</div>
									{/if}
								</div>
							</div>
							<div class="text-field">
								<TextInput
									id="firstName"
									labelText="Primer Nombre"
									placeholder="Ingresa el primer nombre"
									bind:value={newPatient.firstName}
								/>
							</div>

							<div class="text-field">
								<TextInput
									id="id"
									labelText="Número de Identificación"
									placeholder="Ingresa el número de identificación"
									bind:value={newPatient.personalID}
								/>
							</div>

							<div class="text-field">
								<TextInput
									style="margin-bottom: 3rem"
									id="email"
									type="email"
									labelText="Correo Electrónico"
									placeholder="Ingresa el correo Electrónico"
									bind:value={newPatient.email}
								/>
							</div>
						</div>
					</div>
				</TabContent>
				<TabContent
					><div class="other-data">
						<div class="subtitle">
							<p class="bold">
								Datos extra que pueden servir para una mejor identificación del paciente
							</p>
						</div>
						<div class="text-field">
							<TextInput
								id="secondname"
								labelText="Segundo Nombre"
								placeholder="Ingresa el segundo nombre"
								bind:value={newPatient.lastName}
							/>
						</div>

						<div class="text-field">
							<TextInput
								id="firstSurname"
								labelText="Primer Apellido"
								placeholder="Ingresa el primer apellido"
								bind:value={newPatient.firstSurname}
							/>
						</div>

						<div class="text-field">
							<TextInput
								id="secondSurname"
								labelText="Segundo Apellido"
								placeholder="Ingrese el segundo apellido"
								bind:value={newPatient.lastSurname}
							/>
						</div>

						<div class="text-field">
							<DatePicker datePickerType="single" dateFormat="d/m/Y" on:change>
								<DatePickerInput
									labelText="Fecha de nacimiento"
									placeholder="dd/mm/aaaa"
									bind:value={newPatient.birthdate}
								/>
							</DatePicker>
						</div>
					</div>
				</TabContent>
			</svelte:fragment>
		</Tabs>
		<div id="register-button">
			<ButtonSet>
				<Button kind="secondary">Cancelar</Button>
				<Button
					href=""
					on:click={() => {
						handleRegisterPatient();
					}}>Registrar Paciente</Button
				>
			</ButtonSet>
		</div>
	</div>
</main>

<style>
	.other-data {
		margin-bottom: 3rem;
	}

	.basic-data {
		margin-bottom: 3rem;
	}

	.title {
		margin: 0 0 1rem 0;
	}

	.subtitle {
		margin-top: 0.5rem;
	}

	.bold {
		font-weight: bold;
	}

	.text-field {
		margin-top: 3rem;
	}

	.form {
		grid-column: 2 / 6;
		grid-row: 1 / 3;
		margin-right: 10px;
	}

	#register-button {
		margin-bottom: 20px;
		text-align: right;
	}
</style>
