<form method="POST" class="modal-form" action="?/message">
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
