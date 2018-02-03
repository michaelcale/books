Title: Using a context to cancel work
Id: 10386
Score: 0
Body:
Passing a context with a timeout (or with a cancel function) to a long running function can be used to cancel that functions work:

    ctx, _ := context.WithTimeout(context.Background(), 200*time.Millisecond)
    for {
        select {
        case <-ctx.Done():
            return ctx.Err()
        default:
            // Do an iteration of some long running work here!
        }
    }
|======|
