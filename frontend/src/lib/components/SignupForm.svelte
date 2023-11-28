<script lang="ts">
  import { createForm } from 'felte';
  import type { InferType } from 'yup';
  import { validator } from '@felte/validator-yup';
  import { showModal } from '../stores';
  import { signupSchema } from '$lib/schema/signupSchema';
  import FormField from './FormField.svelte';
  import SubmitError from './SubmitError.svelte';

  const { form: felteForm, errors } = createForm<InferType<typeof signupSchema>>({
    initialValues: {
      username: '',
      password: '',
      confirmation: ''
    },
    extend: [validator({ schema: signupSchema })]
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
  <FormField
    name="username"
    type="text"
    required
    placeholder="username"
    maxlength={50}
    inputErrors={$errors.username}
  />
  <FormField
    name="password"
    type="password"
    required
    placeholder="password"
    minlength={8}
    maxlength={50}
    inputErrors={$errors.password}
  />
  <FormField
    name="confirmation"
    type="password"
    required
    placeholder="confirm password"
    minlength={8}
    maxlength={50}
    inputErrors={$errors.password}
  />
  <button class="send-button" type="submit">send</button>
  {#if formError}
    <SubmitError>{formError}</SubmitError>
  {/if}
</form>
