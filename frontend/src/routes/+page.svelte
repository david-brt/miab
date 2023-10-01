<script lang="ts">
	import SendMessageModal from '$lib/SendMessageModal.svelte';
	import MessageForm from '$lib/MessageForm.svelte';

	let showModal = false;
	export let data;
</script>

<div class="content">
	<p class="greeting">
		Your message by {data.message.sender}:
	</p>
	<p class="message">
		{data.message.content}
	</p>
	<button on:click={() => (showModal = true)}>Share your own idea</button>
	<SendMessageModal bind:showModal>
		<MessageForm />
	</SendMessageModal>
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
