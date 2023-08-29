<script lang="ts">
    import {themeStore} from "@stores/theme";
    import {onMount} from "svelte";
    import {get} from "svelte/store";

    function change(theme: string) {
        themeStore.set(theme)
        const body = document.getElementsByTagName("body")[0]

        body.dataset.theme = theme.toLowerCase()
    }

    const Themes: Array<{type: string, name: string, icon: string}> = [
        { type: 'skeleton', name: 'Skeleton', icon: '💀' },
        { type: 'modern', name: 'Modern', icon: '🤖' },
        { type: 'rocket', name: 'Rocket', icon: '🚀' },
        { type: 'seafoam', name: 'Seafoam', icon: '🧜‍♀️' },
        { type: 'vintage', name: 'Vintage', icon: '📺' },
        { type: 'sahara', name: 'Sahara', icon: '🏜️' },
        { type: 'hamlindigo', name: 'Hamlindigo', icon: '👔' },
        { type: 'gold-nouveau', name: 'Gold Nouveau', icon: '💫' },
        { type: 'crimson', name: 'Crimson', icon: '⭕' },
        {type: "wintry", name: "wintry", icon: "🌨️"}
    ]

    let selected: string

    onMount(() => {
        selected = get(themeStore)
    })
</script>

<div class="p-4 w-1/2">
    <label class="label">
        <span>Theme</span>
        <select class="select" bind:value={selected} on:change={() => change(selected)}>
            {#each Themes as theme}
                <option value="{theme.type}">{theme.icon} {theme.name}</option>
            {/each}
        </select>
    </label>
</div>