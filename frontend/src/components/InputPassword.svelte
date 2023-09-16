<script lang="ts">
	import { getModalStore, type ModalSettings } from "@skeletonlabs/skeleton"
	import { IconHelp, IconEye, IconEyeOff } from "@tabler/icons-svelte"

	const modalStore = getModalStore()

	let ModalSett: ModalSettings

	export let value: string
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

	let input: HTMLInputElement

	let hidden = true

	$: {
		if (input)
			input.type = hidden ? "password" : "text"
	}

</script>

<label class="label">
	<span
	>{label}
		{#if modal}
			<button
				on:click={() => {
					modalStore.trigger(ModalSett)
				}}
				class="inline-block p-1"
				type="button"
			>
				<IconHelp class="inline-block" />
			</button>
		{/if}
	</span>
	<div class="input-group input-group-divider grid-cols-[1fr_auto]">

		<input type="password" autocomplete="off" bind:this={input} bind:value tabindex="0" {placeholder} required />
		<div class="input-group-shim">
		<button on:click|preventDefault={() => hidden = !hidden}>
		{#if (hidden)}
			<IconEyeOff />
		{:else }
			<IconEye />
		{/if}
		</button>*
		</div>
	</div>
</label>
