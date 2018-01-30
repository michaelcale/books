Title: Deferred Function Calls
Id: 10474
Score: 0
BodyHtml:
<p>Deferred function calls serve a similar purpose to things like <code>finally</code> blocks in languages like Java: they ensure that some function will be executed when the outer function returns, regardless of if an error occurred or which return statement was hit in cases with multiple returns. This is useful for cleaning up resources that must be closed like network connections or file pointers. The <code>defer</code> keyword indicates a deferred function call, similarly to the <code>go</code> keyword initiating a new goroutine. Like a <code>go</code> call, function arguments are evaluated immediately, but unlike a <code>go</code> call, deferred functions are not executed concurrently.</p>
<pre><code>func MyFunc() {
    conn := GetConnection()    // Some kind of connection that must be closed.
    defer conn.Close()        // Will be executed when MyFunc returns, regardless of how.
    // Do some things...
    if someCondition {
        return                // conn.Close() will be called
    }
    // Do more things
}// Implicit return - conn.Close() will still be called
</code></pre>
<p>Note the use of <code>conn.Close()</code> instead of <code>conn.Close</code> - you're not just passing in a function, you're deferring a full function <em>call</em>, including its arguments. Multiple function calls can be deferred in the same outer function, and each will be executed once in reverse order. You can also defer closures - just don't forget the parens!</p>
<pre><code>defer func(){
    // Do some cleanup
}()
</code></pre>


|======|
