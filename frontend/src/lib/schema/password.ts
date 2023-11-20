import { string } from 'yup'

const createPasswordValidationError = (str: string) => {
  return `Password must include at least 1 ${str}`
}

export const passwordString = string()
  .min(8, "Password must be at least 8 characters long")
  .max(50, "Password can only be 50 characters long")
  .matches(/[0-9]/, createPasswordValidationError("digit"))
  .matches(/[a-z/]/, createPasswordValidationError("lowercase character"))
  .matches(/[A-Z/]/, createPasswordValidationError("uppercase character"))
  .matches(/[#?!@$ %^&*]/, createPasswordValidationError("special cahracter (#?!@$ %^&*)"))
  .required("Please enter a password")
