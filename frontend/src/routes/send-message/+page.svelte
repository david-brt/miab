<div class="content">
	<form method="POST" class="message-form">
		<label for="message-input" class="form-label">Message</label>
		<textarea
			name="sender"
			id="message-input"
			placeholder="your message"
			class="message-input"
			maxlength="1000"
			rows="3"
		/>
		<label for="sender-input" class="form-label">Name</label>
		<input
			name="message"
			id="sender-input"
			type="text"
			placeholder="your name"
			maxlength="255"
			class="message-input"
		/>
		<button>send</button>
	</form>
</div>

<style>
	.content {
		display: flex;
		align-items: center;
		justify-content: center;
		gap: 3em;
		flex-direction: column;
		height: 100vh;
		width: 100%;
		padding: 3em;
	}

	.message-form {
		display: flex;
		flex-direction: column;
		gap: 2em;
		width: 100%;
	}

	.message-input {
		border: none;
		border-radius: var(--border-radius);
		font-family: 'Lato';
		padding: 0.5em;
		resize: none;
		color: #fff;
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
