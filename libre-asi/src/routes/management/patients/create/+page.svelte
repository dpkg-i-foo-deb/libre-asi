<script lang="ts">
	import {
		DatePicker,
		DatePickerInput,
		TextInput,
		Tile,
		FileUploader,
		Button,
		ProgressIndicator,
		ProgressStep,
		Grid,
		Row,
		Column
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
		if (currentIndex == 0) {
			submitButtonText = 'Registrar paciente';
			currentIndex += 1;
			return;
		}

		if (currentIndex > 1) {
			currentIndex = 0;
			return;
		}

		if (currentIndex < 1) {
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

		<Tile>
			<Grid>
				<Row>
					<Column sm={3} md={4} lg={4}>
						<div class="progress-indicator">
							<ProgressIndicator vertical bind:currentIndex preventChangeOnClick>
								<ProgressStep
									label="Datos mínimos"
									description="Datos mínimos para registrar a un paciente"
									complete={currentIndex > 0}
								/>
								<ProgressStep
									label="Datos extra"
									description="Datos extra para identificar mejor al paciente"
									complete={currentIndex > 1}
								/>
							</ProgressIndicator>
						</div>
					</Column>

					<Column sm={13} md={12} lg={12}>
						{#if currentIndex == 0}
							<div class="basic-data">
								<div class="subtitle">
									<p class="bold">Estos son los datos mínimos para registrar un paciente</p>
									<div
										class="image-container"
										style="display: flex; margin-top: 20px; padding: 70px;"
									>
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
						{/if}

						{#if currentIndex == 1}
							<div class="other-data">
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
						{/if}
					</Column>
				</Row>

				<Row>
					<div class="button-set-container">
						<div class="button">
							<Button style="width:100%" size="lg" kind="ghost">Cancelar</Button>
						</div>
						<div class="button">
							<Button
								style="width:100%"
								size="lg"
								kind="secondary"
								on:click={function () {
									handleBack();
								}}>Atrás</Button
							>
						</div>
						<div class="button">
							<Button
								style="width:100%"
								size="lg"
								on:click={function () {
									handleNext();
								}}>{submitButtonText}</Button
							>
						</div>
					</div>
				</Row>
			</Grid>
		</Tile>
	</div>
</main>

<style>
	.progress-indicator {
		margin-top: 3rem;
	}
	.other-data {
		margin-bottom: 3rem;
	}

	.basic-data {
		margin-bottom: 1rem;
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

	.button-set-container {
		display: flex;
		flex-direction: row;
		width: 100%;
		margin: 0 auto;
	}

	.button {
		width: 100%;
		margin-left: 0;
		margin-right: 0;
	}
</style>
