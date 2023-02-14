<script lang="ts">
	import LL from '$lib/i18n/i18n-svelte';
	import { Locale } from '$lib/models/Locale';
	import { Tile, Dropdown, Button, DropdownSkeleton } from 'carbon-components-svelte';
	import type { DropdownItem } from 'carbon-components-svelte/types/Dropdown/Dropdown.svelte';
	import { onMount } from 'svelte';
	import locale from '$lib/stores/localeStore';
	import { setLocale } from '$lib/i18n/i18n-svelte';
	import { loadLocaleAsync } from '$lib/i18n/i18n-util.async';
	import { sendSuccess } from '$lib/util/notifications';

	let languageOptions: ReadonlyArray<DropdownItem>;
	let selectedLanguage = Locale.EN;
	let loading = false;

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

			return { id: value, text: text };
		});
	}

	async function save() {
		loading = true;
		await setupLocale();
		loading = false;

		sendSuccess($LL.general.SUCCESS(), $LL.settings.SUCCESS_TEXT());
	}

	async function setupLocale() {
		locale.set(selectedLanguage);

		await loadLocaleAsync(selectedLanguage);

		setLocale(selectedLanguage);
	}
</script>

<Tile>
	<h1 class="title">{$LL.settings.TITLE()}</h1>
	<br />
	<br />
	<h4>{$LL.settings.LOCAL_SETTINGS()}</h4>

	<div class="bar" />

	<div class="settings-container">
		<h3>{$LL.settings.LANGUAGE()}</h3>
		<div class="settings-element">
			{#if !languageOptions}
				<DropdownSkeleton />
			{:else}
				<Dropdown
					titleText={$LL.settings.PICKER()}
					bind:selectedId={selectedLanguage}
					items={languageOptions}
				/>
			{/if}
		</div>

		<Button on:click={save} bind:disabled={loading}>{$LL.general.SAVE()}</Button>
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
