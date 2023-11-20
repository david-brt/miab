import { object, string, ref } from 'yup';
import { passwordString } from './password';
import { usernameString } from './username';

export const signupSchema = object({
  username: usernameString,
  password: passwordString,
  confirmPassword: string()
    .required('Please re-type your password')
    // oneOf matches one of the values inside the array.
    // "ref" gets the value of password.
    .oneOf([ref('password')], 'Passwords do not match')
});
