<script lang="ts">
	import LL from '$lib/i18n/i18n-svelte';
	import { Locale } from '$lib/models/Locale';
	import {
		Tile,
		Dropdown,
		Button,
		DropdownSkeleton,
		RadioButtonGroup,
		Theme,
		RadioButton
	} from 'carbon-components-svelte';

	import type { DropdownItem } from 'carbon-components-svelte/types/Dropdown/Dropdown.svelte';
	import { onMount } from 'svelte';
	import locale from '$lib/stores/localeStore';
	import { setLocale } from '$lib/i18n/i18n-svelte';
	import { loadLocaleAsync } from '$lib/i18n/i18n-util.async';
	import { sendSuccess } from '$lib/util/notifications';
	import type { CarbonTheme } from 'carbon-components-svelte/types/Theme/Theme.svelte';
	import settings from '$lib/i18n/en/settings';
	import { POST } from '../api/administrators/+server';
	import { handleResponse } from '$lib/util/handleResponse';

	let theme: CarbonTheme = 'g90';
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
		const response = await fetch('/api/lang', {
			method: 'POST',
			body: JSON.stringify(selectedLanguage)
		});

		handleResponse(response.status, false);

		//locale.set(selectedLanguage);

		//await loadLocaleAsync(selectedLanguage);

		//setLocale(selectedLanguage);
	}
</script>

<Theme bind:theme persist persistKey="__carbon-theme" />

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

		<h3>{$LL.settings.THEME()}</h3>
		<div class="settings-element">
			<RadioButtonGroup legendText={$LL.settings.THEME_PICKER()} bind:selected={theme}>
				<RadioButton labelText={$LL.settings.WHITE()} value="white" />
				<RadioButton labelText={$LL.settings.G10()} value="g10" />
				<RadioButton labelText={$LL.settings.G80()} value="g80" />
				<RadioButton labelText={$LL.settings.G90()} value="g90" />
				<RadioButton labelText={$LL.settings.G100()} value="g100" />
			</RadioButtonGroup>
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
