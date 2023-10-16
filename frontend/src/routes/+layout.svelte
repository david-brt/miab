<script>
  import LoginForm from "$lib/LoginForm.svelte";
import Modal from "$lib/Modal.svelte";
  let showModal = false
</script>

<div class="container">
	<nav class="navbar">
    <button on:click={() => (showModal = true)}>Login</button>
    <Modal bind:showModal>
      <LoginForm />
    </Modal>
  </nav>
	<slot />
</div>

<style>
	:global(body) {
		font-family: 'Lato';
		font-size: 12px;
		margin: 0;
		padding: 0;
		background-color: var(--salmon);
		--bg-yellow: #fffce6;
		--pastel-yellow: #fffba6;
		--brownish-yellow: #a37b00;
		--orange: #ffa540;
		--almost-red: #ff605b;
		--border-radius: 0.5em;
		--orange: #f79b24;
	}
	:global(button) {
		background-color: var(--pastel-yellow);
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

  .navbar {
    width: 100%;
    padding: 1em;
  }
</style>
