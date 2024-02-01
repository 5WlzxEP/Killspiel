import {type Writable, writable} from "svelte/store"

export const latestVotes: Writable<{
	d: Array<{ time: Date; name: string; vote: number; color: string }>
}> = writable({d: []})
