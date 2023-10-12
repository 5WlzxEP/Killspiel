<script lang="ts">
	import "../../app.postcss"
	import { onMount } from "svelte"
	import { get } from "svelte/store"
	import { themeStore } from "@stores/theme"
	import Websocket from "../Websocket.svelte"
	import { setModeCurrent } from "@skeletonlabs/skeleton"

	onMount(() => {
		const params = new URLSearchParams(window.location.search)

		// theme
		const theme = params.get("theme")

		if (
			theme &&
			[
				"skeleton",
				"modern",
				"rocket",
				"seafoam",
				"vintage",
				"sahara",
				"hamlindigo",
				"gold-nouveau",
				"crimson",
				"wintry"
			].includes(theme)
		)
			document.getElementsByTagName("body")[0].dataset.theme = theme
		else document.getElementsByTagName("body")[0].dataset.theme = get(themeStore)

		// darkmode
		const dark = params.get("dark")
		setModeCurrent(<boolean>(dark && dark === "0"))
	})
</script>

<slot />
<Websocket />
