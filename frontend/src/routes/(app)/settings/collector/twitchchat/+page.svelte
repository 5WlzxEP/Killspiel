<script lang="ts">
	import { type ToastSettings, getToastStore, TabGroup, Tab } from "@skeletonlabs/skeleton"
	import { onMount } from "svelte"
	import { IconCheck, IconX } from "@tabler/icons-svelte"
	import Basic from "./Basic.svelte"
	import Advanced from "./Advanced.svelte"

	const BACKEND_URL = import.meta.env.VITE_BACKEND_URL

	const toastStore = getToastStore()

	let tabSet = 0

	let ready: boolean

	let d = {
		apiKey: "",
		channel: "",
		channelSender: "",
		prefix: "",
		msgBegin: "",
		msgEnd: "",
		msgFinal: "",
		oauth: {
			clientId: "",
			apiKey: ""
		}
	}

	async function isReady() {
		const url = `${BACKEND_URL}/api/collector/chat/ready/`
		const res = await fetch(url)
		ready = (await res.text()) === "true"
	}

	onMount(async () => {
		const url = `${BACKEND_URL}/api/collector/chat/`
		try {
			const res = await fetch(url)
			const val = await res.json()
			d.apiKey = val.apiKey
			d.channel = val.channel
			d.channelSender = val.channelSender
			d.prefix = val.prefix
			d.msgBegin = val.msgBegin
			d.msgEnd = val.msgEnd
			d.msgFinal = val.msgFinal

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
</script>

<svelte:head>
	<title>Twitchchat | Killspiel</title>
</svelte:head>

<div class="container mx-auto">
	<div class="card w-full">
		<div class="text-center font-bold text-2xl p-3 flex mx-auto justify-center">
			Twitchchat Settings
			{#if ready}
				<IconCheck size={30} color="lime" class="ml-2 mt-0.5" />
			{:else}
				<div class="cursor-help ml-2 mt-0.5" title="ApiKey und/oder TwitchAccount stimmen nicht.">
					<IconX size={30} color="red" />
				</div>
			{/if}
		</div>
		<!--		<hr />-->
		<TabGroup>
			<div class="flex mx-auto">
				<Tab bind:group={tabSet} name="tab1" value={0}>Basic</Tab>
				<Tab bind:group={tabSet} name="tab2" value={1}>Advanced</Tab>
			</div>

			<svelte:fragment slot="panel">
				{#if tabSet === 0}
					<Basic bind:ready bind:d />
				{:else if tabSet === 1}
					<Advanced bind:oauth={d.oauth} />
				{/if}
			</svelte:fragment>
		</TabGroup>
	</div>
</div>
