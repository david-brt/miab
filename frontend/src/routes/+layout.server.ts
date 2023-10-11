import { PUBLIC_DATA_ROUTE } from '$env/static/public';
import type { User } from '../../types/user'
import type { LayoutServerLoad } from './$types';

  export const load: LayoutServerLoad = async ({ cookies }) => {
  const response = await fetch(`${PUBLIC_DATA_ROUTE}/verify-token`, {
    headers: {
      Authorization: `Bearer ${cookies.get("auth-token")}`
    }
  });
  const data: User = await response.json();
  return {
    user: data
  }
}
