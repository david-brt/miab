export async function handleSubmit(e: SubmitEvent, dataRoute: string) {
	const formData = new FormData(e.target as HTMLFormElement);
	const data: { [key: string]: FormDataEntryValue } = {};
	//deconstructs each field and adds it to data object
	for (let field of formData) {
		const [key, value] = field;
		data[key] = value;
	}
	await fetch(dataRoute, {
		method: 'POST',
		mode: 'cors',
		headers: {
			'Content-Type': 'application/json'
		},
		body: JSON.stringify(data),
		credentials: 'include'
	});
}
