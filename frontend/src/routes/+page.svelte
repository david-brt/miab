<script lang="ts">
	import Modal from '$lib/components/Modal.svelte';
	import MessageForm from '$lib/components/MessageForm.svelte';
  import { showModal} from '$lib/stores.js';

	export let data;

</script>

<div class="content">
	<p class="greeting">
		Your message by {data.message.senderName}:
	</p>
	<p class="message">
		{data.message.content}
	</p>
  <button on:click={() => showModal.set('message', true)}>Share your own idea</button>
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
