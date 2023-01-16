import adapter from '@sveltejs/adapter-static';
import { vitePreprocess } from '@sveltejs/kit/vite';
import { optimizeImports } from 'carbon-preprocess-svelte';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	// Consult https://kit.svelte.dev/docs/integrations#preprocessors
	// for more information about preprocessors
	preprocess: [vitePreprocess(), optimizeImports()],

	kit: {
		adapter: adapter({ out: 'build', fallback: 'index.html' }),
		prerender: { entries: ['*'] }
	}
};

export default config;