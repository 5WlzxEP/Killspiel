<script lang="ts">
	import { IconSearch } from "@tabler/icons-svelte"
	import { getToastStore, type ToastSettings } from "@skeletonlabs/skeleton"
	import { goto } from "$app/navigation"

	const BACKEND_URL = import.meta.env.VITE_BACKEND_URL

	const toastStore = getToastStore()
	let name: string

	const toast: ToastSettings = {
		classes: "variant-filled-error",
		autohide: true,
		timeout: 3000,
		message: "Name muss mindestens 3 Zeichen lang sein."
	}

	const toastError: ToastSettings = {
		classes: "variant-filled-error",
		autohide: true,
		timeout: 3000,
		message: "Es ist ein Fehler beim Suchen aufgetreten."
	}

	let results: Array<{ id: number; name: string; points: number }> | undefined = undefined

	async function search() {
		if (name.length < 3) {
			toastStore.trigger(toast)
		}

		try {
			const url = `${BACKEND_URL}/api/user/`
			const res = await fetch(url, {
				method: "POST",
				headers: {
					"content-type": "application/json"
				},
				body: JSON.stringify({ name: name })
			})
			results = await res.json()
		} catch (e) {
			toastStore.trigger(toastError)
		}
	}
</script>

<svelte:head>
	<title>Spieler Suche | Killspiel</title>
</svelte:head>

<div class="container mx-auto w-2/5">
	<form role="search">
		<label class="m-2">
			<div class="p-2">Suche</div>
			<div class="input-group input-group-divider grid-cols-[1fr_auto]">
				<!-- svelte-ignore a11y-autofocus -->
				<input type="search" placeholder="5W_lzxEP" autofocus required bind:value={name} />
				<button class="variant-filled-secondary" on:click={search}>
					<IconSearch />
				</button>
			</div>
		</label>
	</form>
	<div class="card mt-6 m-2 w-2/3 mx-auto">
		{#if Array.isArray(results)}
			<div class="table-container">
				<table class="table table-hover">
					<thead>
						<tr>
							<th>Name</th>
							<th>Punkte</th>
						</tr>
					</thead>
					<tbody class="cursor-pointer">
						{#each results as row}
							<tr on:click={() => goto(row.id.toString())}>
								<td>{row.name}</td>
								<td class="text-right">{row.points}</td>
							</tr>
						{/each}
					</tbody>
					<tfoot>
						<tr>
							<th>Total</th>
							<th class="text-right">{results.length}</th>
						</tr>
					</tfoot>
				</table>
			</div>

			<!--		<Table source={table} text="text-right" interactive={true} on:selected={(e) => {goto("" +(e.detail[0]))}} />-->
		{:else if results === null}
			<div class="p-4">Keine Ergebnisse gefunden</div>
		{/if}
	</div>
</div>
