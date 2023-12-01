import type { LayoutServerLoad } from './$types';
import { INTERNAL_DATA_ROUTE } from '$env/static/private';

export const load: LayoutServerLoad = async ({ cookies }) => {
  const tokenString = cookies.get('auth_token');
  const response = await fetch(`${INTERNAL_DATA_ROUTE}/authorized/verify-token`, {
    method: 'GET',
    mode: 'cors',
    headers: {
      Authorization: `Bearer ${tokenString}`
    },
    credentials: 'include'
  });
  if (response.ok) {
    const { claims } = await response.json();
    return {
      user: {
        id: claims.sub,
        name: claims.name
      },
      token: tokenString
    };
  }
};
