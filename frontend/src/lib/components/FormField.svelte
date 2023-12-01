<script lang="ts">
  import SubmitError from './SubmitError.svelte';
  import type { HTMLInputTypeAttribute } from 'svelte/elements';

  export let name: string;
  export let placeholder: string;
  export let type: HTMLInputTypeAttribute | undefined = 'text';
  export let required: boolean;
  export let maxlength: number;
  export let minlength: number | undefined = 0;

  export let inputErrors: string[] | null;
  export let textArea: boolean | undefined = false;
  export let rows: number | undefined = 3;
</script>

<div class="input-container">
  <label for={`${name}-input`} class="form-label">{name}</label>
  {#if textArea}
    <textarea
      {name}
      id={`${name}-input`}
      {placeholder}
      {required}
      {minlength}
      {maxlength}
      {rows}
      class="form-input"
      spellcheck="false"
    />
  {/if}
  {#if !textArea}
    <input
      {name}
      {placeholder}
      {type}
      id={`${name}-input`}
      {required}
      {minlength}
      {maxlength}
      class="form-input"
    />
  {/if}
  {#if inputErrors}
    {#each inputErrors as inputError}
      <SubmitError>{inputError}</SubmitError>
    {/each}
  {/if}
  {#if !inputErrors}
    <SubmitError></SubmitError>
  {/if}
</div>

<style>
  .input-container {
    display: flex;
    gap: 0.5em;
    flex-direction: column;
  }

  .form-input {
    border: none;
    flex: 1;
    border-radius: var(--border-radius);
    font-family: 'Lato';
    padding: 0.5em;
    resize: none;
    background-color: var(--pastel-yellow);
    font-size: 1.5em;
  }
</style>
