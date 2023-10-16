import { PUBLIC_DATA_ROUTE } from '$env/static/public';
import type { LoadEvent } from '@sveltejs/kit';
import type { Message } from '../../types/message';
import type { Actions } from './$types';

export async function load({ fetch }: LoadEvent) {
	const response = await fetch(PUBLIC_DATA_ROUTE);
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
		const url = `${PUBLIC_DATA_ROUTE}/send-message`;
		const response = await fetch(url, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(jsonData)
		});
	},
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
		}
	}
};
