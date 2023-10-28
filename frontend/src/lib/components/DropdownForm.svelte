<script lang="ts">
  import { clickOutside } from '$lib/utils/clickOutside';
  export let showPopup = false;

  function handleKeyDown(e: KeyboardEvent) {
    console.log('hi');
    if (e.key === 'Escape') {
      showPopup = false;
    }
  }
  function handleClickOutside() {
    showPopup = false;
  }
</script>

<div class="popup" on:keydown={handleKeyDown} use:clickOutside on:clickoutside={handleClickOutside}>
  <ul class="list">
    <li>
      <p class="list-item">change username</p>
    </li>
    <li>
      <p class="list-item">logout</p>
    </li>
    <li>
      <p class="list-item">DELETE</p>
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
    padding: 1em;
    border-radius: var(--border-radius);
    width: 13em;
    border: var(--border-width) none black;
  }
  .list {
    list-style-type: none;
    color: black;
    font-size: 1.5em;
    padding: 0;
  }

  .list-item {
    transition: background-color 0.2s;
    border-radius: var(--border-radius);
  }

  .list-item:hover {
    background-color: var(--col4);
    filter: brightness(95%);
    cursor: pointer;
  }
</style>
