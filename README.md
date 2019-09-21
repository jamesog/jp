# JWT Parser

`jp` reads a [JSON Web Token](https://jwt.io/) and prints the decoded header and payload in pretty-printed JSON.

While working with JWTs I frequently found myself doing something like:

```
echo $token | tr . '\n' | while read line; do echo $token | base64 --decode; echo; done
```

Which is fine the first couple of times you need it but gets tedious after that. It also doesn't account for the lack of padding in each part of the token so the resulting base64-decoded output could be incomplete and thus invalid JSON. This method also prints binary garbage for the final part (the signature).

This tool is just a tiny helper to do the same, while handling the lack of padding in each part and discarding the signature as it's only useful to the service which issued the token.

## Usage

Tokens can be passed either as an argument to the program or by standard input.

```
$ echo eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c | jp
>> Header
{
  "alg": "HS256",
  "typ": "JWT"
}
>> Payload
{
  "sub": "1234567890",
  "name": "John Doe",
  "iat": 1516239022
}
```

```
$ jp eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
>> Header
{
  "alg": "HS256",
  "typ": "JWT"
}
>> Payload
{
  "sub": "1234567890",
  "name": "John Doe",
  "iat": 1516239022
}
```
