import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore
import * as fs from "fs";

export default defineConfig({
	plugins: [sveltekit()],
	server: {
		https: {
			key: fs.readFileSync('localhost-key.pem'),
			cert: fs.readFileSync('localhost.pem')
		}
	}

});
