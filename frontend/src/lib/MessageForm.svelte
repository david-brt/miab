<script lang="ts">
  import { PUBLIC_DATA_ROUTE } from '$env/static/public'
  async function handleSubmit(e: SubmitEvent) {
    const formData = new FormData(e.target as HTMLFormElement)
    const data: {[key: string]: FormDataEntryValue} = {};
    //deconstructs each field and adds it to data object
    for (let field of formData) {
      const [key, value] = field;
      data[key] = value;
    }
    await fetch(`${PUBLIC_DATA_ROUTE}/send-message`, {
      method: 'POST',
      mode: 'cors',
      headers: {
      'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
      credentials: 'include',
    })
  }
</script>

<form class="modal-form" on:submit|stopPropagation={handleSubmit}>
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
