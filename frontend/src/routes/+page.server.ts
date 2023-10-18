import { INTERNAL_DATA_ROUTE } from '$env/static/private';
import type { LoadEvent } from '@sveltejs/kit';
import type { Message } from '../../types/message';

export async function load({ fetch }: LoadEvent) {
	const response = await fetch(INTERNAL_DATA_ROUTE);
	const message: Message = await response.json();
	return {
		message
	};
}
