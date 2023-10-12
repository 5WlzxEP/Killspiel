<script lang="ts">
	import {
		getToastStore,
		localStorageStore,
		Tab,
		TabGroup,
		type ToastSettings
	} from "@skeletonlabs/skeleton"
	import { onMount } from "svelte"
	import { IconCheck, IconX } from "@tabler/icons-svelte"
	import type { Writable } from "svelte/store"
	import LoL from "./LoL.svelte"
	import { _isReady, _ready, type Resp } from "./+page"

	const BACKEND_URL = import.meta.env.VITE_BACKEND_URL

	const toastStore = getToastStore()

	const tabSet: Writable<number> = localStorageStore("tabset", 0)

	let data: Resp = {
		lol: {
			summonerName: "",
			profileIcon: -1,
			kategorie: "",
			server: "euw1",
			apiKey: ""
		},
		general: {
			intervall: 20
		}
	}

	onMount(async () => {
		const url = `${BACKEND_URL}/api/data/riot/`
		try {
			const res = await fetch(url)
			data = await res.json()
			data.general.intervall /= 1000000000

			await _isReady()
		} catch (e) {
			const t: ToastSettings = {
				message: "Es ist ein Fehler beim Abfragen der aktuellen Einstellungen aufgetreten.",
				timeout: 5000,
				background: "variant-filled-error"
			}
			toastStore.trigger(t)
		}
	})

	async function submit() {
		try {
			const res = await fetch(`${BACKEND_URL}/api/data/riot/`, {
				headers: {
					"Content-Type": "application/json"
				},
				method: "POST",
				body: JSON.stringify(data.general)
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

			await _isReady()
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
</script>

<svelte:head>
	<title>RiotApi Einstellungen | Killspiel</title>
</svelte:head>
<div class="container mx-auto">
	<div class="card w-full">
		<div class="text-center font-bold text-2xl p-3 flex mx-auto justify-center">
			RiotApi Settings
			{#if $_ready}
				<IconCheck size={30} color="lime" class="ml-2 mt-0.5" />
			{:else}
				<div class="cursor-help ml-2 mt-0.5" title="ApiKey und/oder TwitchAccount stimmen nicht.">
					<IconX size={30} color="red" />
				</div>
			{/if}
		</div>
		<TabGroup>
			<div class="flex mx-auto">
				<Tab bind:group={$tabSet} name="allgemein" value={0}>
					<span>Allgemein</span>
				</Tab>
				<Tab bind:group={$tabSet} name="LoL" value={1}>
					<span>LoL</span>
				</Tab>
			</div>

			<svelte:fragment slot="panel">
				{#if $tabSet === 0}
					<form class="p-5 pt-0" on:submit|preventDefault={submit}>
						<div class="grid lg:!grid-cols-2 gap-10 w-full m-2">
							<label class="label m-2">
								<span>Frequenz der Api-Abfrage</span>
								<div class="input-group input-group-divider grid-cols-[1fr_auto]">
									<input
										title="20"
										type="number"
										placeholder="20"
										bind:value={data.general.intervall}
									/>
									<div>s</div>
								</div>
							</label>
						</div>
						<div class="flex mt-5">
							Mit * markierte Felder sind Pflicht
							<button class="btn variant-ghost-success ms-auto" type="submit">Speichern</button>
						</div>
					</form>
				{:else if $tabSet === 1}
					<LoL bind:lol={data.lol} />
				{/if}
			</svelte:fragment>
		</TabGroup>
	</div>
</div>
