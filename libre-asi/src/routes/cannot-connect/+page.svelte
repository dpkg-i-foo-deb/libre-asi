<script lang="ts">
	import LL from '$lib/i18n/i18n-svelte';
	import { goto } from '$app/navigation';
	import { Button, ExpandableTile } from 'carbon-components-svelte';
	import session from '$lib/stores/userStore';
	import { SessionRole } from '$lib/models/Session';
	import { onMount } from 'svelte';

	//If the user lands here, their session is invalidated
	onMount(function () {
		$session.active = false;
		$session.role = SessionRole.None;
	});
</script>

<h1>{$LL.cannotConnect.ERROR()}</h1>
<h3>
	{$LL.cannotConnect.PARAGRAPH_1()}
</h3>
<h3>{$LL.cannotConnect.PARAGRAPH_2()}</h3>
<h4>{$LL.cannotConnect.CODE()}</h4>
<Button
	on:click={() => {
		goto('/');
	}}>{$LL.cannotConnect.GO_BACK()}</Button
>

<div class="help">
	<ExpandableTile>
		<div slot="above" style="height:3rem">
			{$LL.cannotConnect.IF_ADMIN()}
		</div>
		<div slot="below">
			{$LL.cannotConnect.HELP()}
		</div>
	</ExpandableTile>
</div>

<style>
	h1 {
		font-weight: 500;
		margin-bottom: 15px;
	}

	h3 {
		margin-bottom: 25px;
	}

	.help {
		margin-top: 40px;
	}

	h4 {
		margin-bottom: 30px;
	}
</style>
