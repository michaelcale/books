---
Title: HTML templates
Id: 222
SOId: 13461
---

Package `html/template` has the same base functionality as `text/template`.

The difference is that `html/template` understands structure of HTML and JavaScript code inside HTML.

Inserted text is escaped based on its surrounding context which eliminates cross-site scripting bugs.

@file html.go output sha1:931ec6a6a6a77a56bdd3bcf5dc342a2250489d04 goplayground:FNOwPIazdD8
