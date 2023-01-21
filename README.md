# JWT Inspect

## Why?

JWT Inspect is an attempt to make an open JWT (JSON Web Token) decoding tool based entirely on golang standard libraries making it easier to reason about the security of the tool. This tool adheres closely to RFC 7519 and can be helpful for detecting tokens which are not correctly encoded as well as allowing the user to view the contents of the header and claims.