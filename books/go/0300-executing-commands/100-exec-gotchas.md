---
Title: os.Exec gotchas
Id: 23356
---

## `exec.Cmd` cannot be reused

An `exec.Cmd` cannot be reused after calling its `Run`, `Output` or `CombinedOutput` methods.

@file reuse_cmd.go output allow_error noplayground

