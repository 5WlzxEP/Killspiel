<script lang="ts">
	import { onMount } from "svelte"
	import { state } from "@stores/state"

	const BACKEND_URL: string = import.meta.env.VITE_BACKEND_URL

	// case Release
	let url: string

	onMount(() => {
		url = `ws://${window.location.host}/ws`
		// case local testing
		if (BACKEND_URL.startsWith("http://")) url = `${BACKEND_URL.replace("http", "ws")}/ws`
		ws()
	})

	let websocket: WebSocket

	function ws() {
		websocket = new WebSocket(url)

		websocket.onmessage = (e: MessageEvent<string>) => {
			if (e.data === "Ping") return

			if (e.data.startsWith("State: ")) handleState(e.data.substring(7))
		}

		websocket.onerror = (e) => {
			// console.error(e)
		}

		websocket.onclose = () => {
			// console.log("Websocket disconnected, triggering websocket reloading")
			setTimeout(ws, 5000)
		}
	}

	function handleState(s: string) {
		state.set(parseInt(s))
	}
</script>
