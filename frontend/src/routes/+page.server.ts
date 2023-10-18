import { INTERNAL_DATA_ROUTE } from '$env/static/private';
import type { LoadEvent } from '@sveltejs/kit';
import type { Message } from '../../types/message';
import type { Actions } from './$types';

export async function load({ fetch }: LoadEvent) {
	const response = await fetch(INTERNAL_DATA_ROUTE);
	const message: Message = await response.json();
	return {
		message
	};
}

export const actions: Actions = {
	message: async ({ request }) => {
		const data = await request.formData();
		const jsonData = {
			senderName: data.get('senderName'),
			content: data.get('content')
		};
		const url = `${INTERNAL_DATA_ROUTE}/send-message`;
		const response = await fetch(url, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(jsonData)
		});
	} /*,
	login: async ({ request, cookies }) => {
		const data = await request.formData();
		const jsonData = {
			username: data.get('username'),
			password: data.get('password')
		};
		const url = `${PUBLIC_DATA_ROUTE}/login`;
		const response = await fetch(url, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(jsonData)
		});
		if (response.status === 202) {
			const { auth_token } = await response.json();
			cookies.set('auth_token', auth_token);
			return { success: true, validCredentials: true };
		} else if (response.status === 401 || response.status === 404) {
			return { success: true, validCredentials: false, formData: jsonData };
		} else {
			return { success: false, validCredentials: false, formData: jsonData };
		}
	}
  */
};
