Title: Separate integration tests
Id: 8607
Score: 2
Body:
Build constraints are commonly used to separate normal unit tests from integration tests that require external resources, like a database or network access. To do this, add a custom build constraint to the top of the test file:

    // +build integration
     
    package main
     
    import (
        "testing"
    )
     
    func TestThatRequiresNetworkAccess(t *testing.T) {
        t.Fatal("It failed!")
    }

The test file will not compile into the build executable unless the following invocation of `go test` is used:

    go test -tags "integration"

Results:

    $ go test
    ?       bitbucket.org/yourname/yourproject    [no test files]
    $ go test -tags "integration"
    --- FAIL: TestThatRequiresNetworkAccess (0.00s)
            main_test.go:10: It failed!
    FAIL
    exit status 1
    FAIL    bitbucket.org/yourname/yourproject    0.003s
|======|
