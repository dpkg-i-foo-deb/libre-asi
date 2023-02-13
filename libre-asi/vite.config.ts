import { sveltekit } from '@sveltejs/kit/vite';
import type { UserConfig } from 'vite';
import mkcert from 'vite-plugin-mkcert';

const config: UserConfig = {
	server: {
		https: true,
		proxy: {}
	},
	plugins: [mkcert(), sveltekit()],
	test: {
		include: ['src/**/*.{test,spec}.{js,ts}']
	}
};

export default config;
