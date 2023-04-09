<script lang="ts">
	import { page } from '$app/stores';

	let slug: string;

    slug = $page.params.slug
</script>

<main>
	<h1>{slug}</h1>
</main>
