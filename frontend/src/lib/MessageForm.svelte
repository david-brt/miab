<form method="POST" class="message-form">
	<label for="message-input" class="form-label">message</label>
	<textarea
		name="content"
		id="message-input"
		placeholder="your message"
		maxlength="240"
		rows="3"
		spellcheck="false"
	/>
	<label for="sender-input" class="form-label">name</label>
	<input
		name="senderName"
		id="sender-input"
		type="text"
		placeholder="your name"
		maxlength="50"
		class="message-input"
	/>
	<button class="send-button">send</button>
</form>

<style>
	.message-form {
		display: flex;
		flex-direction: column;
		gap: 2em;
	}

	#message-input {
		border: none;
		border-radius: var(--border-radius);
		font-family: 'Lato';
		padding: 0.5em;
		resize: none;
		background-color: var(--pastel-yellow);
		font-size: 1.5em;
	}

	.form-label {
		display: none;
	}

	textarea:focus,
	input:focus {
		outline: none;
	}
</style>
