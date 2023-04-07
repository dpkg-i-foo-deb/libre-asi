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
		FileUploader
	} from 'carbon-components-svelte';
	import { ImageLoader } from 'carbon-components-svelte';

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
						<div class="profile-pinture" style="position: relative;">
							<FileUploader
								bind:files
								on:add={handleFileSelect}
								multiple
								labelTitle="Upload files"
								buttonLabel="Add files"
								labelDescription="Only JPEG files are accepted."
								accept={['.jpg', '.jpeg']}
								status="complete"
							/>

							{#if imageUrl}
								<!-- svelte-ignore a11y-img-redundant-alt -->
								<img
									src={imageUrl}
									alt="Profile Picture"
									width="100"
									height="100"
									style="position: absolute; top: 0; right: 0; bottom: 0; left: 0; margin-right: 100px;"
								/>
							{/if}
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
				</div>
			</Tile>
		</div>
	</div>
</main>

<style>
	.other-data {
		margin-top: 3rem;
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

	.profile-pinture {
		float: right;
		margin-top: -70px;
	}
</style>
