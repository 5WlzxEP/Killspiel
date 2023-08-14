<script lang="ts">
    import {Modal, modalStore, type ModalSettings, type ToastSettings, toastStore, Toast} from '@skeletonlabs/skeleton'
    import {IconHelp} from "@tabler/icons-svelte"
    import {onMount} from "svelte";

    const BACKEND_URL = import.meta.env.VITE_BACKEND_URL

    let ApiKey: string
    let Channel: string
    let ChannelSender: string
    let Prefix: string
    let MsgBeginn: string
    let MsgEnd: string
    let MsgFinal: string

    const modalApiKey: ModalSettings = {
        type: 'alert',
        title: 'Twitch Api Key',
        body: 'Der Twitch Api Key ist nötig um im Chat Nachrichten zu schreiben. So kann der Begin angekündigt werden, auf Ende hingewiesen und das Ergebnis bekannt gegeben werden. ' +
            'Am einfachsten kann man den Twitch Api Key über <a class="underline" target="_blank" href="https://twitchapps.com/tmi/">https://twitchapps.com/tmi/</a> möglich.',
        buttonTextCancel: "Schließen",
    }
    const modalPrefix: ModalSettings = {
        type: 'alert',
        title: 'Prefix',
        body: 'Der Prefix gibt an, womit eine Nachricht beginnen muss, damit sie als Guess registriert wird. Der Prefix kann auch leer sein, dann wird jede Nachricht, die nur eine Zahl darstellt ausgewertet.',
        buttonTextCancel: "Schließen",
    }
    const modalChannel: ModalSettings = {
        type: 'alert',
        title: 'Twitch Channel',
        body: 'Dies ist der Twitchchannel mit dessen Chat interagiert wird.',
        buttonTextCancel: "Schließen",
    }
    const modalChannelApi: ModalSettings = {
        type: 'alert',
        title: 'Twitch Account Api',
        body: 'Dies ist der Twitchchannel von dem der Api Key stammt. Dieser Account wird die Nachrichten senden. Wenn keiner gesetzt wird, wird der Twitchchannel genutzt, wessen Chat genutzt wird. <br> Um Twitchchatfeatures wie /announce zu nutzen, muss der Account auf dem Channel Mod-Rechte haben.',
        buttonTextCancel: "Schließen",
    }
    const modalBeginn: ModalSettings = {
        type: 'alert',
        title: 'Nachricht zum Begin der Erhebung',
        body: 'Diese Nachricht wird zum Beginn der Erhebung, also wenn die Zuschauer ihre Schätzungen abgeben können, in den Chat gepostet. <br> Es können Twitchfeatures wie /announce genutzt werden, jedoch sind für diese eventuell Rechte nötig.',
        buttonTextCancel: "Schließen",
    }
    const modalEnd: ModalSettings = {
        type: 'alert',
        title: 'Nachricht zum Ende der Erhebung',
        body: 'Diese Nachricht wird zum Ende der Erhebung in den Chat gepostet. <br> Es können Twitchfeatures wie /announce genutzt werden, jedoch sind für diese eventuell Rechte nötig.',
        buttonTextCancel: "Schließen",
    }
    const modalFinal: ModalSettings = {
        type: 'alert',
        title: 'Nachricht zum Auflösen der richtigen Schätzungen',
        body: 'Diese Nachricht wird versendet, wenn das Ergebnis feststeht. <br> <i>$WINNERS</i> sind die Gewinner <br> <i>$RESULT</i> sind die erzielten Kills',
        buttonTextCancel: "Schließen",
    }

    async function submit(e: Event) {
        console.log(e)
        try {
            const res = await fetch(`${BACKEND_URL}/api/collector/chat/`, {
                headers: {
                    'Content-Type': 'application/json'
                },
                method: "POST",
                body: JSON.stringify({
                    "apiKey": ApiKey,
                    "channel": Channel,
                    "channelSender": ChannelSender,
                    "prefix": Prefix,
                    "msgBeginn": MsgBeginn,
                    "msgEnd": MsgEnd,
                    "msgFinal": MsgFinal,
                })
            })
            if (res.ok) {
                const t: ToastSettings = {
                    message: 'Einstellungen erfolgreich gespeichert.',
                    timeout: 5000,
                    background: 'variant-filled-success',
                }
                toastStore.trigger(t)
            }
        } catch (e) {
            console.error(e)
            const t: ToastSettings = {
                message: 'Es ist ein Fehler beim Abfragen der aktualisieren Einstellungen aufgetreten.',
                timeout: 5000,
                background: 'variant-filled-error',
            }
            toastStore.trigger(t)
        }
    }

    onMount(async () => {
        const url = `${BACKEND_URL}/api/collector/chat/`
        try {
            const res = await fetch(url)
            const val = await res.json()
            ApiKey = val.apiKey
            Channel = val.channel
            ChannelSender = val.channelApi
            Prefix = val.prefix
            MsgBeginn = val.msgBeginn
            MsgEnd = val.msgEnd
            MsgFinal = val.msgFinal
        } catch (e) {
            const t: ToastSettings = {
                message: 'Es ist ein Fehler beim Abfragen der aktuellen Einstellungen aufgetreten.',
                timeout: 5000,
                background: 'variant-filled-error',
            }
            toastStore.trigger(t)
        }
    })

