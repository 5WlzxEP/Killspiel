<script lang="ts">
	import { onMount } from "svelte"
	const BACKEND_URL = import.meta.env.VITE_BACKEND_URL
	onMount(() => {
		const h = window.location.hash.substring(1)

		window.location.replace(`${BACKEND_URL}/api/collector/chat/oauth?${encodeURI(h)}`)
	})
</script>

{window.location.hash}
