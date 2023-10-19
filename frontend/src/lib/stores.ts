import { writable } from 'svelte/store';

export const showModal = writable({
	login: false,
	message: false
});
