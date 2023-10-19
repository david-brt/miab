<script lang="ts">
	import Modal from '$lib/Modal.svelte';
	import MessageForm from '$lib/MessageForm.svelte';
  import { showModal } from '$lib/stores.js';

  function onClick() {
    showModal.update(previousState => ({ ...previousState, message: true}))
  }
	export let data;
</script>

<div class="content">
	<p class="greeting">
		Your message by {data.message.senderName}:
	</p>
	<p class="message">
		{data.message.content}
	</p>
  <button on:click={onClick}>Share your own idea</button>
  <Modal modalType={"message"}>
		<MessageForm />
	</Modal>
</div>

<style>
	.content {
		display: flex;
		align-items: center;
		justify-content: center;
		gap: 3em;
		flex-direction: column;
		height: 100vh;
	}

	.greeting {
		font-size: 1.5em;
	}

	.message {
		margin: 0;
		padding: 0 1em 0 1em;
		font-family: 'Lato';
		font-size: 4em;
	}
</style>
