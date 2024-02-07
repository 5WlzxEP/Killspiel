<script lang="ts">
	import { collectionEnd, printState, state } from "@stores/state"
	import { onDestroy, onMount } from "svelte"
	import CurrentVotes from "@components/CurrentVotes.svelte"
	const BACKEND_URL = import.meta.env.VITE_BACKEND_URL

	async function start() {
		try {
			await fetch(`${BACKEND_URL}/api/data/`, {
				headers: {
					"Content-Type": "application/json"
				},
				method: "POST",
				body: JSON.stringify({
					state: 0
				})
			})
		} catch (e) {
			console.error(e)
		}
	}

	let result: number
	async function stop() {
		clearInterval(interval)
		try {
			await fetch(`${BACKEND_URL}/api/data/`, {
				headers: {
					"Content-Type": "application/json"
				},
				method: "POST",
				body: JSON.stringify({
					state: 1
				})
			})
		} catch (e) {
			console.error(e)
		}
	}

	async function aufloesen() {
		try {
			await fetch(`${BACKEND_URL}/api/data/`, {
				headers: {
					"Content-Type": "application/json"
				},
				method: "POST",
				body: JSON.stringify({
					state: 2,
					result: result
				})
			})
		} catch (e) {
			console.error(e)
		}
	}

	onMount(async () => {
		const url = `${BACKEND_URL}/api/data/`
		try {
			const res = await fetch(url)
			const val = await res.json()
			$state = val.state
			if (val.end !== undefined) {
				$collectionEnd = val.end
				printTime()
			}
		} catch (e) {
			console.error(e)
		}
	})

	$: if ($state === 1) printTime()

	let time: string = ""
	// eslint-disable-next-line no-undef
	let interval: NodeJS.Timeout
	function printTime() {
		const d = new Date($collectionEnd * 1000)
		interval = setInterval(() => {
			const d2 = new Date()
			const diff = d - d2
			if (diff < 0) {
				clearInterval(interval)
				time = ""
			} else {
				time = new Date(diff).toLocaleString("de-DE", { minute: "2-digit", second: "2-digit" })
			}
		}, 1000)
	}

	onDestroy(() => {
		clearInterval(interval)
	})
</script>

<div class="w-full h-full overflow-y-hidden">
	<h1 class="text-lg text-center anchor">
		<a href="/settings/data/manual/">Manuell & Status</a>
	</h1>
	<div class="text-center card p-1">
		Aktueller Status: <h3 class="text-lg">
			{printState($state)}
			{#if $state === 1}
				{time}
			{/if}
		</h3>
	</div>
	<br />
	{#if $state === 0}
		<button class="btn variant-ghost w-full" type="button" on:click={start}> Starten </button>
	{:else if $state === 1}
		<button class="btn variant-ghost w-full" type="button" on:click={stop}> Beenden </button>
		<div class="card h-[15.6rem] m-1 pl-5 pr-5 p-1 mt-3">
			<CurrentVotes />
		</div>
	{:else if $state === 2}
		<form class="p-2 gap-2 grid">
			<input
				class="input lineNumbers"
				bind:value={result}
				placeholder="7.2"
				type="number"
				required
				min="0"
			/>
			<button class="btn variant-ghost" on:click={aufloesen} type="submit">
				Ergebnis speichern
			</button>
		</form>
	{/if}
</div>
