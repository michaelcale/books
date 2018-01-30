Title: Disabling inline expansion
Id: 9094
Score: 1
Body:
Inline expansion can be disabled with the `go:noinline` pragma. For example, if we build the following simple program:

    package main
     
    func printhello() {
        println("Hello")
    }
     
    func main() {
        printhello()
    }

we get output that looks like this (trimmed for readability):

    $ go version
    go version go1.6.2 linux/amd64
    $ go build main.go
    $ ./main
    Hello
    $ go tool objdump main
    TEXT main.main(SB) /home/sam/main.go
            main.go:7       0x401000        64488b0c25f8ffffff      FS MOVQ FS:0xfffffff8, CX
            main.go:7       0x401009        483b6110                CMPQ 0x10(CX), SP
            main.go:7       0x40100d        7631                    JBE 0x401040
            main.go:7       0x40100f        4883ec10                SUBQ $0x10, SP
            main.go:8       0x401013        e8281f0200              CALL runtime.printlock(SB)
            main.go:8       0x401018        488d1d01130700          LEAQ 0x71301(IP), BX
            main.go:8       0x40101f        48891c24                MOVQ BX, 0(SP)
            main.go:8       0x401023        48c744240805000000      MOVQ $0x5, 0x8(SP)
            main.go:8       0x40102c        e81f290200              CALL runtime.printstring(SB)
            main.go:8       0x401031        e89a210200              CALL runtime.printnl(SB)
            main.go:8       0x401036        e8851f0200              CALL runtime.printunlock(SB)
            main.go:9       0x40103b        4883c410                ADDQ $0x10, SP
            main.go:9       0x40103f        c3                      RET
            main.go:7       0x401040        e87b9f0400              CALL runtime.morestack_noctxt(SB)
            main.go:7       0x401045        ebb9                    JMP main.main(SB)
            main.go:7       0x401047        cc                      INT $0x3
            main.go:7       0x401048        cc                      INT $0x3
            main.go:7       0x401049        cc                      INT $0x3
            main.go:7       0x40104a        cc                      INT $0x3
            main.go:7       0x40104b        cc                      INT $0x3
            main.go:7       0x40104c        cc                      INT $0x3
            main.go:7       0x40104d        cc                      INT $0x3
            main.go:7       0x40104e        cc                      INT $0x3
            main.go:7       0x40104f        cc                      INT $0x3
    …

note that there is no `CALL` to `printhello`. However, if we then build the program with the pragma in place:

    package main
     
    //go:noinline
    func printhello() {
        println("Hello")
    }
     
    func main() {
        printhello()
    }

The output contains the printhello function and a `CALL main.printhello`:

    $ go version
    go version go1.6.2 linux/amd64
    $ go build main.go
    $ ./main
    Hello
    $ go tool objdump main
    TEXT main.printhello(SB) /home/sam/main.go
            main.go:4       0x401000        64488b0c25f8ffffff      FS MOVQ FS:0xfffffff8, CX
            main.go:4       0x401009        483b6110                CMPQ 0x10(CX), SP
            main.go:4       0x40100d        7631                    JBE 0x401040
            main.go:4       0x40100f        4883ec10                SUBQ $0x10, SP
            main.go:5       0x401013        e8481f0200              CALL runtime.printlock(SB)
            main.go:5       0x401018        488d1d01130700          LEAQ 0x71301(IP), BX
            main.go:5       0x40101f        48891c24                MOVQ BX, 0(SP)
            main.go:5       0x401023        48c744240805000000      MOVQ $0x5, 0x8(SP)
            main.go:5       0x40102c        e83f290200              CALL runtime.printstring(SB)
            main.go:5       0x401031        e8ba210200              CALL runtime.printnl(SB)
            main.go:5       0x401036        e8a51f0200              CALL runtime.printunlock(SB)
            main.go:6       0x40103b        4883c410                ADDQ $0x10, SP
            main.go:6       0x40103f        c3                      RET
            main.go:4       0x401040        e89b9f0400              CALL runtime.morestack_noctxt(SB)
            main.go:4       0x401045        ebb9                    JMP main.printhello(SB)
            main.go:4       0x401047        cc                      INT $0x3
            main.go:4       0x401048        cc                      INT $0x3
            main.go:4       0x401049        cc                      INT $0x3
            main.go:4       0x40104a        cc                      INT $0x3
            main.go:4       0x40104b        cc                      INT $0x3
            main.go:4       0x40104c        cc                      INT $0x3
            main.go:4       0x40104d        cc                      INT $0x3
            main.go:4       0x40104e        cc                      INT $0x3
            main.go:4       0x40104f        cc                      INT $0x3
     
    TEXT main.main(SB) /home/sam/main.go
            main.go:8       0x401050        64488b0c25f8ffffff      FS MOVQ FS:0xfffffff8, CX
            main.go:8       0x401059        483b6110                CMPQ 0x10(CX), SP
            main.go:8       0x40105d        7606                    JBE 0x401065
            main.go:9       0x40105f        e89cffffff              CALL main.printhello(SB)
            main.go:10      0x401064        c3                      RET
            main.go:8       0x401065        e8769f0400              CALL runtime.morestack_noctxt(SB)
            main.go:8       0x40106a        ebe4                    JMP main.main(SB)
            main.go:8       0x40106c        cc                      INT $0x3
            main.go:8       0x40106d        cc                      INT $0x3
            main.go:8       0x40106e        cc                      INT $0x3
            main.go:8       0x40106f        cc                      INT $0x3
    …
|======|
