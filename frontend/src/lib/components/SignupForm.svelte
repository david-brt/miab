<script lang="ts">
  import { createForm } from 'felte';
  import type { InferType } from 'yup';
  import { validator } from '@felte/validator-yup';
  import { reporter, ValidationMessage } from '@felte/reporter-svelte';
  import { showModal } from '../stores';
  import { signupSchema } from '$lib/schema/signupSchema';
  import SubmitError from './SubmitError.svelte';

  const { form: felteForm } = createForm<InferType<typeof signupSchema>>({
    initialValues: {
      username: '',
      password: '',
      confirmation: ''
    },
    extend: [validator({ schema: signupSchema }), reporter]
  });

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
</script>

<form
  class="modal-form"
  use:felteForm
  action="?/signup"
  method="post"
  on:feltesuccess={handleSuccess}
>
  <div class="input-container">
    <label for="username-input" class="form-label">name</label>
    <input
      name="username"
      id="username-input"
      type="text"
      placeholder="username"
      required
      maxlength="20"
      class="form-input"
    />
    <ValidationMessage for="username" let:messages={message}>
      <SubmitError>{message}</SubmitError>
      <SubmitError slot="placeholder"></SubmitError>
    </ValidationMessage>
  </div>
  <div class="input-container">
    <label for="password-input" class="form-label">password</label>
    <input
      name="password"
      id="password-input"
      type="password"
      required
      placeholder="password"
      minlength="8"
      maxlength="50"
      class="form-input"
    />
    <label for="password-retype" class="form-label">password</label>
    <ValidationMessage for="password" let:messages={message}>
      <SubmitError>{message}</SubmitError>
      <SubmitError slot="placeholder"></SubmitError>
    </ValidationMessage>
  </div>
  <div class="input-container">
    <input
      name="confirmation"
      id="password-retype"
      type="password"
      required
      placeholder="retype password"
      minlength="8"
      maxlength="50"
      class="form-input"
    />
    <ValidationMessage for="confirmation" let:messages={message}>
      <SubmitError>{message}</SubmitError>
      <SubmitError slot="placeholder"></SubmitError>
    </ValidationMessage>
  </div>
  <button class="send-button" type="submit">send</button>
</form>

<style>
  .input-container {
    display: flex;
    gap: 0.5em;
    flex-direction: column;
  }
</style>
