import { localStorageStore } from "@skeletonlabs/skeleton"
import type { Writable } from "svelte/store"

export const themeStore: Writable<string> = localStorageStore("theme", "skeleton")
