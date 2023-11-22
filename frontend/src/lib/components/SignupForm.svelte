<script lang="ts">
  import { createForm } from 'felte';
  import { showModal } from '../stores';
  import SubmitError from './SubmitError.svelte';

  const { form: felteForm } = createForm({});

  async function handleSuccess(event: CustomEvent) {
    const { response } = event.detail;
    const responseBody = await response.json();
    const data = JSON.parse(responseBody.data);
    if (responseBody.type === 'failure') {
      // a little hacky as $page.form only contains the response for the form on the route corresponding to the action name
      // format returned by action is transformed in an odd way. might find a better solution later
      formError = data[2];
    } else {
      showModal.set('signup', false);
    }
  }

  let formError: undefined | string;
  let username = '';
  let password = '';
  let retyped_password = '';
</script>

<form
  class="modal-form"
  use:felteForm
  action="?/signup"
  method="post"
  on:feltesuccess={handleSuccess}
>
  <label for="username-input" class="form-label">name</label>
  <input
    name="username"
    id="username-input"
    type="text"
    placeholder="username"
    required
    bind:value={username}
    maxlength="20"
    class="form-input"
  />
  <label for="password-input" class="form-label">password</label>
  <input
    name="password"
    id="password-input"
    type="password"
    required
    placeholder="password"
    bind:value={password}
    minlength="8"
    maxlength="50"
    class="form-input"
  />
  <label for="password-retype" class="form-label">password</label>
  <input
    name="retyped"
    id="password-retype"
    type="password"
    required
    placeholder="retype password"
    bind:value={retyped_password}
    minlength="8"
    maxlength="50"
    class="form-input"
  />
  {#if formError}
    <SubmitError>{formError}</SubmitError>
  {/if}
  <button class="send-button" type="submit">send</button>
</form>
