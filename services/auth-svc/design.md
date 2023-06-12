# Service Design

## Confirm email address

- On sign-up success, a new unconfirmed account will be created in both auth-svc and user-svc.
  - Trying to create a new user with the same email will result in an error.
  - If someone creates an account using an email address they don't have access to, and later the owner of the email wants to "claim" the account, the email owner can go through the password reset flow. Our user interface (UI) will clearly indicate how to "claim" the account in such a scenario.
  - The password reset functionality is available regardless of the confirmation status.
- If the email address is already confirmed, calling the "send confirm email" function will result in an error. The UI layer can handle this error by ignoring it and treating it as a no-op for users.
- If the email address is already confirmed, calling the "confirm email" function with the provided { confirmToken } will result in an error. The UI layer can handle this error by ignoring it and treating it as a no-op for users.
- Resending the confirmation email will invalidate the previous confirmation token for the email i.e. there will be only one valid verification token for an email at any given time.

---

- The behavior for changing email currently is not specified.
