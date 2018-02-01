Title: Send/receive emails
Id: 5912
Syntax:
- func PlainAuth(identity, username, password, host string) Auth
- func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
|======|
