<script lang="ts">
	import { getToastStore, type ToastSettings } from "@skeletonlabs/skeleton"
	import { onMount } from "svelte"
	import { clipboard } from "@skeletonlabs/skeleton"
	import InputText from "@components/InputText.svelte"
	import InputPassword from "@components/InputPassword.svelte"
	import { IconBrandTwitch } from "@tabler/icons-svelte"
	import { goto } from "$app/navigation"

	const BACKEND_URL = import.meta.env.VITE_BACKEND_URL

	const toastStore = getToastStore()

	const host = window.location.origin.replace("127.0.0.1", "localhost")

	export let oauth: {
		color: number
		clientId: string
		apiKey: string
	}

	let twitch_url: string
	$: twitch_url = `https://id.twitch.tv/oauth2/authorize?response_type=token&client_id=${oauth?.clientId}&redirect_uri=${window.location.origin}/oauth/&scope=chat:read+chat:edit+channel:moderate+whispers:read+whispers:edit+moderator:manage:announcements+moderator:read:blocked_terms`

	async function submit() {
		try {
			const res = await fetch(`${BACKEND_URL}/api/collector/chat/adv/`, {
				headers: {
					"Content-Type": "application/json"
				},
				method: "POST",
				body: JSON.stringify(oauth)
			})
			if (res.ok) {
				const t: ToastSettings = {
					message: "Einstellungen erfolgreich gespeichert.",
					timeout: 5000,
					background: "variant-filled-success"
				}
				toastStore.trigger(t)
				return true
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
		return false
	}

	onMount(async () => {
		const url = `${BACKEND_URL}/api/collector/chat/adv/`
		try {
			const res = await fetch(url)
			oauth = await res.json()

			selectValue = oauth.color.toString()
		} catch (e) {
			const t: ToastSettings = {
				message: "Es ist ein Fehler beim Abfragen der aktuellen Einstellungen aufgetreten.",
				timeout: 5000,
				background: "variant-filled-error"
			}
			toastStore.trigger(t)
		}
	})

	function easy() {
		goto(`${BACKEND_URL}/api/collector/chat/std-oauth/`)
	}

	async function hard() {
		if (oauth.clientId === "") {
			const t: ToastSettings = {
				message: "ClientId darf nicht leer sein.",
				timeout: 5000,
				background: "variant-filled-warning"
			}
			toastStore.trigger(t)
			return
		}
		if (!(await submit())) return

		await goto(twitch_url)
	}

	let selectValue: string

	function setSelected() {
		oauth.color = parseInt(selectValue)
	}
</script>

<svelte:head>
	<title>Twitchchat | Killspiel</title>
</svelte:head>

<div class="pl-5 pr-5 p-2">
	<p>
		Um /announce zu nutzen, ist Zugriff auf die Api nötig. Dafür wird ein Key benötigt mit mehr
		Rechten als <a class="anchor" href="https://twitchapps.com/tmi/">https://twitchapps.com/tmi/</a>
		liefert. Um einen ApiKey zu generieren, wie es TwitchApps macht, ist zudem eine ClientId nötig, welches
		Twitch beim Anlegen einer Anwendung generiert. Für eine ClientId können jedoch nur begrenzte Url
		eine ApiKey anfragen. Wenn dein Server über
		<a class="anchor" href="http://localhost:8088">http://localhost:8088</a>
		oder <a class="anchor" href="http://127.0.0.1:8088">http://127.0.0.1:8088</a> läuft,

		<!-- eslint-disable-next-line svelte/no-at-html-tags -->
		{@html host === "http://localhost:8088"
			? "was er tut"
			: "was er <u class='text-error-500'>nicht</u> tut"}, kannst du die einfache Variante nutzen.
	</p>
	<br />
	<h1 class="text-lg underline">Leichte Variante</h1>
	Klicke
	<a class="anchor" href={`${BACKEND_URL}/api/collector/chat/std-oauth/`}>hier</a> oder unten den
	Button <i>Twitchaccount verknüpfen: leichte Variante</i>.
	<br />
	<br />
	<h1 class="text-lg underline">Komplizierte Variante</h1>
	<p>
		Erstelle <a class="anchor" target="_blank" href="https://dev.twitch.tv/console/apps/create"
			>hier</a
		>
		eine Anwendung. Name und Kategorie sind egal, bei <i>OAuth Redirect URLs</i> gibe
		<code class="italic anchor" use:clipboard={`${host}/oauth/`}>{host}/oauth/</code> ein. <br />
		Klicke auf <i>Erstellen</i> und werde auf eine Übersichtsseite deiner Apps weitergeleitet.
		<br />
		Klicke nun auf Verwalten und kopiere die <i>Client-ID</i> in das untenliegende Feld und
		speichere. <br />
		Klicke <a class="anchor" href={twitch_url}> hier</a> oder unten den Button
		<i>Twitchaccount verknüpfen: komplizierte Variante</i>.
	</p>
</div>
<form class="p-5" on:submit|preventDefault={submit}>
	<div class="grid lg:!grid-cols-2 gap-10 w-full m-2">
		<label>
			<button class="btn variant-ghost w-full" type="button" on:click|preventDefault={easy}>
				<IconBrandTwitch />
				Twitchaccount verknüpfen: leichte Variante
			</button>
		</label>
		<label>
			<button class="btn variant-ghost w-full" type="button" on:click|preventDefault={hard}>
				<IconBrandTwitch />
				Twitchaccount verknüpfen: komplizierte Variante
			</button>
		</label>
		<InputText
			bind:value={oauth.clientId}
			label="ClientId"
			required={true}
			placeholder="6yqxsptsgw5bffjno9b1q9dtr2nlf8"
			modal={{
				title: "ClientId",
				body: "ClientId, wie oben beschrieben, oder bei der leichten Variante einfach den Button <i>Leichte Variante</i> klicken."
			}}
		/>

		<InputPassword
			bind:value={oauth.apiKey}
			label="Twitch Access Token"
			placeholder="bcgf6ogc8swu329nmnqprwgdodizgw"
			prefix="oauth:"
			required={false}
			modal={{
				title: "Twitch Access Token",
				body: "Sollte automatisch ausgefüllt werden durch klicken der obigen Buttons."
			}}
		/>

		<label class="label">
			<span class="">/announce Farbe oder deaktivieren</span>
			<select class="select" bind:value={selectValue} on:change={setSelected}>
				<option value="0" selected>Deaktivieren</option>
				<option value="1" class="capitalize">primary</option>
				<option value="2" class="capitalize">blue</option>
				<option value="3" class="capitalize">green</option>
				<option value="4" class="capitalize">orange</option>
				<option value="5" class="capitalize">purple</option>
			</select>
		</label>
	</div>
	<div class="flex">
		Mit * markierte Felder sind Pflicht
		<button class="btn variant-ghost-success ms-auto" type="submit">Speichern</button>
	</div>
</form>
<hr />
<div class="p-5">
	<h1 class="text-lg underline">Erklärung der Berechtigungen</h1>
	<table class="table rounded-none">
		<tr>
			<td>
				<i>chat:read</i>,
				<i>chat:edit</i>,
				<i>channel:moderate</i>,
				<i>whispers:read</i>,
				<i>whispers:edit</i></td
			>
			<td> Für den Client benötigt, um Nachrichten zu lesen und zu schreiben.</td>
		</tr>
		<tr>
			<td>
				<i>moderator:manage:announcements</i>
			</td>
			<td> Wird benötigt, um /announce zu nutzen. </td>
		</tr>
		<tr>
			<td>
				<i>moderator:read:blocked_terms</i>
			</td>
			<td>
				Wird genutzt um zu überprüfen, das /announce genutzt werden kann ohne eine Nachricht
				schreiben zu müssen.
			</td>
		</tr>
	</table>
	<br />
	<h1 class="text-lg underline">Verbindung trennen</h1>
	Auf
	<a class="anchor" href="https://www.twitch.tv/settings/connections"
		>https://www.twitch.tv/settings/connections</a
	> die "Verbindung" Killspiel-Server trennen. Bei eigener Twitch-App ist der Name entsprechend anders.
</div>
