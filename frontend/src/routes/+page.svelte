<script lang="ts">
	import Leaderboard from "./Leaderboard.svelte"
	import Games from "./Games.svelte"
</script>

<svelte:head>
	<title>Killspiel</title>
</svelte:head>

<div class="container mx-auto mt-2 grid grid-cols-2 gap-3">
	<div class="card p-3">Main Page, still to fill</div>
	<div>
		<h1 class="text-lg text-center anchor"><a href="/">Aktuelles Spiel</a></h1>
	</div>
	<div>
		<h1 class="text-lg text-center anchor"><a href="/leaderboard">Leaderboard</a></h1>
		<Leaderboard />
	</div>
	<div>
		<h1 class="text-lg text-center anchor"><a href="/game">Letzte Spiele</a></h1>
		<Games />
	</div>
</div>
