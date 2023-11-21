import { object } from 'yup';
import { passwordString } from './password';
import { usernameString } from './username';

export const loginSchema = object({
  username: usernameString,
  password: passwordString
});
