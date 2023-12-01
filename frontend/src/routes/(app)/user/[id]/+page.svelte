<script lang="ts">
	import type { User } from "./+page"
	import {
		getModalStore,
		getToastStore,
		type ModalSettings,
		type PaginationSettings,
		Paginator,
		type ToastSettings
	} from "@skeletonlabs/skeleton"
	import Check from "@components/Check.svelte"
	import X from "@components/X.svelte"
	import { goto } from "$app/navigation"
	import { modalButtonPositive } from "@stores/modal"

	const BACKEND_URL = import.meta.env.VITE_BACKEND_URL
	const modalStore = getModalStore()
	const toastStore = getToastStore()
	/** @type {import("./$types").PageData} */
	export let data: User

	let meta = {
		page: 0,
		limit: 25,
		size: data.guesses,
		amounts: [10, 15, 25]
	} satisfies PaginationSettings

	function setIcon() {
		data.history.forEach((h) => {
			h.icon = Math.abs(h.guess - h.correct) < h.precision ? Check : X
		})
	}
	setIcon()

	async function deleteConfirmed(key: string) {
		const url = `${BACKEND_URL}/api/user/${data.id}/`
		await fetch(url, {
			method: "DELETE",
			body: JSON.stringify({ key: key }),
			headers: {
				"content-type": "application/json"
			}
		})

		const t: ToastSettings = {
			autohide: true,
			message: `User ${data.name} wurde gelöscht`,
			timeout: 3000
		}

		toastStore.trigger(t)
		await goto("/")
	}

	async function deleteUser() {
		modalButtonPositive.set("variant-filled-error")
		const url = `${BACKEND_URL}/api/user/${data.id}/`
		const res = await fetch(url, {
			method: "DELETE"
		})
		const j = await res.json()

		if (!j.key) {
			console.log("ERROR")
			return
		}

		const modal: ModalSettings = {
			type: "confirm",
			title: `Lösche User ${data.name}`,
			body: `Bist du dir sicher, dass du ${data.name} löschen möchtest?`,
			buttonTextConfirm: "Löschen",
			modalClasses: "variant-ghost-error",
			response: async (r: boolean) => {
				modalButtonPositive.set("variant-filled")
				if (!r) {
					return
				}
				await deleteConfirmed(j.key)
			}
		}

		modalStore.trigger(modal)
	}

	async function update() {
		const url = `${BACKEND_URL}/api/user/${data.id}/history/?offset=${
			meta.page * meta.limit
		}&limit=${meta.limit}`
		try {
			const resp = await fetch(url)
			data.history = await resp.json()
			setIcon()
		} catch (e) {
			/* empty */
		}
	}
</script>

<!--<Modal buttonPositive="variant-ghost-error" />-->
<svelte:head>
	<title>{data.name} | Killspiel</title>
</svelte:head>

<div class="container mx-auto p-2">
	<h1 class="text-6xl text-center">{data.name}</h1>
	<hr class="p-2 mt-3" />

	<div class="grid grid-cols-3 gap-4">
		<div class="card p-2 text-xl">
			Punkte:
			<div class="text-5xl text-center">{data.points}</div>
		</div>
		<div class="card p-2 text-xl">
			Versuche:
			<div class="text-5xl text-center">{data.guesses}</div>
		</div>
		<div class="card p-2 text-xl">
			Trefferrate:
			<div class="text-5xl text-center">
				{data.guesses !== 0 ? ((data.points / data.guesses) * 100).toFixed(2) : 0}%
			</div>
		</div>
	</div>

	<div class="mt-6 {data.history.length > 0 ? '' : 'hidden'}">
		<div class="table-container">
			<!-- Native Table Element -->
			<table class="table table-hover">
				<thead>
					<tr>
						<th class="w-[10%] text-center">Game</th>
						<th class="text-right">Schätzung</th>
						<th class="text-right">Ergebnis</th>
						<th class="text-right">Toleranz</th>
						<th class="text-right">Treffer</th>
						<th class="w-[15%] text-right">Zeit</th>
					</tr>
				</thead>
				<tbody>
					{#each data.history as row}
						<tr>
							<td class="text-right">
								<a href="/game/{row.id}" class="w-full p-4" data-sveltekit-preload-data="tap">
									{row.id}
								</a>
							</td>
							<td class="text-right">{row.guess.toFixed(2)}</td>
							<td class="text-right">{row.correct.toFixed(2)}</td>
							<td class="text-right">{row.precision.toFixed(2)}</td>
							<td class="flex">
								<div class="ms-auto">
									<svelte:component this={row.icon} />
								</div>
							</td>
							<td class="text-right">{new Date(Date.parse(row.time)).toLocaleString()}</td>
						</tr>
					{/each}
				</tbody>
			</table>
			<div class="mt-2">
				<Paginator
					bind:settings={meta}
					showFirstLastButtons={true}
					showPreviousNextButtons={true}
					showNumerals={true}
					on:page={update}
					on:amount={update}
				/>
			</div>
		</div>
	</div>

	<div class="flex mt-4">
		<button class="ms-auto mr-3 btn variant-ghost" on:click={deleteUser}>User löschen</button>
	</div>
</div>
