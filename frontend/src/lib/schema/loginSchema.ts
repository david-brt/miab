import { object } from 'yup';
import { passwordString } from './password';
import { usernameString } from './username';

export const signupSchema = object({
  username: usernameString,
  password: passwordString
});
