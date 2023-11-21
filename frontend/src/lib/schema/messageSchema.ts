import { object, string } from 'yup';

export const messageSchema = object({
  message: string()
    .min(5, 'Type at least 5 characters')
    .max(240, 'Please limit yourself to 240 characters')
    .defined('Please enter a message')
    .required('Please enter a message'),
  sender: string()
    .max(20, 'Username is limited to 20 characters')
    .defined('Please enter your name')
    .required('Please enter your name')
});
