<script lang="ts">
	import { page } from '$app/stores';
	import { API_URL, GET_PATIENT } from '$lib/api/constants';
	import type Patient from '$lib/models/Patient';
	import { fetchWithRefresh } from '$lib/util/fetchRefresh';
	import { handleResponse } from '$lib/util/handleResponse';
	import { onMount } from 'svelte';

	let id = $page.params.id;

	let patient: Patient;

	onMount(async function () {
		loadPatient();
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
</script>

<main>
	<h1>{id}</h1>
</main>
