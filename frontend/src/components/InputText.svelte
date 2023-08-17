<script lang="ts">
    import {modalStore, type ModalSettings} from "@skeletonlabs/skeleton";
    import {IconHelp} from "@tabler/icons-svelte";

    let ModalSett: ModalSettings

    export let value: string
    export let prefix = ''
    export let required = false
    export let placeholder = ""
    export let label: string
    export let modal: {
        title: string,
        body: string
    } | undefined = undefined

    if (modal) {
        ModalSett = {
            type: 'alert',
            title: modal.title,
            body: modal.body,
            buttonTextCancel: "Schließen",
        }
    }
</script>

<label class="label">
    <span>{label}
        {#if modal}
            <button on:click={() => {modalStore.trigger(ModalSett)}} class="inline-block p-1"
                    type="button">
            <IconHelp class="inline-block"/>
            </button>
        {/if}
    </span>
    <div class="input-group input-group-divider grid-cols-[{prefix ? 'auto_': ''}1fr_auto]">
        {#if prefix}
            <div class="input-group-shim">{prefix}</div>
        {/if}
        <input type="text" bind:value={value} tabindex="0"
               placeholder="{placeholder}"
               {required}
        />
        {#if required}
            <div class="">*</div>
        {/if}
    </div>
</label>