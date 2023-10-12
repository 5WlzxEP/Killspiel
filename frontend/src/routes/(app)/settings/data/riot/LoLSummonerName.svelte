<script lang="ts">
	import { getModalStore, type ModalSettings } from "@skeletonlabs/skeleton"
	import { IconHelp } from "@tabler/icons-svelte"

	const modalStore = getModalStore()

	let ModalSett: ModalSettings

	export let value: string
	export let prefix = -1
	export let placeholder = ""
	export let label: string
	export let modal: {
		title: string
		body: string
	}

	if (modal) {
		ModalSett = {
			type: "alert",
			title: modal.title,
			body: modal.body,
			buttonTextCancel: "Schließen"
		}
	}
</script>

<label class="label">
	<span
		>{label}
		{#if modal}
			<div
				on:keypress
				on:click={() => {
					modalStore.trigger(ModalSett)
				}}
				class="inline-block p-1 cursor-pointer"
			>
				<IconHelp class="inline-block" />
			</div>
		{/if}
	</span>
	<div class="input-group input-group-divider grid-cols-[auto_1fr_auto]">
		{#if prefix === -1 || !Number.isInteger(prefix)}
			<div class="input-group-shim" />
		{:else}
			<div class="input-group-shim">
				<img
					src={`https://ddragon.leagueoflegends.com/cdn/13.18.1/img/profileicon/${prefix}.png`}
					alt="Profile icon"
					class="w-10 rounded-full"
				/>
			</div>
		{/if}
		<input type="text" bind:value tabindex="0" {placeholder} required />
		<div class="input-group-shim">*</div>
	</div>
</label>
