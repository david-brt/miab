<script lang="ts">
  import { PUBLIC_DATA_ROUTE } from '$env/static/public';
  import { user } from '$lib/stores';
  import { clickOutside } from '$lib/utils/clickOutside';
  export let showPopup = false;

  function handleKeyDown(e: KeyboardEvent) {
    if (e.key === 'Escape') {
      showPopup = false;
    }
  }
  function handleClickOutside() {
    showPopup = false;
  }

  async function logout() {
    const response = await fetch(`${PUBLIC_DATA_ROUTE}/logout`, {
      method: 'GET',
      mode: 'cors',
      credentials: 'include'
    });

    user.set(undefined);
  }
</script>

<div class="popup" on:keydown={handleKeyDown} use:clickOutside on:clickoutside={handleClickOutside}>
  <ul class="list">
    <li class="list-item">
      <p>change username</p>
    </li>
    <li>
      <p class="list-item" on:click={logout}>logout</p>
    </li>
  </ul>
</div>

<style>
  .popup {
    background-color: var(--col4);
    display: block;
    position: fixed;
    top: 7em; /* Adjust as needed */
    right: 1em;
    border-radius: var(--border-radius);
    width: 16em;
  }
  .list {
    list-style-type: none;
    color: black;
    font-size: 1.5em;
    margin: 0;
    padding: 0;
  }

  .list-item {
    transition: background-color 0.2s;
    padding: 0.5em 1em 0.5em 1em;
    border-radius: 0.5em;
  }

  .list-item p {
    margin: 0;
  }

  .list-item:hover {
    background-color: var(--col4);
    filter: brightness(95%);
    cursor: pointer;
  }
</style>
