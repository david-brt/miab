import { writable } from 'svelte/store';

type ModalType = 'login' | 'signup' | 'message';

function createShowModal() {
	const { subscribe, update } = writable({
		login: false,
		signup: false,
		message: false
	});
	return {
		subscribe,
		set: (modalType: ModalType, value: boolean) =>
			update((previousState) => ({ ...previousState, [modalType]: value }))
	};
}

export const user = writable();

export const showModal = createShowModal();
