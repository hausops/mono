# Service Design

## Email verification

- on sign-up success, a new unverified account will be created (in both auth-svc and user-svc). Trying to create a new user with the same email will return an error. If someone creates an account using an email they don't have access and later the owner of the email want to "claim" the account, the email owner can go through the password reset flow. Our UI will make it obvious how to "claim" the account in such scenario.

  - the password reset is allowed regardless of the verification status

- if the email is already verified, calling send verification email will return an error; the UI layer might handle the error by
- if the email is already verified, calling verify email { verificationToken } will return an error; the UI layer might handle the error by ...
- re-sending verification email will void the previous verification token for the email i.e. there will be only one valid verification token for an email at a time.

---

- the behavior in changing email scenario is not defined.
