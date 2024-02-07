<script lang="ts">
	import { onMount } from "svelte"
	import { state } from "@stores/state"
	import { collectionEnd } from "@stores/state.js"
	import { latestVotes } from "@stores/vote"

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
			if (e.data.startsWith("Vote: ")) handleVote(e.data.substring(6))
		}

		websocket.onerror = () => {
			// console.error(e)
		}

		websocket.onclose = () => {
			// console.log("Websocket disconnected, triggering websocket reloading")
			setTimeout(ws, 5000)
		}
	}

	function handleState(s: string) {
		state.set(parseInt(s[0]))
		if (s.length > 2) {
			collectionEnd.set(parseInt(s.substring(3)))
		}
	}

	function handleVote(s: string) {
		try {
			const d: { name: string; vote: number; color: string } = JSON.parse(s)
			d.vote = parseFloat(String(d.vote))
			$latestVotes.d.push({ time: new Date(), ...d })
			if ($latestVotes.d.length > 200)
				$latestVotes.d = $latestVotes.d.slice($latestVotes.d.length - 200)
			else $latestVotes.d = $latestVotes.d
		} catch (e) {
			/* empty */
		}
	}
</script>
