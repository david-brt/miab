<script lang="ts">
  import { PUBLIC_DATA_ROUTE } from '$env/static/public'
  import { handleSubmit } from '$lib/utils/form'
  //closure to give handleSubmit access to the data route
  async function callHandleSubmit(e: SubmitEvent) {
    handleSubmit(e, `${PUBLIC_DATA_ROUTE}/send-message`)
  }
</script>

<form class="modal-form" on:submit|stopPropagation={callHandleSubmit}>
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
