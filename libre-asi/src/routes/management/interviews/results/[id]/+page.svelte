<script lang="ts">
	import { page } from '$app/stores';
	import { API_URL, COMPUTE_RESULTS } from '$lib/api/constants';
	import type Interview from '$lib/models/Interview';
	import { fetchWithRefresh } from '$lib/util/fetchRefresh';
	import { handleResponse } from '$lib/util/handleResponse';
	import { InlineLoading } from 'carbon-components-svelte';
	import type { Result } from '$lib/models/Result';
	import { onMount } from 'svelte';

	const id = $page.params.id;

	let result: Result;

	let loading = false;

	onMount(function () {
		computeResults(parseInt(id));
	});

	async function computeResults(id: number) {
		loading = true;

		const interview: Interview = { id: id };

		const response = await fetchWithRefresh(API_URL + COMPUTE_RESULTS, {
			body: JSON.stringify(interview),
			method: 'POST'
		});

		if (response.ok) {
			result = (await response.json()) as Result;

			console.log(result);
		}

		handleResponse(response.status, false);

		loading = false;
	}
</script>

<main>
	{#if loading}
		<InlineLoading description="Computando resultados..." />
	{/if}
</main>

<style></style>
