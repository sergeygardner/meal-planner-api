
```graphql
query($input: AuthCredentialsDTO!) {
  AuthCredentials(input: $input) {
    message
    status
  }
}
```
```json
{
  "input": {
  	"username": "username",
  	"password":"password"
  }
}
```
```graphql
query($input: AuthConfirmationDTO!) {
    AuthConfirmation(input: $input) {
        access_token
        refresh_token
    }
}
```

```json
{
  "input": {
    "username": "username",
    "password": "password",
    "code": "293445"
  }
}
```
