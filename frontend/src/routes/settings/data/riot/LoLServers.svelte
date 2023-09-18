<script lang="ts">
	import { IconHelp } from "@tabler/icons-svelte"
	import { getModalStore, type ModalSettings } from "@skeletonlabs/skeleton"
	import type { Server } from "./+page"

	const options = [
		"BR1",
		"EUN1",
		"EUW1",
		"JP1",
		"KR",
		"LA1",
		"LA2",
		"NA1",
		"OC1",
		"PH2",
		"RU",
		"SG2",
		"TH2",
		"TR1",
		"TW2",
		"VN2"
	]
	const modalStore = getModalStore()

	const ModalSett: ModalSettings = {
		type: "alert",
		title: "Server",
		body: "Server auf dem der Account ist",
		buttonTextCancel: "Schließen"
	}

	export let value: Server
</script>

<label class="label">
	<span>
		Server
		<button
			on:click={() => {
				modalStore.trigger(ModalSett)
			}}
			class="inline-block p-1"
			type="button"
		>
			<IconHelp class="inline-block" />
		</button>
	</span>
	<div class="input-group input-group-divider grid-cols-[1fr_auto]">
		<select class="select" bind:value>
			{#each options as option}
				<option value={option.toLowerCase()}>{option}</option>
			{/each}
		</select>
		<div class="input-group-shim">*</div>
	</div>
</label>
