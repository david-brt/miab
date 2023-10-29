<script lang="ts">
  import SubmitError from './SubmitError.svelte';
  import { PUBLIC_DATA_ROUTE } from '$env/static/public';
  import { handleSubmit } from '$lib/utils/form';
  import { showModal, user } from '../stores';

  const ACCEPTED = 202;
  const UNAUTHORIZED = 401;
  const TOOMANYREQUESTS = 429;
  const INTERNALSERVERERROR = 500;

  let loginStatus: number;

  async function callHandleSubmit(e: SubmitEvent) {
    const response = await handleSubmit(e, `${PUBLIC_DATA_ROUTE}/rename`);
    loginStatus = response.status;
    if (loginStatus === ACCEPTED) {
      const responseBody = await response.json();
      user.set(responseBody.user);
      showModal.set('namechange', false);
    }
  }
</script>

<form class="modal-form" on:submit|preventDefault={callHandleSubmit}>
  <label for="new-username-input" class="form-label">name</label>
  <input
    name="username"
    id="new-username-input"
    type="text"
    placeholder="new username"
    value=""
    maxlength="34"
    class="form-input"
  />
  <label for="old-username-input" class="form-label">name</label>
  <input
    name="username"
    id="old-username-input"
    type="text"
    placeholder="old username"
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
  {#if loginStatus === UNAUTHORIZED}
    <SubmitError>Invalid credentials</SubmitError>
  {/if}
  {#if loginStatus === TOOMANYREQUESTS}
    <SubmitError>Too many attempts. Please wait a few moments.</SubmitError>
  {/if}
  {#if loginStatus === INTERNALSERVERERROR}
    <SubmitError>Something went wrong. Try again later.</SubmitError>
  {/if}
  <button class="send-button">send</button>
</form>
