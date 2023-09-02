<script lang="ts">
	import { type ToastSettings, getToastStore } from "@skeletonlabs/skeleton"
	import { onMount } from "svelte"
	import InputText from "@components/InputText.svelte"
	import InputArea from "@components/InputArea.svelte"

	const BACKEND_URL = import.meta.env.VITE_BACKEND_URL

	const toastStore = getToastStore()

	const d = {
		apiKey: "",
		channel: "",
		channelSender: "",
		prefix: "",
		msgBegin: "",
		msgEnd: "",
		msgFinal: ""
	}

	async function submit() {
		try {
			const res = await fetch(`${BACKEND_URL}/api/collector/chat/`, {
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
		<h1 class="text-center text-2xl p-3">Twitchchat Settings</h1>
		<hr />
		<form class="p-5" on:submit|preventDefault={submit}>
			<div class="grid lg:!grid-cols-2 gap-10 w-full m-2">
				<InputText
					bind:value={d.channel}
					label="Twitch Channel"
					prefix="twitch.tv/"
					required={true}
					placeholder="5W_lzxEP"
					modal={{
						title: "Twitch Channel",
						body: "Dies ist der Twitchchannel mit dessen Chat interagiert wird."
					}}
				/>

				<InputText
					bind:value={d.prefix}
					label="Prefix"
					placeholder="/guess"
					modal={{
						title: "Prefix",
						body:
							"Der Prefix gibt an, womit eine Nachricht beginnen muss, damit sie als Guess registriert wird. " +
							"Der Prefix kann auch leer sein, dann wird jede Nachricht, die nur eine Zahl darstellt ausgewertet."
					}}
				/>

				<InputText
					bind:value={d.apiKey}
					label="Twitch Api Key"
					placeholder="oauth:bcgf6ogc8swu329nmnqprwgdodizgw"
					required={true}
					modal={{
						title: "Twitch Api Key",
						body:
							"Der Twitch Api Key ist nötig um im Chat Nachrichten zu schreiben. So kann der Begin angekündigt werden, auf Ende hingewiesen und das Ergebnis bekannt gegeben werden. " +
							'Am einfachsten kann man den Twitch Api Key über <a class="anchor" target="_blank" href="https://twitchapps.com/tmi/">https://twitchapps.com/tmi/</a> möglich.'
					}}
				/>

				<InputText
					bind:value={d.channelSender}
					label="Twitch Account (selber wie vom Api Key)"
					placeholder="5W_lzxEP"
					prefix="twitch.tv/"
					modal={{
						title: "Twitch Account Api",
						body:
							"Dies ist der Twitchchannel von dem der Api Key stammt. Dieser Account wird die Nachrichten senden. " +
							"Wenn keiner gesetzt wird, wird der Twitchchannel genutzt, wessen Chat genutzt wird. <br> " +
							"Um Twitchchatfeatures wie /announce zu nutzen, muss der Account auf dem Channel Mod-Rechte haben."
					}}
				/>

				<InputArea
					bind:value={d.msgBegin}
					label="Nachricht zum Begin der Erhebung"
					placeholder="/announceblue Das Killspiel hat begonnen. Nimm jetzt Teil mit /guess <Dein Guess>."
					modal={{
						title: "Nachricht zum Begin der Erhebung",
						body:
							"Diese Nachricht wird zum Beginn der Erhebung, also wenn die Zuschauer ihre Schätzungen abgeben können, in den Chat gepostet. <br> " +
							"Es können Twitchfeatures wie /announce genutzt werden, jedoch sind für diese eventuell Rechte nötig."
					}}
				/>

				<InputArea
					bind:value={d.msgEnd}
					label="Nachricht zum Ende der Erhebung"
					placeholder="/announceorange Das Voten ist abgeschlossen. Ab jetzt bitte keine Stimmen in der Chat mehr."
					modal={{
						title: "Nachricht zum Ende der Erhebung",
						body:
							"Diese Nachricht wird zum Ende der Erhebung in den Chat gepostet. <br> " +
							"Es können Twitchfeatures wie /announce genutzt werden, jedoch sind für diese eventuell Rechte nötig."
					}}
				/>

				<InputArea
					bind:value={d.msgFinal}
					label="Nachricht zum Auflösen der richtigen Schätzungen"
					placeholder="/announcegreen Das Killspiel ist beendet. Es wurden $RESULT Kills erzielt und somit haben $WINNERS gewonnen."
					modal={{
						title: "Nachricht zum Auflösen der richtigen Schätzungen",
						body:
							"Diese Nachricht wird versendet, wenn das Ergebnis feststeht. <br> " +
							"<i>$WINNERS</i> sind die Gewinner <br> <i>$RESULT</i> sind die erzielten Kills"
					}}
				/>
			</div>
			<div class="flex">
				Mit * markierte Felder sind Pflicht
				<button class="btn variant-ghost-success ms-auto" type="submit">Speichern</button>
			</div>
		</form>
	</div>
</div>
