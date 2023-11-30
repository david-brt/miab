import { INTERNAL_DATA_ROUTE } from '$env/static/private';
import type { LoadEvent } from '@sveltejs/kit';
import type { Message } from '../../types/message';
import type { Actions } from '@sveltejs/kit';
import { submitForm } from '$lib/utils/submitForm';

const CREATED = 201;
const ACCEPTED = 202;

export async function load({ fetch }: LoadEvent) {
  const response = await fetch(`${INTERNAL_DATA_ROUTE}/global-message`);
  const message: Message = await response.json();
  return {
    message
  };
}

export const actions: Actions = {
  signup: async ({ request }) => {
    const response = await submitForm(request, '/signup', CREATED);
    return response;
  },

  login: async ({ request }) => {
    const response = await submitForm(request, '/login', ACCEPTED);
    return response;
  },

  'send-message': async ({ request }) => {
    const response = await submitForm(request, '/send-message', ACCEPTED);
    return response;
  },

  rename: async ({ request, cookies }) => {
    const token = cookies.get('auth_token');
    const response = await submitForm(request, '/authorized/rename', ACCEPTED, token);
    return response;
  }
};
