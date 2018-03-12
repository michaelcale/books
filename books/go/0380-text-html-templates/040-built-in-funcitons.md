---
Title: Built-in functions
Id: 220
SOId: 13459
---

Templating engine supports calling functions like `{{ len .Tweet }}` where `len` is a function that returns length of an array or slice.

## and, or, not

`and`, `or`, `not` are for logical operations:

@file and_or_not.go output sha1:48310211c3e5bd7ace218c0dbf2cc91a790e37a3 goplayground:SpbYdKSbVd4

## index

`index` is for accessing elements of a slice by index or values in a map by key.

@file index_func.go output sha1:2f085106143925cac8fd7d9700a3610f1403b804 goplayground:SuOa-mIN4VQ

## len

`len` returns length of an array of map.

@file len.go output sha1:0244b93bc93fa498c96b3e0937c493002e49749d goplayground:TAoEu60ToBS

## print, printf, println

`print` is like `fmt.Sprint`.

`printf` is like `fmt.Sprintf`.

`println` is like `fmt.Sprintln`.

@file print.go output sha1:c709d51fec4fec036b76c635993236e3bbc8b7b1 goplayground:CB015PVwVVD

## js, html, urlquery

`js`, `html` and `urlquery` is for escaping text so that it can be safely inserted in a JavaScript, HTML and URL context:

@file js_html_url_escape.go output sha1:2a41b84529fb31cc4f6c8c2531960f4e7c763a4e goplayground:72IyxiQ9iYb