</script>

<Toast/>
<div class="container mx-auto">
    <div class="card w-full">
        <h1 class="text-center text-2xl p-3">Twitchchat Settings</h1>
        <hr>
        <form class="p-5" on:submit|preventDefault={submit}>
            <div class="grid lg:!grid-cols-2 gap-10 w-full m-2">
                <label class="label">
                        <span>Twitch Channel
                            <button on:click={() => {modalStore.trigger(modalChannel)}} class="inline-block p-1"
                                    type="button">
                            <IconHelp class="inline-block"/>
                            </button>
                        </span>
                    <div class="input-group input-group-divider grid-cols-[auto_1fr_auto]">
                        <div class="input-group-shim">twitch.tv/</div>
                        <input type="text" bind:value={Channel} tabindex="0"
                               placeholder="5W_lzxEP" required/>
                        <div class="">*</div>

                    </div>
                </label>
                <label class="label">
                        <span>Prefix
                            <button on:click={() => {modalStore.trigger(modalPrefix)}} class="inline-block p-1"
                                    type="button">
                            <IconHelp class="inline-block"/>
                            </button>
                        </span>
                        <input class="input" type="text" bind:value={Prefix} tabindex="0"
                               placeholder="/guess"/>

                </label>
                <label class="label">
                    <span>Twitch Api Key
                        <button on:click={() => {modalStore.trigger(modalApiKey)}} class="inline-block p-1"
                                type="button">
                            <IconHelp class="inline-block mx-auto"/>
                        </button>
                    </span>
                    <div class="input-group input-group-divider grid-cols-[1fr_auto]">
                        <input class="input" type="text" bind:value={ApiKey}
                               placeholder="oauth:bcgf6ogc8swu329nmnqprwgdodizgw" required/>
                        <div class="">*</div>

                    </div>
                </label>

                <label class="label">
                        <span>Twitch Account (selber wie vom Api Key)
                            <button on:click={() => {modalStore.trigger(modalChannelApi)}} class="inline-block p-1"
                                    type="button">
                            <IconHelp class="inline-block"/>
                            </button>
                        </span>
                    <div class="input-group input-group-divider grid-cols-[auto_1fr_auto]">
                        <div class="input-group-shim">twitch.tv/</div>
                        <input type="text" bind:value={ChannelSender}
                               placeholder="5W_lzxEP"/>
                    </div>

                </label>
                <label class="label">
                    <span>
                        <div>Nachricht zum Begin der Erhebung
                            <button on:click={() => {modalStore.trigger(modalBeginn)}} class="inline-block p-1"
                                    type="button">
                            <IconHelp class="inline-block"/>
                        </button>
                        </div>
                    </span>
                    <textarea class="textarea" rows="4" bind:value={MsgBeginn}
                              placeholder="/announceblue Das Killspiel hat begonnen. Nimm jetzt Teil mit /guess <Dein Guess>."/>
                </label>
                <label class="label">
                    <span>Nachricht zum Ende der Erhebung
                        <button on:click={() => {modalStore.trigger(modalEnd)}} class="inline-block p-1" type="button">
                            <IconHelp class="inline-block"/>
                        </button>
                    </span>

                    <textarea class="textarea" rows="4" bind:value={MsgEnd}
                              placeholder="/announceorange Das Voten ist abgeschloßen. Ab jetzt bitte keine Stimmen in der Chat mehr."/>
                </label>
                <label class="label">
                    <span>Nachricht zum Auflösen der richtigen Schätzungen
                        <button on:click={() => {modalStore.trigger(modalFinal)}} class="inline-block p-1" type="button">
                            <IconHelp class="inline-block"/>
                        </button>
                    </span>

                    <textarea class="textarea" rows="4" bind:value={MsgFinal}
                              placeholder="/announcegreen Das Killspiel ist beendet. Es wurden $RESULT Kills erzielt und somit haben $WINNERS gewonnen."/>
                </label>
            </div>
            <div class="flex">
                Mit * markierte Felder sind Pflicht
                <button class="btn variant-ghost-success ms-auto" type="submit">Speichern</button>
            </div>
        </form>
    </div>
</div>
<Modal/>