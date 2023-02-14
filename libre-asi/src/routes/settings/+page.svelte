<script lang="ts">
	import { Locale } from '$lib/models/Locale';
	import { Tile, Dropdown, Button } from 'carbon-components-svelte';
	import type { DropdownItem } from 'carbon-components-svelte/types/Dropdown/Dropdown.svelte';
	import { onMount } from 'svelte';

	let languageOptions: ReadonlyArray<DropdownItem>;
	let selectedLanguage = 0;

	onMount(function () {
		loadLanguages();
	});

	function loadLanguages() {
		const languages = Object.values(Locale);

		languageOptions = languages.map(function (value, index) {
			let text: string = '';

			switch (value) {
				case Locale.EN:
					text = 'English';
					break;
				case Locale.ES:
					text = 'Español';
					break;
				default:
					'Undefined';
			}

			return { id: index, text: text };
		});
	}

	function save() {
		console.log('owo');
	}
</script>

<Tile>
	<h1 class="title">Libre-ASI Settings</h1>
	<br />
	<br />
	<h4>Settings will only be applied on this browser</h4>

	<div class="bar" />

	<div class="settings-container">
		<h3>Language</h3>
		<div class="settings-element">
			<Dropdown
				titleText="Pick a language"
				bind:selectedId={selectedLanguage}
				items={languageOptions}
			/>
		</div>

		<Button on:click={save}>Save</Button>
	</div>
</Tile>

<style>
	.title {
		font-weight: 400;
	}

	/*https://github.com/carbon-design-system/carbon-components-svelte/blob/master/docs/src/layouts/ComponentLayout.svelte*/
	.bar {
		display: flex;
		justify-content: space-between;
		margin-bottom: var(--cds-layout-02);
		border-bottom: 1px solid var(--cds-ui-03);
		margin-top: 20px;
	}

	.settings-container {
		margin-top: 10px;
	}

	.settings-element {
		margin-top: 30px;
		margin-bottom: 30px;
	}
</style>
