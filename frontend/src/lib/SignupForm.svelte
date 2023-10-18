<script lang="ts">
  import { PUBLIC_DATA_ROUTE } from '$env/static/public'
  import { handleSubmit } from '$lib/utils/form'
  //closure to give handleSubmit access to the data route
  async function callHandleSubmit(e: SubmitEvent) {
    handleSubmit(e, `${PUBLIC_DATA_ROUTE}/signup`)
  }
</script>

<form class="modal-form" on:submit|preventDefault={callHandleSubmit}>
	<label for="username-input" class="form-label">name</label>
	<input
		name="username"
		id="username-input"
		type="text"
		placeholder="username"
    value=""
		maxlength="34"
		class="form-input"
	/>
	<label for="password-input" class="form-label">password</label>
	<input
		name="password"
		id="password-input"
		type="password"
		placeholder="password"
    value=""
		maxlength="72"
		class="form-input"
	/>
	<label for="password-retype" class="form-label">password</label>
	<input
		name="dontcare"
		id="password-retype"
		type="password"
		placeholder="retype password"
    value=""
		maxlength="72"
		class="form-input"
	/>
	<button class="send-button">send</button>
  <p style="display: none;"><strong>something about username already taken</strong></p>
</form>
