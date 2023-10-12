<script lang="ts">
	import { page } from "$app/stores"
</script>

<svelte:head>
	<title>{$page.status} Game does not exist | Killspiel</title>
</svelte:head>

<div class="container mx-auto p-2">
	<h1 class="text-6xl text-center">{$page.status} {$page.error?.message}</h1>
	<hr class="p-2 mt-3" />
	<div class="text-center text-2xl">
		<p>Game with the id {$page.params.id} does not exist.</p>
	</div>
</div>
