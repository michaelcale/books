Title: Go Interfaces from a Mathematical Aspect
Id: 28599
Score: 4
Body:
In mathematics, especially *Set Theory*, we have a collection of things which is called *set* and we name those things as *elements*. We show a set with its name like A, B, C, ... or explicitly with putting its member on brace notation: {a, b, c, d, e}. Suppose we have an arbitrary element x and a set Z, The key question is: "How we can understand that x is member of Z or not?". Mathematician answer to this question with a concept: **Characteristic Property** of a set. *Characteristic Property* of a set is an expression which describe set completely. For example we have set of *Natural Numbers* which is {0, 1, 2, 3, 4, 5, ...}. We can describe this set with this expression: {a<sub>n</sub> | a<sub>0</sub> = 0, a<sub>n</sub> = a<sub>n-1</sub>+1}. In last expression a<sub>0</sub> = 0, a<sub>n</sub> = a<sub>n-1</sub>+1 is the characteristic property of set of natural numbers. **If we have this expression, we can build this set completely**. Let describe the set of *even numbers* in this manner. We know that this set is made by this numbers: {0, 2, 4, 6, 8, 10, ...}. With a glance we understand that all of this numbers are also a *natural number*, in other words *if we add some extra conditions to characteristic property of natural numbers, we can build a new expression which describe this set*. So we can describe with this expression: {n | n is a member of natural numbers *and* the reminder of n on 2 is zero}. Now we can create a filter which get the characteristic property of a set and filter some desired elements to return elements of our set. For example if we have a natural number filter, both of natural numbers and even numbers can pass this filter, but if we have a even number filter, then some elements like 3 and 137871 can't pass the filter.

Definition of interface in Go is like defining the characteristic property and mechanism of using interface as an argument of a function is like a filter which detect the element is a member of our desired set or not. Lets describe this aspect with code:

    type Number interface {
        IsNumber() bool // the implementation filter "meysam" from 3.14, 2 and 3
    }
    
    type NaturalNumber interface {
        Number
        IsNaturalNumber() bool // the implementation filter 3.14 from 2 and 3
    }
    
    type EvenNumber interface {
        NaturalNumber
        IsEvenNumber() bool // the implementation filter 3 from 2
    }

The characteristic property of `Number` is all structures that have `IsNumber` method, for `NaturalNumber` is all ones that have `IsNumber` and `IsNaturalNumber` methods and finally for `EvenNumber` is all types which have `IsNumber`, `IsNaturalNumber` and `IsEvenNumber` methods. Thanks to this interpretation of interface, easily we can understand that since `interface{}` doesn't have any characteristic property, accept all types (because it doesn't have any filter for distinguishing between values).
|======|
