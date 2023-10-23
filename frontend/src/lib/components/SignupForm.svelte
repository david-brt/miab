<script lang="ts">
  import { PUBLIC_DATA_ROUTE } from '$env/static/public'
  import { handleSubmit } from '$lib/utils/form'
  import { showModal } from '../stores'
	import SubmitError from './SubmitError.svelte';

  const CREATED = 201
  const ACCEPTED = 202
  const INTERNALSERVERERROR = 500

  let signupStatus: number;
  let username = '';
  let password = '';
  let retyped_password= '';

  async function onSubmit(e: SubmitEvent) {
    if(password === retyped_password){
      const response = await handleSubmit(e, `${PUBLIC_DATA_ROUTE}/signup`)
      signupStatus = response.status
	  console.log(signupStatus)
      if(signupStatus === CREATED) {
        username = ''
        password = ''
        retyped_password = ''
		showModal.set('signup', false)
      }
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
  {#if signupStatus === INTERNALSERVERERROR}
    <SubmitError>Something went wrong. Try again later.</SubmitError>
  {/if}
</form>
