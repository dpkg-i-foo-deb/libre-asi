<script lang="ts">
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

<h1>Error.</h1>
<h3>
	Something went wrong while processing your request, try again later, if the error persists.
	Contact your administrator.
</h3>
<h3>For your security, your session has been invalidated.</h3>
<h4>Code: 503.</h4>
<Button
	on:click={() => {
		goto('/');
	}}>Go back to welcome page</Button
>

<div class="help">
	<ExpandableTile>
		<div slot="above" style="height:3rem">
			If you're the administrator... Click here to see some help
		</div>
		<div slot="below">
			Status code 503 is thrown when Libre-ASI API refused the connection, this can be caused by
			either misconfiguring environment variables or stopping the API server. Make sure it is
			running, check server logs and its address is correctly configured on the server and try again
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
