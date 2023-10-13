<script lang="ts">
	import { onMount } from "svelte"
	import { IconCheck, IconX } from "@tabler/icons-svelte"
	import { state } from "@stores/state"

	const BACKEND_URL = import.meta.env.VITE_BACKEND_URL

	type ascDesc = "asc" | "desc"

	let page = 0
	let limit = 10
	let sortedBy: "p" | "n" | "g" = "p"
	let order: ascDesc = "asc"
	let latest = false

	type Leaderboard = {
		id: number
		rank: number
		name: string
		points: number
		guesses: number
		latest: number
	}

	type result = {
		meta: {
			size: number
		}
		data: Array<Leaderboard>
	}

	async function fetchLeaderboard(): Promise<void> {
		const url = `${BACKEND_URL}/api/leaderboard?p=${page * limit}&l=${limit}&s=${sortedBy}&o=${
			order === "asc" ? "0" : "1"
		}`
		const resp = await fetch(url)
		data = resp.json()
	}

	let data: Promise<result> = new Promise((_, r) => setTimeout(r, 3000))

	let name: HTMLTableCellElement
	let points: HTMLTableCellElement
	let teilnahmen: HTMLTableCellElement

	onMount(() => {
		init()
		fetchLeaderboard()
	})

	function init() {
		const params = new URLSearchParams(location.search)

		order = params.get("order") === "desc" ? "desc" : "asc"

		setSorting(params.get("sorting"))

		latest = params.get("latest") === "true"

		const l = params.get("limit")
		limit = l ? parseInt(l) : limit
	}

	function setSorting(sorting: string | null) {
		let target: HTMLTableCellElement
		switch (sorting) {
			case "name":
				sortedBy = "n"
				target = name
				break
			case "teilnahmen":
				sortedBy = "g"
				target = teilnahmen
				break
			default:
				sortedBy = "p"
				target = points
		}

		if (order === "asc") {
			target.classList.add("table-sort-asc")
		} else {
			target.classList.add("table-sort-dsc")
		}
	}

	$: if ($state === 3) fetchLeaderboard()
</script>

<h1 class="text-primary-500 text-center text-lg card rounded-none">Leaderboard</h1>

<div class="table-container mx-auto w-full rounded-none">
	<table class="table table-hover table-cell-fit mx-auto w-full rounded-none">
		<thead>
			<tr>
				<th>Rank</th>
				<th class="text-center" bind:this={name}>Name </th>
				<th class="text-center" bind:this={points}>Punkte </th>
				<th class="text-center" bind:this={teilnahmen}>Teilnahmen </th>
				<th class="text-center">Rate</th>
				{#if latest}
					<th class="text-right">letzten 8 Versuche</th>
				{/if}
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
					{#if latest}
						<td>
							<div class="placeholder animate-pulse" />
						</td>
					{/if}
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
						{#if latest}
							<td class="grid-cols-8 grid">
								{#each row.latest.toString(2).padStart(8, "0") as n}
									{#if n === "0"}
										<IconX color="red" />
									{:else}
										<IconCheck color="lime" />
									{/if}
								{/each}
							</td>
						{/if}
					</tr>
				{/each}
			{/await}
		</tbody>
	</table>
</div>
