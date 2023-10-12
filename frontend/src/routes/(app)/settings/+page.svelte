<script lang="ts">
	import Theme from "./Theme.svelte"
	import { onMount } from "svelte"
	import { getToastStore, type ToastSettings } from "@skeletonlabs/skeleton"

	const toastStore = getToastStore()

	const BACKEND_URL = import.meta.env.VITE_BACKEND_URL

	let precision: {
		precision: number
	} = {
		precision: 0
	}

	async function submit() {
		try {
			const res = await fetch(`${BACKEND_URL}/api/`, {
				headers: {
					"Content-Type": "application/json"
				},
				method: "POST",
				body: JSON.stringify(precision)
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
				return
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

	onMount(async () => {
		const url = `${BACKEND_URL}/api/`
		try {
			const res = await fetch(url)
			precision = await res.json()
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

<svelte:head>
	<title>Settings | Killspiel</title>
</svelte:head>

<div class="mx-auto container p-2 w-3/5">
	<div class="card container">
		<h1 class="text-center text-2xl p-3">Settings</h1>
		<hr />

		<section class="p-4">
			<div class="text-center">
				<a href="/settings/collector" class="btn variant-ghost">Guess Collectors</a>
				<a href="/settings/data" class="btn variant-ghost">Daten Collectors</a>
			</div>
		</section>
		<hr />
		<section class="p-4">
			<h2 class="text-xl">Allgemeine Einstellungen</h2>
		</section>
		<Theme />

		<div class="p-4 w-full">
			<a class="btn w-full variant-ghost" href="/overlay">Overlays</a>
		</div>
		<hr class="mt-2 mb-2" />
		<form class="p-4" on:submit|preventDefault={submit}>
			<label class="label m-2">
				<span>Genauigkeit</span>
				<div class="input-group input-group-divider grid-cols-[1fr_auto]">
					<input
						type="number"
						placeholder="0.01"
						min="0.01"
						step="0.01"
						bind:value={precision.precision}
						required
					/>
				</div>
			</label>

			<div class="flex mt-4">
				<button class="btn variant-ghost-success ms-auto" type="submit">Speichern</button>
			</div>
		</form>
	</div>
</div>
