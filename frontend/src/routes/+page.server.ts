import { PUBLIC_DATA_ROUTE } from '$env/static/public';
import type { LoadEvent } from '@sveltejs/kit';
import type { UserMessage } from '../../types/userMessage';
import type { Actions } from './$types';

export async function load({ fetch }: LoadEvent) {
	const response = await fetch(PUBLIC_DATA_ROUTE);
	const message: UserMessage = await response.json();
	return {
		message
	};
}

export const actions: Actions = {
	default: async ({ request }) => {
		const data = await request.formData();
		const jsonData = {
			senderName: data.get('senderName'),
			messageContent: data.get('content')
		};
		const url = `${PUBLIC_DATA_ROUTE}/update`;
		const response = await fetch(url, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(jsonData)
		});
		return response.json();
	}
};
