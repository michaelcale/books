---
Title: Context is a tree of values
Id: 234
SOId: 9229
---
Context is created by wrapping existing immutable context and adding additional information.

Since you can "branch" the same context multiple times, context value can be thought of as tree of values.

The following tree:

```go
ctx := context.WithValue(
    context.WithDeadline(
        context.WithValue(context.Background(), sidKey, sid),
        time.Now().Add(30 * time.Minute),
    ),
    ridKey, rid,
)
trCtx := trace.NewContext(ctx, tr)
logCtx := myRequestLogging.NewContext(ctx, myRequestLogging.NewLogger())
```

can be visualized as:

![Context represented as a directed graph](90100016-context-tree.png)

Each child context has access to values of its parent contexts.

Data access flows upwards in the tree (represented by black edges).

Cancelation signals travel down the tree. If a context is canceled, all of its children are also canceled.

The cancelation signal flow is represented by the grey edges.
