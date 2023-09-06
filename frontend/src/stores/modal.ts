import { localStorageStore } from "@skeletonlabs/skeleton"
import type { Writable } from "svelte/store"

export const modalButtonPositive: Writable<string> = localStorageStore(
	"modalButtonP",
	"variant-filled-error"
)
