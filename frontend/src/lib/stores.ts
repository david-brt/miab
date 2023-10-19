import { writable } from 'svelte/store';

export const showModal = writable({
	login: false,
  signup: false,
	message: false
});
