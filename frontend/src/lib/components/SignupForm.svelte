<script lang="ts">
  import { PUBLIC_DATA_ROUTE } from '$env/static/public';
  import { handleSubmit } from '$lib/utils/form';
  import { showModal } from '../stores';
  import SubmitError from './SubmitError.svelte';

  const CREATED = 201;

  let signupStatus: number;
  let responseData: { [index: string]: string };
  let username = '';
  let password = '';
  let retyped_password = '';

  async function onSubmit(e: SubmitEvent) {
    if (password === retyped_password) {
      const response = await handleSubmit(e, `${PUBLIC_DATA_ROUTE}/signup`);

      responseData = await response.json();
      signupStatus = response.status;

      if (signupStatus === CREATED) {
        username = '';
        password = '';
        retyped_password = '';
        showModal.set('signup', false);
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
    pattern="^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$ %^&*])$ "
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
    pattern="^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$ %^&*])$ "
    class="form-input"
  />
  <SubmitError>{responseData?.errorMessage || ''}</SubmitError>
  <button class="send-button">send</button>
</form>
