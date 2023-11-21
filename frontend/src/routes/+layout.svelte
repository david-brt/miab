<script lang="ts">
  import LoginForm from '$lib/components/LoginForm.svelte';
  import SignupForm from '$lib/components/SignupForm.svelte';
  import Modal from '$lib/components/Modal.svelte';
  import type { LayoutData } from './$types';
  import { showModal, user, token } from '$lib/stores';
  import Dropdown from '$lib/components/Dropdown.svelte';
  import RenameForm from '$lib/components/RenameForm.svelte';

  export let data: LayoutData;
  user.set(data.user);
  token.set(data.token);

  let showPopup = false;

  function onClick(modalType: keyof typeof $showModal) {
    showModal.set(modalType, true);
  }
</script>

<div class="container">
  <nav class="navbar">
    <div>
      {#if !$user}
        <button on:click={() => onClick('login')}>Login</button>
        <Modal modalType={'login'}>
          <LoginForm />
        </Modal>
        <button on:click={() => onClick('signup')}>Signup</button>
        <Modal modalType={'signup'}>
          <SignupForm />
        </Modal>
      {/if}
    </div>

    <div>
      {#if $user}
        <button on:click={() => (showPopup = true)}>{$user.name}</button>
      {/if}
    </div>
  </nav>
  {#if showPopup}
    <Dropdown bind:showPopup />
  {/if}
  <Modal modalType={'namechange'}>
    <RenameForm />
  </Modal>
  <slot />
</div>

<style>
  :global(body) {
    font-family: 'Lato';
    font-size: 12px;
    color: var(--col2);
    margin: 0;
    padding: 0;
    overflow: hidden;
    background-color: var(--col1);
    --border-radius: 0.5em;
    --border-width: 0.2em;
    --col1: #264653;
    --col2: #e9c46a;
    --col3: #e76f51;
    --col4: #457b9d;
    --pastel-yellow: #fffba6;

    --error-red: #c1121f;
  }
  :global(button) {
    background-color: var(--col3);
    border: none;
    border-radius: var(--border-radius);
    padding: 1em;
    font-size: 1.5em;
  }

  :global(button:hover) {
    cursor: pointer;
  }

  :global(::placeholder) {
    color: #a37b00;
  }

  :global(.modal-form) {
    display: flex;
    flex-direction: column;
    gap: 2em;
    width: 70%;
  }

  :global(.form-input) {
    border: none;
    border-radius: var(--border-radius);
    font-family: 'Lato';
    padding: 0.5em;
    resize: none;
    background-color: var(--pastel-yellow);
    font-size: 1.5em;
  }

  :global(.form-label) {
    display: none;
  }

  :global(.container) {
    margin: auto;
    display: flex;
    flex-direction: column;
    justify-content: center;
    padding: 1em;
  }

  :global(text) {
    color: white;
  }

  .navbar {
    display: flex;
    gap: 1em;
    justify-content: space-between;
  }

  :global(.errortext) {
    margin: 0;
    padding: 0;
    color: red;
  }
</style>
