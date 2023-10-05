<script lang="ts">
	import { onMount } from "svelte"
	import {
		getToastStore,
		type ToastSettings,
		type PaginationSettings
	} from "@skeletonlabs/skeleton"

	const toastStore = getToastStore()

	const BACKEND_URL = import.meta.env.VITE_BACKEND_URL

	type ascDesc = "asc" | "desc"

	let page = 0
	let limit = 10
	let sortedBy: "p" | "n" | "g" = "p"
	let order: { p: ascDesc; n: ascDesc; g: ascDesc } = {
		p: "asc",
		n: "asc",
		g: "asc"
	}
	const orderDefault = {
		p: "asc",
		n: "desc",
		g: "asc"
	} as const

	let currentOrder: ascDesc = "asc"

	type Leaderboard = {
		id: number
		rank: number
		name: string
		points: number
		guesses: number
		latest: number
	}

	type result = {
		meta: PaginationSettings
		data: Array<Leaderboard>
	}

	async function fetchLeaderboard(): Promise<result> {
		const url = `${BACKEND_URL}/api/leaderboard?p=${page * limit}&l=${limit}&s=${sortedBy}&o=${
			currentOrder === orderDefault[sortedBy] ? "0" : "1"
		}`
		try {
			const resp = await fetch(url)
			return resp.json()
		} catch (e) {
			const t: ToastSettings = {
				message: "Es ist ein Fehler beim Abfragen der Daten aufgetreten.",
				timeout: 5000,
				background: "variant-filled-error"
			}
			toastStore.trigger(t)
		}
		return new Promise((_, r) => r())
	}

	let data: Promise<result> = new Promise((_, r) => setTimeout(r, 2000))

	function update() {
		data = fetchLeaderboard()
	}

	let first: HTMLTableCellElement

	onMount(async () => {
		update()

		latest = first
	})

	function toggleAscDesc(s: ascDesc, target: HTMLTableCellElement): ascDesc {
		if (s === "asc") {
			target.classList.remove("table-sort-asc")
			target.classList.add("table-sort-dsc")
			return "desc"
		}

		target.classList.replace("table-sort-dsc", "table-sort-asc")
		return "asc"
	}

	let latest: HTMLTableCellElement

	function changeSorting(
		s: "g" | "p" | "n",
		event: MouseEvent & { currentTarget: EventTarget & HTMLTableCellElement }
	) {
		const target: HTMLTableCellElement = event.currentTarget
		if (sortedBy === s) {
			order[s] = toggleAscDesc(order[s], target)
		} else {
			latest.classList.remove("table-sort-dsc")
			latest.classList.remove("table-sort-asc")

			latest = target

			sortedBy = s
			if (order[s] === "desc") {
				target.classList.add("table-sort-dsc")
			} else {
				target.classList.add("table-sort-asc")
			}
		}
		currentOrder = order[s]
		update()
	}
</script>

<div class="table-container mx-auto w-full">
	<table class="table table-hover table-cell-fit mx-auto w-full m-1">
		<thead>
			<tr>
				<th class="p-1 m-1">Rank</th>
				<th class="text-center cursor-pointer" on:click={(e) => changeSorting("n", e)}>Name </th>
				<th
					class="text-center table-sort-asc cursor-pointer"
					bind:this={first}
					on:click={(e) => changeSorting("p", e)}
					>Punkte
				</th>
				<th class="text-center cursor-pointer" on:click={(e) => changeSorting("g", e)}
					>Teilnahmen
				</th>
				<th class="text-center">Rate</th>
			</tr>
		</thead>
		<tbody>
			{#await data}
				<tr>
					<td>
						<div class="placeholder animate-pulse" />
					</td>
					<td>
						<div class="placeholder animate-pulse" />
					</td>
					<td>
						<div class="placeholder animate-pulse" />
					</td>
					<td>
						<div class="placeholder animate-pulse" />
					</td>
					<td>
						<div class="placeholder animate-pulse" />
					</td>
				</tr>
			{:then res}
				{#each res.data as row}
					<tr>
						<td>{row.rank}</td>
						<td
							><a href="/user/{row.id}" class="w-full p-4" data-sveltekit-preload-data="tap"
								>{row.name}</a
							></td
						>
						<td class="text-right">{row.points}</td>
						<td class="text-right">{row.guesses}</td>
						<td class="text-right"
							>{((row.guesses > 0 ? row.points / row.guesses : 0) * 100).toFixed(2)} %
						</td>
					</tr>
				{/each}
			{:catch error}
				<tr>
					<td class="text-center" colspan="7"
						>Keine Daten oder keine Verbindung
						{error ? ": " + error : ""}
					</td>
				</tr>
			{/await}
		</tbody>
	</table>
</div>
