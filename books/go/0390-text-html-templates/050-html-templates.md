---
Title: HTML templates
Id: 222
SOId: 13461
---

Package `html/template` has the same base functionality as `text/template`.

The difference is that `html/template` understands structure of HTML and JavaScript code inside HTML.

Inserted text is escaped based on its surrounding context which eliminates cross-site scripting bugs.

@file html.go output sha1:59b81836761f5efa1451433fb46648da7e27c678 goplayground:BC\_-ZpGvIGD

## Inserting unescaped HTML

Sometimes you need to subvert escaping of text:

@file html_raw.go output sha1:2c1d354582f6b03300f1560921ff021465146dd8 goplayground:dl8dRoGbpWU

`template.HTML` and `template.JS` are type alises for `string` so you can assign string values to them.

Templating engine recognizes those types and disables escaping for them.
