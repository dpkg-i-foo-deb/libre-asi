<script lang="ts">
	import { Button, Tile } from 'carbon-components-svelte';
</script>

<main>
	<Tile>
		<h1>El servidor no está respondiendo.</h1>
		<h2>Intenta de nuevo en unos minutos.</h2>
		<h4>Código de estado: 504.</h4>

		<Button
			size="lg"
			on:click={function () {
				history.back();
			}}>Intentar nuevamente</Button
		>
	</Tile>
</main>

<style>
	h1 {
		font-weight: 500;
		margin-bottom: 15px;
	}

	h4 {
		margin-top: 20px;
		margin-bottom: 30px;
	}
</style>
