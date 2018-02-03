Title: Implement a Unit System with Types
Id: 31206
Score: 0
Body:
This example illustrates how Go's type system can be used to implement some unit system.

    package main
    
    import (
        "fmt"
    )
    
    type MetersPerSecond float64
    type KilometersPerHour float64
    
    func (mps MetersPerSecond) toKilometersPerHour() KilometersPerHour {
        return KilometersPerHour(mps * 3.6)
    }
    
    func (kmh KilometersPerHour) toMetersPerSecond() MetersPerSecond {
        return MetersPerSecond(kmh / 3.6)
    }
    
    func main() {
        var mps MetersPerSecond
        mps = 12.5
        kmh := mps.toKilometersPerHour()
        mps2 := kmh.toMetersPerSecond()
        fmt.Printf("%vmps = %vkmh = %vmps\n", mps, kmh, mps2)
    }

[Open in Playground](https://play.golang.org/p/bhtAQWt5ci)

|======|
