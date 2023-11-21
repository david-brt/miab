<script lang="ts">
  import { PUBLIC_DATA_ROUTE } from '$env/static/public';
  import { handleSubmit } from '$lib/utils/form';
  import SubmitError from './SubmitError.svelte';
  import { showModal } from '../stores';

  const INTERNALSERVERERROR = 500;
  const ACCEPTED = 202;
  let messageStatus: number;

  async function callHandleSubmit(e: SubmitEvent) {
    const response = await handleSubmit(e, `${PUBLIC_DATA_ROUTE}/send-message`);
    messageStatus = response.status;
    if (messageStatus === ACCEPTED) {
      showModal.set('message', false);
    }
  }
</script>

<form class="modal-form" on:submit|preventDefault={callHandleSubmit}>
  <label for="message-input" class="form-label">message</label>
  <textarea
    name="content"
    id="message-input"
    placeholder="your message"
    required
    maxlength="240"
    rows="3"
    class="form-input"
    spellcheck="false"
  />
  <label for="sender-input" class="form-label">name</label>
  <input
    name="senderName"
    id="sender-input"
    type="text"
    required
    placeholder="your name"
    maxlength="50"
    class="form-input"
  />
  {#if messageStatus === INTERNALSERVERERROR}
    <SubmitError>Something went wrong. Try again later</SubmitError>
  {/if}
  <button class="send-button">send</button>
</form>

<style>
  textarea:focus,
  input:focus {
    outline: none;
  }
</style>
