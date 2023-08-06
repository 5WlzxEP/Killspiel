<script>
    import  {TabGroup, Tab} from "@skeletonlabs/skeleton"

    let tabSet = 0;

</script>

<div class="mx-auto container p-2">
    <div class="card container">
            <h1 class="text-center text-2xl p-3">Settings</h1>

<!--        <hr>-->
        <TabGroup>
            <Tab bind:group={tabSet} name="tab1" value={0}>
<!--                <svelte:fragment slot="lead">Allgemein</svelte:fragment>-->
                <span>Allgemein</span>
            </Tab>
            <Tab bind:group={tabSet} name="tab2" value={1}>(label 2)</Tab>
            <Tab bind:group={tabSet} name="tab3" value={2}>(label 3)</Tab>
            <!-- Tab Panels --->
            <svelte:fragment slot="panel">
                {#if tabSet === 0}
                    <a href="settings/collector">Collectors</a>
                {:else if tabSet === 1}
                    (tab panel 2 contents)
                {:else if tabSet === 2}
                    (tab panel 3 contents)
                {/if}
            </svelte:fragment>
        </TabGroup>
    </div>
</div>