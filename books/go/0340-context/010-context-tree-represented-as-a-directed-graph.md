Title: Context tree represented as a directed graph
Id: 9229
Score: 0
Body:
A simple context tree (containing some common values that might be request scoped and included in a context) built from Go code like the following:

    // Pseudo-Go
    ctx := context.WithValue(
        context.WithDeadline(
            context.WithValue(context.Background(), sidKey, sid),
            time.Now().Add(30 * time.Minute),
        ),
        ridKey, rid,
    )
    trCtx := trace.NewContext(ctx, tr)
    logCtx := myRequestLogging.NewContext(ctx, myRequestLogging.NewLogger())

Is a tree that can be represented as a directed graph that looks like this:

[![Context represented as a directed graph][1]][1]

Each child context has access to the values of its parent contexts, so the data access flows upwards in the tree (represented by black edges). Cancelation signals on the other hand travel down the tree (if a context is canceled, all of its children are also canceled). The cancelation signal flow is represented by the grey edges.

  [1]: http://i.stack.imgur.com/R0CED.png

|======|
