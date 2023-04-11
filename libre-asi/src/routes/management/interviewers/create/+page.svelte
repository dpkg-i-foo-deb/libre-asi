<script lang="ts">
	import type Interviewer from '$lib/models/Interviewer';
	import {
		Button,
		DatePicker,
		DatePickerInput,
		FileUploader,
		ProgressIndicator,
		ProgressStep,
		TextInput,
		Tile
	} from 'carbon-components-svelte';

	let currentIndex = 0;

	let file: File | null = null;
	let imageUrl: string | null = null;
	let files: File[] = [];

	let newInterviewer: Interviewer = {
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

	let submitButtonText = 'Siguiente';

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
</script>

<main>
	<div class="form">
		<div class="title">
			<h1>Creando un nuevo entrevistador</h1>
		</div>
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
					<p class="bold">Estos son los datos mínimos para crear un entrevistador</p>
				</div>

				<div class="text-field">
					<TextInput
						id="firstName"
						labelText="Primer Nombre"
						placeholder="Ingresa el primer nombre"
						bind:value={newInterviewer.firstName}
					/>
				</div>

				<div class="text-field">
					<TextInput
						id="id"
						labelText="Número de identificación"
						placeholder="Ingresa el  número de identificación"
						bind:value={newInterviewer.personalID}
					/>
				</div>

				<div class="text-field">
					<TextInput
						id="email"
						type="email"
						labelText="Correo electrónico"
						bind:value={newInterviewer.email}
					/>
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
									<Button kind="danger" on:click={removeImage} style="margin: auto;">Remove</Button>
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
					<p class="bold">Datos extra para identificar el entrevistador</p>
				</div>

				<div class="text-field">
					<TextInput
						id="lastName"
						labelText="Segundo Nombre"
						placeholder="Ingresa el segundo nombre"
						bind:value={newInterviewer.lastName}
					/>
				</div>

				<div class="text-field">
					<TextInput
						id="firstSurname"
						labelText="Primer apellido"
						placeholder="Ingresa el primer apellido"
						bind:value={newInterviewer.firstSurname}
					/>
				</div>

				<div class="text-field">
					<TextInput
						id="lastSurname"
						labelText="Segundo apellido"
						placeholder="Ingrese el segundo apellido"
						value={newInterviewer.lastSurname}
					/>
				</div>

				<div class="text-field">
					<DatePicker datePickerType="single">
						<DatePickerInput
							labelText="Fecha de nacimiento"
							placeholder="dd/mm/aaaa"
							bind:value={newInterviewer.birthdate}
						/>
					</DatePicker>
				</div>
			</div>
		{/if}

		<div class="button-set-container">
			<Button size="lg" kind="ghost">Cancelar</Button>

			<Button size="lg" kind="secondary">Atrás</Button>

			<Button size="lg">{submitButtonText}</Button>
		</div>
	</Tile>
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
