<script lang="ts">

	import { onMount } from "svelte"

	const BACKEND_URL = import.meta.env.VITE_BACKEND_URL

	let data: { modes: Array<string> | null, champs: Array<string> | null, types: Array<string> | null}

	onMount(async () => {
		const url = `${BACKEND_URL}/api/data/riot/filters`
		const res = await fetch(url)
		data = await res.json()
		console.log(data)
	})

</script>

