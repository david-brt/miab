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
	default: async ({ request }) => {
		const data = await request.formData();
		const jsonData = {
			sender: data.get('sender'),
			content: data.get('content')
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
