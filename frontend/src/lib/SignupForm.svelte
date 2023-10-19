<script lang="ts">
  import { PUBLIC_DATA_ROUTE } from '$env/static/public'
  import { handleSubmit } from '$lib/utils/form'
  let username = '';
  let password = '';
  let retyped_password= '';
  let wrong_password = false;


  //closure to give handleSubmit access to the data route
  async function onSubmit(e: SubmitEvent) {
    if(password === retyped_password){
      await handleSubmit(e, `${PUBLIC_DATA_ROUTE}/signup`)
      username = ''
      password = ''
      retyped_password = ''
    }
    else{
      wrong_password=true
    }
  }

</script>

<form class="modal-form" on:submit|preventDefault={onSubmit}>
	<label for="username-input" class="form-label">name</label>
	<input
		name="username"
		id="username-input"
		type="text"
		placeholder="username"
    	bind:value={username}
		maxlength="34"
		class="form-input"
	/>
	<label for="password-input" class="form-label">password</label>
	<input
		name="password"
		id="password-input"
		type="password"
		placeholder="password"
    	bind:value={password}
		maxlength="72"
		class="form-input"

	/>
	<label for="password-retype" class="form-label">password</label>
	{#if wrong_password}
		<p class="errortext">joa is blöd</p>
	{/if}
	<input
		name="retyped"
		id="password-retype"
		type="password"
		placeholder="retype password"
    	bind:value={retyped_password}
		maxlength="72"
		class="form-input"
	/>
	<button class="send-button">send</button>
  <p style="display: none;"><strong>something about username already taken</strong></p>
</form>
