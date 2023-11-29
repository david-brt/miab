import { fail } from '@sveltejs/kit';
import { INTERNAL_DATA_ROUTE } from '$env/static/private';

/**
 * for use in form actions
 * @param request the request recieved from the html form
 * @param route route relative to INTERNAL_DATA_ROUTE environment variable (needs a "/" as the first character)
 * @param successStatus the HTTP status code which indicates a successful response
 * @param token optional: jwt auth token from HTTP-only cookie for authorized routes
 */
export async function submitForm(
  request: Request,
  route: string,
  successStatus: number,
  token = ''
) {
  const formData = await request.formData();
  const formJSON = JSON.stringify(Object.fromEntries(formData));

  const headers: Record<string, string> = {
    'Content-Type': 'application/json',
    Authorization: `Bearer ${token}`
  };

  const response = await fetch(`${INTERNAL_DATA_ROUTE}${route}`, {
    method: 'POST',
    mode: 'cors',
    headers: headers,
    body: formJSON
  });

  const responseData = await response.json();

  if (response.status === successStatus) {
    return { success: true };
  }

  return fail(response.status, responseData);
}
