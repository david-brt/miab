<script lang="ts">
  import Modal from '$lib/components/Modal.svelte';
  import type { LayoutData } from './$types';
  import { showModal, user, token } from '$lib/stores';
  import Dropdown from '$lib/components/Dropdown.svelte';
  import Form from '$lib/components/Form.svelte';
  import { loginSchema } from '$lib/schema/loginSchema';
  import { signupSchema } from '$lib/schema/signupSchema';
  import { renameSchema } from '$lib/schema/renameSchema';

  export let data: LayoutData;
  user.set(data.user);
  token.set(data.token);

  let showPopup = false;

  const formFields = {
    login: [
      { name: 'username', placeholder: 'username', type: 'text', required: true, maxlength: 20 },
      {
        name: 'password',
        placeholder: 'password',
        type: 'password',
        required: true,
        maxlength: 50,
        minlength: 8
      }
    ],
    signup: [
      { name: 'username', placeholder: 'username', type: 'text', required: true, maxlength: 20 },
      {
        name: 'password',
        placeholder: 'password',
        type: 'password',
        required: true,
        maxlength: 50,
        minlength: 8
      },
      {
        name: 'confirmation',
        placeholder: 'confirm password',
        type: 'password',
        required: true,
        maxlength: 50,
        minlength: 8
      }
    ],
    rename: [
      {
        name: 'username',
        placeholder: 'new username',
        type: 'text',
        required: true,
        maxlength: 20
      }
    ]
  };

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
          <Form
            fields={formFields.login}
            validationSchema={loginSchema}
            actionRoute="?/login"
            modalType="login"
          />
        </Modal>
        <button on:click={() => onClick('signup')}>Signup</button>
        <Modal modalType={'signup'}>
          <Form
            fields={formFields.signup}
            validationSchema={signupSchema}
            actionRoute="?/signup"
            modalType="signup"
          />
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
    <Form
      fields={formFields.rename}
      validationSchema={renameSchema}
      actionRoute="?/rename"
      modalType="namechange"
    />
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
    gap: 1em;
    width: 70%;
  }

  :global(.form-input) {
    border: none;
    flex: 1;
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
