<script lang="ts">
  import { showModal } from '$lib/stores';

  type ShowModalKey = keyof typeof $showModal;
  export let modalType: ShowModalKey;
  let dialog: HTMLDialogElement;

  function handleKeyDown(e: KeyboardEvent) {
    if (e.key === 'Escape') {
      showModal.set(modalType, false);
    }
  }

  $: if (dialog && $showModal[modalType]) dialog.showModal();
  $: if (dialog && !$showModal[modalType]) dialog.close();
</script>

<dialog
  bind:this={dialog}
  on:close={() => showModal.set(modalType, false)}
  on:click|self={() => showModal.set(modalType, false)}
  on:keydown={handleKeyDown}
  class="modal-dialog"
  on:click|stopPropagation
>
  <div class="content">
    <button class="close-button" on:click={() => showModal.set(modalType, false)}>
      <img width="30px" src="icons/x-icon.svg" alt="" />
    </button>
    <slot />
  </div>
</dialog>

<style>
  .modal-dialog {
    background-color: var(--col2);
    border-radius: var(--border-radius);
    border: none;
    padding: 0;
    width: 95vw;
  }

  .content {
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
    padding: 1em;
    width: 90%;
    margin: auto;
  }

  .close-button {
    padding: 0;
    background: none;
    align-self: flex-end;
  }

  @media screen and (min-width: 330px) {
    .modal-dialog {
      width: 80vw;
    }
  }

  @media screen and (min-width: 768px) {
    .modal-dialog {
      width: 60vw;
    }
  }

  @media screen and (min-width: 1280px) {
    .modal-dialog {
      width: 30vw;
    }
  }
</style>
