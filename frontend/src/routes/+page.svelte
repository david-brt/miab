<script lang="ts">
  import Modal from '$lib/components/Modal.svelte';
  import { showModal } from '$lib/stores';
  import { PUBLIC_DATA_ROUTE } from '$env/static/public';
  import Form from '$lib/components/Form.svelte';
  import { messageSchema } from '$lib/schema/messageSchema.js';

  export let data;

  const messageFields = [
    {
      name: 'message',
      placeholder: 'your message',
      required: true,
      maxlength: 240,
      textArea: true
    },
    {
      name: 'sender',
      placeholder: 'your name',
      type: 'text',
      required: true,
      maxlength: 20
    }
  ];

  async function fetchNewMessage() {
    const response = await fetch(`${PUBLIC_DATA_ROUTE}/global-message`);
    data.message = await response.json();
  }
</script>

<div class="content">
  <p class="greeting">
    Your message by {data.message.sender}:
  </p>
  <p class="message">
    {data.message.message}
  </p>
  <button on:click={() => showModal.set('message', true)}>Share your own idea</button>
  <Modal modalType={'message'}>
    <Form
      fields={messageFields}
      validationSchema={messageSchema}
      actionRoute="?/send-message"
      modalType="message"
    />
  </Modal>
  <button on:click={() => fetchNewMessage()}>Get new Message</button>
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
