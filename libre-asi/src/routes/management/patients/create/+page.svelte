<script lang="ts">
	import {
		DatePicker,
		DatePickerInput,
		TextInput,
		Tile,
		FileUploader,
		Button,
		ProgressIndicator,
		ProgressStep
	} from 'carbon-components-svelte';
	import { goto } from '$app/navigation';
	import type Patient from '$lib/models/Patient';
	import { API_URL, REGISTER_PATIENT } from '$lib/api/constants';
	import { fetchWithRefresh } from '$lib/util/fetchRefresh';
	import { sendSuccess } from '$lib/util/notifications';
	import { handleResponse } from '$lib/util/handleResponse';

	let submitButtonText = 'Siguiente';

	let file: File | null = null;
	let imageUrl: string | null = null;
	let files: File[] = [];
	let currentIndex = 0;

	let newPatient: Patient = {
		email: '',
		username: '',
		password: '',
		firstName: '',
		lastName: '',
		firstSurname: '',
		lastSurname: '',
		birthdate: new Date(),
		personalID: '',
		age: 0
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

	async function handleNext() {
		if (currentIndex > 2) {
			currentIndex = 0;
			return;
		}

		if (currentIndex == 1 && submitButtonText == 'Siguiente') {
			currentIndex += 1;
			submitButtonText = 'Registrar paciente';
			return;
		}

		if (currentIndex < 2) {
			currentIndex += 1;
			return;
		}

		if (currentIndex == 1) {
			handleRegisterPatient();
		}
	}

	function handleBack() {
		if (currentIndex > 0 && submitButtonText == 'Registrar paciente') {
			currentIndex -= 1;
			submitButtonText = 'Siguiente';
			return;
		}

		if (currentIndex > 0) {
			currentIndex -= 1;
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

		<div class="stepper">
			<ProgressIndicator spaceEqually preventChangeOnClick bind:currentIndex>
				<ProgressStep complete={currentIndex > 0} label="Datos mínimos" />
				<ProgressStep complete={currentIndex > 1} label="Foto" />
				<ProgressStep complete={currentIndex > 2} label="Datos extra" />
			</ProgressIndicator>
		</div>

		<Tile style="padding-bottom:0; padding-left:0; padding-right:0;">
			{#if currentIndex == 0}
				<div class="basic-data">
					<div class="subtitle">
						<p class="bold">Estos son los datos mínimos para registrar un paciente</p>

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
								id="email"
								type="email"
								labelText="Correo electrónico"
								placeholder="Ingresa el correo electrónico"
								bind:value={newPatient.email}
							/>
						</div>
					</div>
				</div>
			{/if}

			{#if currentIndex == 1}
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
								<div style="position: absolute; bottom: 0; left: 0; right: 0; text-align: center;">
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
			{/if}

			{#if currentIndex == 2}
				<div class="other-data">
					<div class="subtitle">
						<p class="bold">
							Datos extra que pueden servir para una mejor identificación del paciente
						</p>
					</div>
					<div class="text-field">
						<TextInput
							id="lastName"
							labelText="Segundo nombre"
							placeholder="Ingresa el segundo nombre"
							bind:value={newPatient.lastName}
						/>
					</div>

					<div class="text-field">
						<TextInput
							id="firstSurname"
							labelText="Primer apellido"
							placeholder="Ingresa el primer apellido"
							bind:value={newPatient.firstSurname}
						/>
					</div>

					<div class="text-field">
						<TextInput
							id="lastSurname"
							labelText="Segundo apellido"
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
			{/if}

			<div class="button-set-container">
				<Button size="lg" kind="ghost">Cancelar</Button>

				<Button
					size="lg"
					kind="secondary"
					on:click={function () {
						handleBack();
					}}>Atrás</Button
				>

				<Button
					size="lg"
					on:click={function () {
						handleNext();
					}}>{submitButtonText}</Button
				>
			</div></Tile
		>
	</div>
</main>

<style>
	.other-data {
		margin-left: 1rem;
		margin-right: 1rem;
	}

	.basic-data {
		margin-left: 1rem;
		margin-right: 1rem;
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
		margin-top: 2rem;
		margin-bottom: 2rem;
	}

	.button-set-container {
		display: flex;
		align-items: flex-end;
		justify-content: flex-end;
	}

	.stepper {
		margin-top: 1rem;
		margin-bottom: 1rem;
	}
</style>
