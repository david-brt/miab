<script lang="ts">
	export let showModal: boolean;
	let dialog: HTMLDialogElement;
	function handleKeyDown(e: KeyboardEvent) {
		if (e.key === 'Escape') {
			showModal = false;
		}
	}

	$: if (dialog && showModal) dialog.showModal();
</script>

<dialog
	bind:this={dialog}
	on:close={() => (showModal = false)}
	on:click|self={() => dialog.close()}
	on:keydown={handleKeyDown}
	class="modal-dialog"
	on:click|stopPropagation
>
	<div class="content">
		<button class="close-button" on:click={() => dialog.close()}>
			<img width="30px" src="icons/x-icon.svg" alt="" />
		</button>
		<slot />
	</div>
</dialog>

<style>
	.modal-dialog {
		background-color: var(--col2);
		border-radius: var(--border-radius);
		border: none;
		padding: 0;
		width: 95vw;
	}

	.content {
		display: flex;
		align-items: center;
		justify-content: center;
		flex-direction: column;
		padding: 1em;
		width: 90%;
		margin: auto;
	}

	.close-button {
		padding: 0;
		background: none;
		align-self: flex-end;
	}

	@media screen and (min-width: 330px) {
		.modal-dialog {
			width: 80vw;
		}
	}

	@media screen and (min-width: 768px) {
		.modal-dialog {
			width: 60vw;
		}
	}

	@media screen and (min-width: 1280px) {
		.modal-dialog {
			width: 30vw;
		}
	}
</style>
