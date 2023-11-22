import { INTERNAL_DATA_ROUTE } from '$env/static/private';
import { fail, type LoadEvent } from '@sveltejs/kit';
import type { Message } from '../../types/message';
import type { Actions } from '@sveltejs/kit';

const CREATED = 201;

export async function load({ fetch }: LoadEvent) {
  const response = await fetch(`${INTERNAL_DATA_ROUTE}/global-message`);
  const message: Message = await response.json();
  return {
    message
  };
}

export const actions: Actions = {
  signup: async ({ request }) => {
    const formData = await request.formData();
    const formBody = {
      username: formData.get('username'),
      password: formData.get('password')
    };

    const headers: Record<string, string> = {
      'Content-Type': 'application/json'
    };

    const response = await fetch(`${INTERNAL_DATA_ROUTE}/signup`, {
      method: 'POST',
      mode: 'cors',
      headers: headers,
      body: JSON.stringify(formBody),
      credentials: 'include'
    });

    const responseData = await response.json();
    const signupStatus = response.status;

    if (signupStatus === CREATED) {
      return { success: true };
    }

    return fail(signupStatus, responseData);
  }
};
