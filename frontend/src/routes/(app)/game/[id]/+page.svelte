<script lang="ts">
	import type { Game } from "./+page"
	import GameChart from "@components/GameChart.svelte"
	import { writable } from "svelte/store"
	import {
		getModalStore,
		getToastStore,
		type ModalSettings,
		type ToastSettings
	} from "@skeletonlabs/skeleton"
	const BACKEND_URL = import.meta.env.VITE_BACKEND_URL

	/** @type {import("./$types").PageData} */
	export let data: Game

	const modalStore = getModalStore()
	const toastStore = getToastStore()

	let vert = writable(JSON.parse(data.verteilung))

	async function aufloesen() {
		const modal: ModalSettings = {
			type: "prompt",
			title: `Spiel ${data.id} auflösen`,
			body: "Korrektes Ergebnis eintragen:",
			valueAttr: { type: "number", required: true, step: 0.01, min: 0 },
			value: "0.00",
			response: async (r: number) => {
				if (!r) {
					return
				}
				fetch(`${BACKEND_URL}/api/game/${data.id}/`, {
					method: "PUT",
					body: r
				})
					.catch((e) => {
						const toast: ToastSettings = {
							message: e,
							autohide: true,
							timeout: 3000,
							background: "variant-ghost-error"
						}
						toastStore.trigger(toast)
					})
					.then(async (r: Response) => {
						if (r.ok) {
							const toast: ToastSettings = {
								message: "Erfolgreich gespeichert",
								autohide: true,
								timeout: 3000,
								background: "variant-ghost-success"
							}
							toastStore.trigger(toast)

							setTimeout(location.reload, 1000)
						} else {
							const toast: ToastSettings = {
								message: await r.text(),
								autohide: true,
								timeout: 3000,
								background: "variant-ghost-error"
							}
							toastStore.trigger(toast)
						}
					})
			}
		}
		modalStore.trigger(modal)
	}
</script>

<svelte:head>
	<title>Game {data.id} | Killspiel</title>
</svelte:head>

<div class="container mx-auto p-2">
	<h1 class="text-6xl text-center">Spiel {data.id}</h1>
	<hr class="p-2 mt-3" />

	<div class="grid grid-cols-3 gap-4">
		<div class="card p-2 text-xl">
			Teilnehmer:
			<div class="text-5xl text-center">{data.userCount}</div>
		</div>
		<div class="card p-2 text-xl">
			Korrekte Schätzungen:
			<div class="text-5xl text-center">{data.correctCount}</div>
		</div>
		<div class="card p-2 text-xl">
			Trefferrate:
			<div class="text-5xl text-center">
				{data.userCount !== 0 ? ((data.correctCount / data.userCount) * 100).toFixed(2) : 0}%
			</div>
		</div>
		<div class="card p-2 text-xl">
			Richtiges Ergebnis:
			<div class="text-5xl text-center">{data.correct}</div>
		</div>
		<div class="card p-2 text-xl">
			Präzision:
			<div class="text-5xl text-center">{data.precision}</div>
		</div>
		<div class="card p-2 text-xl">
			Erstellt:
			<div class="text-5xl text-center">
				{new Date(Date.parse(data.time)).toLocaleString()}
			</div>
		</div>
	</div>

	<div class="p-2 card mt-2">
		<GameChart verteilung={vert} gameid={data.id} />
	</div>

	{#if data.correct === null}
		<div class="flex mt-4">
			<button class="ms-auto mr-3 btn variant-ghost" on:click={aufloesen}>Auflösen</button>
		</div>
	{/if}
</div>
