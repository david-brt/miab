<div class="container">
	<nav class="navbar" />
	<slot />
</div>

<style>
	:global(body) {
		font-family: 'Lato';
		font-size: 12px;
		margin: 0;
		padding: 0;
		background-color: var(--salmon);
		--bg-yellow: #fffce6;
		--pastel-yellow: #fffba6;
		--brownish-yellow: #a37b00;
		--orange: #ffa540;
		--almost-red: #ff605b;
		--border-radius: 0.5em;
		--orange: #f79b24;
	}
	:global(button) {
		background-color: var(--pastel-yellow);
		border: none;
		border-radius: var(--border-radius);
		padding: 1em;
		font-size: 1.5em;
	}

	:global(button:hover) {
		cursor: pointer;
	}

	:global(::placeholder) {
		color: #a37b00;
	}

	.container {
		margin: auto;
		display: flex;
		justify-content: center;
		width: 90vw;
	}
</style>
