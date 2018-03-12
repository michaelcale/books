---
Title: Middleware
Id: 296
SOId: 9343
---
## Introduction

In Go Middleware can be used to execute code before and after handler function. It uses the power of Single Function Interfaces.
Can be introduced at any time without affecting the other middleware.
For Ex: Authentication logging can be added in later stages of development without disturbing the existing code.

## Remarks
The **Signature of middleware** should be (http.ResponseWriter, *http.Request) i.e. of
**http.handlerFunc** type.
