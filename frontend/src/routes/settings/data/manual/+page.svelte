<script lang="ts">
	import { type ToastSettings, getToastStore } from "@skeletonlabs/skeleton"
	import { onMount } from "svelte"

	const toastStore = getToastStore()
	const BACKEND_URL = import.meta.env.VITE_BACKEND_URL

	let result = 0
	let state = 0
	let start: HTMLButtonElement
	let dissolve: HTMLButtonElement
	let end: HTMLButtonElement

	let r: number | string

	$: {
		if (typeof r === "number") r = Math.abs(r)
	}

	function change(e: MouseEvent & { currentTarget: EventTarget & HTMLButtonElement }) {
		setState(++state)

		if (state === 1) {
			setTimeout(change, 5000)
		}
	}

	function setState(change: number) {
		state = change % 3

		start.disabled = state != 0 // none
		end.disabled = state !== 1 // collecting
		dissolve.disabled = state !== 2 || result == undefined // running
	}

	async function submit(e: Event) {
		// console.log(e)
		try {
			const res = await fetch(`${BACKEND_URL}/api/collector/chat/`, {
				headers: {
					"Content-Type": "application/json"
				},
				method: "POST",
				body: JSON.stringify("")
			})
			if (res.ok) {
				const t: ToastSettings = {
					message: "Einstellungen erfolgreich gespeichert.",
					timeout: 5000,
					background: "variant-filled-success"
				}
				toastStore.trigger(t)
			}
		} catch (e) {
			console.error(e)
			const t: ToastSettings = {
				message: "Es ist ein Fehler beim Abfragen der aktualisieren Einstellungen aufgetreten.",
				timeout: 5000,
				background: "variant-filled-error"
			}
			toastStore.trigger(t)
		}
	}

	async function submitEnd() {
		const url = `${BACKEND_URL}/api/end`
		try {
			const res = await fetch(url)
			const val = await res.json()
		} catch (e) {
			console.error(e)
			const t: ToastSettings = {
				message: "Es ist ein Fehler beim Abfragen der aktuellen Einstellungen aufgetreten.",
				timeout: 5000,
				background: "variant-filled-error"
			}

			toastStore.trigger(t)
		}
	}

	onMount(async () => {
		setState(0)

		const url = `${BACKEND_URL}/api/data/manual/`
		try {
			const res = await fetch(url)
			const val = await res.json()
			setState(val.state)
		} catch (e) {
			const t: ToastSettings = {
				message: "Es ist ein Fehler beim Abfragen der aktuellen Einstellungen aufgetreten.",
				timeout: 5000,
				background: "variant-filled-error"
			}
			toastStore.trigger(t)
		}
	})
</script>

<div class="container mx-auto">
	<div class="card w-full">
		<h1 class="text-center text-2xl p-3">Manuell</h1>
		<hr />
		<div class="grid gap-10 w-full m-2 grid-cols-[2fr_auto_1fr_auto_2fr] p-2">
			<div class="p-2 grid gap-2 h-3">
				<button
					class="btn variant-ghost"
					type="button"
					bind:this={start}
					on:click={(e) => change(e)}
				>
					Starten
				</button>
				<button class="btn variant-ghost" type="button" bind:this={end} on:click={submitEnd}>
					Beenden
				</button>
			</div>

			<span class="divider-vertical h-full" />

			<div class="p-2 grid gap-2 h-3">
				Current State:
				{#if state === 0}
					Ready
				{:else if state === 1}
					Sammele Schätzungen
				{:else if state === 2}
					Warte auf Auflösung
				{/if}
			</div>

			<span class="divider-vertical h-full" />

			<form class="p-2 gap-2 grid">
				<input
					class="input"
					bind:value={r}
					on:keyup={() => (dissolve.disabled = state !== 2 || result === null)}
					placeholder="7.2"
					type="number"
					required
					min="0"
				/>
				<button
					class="btn variant-ghost"
					bind:this={dissolve}
					on:click={(e) => change(e)}
					type="submit"
				>
					Ergebnis speichern
				</button>
			</form>
		</div>
	</div>
</div>
