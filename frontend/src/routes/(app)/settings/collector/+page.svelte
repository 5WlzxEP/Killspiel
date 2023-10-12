<script lang="ts">
	import { IconBrandTwitch } from "@tabler/icons-svelte"
	import { onMount } from "svelte"
	import { getToastStore, type ToastSettings } from "@skeletonlabs/skeleton"

	const BACKEND_URL = import.meta.env.VITE_BACKEND_URL
	const toastStore = getToastStore()

	let scale: number
	let time: number
	let options: Array<{ name: string; ready: boolean }> = []
	let select: HTMLSelectElement
	let selectCollector: HTMLSelectElement

	$: scale = select?.value as number

	async function send() {
		try {
			const res = await fetch(`${BACKEND_URL}/api/collector/`, {
				method: "POST",
				headers: {
					"content-type": "application/json"
				},
				body: JSON.stringify({
					collector: selectCollector.value,
					duration: scale * time
				})
			})
			if (res.ok) {
				const t: ToastSettings = {
					message: "Erfolgreich gespeichert.",
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

	onMount(async () => {
		try {
			const res = await fetch(`${BACKEND_URL}/api/collector/`)
			if (res.ok) {
				const j: {
					all: Array<{ name: string; ready: boolean }>
					collector: { name: string }
					time: number
				} = await res.json()
				options = j.all
				options.push({ name: "", ready: true })

				if (j.collector !== undefined) selectCollector.value = j.collector.name
				else selectCollector.value = ""

				if (j.time % 60 === 0) {
					time = Math.floor(j.time / 60)
					scale = 60
				} else {
					time = j.time
					scale = 1
				}

				select.value = String(scale)
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
	})
</script>

<svelte:head>
	<title>Guess Collector | Killspiel</title>
</svelte:head>

<div class="container grid grid-cols-2 mx-auto gap-4">
	<a href="/settings/collector/twitchchat" class="card p-2">
		<div class="p-1">
			<div class="flex gap-2">
				<IconBrandTwitch />
				<h1 class="text-lg font-bold">Twitchchat</h1>
			</div>
			Lass die Teilnehmer über den Twitchchat abstimmen.
		</div>
	</a>
</div>

<div class="container mx-auto m-2 justify-center items-center flex mt-6">
	<div class="card p-4 w-2/3">
		<div class="m-2">
			<h2 class="text-xl">Einstellungen</h2>
		</div>
		<hr />
		<form>
			<label class="label m-2">
				<span>Dauer</span>
				<div class="input-group input-group-divider grid-cols-[1fr_auto]">
					<input title="3" type="number" placeholder="3" bind:value={time} />
					<select bind:this={select}>
						<option value="60">min</option>
						<option value="1">s</option>
					</select>
				</div>
			</label>
			<label class="label m-2">
				<span>Aktiver Collector</span>
				<select bind:this={selectCollector} class="input select">
					{#each options as option}
						<option value={option.name} disabled={!option.ready}>{option.name}</option>
					{/each}
				</select>
			</label>
			<div class="flex mt-4">
				<button class="btn variant-ghost-success ms-auto" type="submit" on:click={send}
					>Speichern
				</button>
			</div>
		</form>
	</div>
</div>
