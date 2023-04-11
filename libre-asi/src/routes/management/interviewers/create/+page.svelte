<script lang="ts">
	import type Interviewer from '$lib/models/Interviewer';
	import { Column, Grid, Row, Tile } from 'carbon-components-svelte';

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

<main />

<style></style>
