<script lang="ts">
  import { createForm } from 'felte';
  import FormField from './FormField.svelte';
  import { validator } from '@felte/validator-yup';
  import { showModal } from '../stores';
  import type { ComponentProps } from 'svelte';
  import type { InferType } from 'yup';
  import SubmitError from './SubmitError.svelte';
  import { writable } from 'svelte/store';

  export let fields: Omit<ComponentProps<FormField>, 'inputErrors'>[] = [];
  export let validationSchema;
  export let actionRoute: string;
  export let modalType: keyof typeof $showModal;

  const formError = writable();

  const { form: felteForm, errors } = createForm<InferType<typeof validationSchema>>({
    extend: [validator({ schema: validationSchema })]
  });

  async function handleSuccess(event: CustomEvent) {
    const { response } = event.detail;
    const responseBody = await response.json();
    const data = JSON.parse(responseBody.data);
    if (responseBody.type === 'failure') {
      // a little hacky as $page.form only contains the response for the form on the route corresponding to the action name
      // format returned by action is transformed in an odd way. might find a better solution later
      formError.set(data[1]);
    } else {
      showModal.set(modalType, false);
    }
  }
</script>

<form
  use:felteForm
  class="modal-form"
  action={actionRoute}
  method="post"
  on:feltesuccess={handleSuccess}
>
  {#each fields as field}
    <FormField {...field} inputErrors={$errors[field.name]} />
  {/each}

  <button class="send-button" type="submit">submit</button>

  {#if $formError}
    <SubmitError>{$formError}</SubmitError>
  {/if}
</form>
