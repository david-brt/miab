<script lang='ts'>
  import LoginForm from "$lib/components/LoginForm.svelte";
  import SignupForm from "$lib/components/SignupForm.svelte";
  import Modal from "$lib/components/Modal.svelte";
  import type { LayoutData } from "./$types";
  import { showModal, user } from "$lib/stores";

  export let data: LayoutData;
  user.set(data.user)

  function onClick(modalType: keyof typeof $showModal) {
    showModal.set(modalType, true)
  }
</script>

<div class="container">
	<nav class="navbar">
    {#if !$user}
    <button on:click={() => onClick("login")}>Login</button>
    <Modal modalType={"login"}>
      <LoginForm />
    </Modal>
		<button on:click={() => onClick("signup")}>Signup</button>
    <Modal modalType={"signup"}>
      <SignupForm />
    </Modal>
    {/if}
    {#if $user}
      <p>you are logged in as {$user.name}</p>
    {/if}

  </nav>
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
		--bg-yellow: #fffce6;
		--pastel-yellow: #fffba6;
		--brownish-yellow: #a37b00;
		--orange: #ffa540;
		--almost-red: #ff605b;
		--border-radius: 0.5em;
		--orange: #f79b24;
		--col1: #264653;
		--col2: #e9c46a;
		--col3: #e76f51;
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
	}

	:global(text){
		color: white;
	}

  .navbar {
    width: 100%;
    padding: 1em;
  }

	:global(.errortext){
		margin: 0;
		padding: 0;
		color: red;
	}

</style>
