import { object } from 'yup';
import { usernameString } from './username';

export const signupSchema = object({
  username: usernameString
});
