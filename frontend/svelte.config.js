import adapter from "@sveltejs/adapter-static"
import { vitePreprocess } from "@sveltejs/kit/vite"
import path from "path"

/** @type {{preprocess: *, kit: {adapter: Adapter, alias: {"@components": string}}, allowImportingTsExtensions: boolean, serverSideRendering: boolean}} */
const config = {
	// Consult https://kit.svelte.dev/docs/integrations#preprocessors
	// for more information about preprocessors
	preprocess: vitePreprocess(),

	kit: {
		// Adapter-auto only supports some environments, see https://kit.svelte.dev/docs/adapter-auto for a list.
		// If your environment is not supported, or you settled on a specific environment, switch out the adapter.
		// See https://kit.svelte.dev/docs/adapters for more information about adapters.
		adapter: adapter({
			// Default options are shown. On some platforms
			// these options are set automatically — see below
			pages: "build",
			assets: "build",
			fallback: undefined,
			strict: true
		}),
		alias: {
			"@components": path.resolve("./src/components"),
			"@stores": path.resolve("./src/stores")
		}
	},
	serverSideRendering: false,
	allowImportingTsExtensions: true
}
export default config
