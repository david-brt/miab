<script lang="ts">
  import { PUBLIC_DATA_ROUTE } from '$env/static/public';
  import { showModal, user } from '$lib/stores';
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

  function openRenameModal() {
    showModal.set('namechange', true);
  }

  async function logout() {
    await fetch(`${PUBLIC_DATA_ROUTE}/authorized/logout`, {
      method: 'GET',
      mode: 'cors',
      credentials: 'include'
    });
    showPopup = false;
    user.set(undefined);
  }
</script>

<div class="popup" on:keydown={handleKeyDown} use:clickOutside on:clickoutside={handleClickOutside}>
  <ul class="list">
    <li class="list-item">
      <button class="dropdown-option" on:click={openRenameModal}>
        <span>change username</span>
      </button>
    </li>
    <li class="list-item">
      <button class="dropdown-option" on:click={logout}>
        <span>logout</span>
      </button>
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
    border-radius: 0.5em;
  }

  .dropdown-option {
    all: unset;
    box-sizing: border-box;
    padding: 0.5em 1em;
    width: 100%;
    height: 100%;
  }

  .list-item:hover {
    background-color: var(--col4);
    filter: brightness(95%);
    cursor: pointer;
  }
</style>
