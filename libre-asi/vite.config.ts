import { sveltekit } from '@sveltejs/kit/vite';
import type { UserConfig } from 'vite';
import path from 'path'

const config: UserConfig = {
	server: {
		https: {
      key:path.resolve('./key.pem'),
      cert:path.resolve('./cert.pem')
    },
		proxy: {}
	},
	plugins: [sveltekit()],
	test: {
		include: ['src/**/*.{test,spec}.{js,ts}']
	}
};

export default config;
