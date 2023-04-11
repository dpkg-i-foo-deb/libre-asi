<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { API_URL, EDIT_PATIENT, GET_PATIENT } from '$lib/api/constants';
	import type Patient from '$lib/models/Patient';
	import { fetchWithRefresh } from '$lib/util/fetchRefresh';
	import { handleResponse } from '$lib/util/handleResponse';
	import { sendSuccess } from '$lib/util/notifications';
	import {
		Button,
		DatePicker,
		DatePickerInput,
		FileUploader,
		Tab,
		TabContent,
		Tabs,
		TextInput,
		Tile
	} from 'carbon-components-svelte';
	import { onMount } from 'svelte';

	let id = $page.params.id;

	let files: File[] = [];

	let imageUrl: string;

	let patient: Patient = {
		firstName: '',
		lastName: '',
		email: '',
		firstSurname: '',
		lastSurname: '',
		birthdate: new Date(),
		username: '',
		password: '',
		personalID: '',
		age: 0
	};

	onMount(async function () {
		await loadPatient();
	});

	async function loadPatient() {
		const response = await fetchWithRefresh(API_URL + GET_PATIENT + id, {
			method: 'GET'
		});

		if (response.ok) {
			patient = (await response.json()) as Patient;
			console.log(patient);
		}

		handleResponse(response.status, false);
	}

	async function handleSave() {
		const response = await fetchWithRefresh(API_URL + EDIT_PATIENT, {
			method: 'PATCH',
			body: JSON.stringify(patient)
		});

		if (response.ok) {
			sendSuccess('Éxito', 'Paciente editado correctamente, puede volver a la página anterior');
		}

		handleResponse(response.status, false);
	}

	function handleFileSelect() {}

	function handleFileRemove() {}

	function removeImage() {}
</script>

<main>
	<div class="form">
		<div class="title">
			<h1>Editando paciente</h1>
		</div>

		<Tabs autoWidth>
			<Tab label="Datos mínimos" />
			<Tab label="Foto" />
			<Tab label="Datos extra" />

			<svelte:fragment slot="content">
				<TabContent style="padding:0;">
					<Tile style="padding-bottom:0; padding-left:0; padding-right:0;">
						<div class="content-container">
							<div class="attributes">
								<div class="basic-data">
									<div class="subtitle">
										<p class="bold">Estos son los datos mínimos para registrar un paciente</p>
										<div class="text-field">
											<TextInput
												id="firstName"
												labelText="Primer Nombre"
												placeholder="Ingresa el primer nombre"
												bind:value={patient.firstName}
											/>
										</div>

										<div class="text-field">
											<TextInput
												id="id"
												labelText="Número de Identificación"
												placeholder="Ingresa el número de identificación"
												bind:value={patient.personalID}
											/>
										</div>

										<div class="text-field">
											<TextInput
												id="email"
												type="email"
												labelText="Correo Electrónico"
												placeholder="Ingresa el correo Electrónico"
												bind:value={patient.email}
											/>
										</div>
									</div>
								</div>
							</div>
						</div>

						<div class="buttons-container">
							<Button
								kind="ghost"
								size="lg"
								on:click={function () {
									goto('/management/patients');
								}}>Cancelar</Button
							>
							<Button size="lg" on:click={handleSave}>Guardar</Button>
						</div>
					</Tile>
				</TabContent>

				<TabContent style="padding:0">
					<Tile style="padding-bottom:0; padding-left:0; padding-right:0;">
						<div class="content-container">
							<div class="other-data">
								<div class="subtitle">
									<p class="bold">Establece una foto del paciente</p>

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
								</div>
							</div>
						</div>

						<div class="buttons-container">
							<Button
								kind="ghost"
								size="lg"
								on:click={function () {
									goto('/management/patients');
								}}>Cancelar</Button
							>
							<Button size="lg" on:click={handleSave}>Guardar</Button>
						</div>
					</Tile>
				</TabContent>

				<TabContent style="padding:0;"
					><Tile style="padding-bottom:0; padding-left:0; padding-right:0;">
						<div class="content-container">
							<div class="attributes">
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
											bind:value={patient.lastName}
										/>
									</div>

									<div class="text-field">
										<TextInput
											id="firstSurname"
											labelText="Primer Apellido"
											placeholder="Ingresa el primer apellido"
											bind:value={patient.firstSurname}
										/>
									</div>

									<div class="text-field">
										<TextInput
											id="secondSurname"
											labelText="Segundo Apellido"
											placeholder="Ingrese el segundo apellido"
											bind:value={patient.lastSurname}
										/>
									</div>

									<div class="text-field">
										<DatePicker datePickerType="single" dateFormat="d/m/Y" on:change>
											<DatePickerInput
												labelText="Fecha de nacimiento"
												placeholder="dd/mm/aaaa"
												bind:value={patient.birthdate}
											/>
										</DatePicker>
									</div>
								</div>
							</div>
						</div>

						<div class="buttons-container">
							<Button
								kind="ghost"
								size="lg"
								on:click={function () {
									goto('/management/patients');
								}}>Cancelar</Button
							>
							<Button size="lg" on:click={handleSave}>Guardar</Button>
						</div>
					</Tile></TabContent
				>
			</svelte:fragment>
		</Tabs>
	</div>
</main>

<style>
	.attributes {
		margin-top: 2rem;
	}

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

	.content-container {
		display: block;
		padding-left: 1rem;
		padding-right: 1rem;
	}

	.buttons-container {
		display: flex;
		flex-direction: row;
		width: 100%;
		align-items: right;
		justify-content: right;
		margin-left: 0;
		margin-right: 0;
	}
</style>
