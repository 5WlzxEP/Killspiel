import { localStorageStore } from '@skeletonlabs/skeleton'
import type {Writable} from "svelte/store"

export const storeTheme: Writable<string> = localStorageStore("storeTheme", "skeleton")