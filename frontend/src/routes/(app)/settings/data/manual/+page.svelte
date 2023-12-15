<script lang="ts">
	import { type ToastSettings, getToastStore } from "@skeletonlabs/skeleton"
	import { onDestroy, onMount } from "svelte"
	import { state, printState, collectionEnd } from "@stores/state"

	const toastStore = getToastStore()
	const BACKEND_URL = import.meta.env.VITE_BACKEND_URL

	let dissolve: boolean

	let r: number | string | null

	$: {
		if (typeof r === "number") r = Math.abs(r)
		else if (r !== null) r = 0
	}

	$: dissolve = $state !== 2 || typeof r !== "number"

	async function change() {
		try {
			const res = await fetch(`${BACKEND_URL}/api/data/`, {
				headers: {
					"Content-Type": "application/json"
				},
				method: "POST",
				body: JSON.stringify({
					state: $state,
					result: r
				})
			})
			if (res.ok) {
				const t: ToastSettings = {
					message: "Einstellungen erfolgreich gespeichert.",
					timeout: 5000,
					background: "variant-filled-success"
				}
				toastStore.trigger(t)
			} else {
				const t: ToastSettings = {
					message: await res.text(),
					timeout: 5000,
					background: "variant-filled-error"
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

	async function update() {
		const url = `${BACKEND_URL}/api/data/`
		try {
			const res = await fetch(url)
			const val = await res.json()
			$state = val.state
		} catch (e) {
			console.error(e)
		}
	}

	onMount(async () => {
		await update()
	})

	let printSt: string
	$: printSt = printState($state)

	$: if ($state === 1) printTime()
	else {
		clearInterval(interval)
		time = ""
	}

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

<svelte:head>
	<title>Daten Collector - Manuel | Killspiel</title>
</svelte:head>

<div class="container mx-auto">
	<div class="card w-full">
		<h1 class="text-center text-2xl p-3">Manuell</h1>
		<hr />
		<div class="grid gap-10 w-full m-2 grid-cols-[2fr_auto_1fr_auto_2fr] p-2">
			<div class="p-2 grid gap-2 h-3">
				<button class="btn variant-ghost" type="button" disabled={$state !== 0} on:click={change}>
					Starten
				</button>
				<button class="btn variant-ghost" type="button" disabled={$state !== 1} on:click={change}>
					Beenden
				</button>
			</div>

			<span class="divider-vertical h-full" />

			<div class="p-2 grid gap-2 h-3 text-center">
				{printSt}
				<br />
				{time}
			</div>

			<span class="divider-vertical h-full" />

			<form class="p-2 gap-2 grid">
				<input
					class="input lineNumbers"
					bind:value={r}
					placeholder="7.2"
					type="number"
					required
					min="0"
					step="0.1"
				/>
				<button class="btn variant-ghost" disabled={dissolve} on:click={change} type="submit">
					Ergebnis speichern
				</button>
			</form>
		</div>
	</div>
</div>
