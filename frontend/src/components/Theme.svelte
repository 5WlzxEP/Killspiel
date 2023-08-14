<script lang="ts">
    import { storeTheme } from "@stores/theme"
    import { popup } from '@skeletonlabs/skeleton'
    import type { PopupSettings } from '@skeletonlabs/skeleton'


    const themes = [
        { type: 'skeleton', name: 'Skeleton', icon: '💀' },
        { type: 'modern', name: 'Modern', icon: '🤖' },
        { type: 'rocket', name: 'Rocket', icon: '🚀' },
        { type: 'seafoam', name: 'Seafoam', icon: '🧜‍♀️' },
        { type: 'vintage', name: 'Vintage', icon: '📺' },
        { type: 'sahara', name: 'Sahara', icon: '🏜️' },
        { type: 'hamlindigo', name: 'Hamlindigo', icon: '👔' },
        { type: 'gold-nouveau', name: 'Gold Nouveau', icon: '💫' },
        { type: 'crimson', name: 'Crimson', icon: '⭕' },
        { type: 'seasonal', name: 'Seasonal', icon: '🎆' },
    ]

    function setTheme(e: string) {
        console.log(e)
        const theme = e as string
        storeTheme.set(theme)
        document.cookie = `theme=${theme}; path:/`

        const r = document.querySelector("body")
        r.dataset.theme = theme
        // return async ({ result, update }) => {
        //     await update();
        //     if (result.type === 'success') {
        //         const theme = result.data?.theme as string;
        //         storeTheme.set(theme);
        //     }
        //     else {
        //         console.log("An error occurred")
        //     }
        // };
    }

    const popupFeatured: PopupSettings = {
        // Represents the type of event that opens/closed the popup
        event: 'click',
        // Matches the data-popup value on your popup element
        target: 'theme',
        // Defines which side of your trigger the popup will appear
        placement: 'bottom',
    }

</script>

<div>
    <!-- trigger -->
    <button class="btn hover:variant-soft-primary" use:popup={popupFeatured}>
        <i class="fa-solid fa-palette text-lg md:!hidden" />
        <span class="hidden md:inline-block">Theme</span>
        <i class="fa-solid fa-caret-down opacity-50" />
    </button>
    <!-- popup -->
    <div class="card p-4 w-60 shadow-xl" data-popup="theme">
        <div class="space-y-4">
            <nav class="list-nav p-4 -m-4 max-h-64 lg:max-h-[500px] overflow-y-auto">
<!--                <form use:enhance={setTheme}>-->
                    <ul>
                        {#each themes as { icon, name, type }}
                            <li>
                                <button
                                        class="option w-full h-full"
                                        name="theme"
                                        value={type}
                                        class:bg-primary-active-token={$storeTheme === type}
                                        on:click={() => {setTheme(type)}}
                                >
                                    <span>{icon}</span>
                                    <span>{name}</span>
                                </button>
                            </li>
                        {/each}
                    </ul>
<!--                </form>-->
            </nav>
        </div>
        <div class="arrow bg-surface-100-800-token" />
    </div>
    <div class="card p-4 w-72 shadow-xl" data-popup="popupFeatured">
        <div><p>Demo Content</p></div>
        <div class="arrow bg-surface-100-800-token" />
    </div>
</div>

