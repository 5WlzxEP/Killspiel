<script lang="ts">
	import { getToastStore, Tab, TabGroup, type ToastSettings } from "@skeletonlabs/skeleton"
	import { onMount } from "svelte"
	import InputText from "@components/InputText.svelte"
	import { IconCheck, IconX } from "@tabler/icons-svelte"
	import LoLSummonerName from "./LoLSummonerName.svelte"
	import { writable, type Writable } from "svelte/store"
	import LoLCategorie from "./LoLCategorie.svelte"
	import InputPassword from "@components/InputPassword.svelte"

	// setup

	const BACKEND_URL = import.meta.env.VITE_BACKEND_URL

	const toastStore = getToastStore()

	const tabSet: Writable<number> = writable(0);

	let ready: boolean

	let d: {
		lol: {
			summonerName: string,
			profilePic: number,
			kategorie: string,
		}
		apiKey: "",
		periode: 20,
	} = {
		lol: {
			summonerName: "",
			profilePic: -1,
			kategorie: "",
		},
		apiKey: "",
		periode: 20,
	}

	async function isReady() {
		const url = `${BACKEND_URL}/api/data/riot/ready`
		const res = await fetch(url)
		ready = (await res.text()) === "true"
	}

	onMount(async () => {
		const url = `${BACKEND_URL}/api/data/riot/`
		try {
			const res = await fetch(url)
			const val = await res.json()
			console.log(val)
			d = val

			await isReady()
		} catch (e) {
			const t: ToastSettings = {
				message: "Es ist ein Fehler beim Abfragen der aktuellen Einstellungen aufgetreten.",
				timeout: 5000,
				background: "variant-filled-error"
			}
			toastStore.trigger(t)
		}
	})

	// allgemein

	// lol

	// tft
	async function submit() {
		try {
			const res = await fetch(`${BACKEND_URL}/api/data/riot/`, {
				headers: {
					"Content-Type": "application/json"
				},
				method: "POST",
				body: JSON.stringify(d)
			})
			if (res.ok) {
				const t: ToastSettings = {
					message: "Einstellungen erfolgreich gespeichert.",
					timeout: 5000,
					background: "variant-filled-success"
				}
				toastStore.trigger(t)
			}

			await isReady()
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

<div class="container mx-auto">
	<div class="card w-full">
		<div class="text-center font-bold text-2xl p-3 flex mx-auto justify-center">
			RiotApi Settings
			{#if ready}
				<IconCheck size={30} color="lime" class="ml-2 mt-0.5" />
			{:else}
				<div class="cursor-help ml-2 mt-0.5" title="ApiKey und/oder TwitchAccount stimmen nicht.">
					<IconX size={30} color="red" />
				</div>
			{/if}
		</div>
		<hr class="mb-2" />
		<TabGroup>
			<Tab bind:group={$tabSet} name="tab1" value={0}>
				<span>Allgemein</span>
			</Tab>
			<Tab bind:group={$tabSet} name="tab1" value={1}>
				<span>LoL</span>
			</Tab>
			<Tab bind:group={$tabSet} name="tab1" value={2}>
				<span>TfT</span>
			</Tab>
			<svelte:fragment slot="panel">
				{#if $tabSet === 0}
					<form class="p-5 pt-0" on:submit|preventDefault={submit}>
						<div class="grid lg:!grid-cols-2 gap-10 w-full m-2">
							<label class="label m-2">
								<span>Dauer</span>
								<div class="input-group input-group-divider grid-cols-[1fr_auto]">
									<input title="20" type="number" placeholder="20" bind:value={d.periode} />
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
					<form class="p-5 pt-0" on:submit|preventDefault={submit}>
						<div class="grid lg:!grid-cols-2 gap-10 w-full m-2">
							<InputPassword
								bind:value={d.apiKey}
								label="Riot ApiKey"
								placeholder="RGAPI-dab3c199-0558-4652-aa2a-830884c7a215"
								modal={{
						title: "Riot Api key",
						body: "Dies ist der Api Key, der benutzt wird, um die Riot-Api abzufragen, ob ein Spieler gerade spielt und das Ergenis liefert, wenn das Spiel vorbei ist. <br>" +
						 'Ein ApiKey kann über <a class="anchor" href="https://developer.riotgames.com/" target="_blank">developer.riotgames.com</a> angefordert werden. Es wird empfolen, ein Produkt zu registrieren und einen dauerhaften Apikey zu erhalten. Ein temporer funktioniert aber auch. (Läuft halt nur nach 24h ab)'
					}}
							/>

							<LoLSummonerName
								bind:value={d.lol.summonerName}
								label="Beschwörername"
								placeholder="5WlzxEp"
								prefix={d.lol.profilePic}
								modal={{
						title: "Beschwörername",
						body:
							"Der Beschwörername des Accounts, dessen Spiele \"überwacht\" werden sollen."
					}}
							/>
							<LoLCategorie bind:value={d.lol.kategorie}/>
						</div>
						<div class="flex mt-5">
							Mit * markierte Felder sind Pflicht
							<button class="btn variant-ghost-success ms-auto" type="submit">Speichern</button>
						</div>
					</form>
				{:else if $tabSet === 2}
					(tab panel 3 contents)
				{/if}
			</svelte:fragment>
		</TabGroup>
	</div>
</div>
