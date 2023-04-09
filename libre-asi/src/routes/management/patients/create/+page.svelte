<script lang="ts">
	import {
		Column,
		DatePicker,
		DatePickerInput,
		Grid,
		Row,
		TextInput,
		Tile,
		Breadcrumb,
		Toggle,
		FileUploader,
		Button,
		ButtonSet
	} from 'carbon-components-svelte';
	import { ImageLoader } from 'carbon-components-svelte';
	import { goto } from '$app/navigation';

	let showInfo = false;
	let file: File | null = null;
	let imageUrl: string | null = null;
	let files: File[] = [];

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

	function handleRegisterPatient() {
		goto('/management/patients/');
	}
</script>

<main>
	<div class="wrapper">
		<div class="toggle">
			<Toggle labelText="Mostrar información adicional" bind:toggled={showInfo} />
		</div>

		<div class="form">
			<div class="title">
				<h1>Crear un nuevo paciente</h1>
			</div>

			<Tile>
				<div class="basic-data">
					<h3>Datos básicos</h3>
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
							/>
						</div>

						<div class="text-field">
							<TextInput
								id="id"
								labelText="Número de Identificación"
								placeholder="Ingresa el número de identificación"
							/>
						</div>

						<div class="text-field">
							<TextInput
								style="margin-bottom: 3rem"
								id="email"
								type="email"
								labelText="Correo Electrónico"
								placeholder="Ingresa el correo Electrónico"
							/>
						</div>
					</div>

					{#if showInfo}
						<div class="other-data">
							<h3>Otros datos</h3>
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
								/>
							</div>

							<div class="text-field">
								<TextInput
									id="firstSurname"
									labelText="Primer Apellido"
									placeholder="Ingresa el primer apellido"
								/>
							</div>

							<div class="text-field">
								<TextInput
									id="secondSurname"
									labelText="Segundo Apellido"
									placeholder="Ingrese el segundo apellido"
								/>
							</div>

							<div class="text-field">
								<DatePicker datePickerType="single" dateFormat="d/m/Y" on:change>
									<DatePickerInput labelText="Fecha de nacimiento" placeholder="dd/mm/aaaa" />
								</DatePicker>
							</div>
							<div class="text-field">
								<TextInput
									id="phoneNumber"
									labelText="Teléfono"
									placeholder="Ingrese el teléfono"
								/>
							</div>
						</div>
					{/if}
					<div id="register-button">
						<ButtonSet>
							<Button kind="secondary">Cancel</Button>
							<Button
								href=""
								on:click={() => {
									handleRegisterPatient();
								}}>Register Patient</Button
							>
						</ButtonSet>
					</div>
				</div>
			</Tile>
		</div>
	</div>
</main>

<style>
	.other-data {
		margin-top: 3rem;
		margin-bottom: 5rem;
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

	.wrapper {
		display: grid;
		grid-template-columns: 10% 75%;
		grid-gap: 50px;
		grid-auto-rows: minmax(100px, auto);
	}

	.toggle {
		grid-column: 1;
		grid-row: 1 / 5;
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
