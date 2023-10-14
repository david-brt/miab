import type { LayoutServerLoad } from './$types';
import { PUBLIC_DATA_ROUTE } from '$env/static/public';

export const load: LayoutServerLoad = async ({ cookies }) => {
	const tokenString = cookies.get('auth_token');
	const response = await fetch(`${PUBLIC_DATA_ROUTE}/verify-token`, {
		headers: {
			Authorization: `Bearer ${tokenString}`
		}
	});
	if (response.ok) {
		let { claims } = await response.json();
		return {
			user: {
				id: claims.sub,
				usename: claims.name
			}
		};
	}
};
