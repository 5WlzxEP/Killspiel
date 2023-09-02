<script lang="ts">
	import { IconBrandTwitch } from "@tabler/icons-svelte"

	let scale: number
</script>

<div class="container grid grid-cols-2 mx-auto gap-4">
	<a href="/settings/collector/twitchchat" class="card p-2">
		<div class="flex gap-2">
			<IconBrandTwitch />
			<h1>Twitchchat</h1>
		</div>
		Lass die Teilnehmer über den Twitchchat abstimmen.
	</a>
</div>

<div class="container mx-auto m-2 justify-center items-center flex mt-6">
	<div class="card p-4 w-2/3">
		<div class="m-2">
			<h2 class="text-xl">Einstellungen</h2>
		</div>
		<hr />
		<form>
			<label class="label m-2">
				<span>Dauer</span>
				<div class="input-group input-group-divider grid-cols-[1fr_auto]">
					<input title="3" type="number" placeholder="3" />
					<select bind:value={scale}>
						<option value="60">min</option>
						<option value="1">s</option>
					</select>
				</div>
			</label>
			<div class="flex">
				<button class="btn variant-ghost-success ms-auto" type="submit">Speichern</button>
			</div>
		</form>
	</div>
</div>
