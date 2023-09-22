<script lang="ts">
	import { getModalStore, type ModalSettings } from "@skeletonlabs/skeleton"
	import { IconHelp } from "@tabler/icons-svelte"
	import { onMount } from "svelte"

	const modalStore = getModalStore()

	let ModalSett: ModalSettings

	export let value: string
	export let prefix = ""
	export let required = false
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

	let div: HTMLDivElement

	function triggerModal() {
		modalStore.trigger(ModalSett)
	}

	onMount(() => {
		div.classList.add(prefix ? "grid-cols-[auto_1fr_auto]" : "grid-cols-[1fr_auto]")
	})
</script>

<label class="label">
	<span>
		{label}
		{#if modal}
			<div
				on:keypress={triggerModal}
				on:click={triggerModal}
				class="inline-block p-1 cursor-pointer"
			>
				<IconHelp class="inline-block" />
			</div>
		{/if}
	</span>
	<div bind:this={div} class="input-group input-group-divider">
		{#if prefix}
			<div class="input-group-shim">{prefix}</div>
		{/if}
		<input type="text" bind:value tabindex="0" {placeholder} {required} />
		{#if required}
			<div class="input-group-shim">*</div>
		{/if}
	</div>
</label>
