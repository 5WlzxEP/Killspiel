<script lang="ts">
	import "../../app.postcss"
	import {
		AppBar,
		AppShell,
		autoModeWatcher,
		initializeStores,
		LightSwitch,
		Toast
	} from "@skeletonlabs/skeleton"
	import { IconBadges, IconBrandGithub, IconSettings, IconUserSearch } from "@tabler/icons-svelte"
	import { onMount } from "svelte"
	import { get } from "svelte/store"
	import { themeStore } from "@stores/theme"
	import Modal from "@components/Modal.svelte"
	import { computePosition, autoUpdate, offset, shift, flip, arrow } from "@floating-ui/dom"
	import { storePopup } from "@skeletonlabs/skeleton"
	import Breadcrumbs from "@components/Breadcrumbs.svelte"
	import Websocket from "../Websocket.svelte"
	initializeStores()
	storePopup.set({ computePosition, autoUpdate, offset, shift, flip, arrow })

	// set theme
	onMount(() => {
		document.getElementsByTagName("body")[0].dataset.theme = get(themeStore)
	})
</script>

<svelte:head>
	{@html `<script>${autoModeWatcher.toString()} autoModeWatcher();</script>`}
</svelte:head>
<!-- App Shell -->
<AppShell>
	<svelte:fragment slot="pageHeader">
		<!-- App Bar -->
		<AppBar>
			<svelte:fragment slot="lead">
				<a href="/" class="p-1">
					<strong class="text-2xl uppercase">Killspiel</strong>
				</a>
			</svelte:fragment>
			<svelte:fragment slot="trail">
				<a href="/user" class="btn-icon variant-ghost" data-sveltekit-preload-code="viewport">
					<IconUserSearch />
				</a>
				<a
					href="/leaderboard"
					class="btn items-center variant-ghost"
					data-sveltekit-preload-code="eager"
					data-sveltekit-preload-data
				>
					<IconBadges />
					<p>Leaderboard</p>
				</a>
				<a
					href="/settings"
					class="p-2 btn-icon variant-ghost"
					data-sveltekit-preload-code="viewport"
				>
					<IconSettings />
				</a>
			</svelte:fragment>
		</AppBar>
	</svelte:fragment>
	<!-- Page Route Content -->
	<Breadcrumbs />

	<div class="mb-2">
		<slot />
	</div>

	<svelte:fragment slot="pageFooter">
		<AppBar>
			<svelte:fragment slot="lead">
				<a href="https://github.com/5WlzxEP/Killspiel"><IconBrandGithub class="inline-block m-1 -translate-y-0.5"/>Github</a>
			</svelte:fragment>
			<svelte:fragment slot="trail">
				<div class="p-4">
					<LightSwitch />
				</div>
			</svelte:fragment>
		</AppBar>
	</svelte:fragment>
</AppShell>
<Modal />
<Toast />
<Websocket />
