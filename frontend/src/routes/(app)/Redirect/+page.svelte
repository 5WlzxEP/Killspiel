<script lang="ts">
	import { goto } from "$app/navigation"
	import { onMount } from "svelte"

	onMount(() => {
		const loc = window.location.search
		const decoded = decodeURI(loc).substring(1)
		if (decoded === "" || decoded === "/Redirect" || decoded.startsWith("/r.html")) {
			goto("/", { replaceState: true })
			return
		}
		goto(decoded, { replaceState: true })
	})
</script>
