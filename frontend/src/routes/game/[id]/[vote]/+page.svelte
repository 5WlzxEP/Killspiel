<script lang="ts">
	import { goto } from "$app/navigation"
	import type { User } from "./+page"

	/** @type {import("./$types").PageData} */
	export let data: { data: Array<User>, params: {id: number, vote: number} }
</script>

<svelte:head>
	<title>Spieler mit Schätzung {data.params.vote} in Game {data.params.id} | Killspiel</title>
</svelte:head>

<div class="container mx-auto w-2/5">
	<div class="card mt-6 m-2 w-2/3 mx-auto">
		{#if Array.isArray(data.data)}
			<div class="table-container">
				<table class="table table-hover">
					<thead>
						<tr>
							<th>Name</th>
							<th class="text-right">Vote</th>
						</tr>
					</thead>
					<tbody class="cursor-pointer">
						{#each data.data as row}
							<tr on:click={() => goto("/user/" + row.id.toString())}>
								<td>{row.name}</td>
								<td class="text-right">{row.vote.toFixed(2)}</td>
							</tr>
						{/each}
					</tbody>
					<tfoot>
						<tr>
							<th>Total</th>
							<th class="text-right">{data.data?.length}</th>
						</tr>
					</tfoot>
				</table>
			</div>
		{:else if data.data === null}
			<div class="p-4">Keine Ergebnisse gefunden</div>
		{/if}
	</div>
</div>
