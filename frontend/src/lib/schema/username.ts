import { string } from 'yup'

export const usernameString = string()
  .max(20, "Username can only be 20 characters long")
	.matches(/^[a-zA-Z0-9_]+$/, 'Username can only contain letters, numbers and underscores')
  .required("Please enter your username")
