import { object } from 'yup';
import { usernameString } from './username';

export const renameSchema = object({
  username: usernameString
});
