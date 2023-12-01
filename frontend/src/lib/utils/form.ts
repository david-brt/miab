/**
 * handles form submissions
 * @returns http status of the received response
 */
export async function handleSubmit(
  e: SubmitEvent,
  dataRoute: string,
  token?: string
): Promise<Response> {
  const formData = new FormData(e.target as HTMLFormElement);
  const data: { [key: string]: FormDataEntryValue } = {};
  //deconstructs each field and adds it to data object
  for (let field of formData) {
    const [key, value] = field;
    data[key] = value;
  }

  const headers: Record<string, string> = {
    'Content-Type': 'application/json'
  };

  if (token !== undefined) {
    headers['Authorization'] = `Bearer ${token}`;
  }
  const response = await fetch(dataRoute, {
    method: 'POST',
    mode: 'cors',
    headers: headers,
    body: JSON.stringify(data),
    credentials: 'include'
  });
  return response;
}
