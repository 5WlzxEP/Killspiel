<script lang="ts">
	import { getModalStore, type ModalSettings } from "@skeletonlabs/skeleton"
	import { IconHelp, IconEye, IconEyeOff } from "@tabler/icons-svelte"
	import { onMount } from "svelte"

	const modalStore = getModalStore()

	let ModalSett: ModalSettings

	export let value: string
	export let prefix = ""

	export let required = true
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
	let div: HTMLDivElement

	function trigger() {
		modalStore.trigger(ModalSett)
	}

	$: {
		if (input) input.type = hidden ? "password" : "text"
	}

	onMount(() => {
		div.classList.add(prefix ? "grid-cols-[auto_1fr_auto]" : "grid-cols-[1fr_auto]")
	})
</script>

<label class="label">
	<span
		>{label}
		{#if modal}
			<div
				on:keypress={trigger}
				on:click={trigger}
				role="button"
				tabindex="-1"
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
		<input
			type="password"
			autocomplete="off"
			bind:this={input}
			bind:value
			tabindex="0"
			{placeholder}
			{required}
		/>
		<div class="input-group-shim">
			<button on:click|preventDefault={() => (hidden = !hidden)} type="button">
				{#if hidden}
					<IconEyeOff />
				{:else}
					<IconEye />
				{/if}
			</button>{required ? "*" : ""}
		</div>
	</div>
</label>
