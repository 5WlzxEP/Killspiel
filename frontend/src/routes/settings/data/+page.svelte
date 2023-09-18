<svelte:head>
	<title>Daten Collector | Killspiel</title>
</svelte:head>

<div class="container grid grid-cols-2 mx-auto gap-4">
	<a href="/settings/data/riot" class="card p-2">
		<div class="p-1">
			<h1 class="text-lg font-bold">Rito Api</h1>
			Frage die Riot Api ab, ob gerade ein Spiel beginnt, gespielt wird und wie viele Kills erzielt wurden
			oder welche KDA erreicht wurde.
		</div>
	</a>

	<a href="/settings/data/manual" class="card p-2">
		<div class="p-1">
			<h1 class="text-lg font-bold">Manuell</h1>
			Entscheide selbstständig wann das Killspiel beginnt und was das Ergebnis ist.
		</div>
	</a>
</div>
