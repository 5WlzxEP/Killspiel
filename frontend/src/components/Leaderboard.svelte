<script lang="ts">
	import { onMount, type SvelteComponentTyped } from "svelte"
	import {
		getToastStore,
		type ToastSettings,
		type PaginationSettings,
		Paginator
	} from "@skeletonlabs/skeleton"
	import Check from "@components/Check.svelte"
	import X from "@components/X.svelte"

	const toastStore = getToastStore()

	const BACKEND_URL = import.meta.env.VITE_BACKEND_URL

	type ascDesc = "asc" | "desc"

	export let page = 0
	export let limit = 25
	export let sortedBy: "p" | "n" | "g" = "p"
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

	let meta: PaginationSettings = {
		page: page,
		limit: limit,
		size: 1,
		amounts: [10, 25, 50, 100]
	}

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

	async function fetchLeaderboard(): Promise<Array<Leaderboard>> {
		const url = `${BACKEND_URL}/api/leaderboard?p=${meta.page * meta.limit}&l=${
			meta.limit
		}&s=${sortedBy}&o=${currentOrder === orderDefault[sortedBy] ? "0" : "1"}`
		try {
			const resp = await fetch(url)
			const ob: result = await resp.json()
			meta.size = ob.meta.size
			return ob.data
		} catch (e) {
			const t: ToastSettings = {
				message: "Es ist ein Fehler beim Abfragen der Daten aufgetreten.",
				timeout: 5000,
				background: "variant-filled-error"
			}
			toastStore.trigger(t)
		}
		return []
	}

	let data: Array<Leaderboard> = []

	async function update() {
		// data = []
		data = await fetchLeaderboard()
	}

	let first: HTMLTableCellElement

	onMount(async () => {
		await update()

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

	function changeSorting(s: "g" | "p" | "n", event: MouseEvent & { currentTarget: EventTarget & HTMLTableCellElement; }) {
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

	function* generateLatest(n: number): Generator<SvelteComponentTyped> {
		for (let i = 0; i < 8; i++) {
			yield n % 2 === 1 ? Check : X
			n = n >> 1
		}
	}
</script>

<div class="table-container mx-auto w-full m-1">
	<table class="table table-hover table-cell-fit mx-auto w-full m-1">
		<thead>
			<tr>
				<th class="p-1 m-1 w-[10%]">Rank</th>
				<th class="text-center w-[15%] cursor-pointer" on:click={(e) => changeSorting("n", e)}
					>Name
				</th>
				<th
					class="text-center table-sort-asc w-[10%] cursor-pointer"
					bind:this={first}
					on:click={(e) => changeSorting("p", e)}
					>Punkte
				</th>
				<th class="text-center w-[10%] cursor-pointer" on:click={(e) => changeSorting("g", e)}
					>Teilnahmen
				</th>
				<th class="text-center w-[10%]">Rate</th>
				<th class="text-right w-[15%] min-w-[200px]">letzten 8 Versuche</th>
			</tr>
		</thead>
		<tbody>
			{#each data as row}
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
					<td class="grid-cols-8 grid">
						{#each [...generateLatest(row.latest)] as icon}
							<svelte:component this={icon} />
						{/each}
					</td>
				</tr>
			{:else}
				{#each { length: meta.size } as _}
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
						<td>
							<div class="placeholder animate-pulse" />
						</td>
					</tr>
				{/each}
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
