<script lang="ts">
	import type { Question } from '$lib/models/Question';
	import { RadioButton, RadioButtonGroup } from 'carbon-components-svelte';

	export let question: Question;
</script>

<main>
	<RadioButtonGroup>
		{#each question.options ?? [] as option}
			<RadioButton labelText={option.description} />
		{/each}
	</RadioButtonGroup>
</main>

<style>
</style>
