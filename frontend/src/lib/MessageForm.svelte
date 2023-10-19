<script lang="ts">
  import { PUBLIC_DATA_ROUTE } from '$env/static/public'
  import { handleSubmit } from '$lib/utils/form'
  import { showModal } from './stores';

  async function callHandleSubmit(e: SubmitEvent) {
    const response = await handleSubmit(e, `${PUBLIC_DATA_ROUTE}/send-message`)
    if(response.status === 202) {
      showModal.update(previousState => ({ ...previousState, message: false }))
    }
  }
</script>

<form class="modal-form" on:submit|preventDefault={callHandleSubmit}>
	<label for="message-input" class="form-label">message</label>
	<textarea
		name="content"
		id="message-input"
		placeholder="your message"
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
		placeholder="your name"
		maxlength="50"
		class="form-input"
	/>
	<button class="send-button">send</button>
</form>

<style>
	textarea:focus,
	input:focus {
		outline: none;
	}
</style>
