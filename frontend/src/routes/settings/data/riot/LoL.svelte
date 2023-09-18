<script lang="ts">
	import InputPassword from "@components/InputPassword.svelte"
	import LoLSummonerName from "./LoLSummonerName.svelte"
	import LoLCategorie from "./LoLCategorie.svelte"
	import { getToastStore, type ToastSettings } from "@skeletonlabs/skeleton"
	import type { LoL } from "./+page"
	import LoLServers from "./LoLServers.svelte"
	import { _isReady } from "./+page"

	const BACKEND_URL = import.meta.env.VITE_BACKEND_URL

	const toastStore = getToastStore()
	export let lol: LoL

	async function submit() {
		try {
			const res = await fetch(`${BACKEND_URL}/api/data/riot/lol`, {
				headers: {
					"Content-Type": "application/json"
				},
				method: "POST",
				body: JSON.stringify(lol)
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

<form class="p-5 pt-0" on:submit|preventDefault={submit}>
	<div class="grid lg:!grid-cols-2 gap-10 w-full m-2">
		<InputPassword
			bind:value={lol.apiKey}
			label="Riot ApiKey"
			placeholder="RGAPI-dab3c199-0558-4652-aa2a-830884c7a215"
			modal={{
				title: "Riot Api key",
				body:
					"Dies ist der Api Key, der benutzt wird, um die Riot-Api abzufragen, ob ein Spieler gerade spielt und das Ergenis liefert, wenn das Spiel vorbei ist. <br> Riot empfiehlt, für LoL, TfT und LoR unterschiedliche Keys zu nutzen, es können aber auch dieselben genutzt werden.<br>" +
					'Ein ApiKey kann über <a class="anchor" href="https://developer.riotgames.com/" target="_blank">developer.riotgames.com</a> angefordert werden. Es wird empfolen, ein Produkt zu registrieren und einen dauerhaften Apikey zu erhalten. Ein temporer funktioniert aber auch. (Läuft halt nur nach 24h ab)'
			}}
		/>

		<LoLSummonerName
			bind:value={lol.summonerName}
			label="Beschwörername"
			placeholder="5WlzxEp"
			prefix={lol.profileIcon}
			modal={{
				title: "Beschwörername",
				body: 'Der Beschwörername des Accounts, dessen Spiele "überwacht" werden sollen.'
			}}
		/>
		<LoLCategorie bind:value={lol.kategorie} />

		<LoLServers bind:value={lol.server} />
	</div>
	<div class="flex mt-5">
		Mit * markierte Felder sind Pflicht
		<button class="btn variant-ghost-success ms-auto" type="submit">Speichern</button>
	</div>
</form>
