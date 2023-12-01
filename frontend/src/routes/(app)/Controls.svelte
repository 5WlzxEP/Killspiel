<script lang="ts">
	import { printState, state } from "@stores/state"
	import { onMount } from "svelte"
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
		} catch (e) {
			console.error(e)
		}
	})
</script>

<div class=" w-full h-full p-2">
	<h1 class="text-lg text-center anchor">
		<a href="/settings/data/manual/">Manuell & Status</a>
	</h1>
	<div class="text-center card p-2">
		Aktueller Status: <h3 class="text-lg">{printState($state)}</h3>
	</div>
	<br />
	{#if $state === 0}
		<button class="btn variant-ghost w-full mt-2" type="button" on:click={start}> Starten </button>
	{:else if $state === 1}
		<button class="btn variant-ghost w-full mt-2" type="button" on:click={stop}> Beenden </button>
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
