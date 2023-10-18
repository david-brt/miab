<script lang="ts">
  import { PUBLIC_DATA_ROUTE } from '$env/static/public'
  async function handleSubmit(e: SubmitEvent) {
    const formData = new FormData(e.target as HTMLFormElement)
    const data: {[key: string]: FormDataEntryValue} = {};
    //deconstructs each field and adds it to data object
    for (let field of formData) {
      const [key, value] = field;
      data[key] = value;
    }
    await fetch(`${PUBLIC_DATA_ROUTE}/login`, {
      method: 'POST',
      mode: 'cors',
      headers: {
      'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
      credentials: 'include',
    })
  }
</script>

<form class="modal-form" on:submit|preventDefault={handleSubmit}>
	<label for="username-input" class="form-label">name</label>
	<input
		name="username"
		id="username-input"
		type="text"
		placeholder="username"
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
	<button class="send-button">send</button>
  <p style="display: none;"><strong>Invalid credentials</strong></p>
</form>
