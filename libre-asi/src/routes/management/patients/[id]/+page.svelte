<script lang="ts">
	import { page } from '$app/stores';

	let id: string;

	id = $page.params.id;
</script>

<main>
	<h1>{id}</h1>
</main>
