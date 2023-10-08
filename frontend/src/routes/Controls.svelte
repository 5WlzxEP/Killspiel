<script lang="ts">
	import { printState, state } from "@stores/state"
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
</script>

<div class=" w-full h-full p-2">
	<div class="text-center card p-2">
		Aktueller Status: {printState($state)}
	</div>
	<br />
	{#if $state === 0}
		<button
			class="btn variant-ghost w-full mt-2"
			type="button"
			disabled={$state !== 0}
			on:click={start}
		>
			Starten
		</button>
	{:else if $state === 1}
		<button
			class="btn variant-ghost w-full mt-2"
			type="button"
			disabled={$state !== 1}
			on:click={stop}
		>
			Beenden
		</button>
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
