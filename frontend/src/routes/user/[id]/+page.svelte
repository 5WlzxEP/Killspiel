<script lang="ts">
	import type { User } from "./+page"
	import {
		getModalStore,
		getToastStore,
		type ModalSettings,
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

	data.history.forEach((h) => {
		h.icon = h.guess === h.correct ? Check : X
	})

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
		modalButtonPositive.set("variant-filled-error")
		modalStore.trigger(modal)
	}
</script>

<!--<Modal buttonPositive="variant-ghost-error" />-->
<svelte:head>
	<title>{data.name} | Killspiel</title>
</svelte:head>

<div class="container mx-auto">
	<div class=" p-2">
		<h1 class="text-6xl text-center">{data.name}</h1>
		<hr class="p-2" />

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
				Rate:
				<div class="text-5xl text-center">
					{data.guesses !== 0 ? ((data.points / data.guesses) * 100).toFixed(2) : 0}%
				</div>
			</div>
		</div>

		<div class="mt-6">
			<div class="table-container">
				<!-- Native Table Element -->
				<table class="table table-hover">
					<thead>
						<tr>
							<th>Game</th>
							<th>Schätzung</th>
							<th>Ergebnis</th>
							<th>Treffer</th>
						</tr>
					</thead>
					<tbody>
						{#each data.history as row}
							<tr>
								<td>{row.game}</td>
								<td>{row.guess}</td>
								<td>{row.correct}</td>
								<td>
									<svelte:component this={row.icon} />
								</td>
							</tr>
						{/each}
					</tbody>
				</table>
			</div>
		</div>
	</div>

	<div class="flex">
		<button class="ms-auto mr-3 btn variant-ghost" on:click={deleteUser}>Lösche User</button>
	</div>
</div>
