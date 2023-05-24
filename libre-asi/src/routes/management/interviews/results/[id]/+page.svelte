<script lang="ts">
	import { page } from '$app/stores';
	import { API_URL, COMPUTE_RESULTS } from '$lib/api/constants';
	import type Interview from '$lib/models/Interview';
	import { fetchWithRefresh } from '$lib/util/fetchRefresh';
	import { handleResponse } from '$lib/util/handleResponse';
	import {
		DataTable,
		DataTableSkeleton,
		InlineLoading,
		SkeletonText
	} from 'carbon-components-svelte';
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
	{#if loading || !result}
		<div class="load">
			<InlineLoading description="Computando resultados..." />
		</div>

		<div class="title">
			<SkeletonText heading />
		</div>

		<DataTableSkeleton showHeader={false} showToolbar={false} />
	{:else}
		<h1 class="title">Resultados</h1>

		<DataTable
			headers={[
				{ key: 'scale', value: 'Índices de Severidad' },
				{ key: 'value', value: 'Valor' }
			]}
			rows={[
				{
					id: 'med',
					scale: 'Médico',
					value: result.medicalScale
				},
				{
					id: 'emp',
					scale: 'Empleo',
					value: result.employmentScale
				},
				{
					id: 'alc',
					scale: 'Alcohol',
					value: result.alcoholScale
				},
				{
					id: 'dru',
					scale: 'Drogas',
					value: result.drugScale
				},
				{
					id: 'law',
					scale: 'Legal',
					value: result.legalScale
				},
				{
					id: 'famSupport',
					scale: 'Familia (Apoyo)',
					value: result.familySocialSupportScale
				},
				{
					id: 'famChild',
					scale: 'Familia (Niños)',
					value: result.familyChildScale
				},
				{
					id: 'famProb',
					scale: 'Familia (Problemas)',
					value: result.familySocialProblemScale
				},
				{
					id: 'psy',
					scale: 'Psicológico',
					value: result.psychScale
				}
			]}
		/>
	{/if}
</main>

<style>
	.load {
		margin-bottom: 2rem;
	}

	.title {
		margin-top: 1.5rem;
		margin-bottom: 1.5rem;
	}
</style>
