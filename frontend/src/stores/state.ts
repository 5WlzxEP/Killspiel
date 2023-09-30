import type { Writable } from "svelte/store"
import { writable } from "svelte/store"

export const state: Writable<number> = writable(0)
