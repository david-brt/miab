import { writable, type Writable } from 'svelte/store';

type ModalType = 'login' | 'signup' | 'message' | 'namechange';

function createShowModal() {
  const { subscribe, update } = writable({
    login: false,
    signup: false,
    message: false,
    namechange: false
  });
  return {
    subscribe,
    set: (modalType: ModalType, value: boolean) =>
      update((previousState) => ({ ...previousState, [modalType]: value }))
  };
}

type UserType = {
  id: string;
  name: string;
};

export const user: Writable<UserType | undefined> = writable();

export const showModal = createShowModal();
