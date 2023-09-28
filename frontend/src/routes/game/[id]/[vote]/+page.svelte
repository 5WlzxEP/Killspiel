<script lang="ts">
	import { goto } from "$app/navigation"
	import type { User } from "./+page"

	/** @type {import("./$types").PageData} */
	export let data: { data: Array<User> }
</script>

<svelte:head>
	<title>Spieler Suche | Killspiel</title>
</svelte:head>

<div class="container mx-auto w-2/5">
	<div class="card mt-6 m-2 w-2/3 mx-auto">
		{#if Array.isArray(data.data)}
			<div class="table-container">
				<table class="table table-hover">
					<thead>
						<tr>
							<th>Id</th>
							<th class="text-right">Name</th>
						</tr>
					</thead>
					<tbody class="cursor-pointer">
						{#each data.data as row}
							<tr on:click={() => goto("/user/" + row.id.toString())}>
								<td>{row.id}</td>
								<td class="text-right">{row.name}</td>
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
