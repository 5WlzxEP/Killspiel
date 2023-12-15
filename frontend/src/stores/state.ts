import type {Writable} from "svelte/store"
import {writable} from "svelte/store"

export const state: Writable<number> = writable(0)

export const collectionEnd: Writable<number> = writable(0)

export function printState(state: number): string {
	switch (state) {
		case 0:
			return "Bereit"
		case 1:
			return "Sammle Schätzungen"
		case 2:
			return "Warte auf Auflösung"
		case 3:
			return "Verkünde Ergebnis"
	}
	return "Unbekannter Status"
}
