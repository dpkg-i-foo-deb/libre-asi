<script lang="ts">
	import LL from '$lib/i18n/i18n-svelte';
</script>

<h1 class="title">{$LL.welcome.TITLE()}</h1>

<h4 class="subtitle">
	{$LL.welcome.PARAGRAPH_1()}
</h4>

<h7 class="subtitle">{$LL.welcome.PARAGRAPH_2()}</h7>

<p class="text">
	{$LL.welcome.VISIT()}
	<a target="_blank" rel="noreferrer" href="https://github.com/dpkg-i-foo-deb/libre-asi"
		>{$LL.welcome.GITHUB_REPOSITORY()}
	</a>
	{$LL.welcome.KNOW()}
</p>

<style>
	.title {
		font-weight: 400;
	}

	.subtitle,
	h7 {
		margin-top: 20px;
		margin-bottom: 20px;
		font-size: 18pt;
	}

	.text {
		margin-top: 20px;
	}
</style>
