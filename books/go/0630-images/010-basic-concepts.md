---
Title: Basic concepts
Id: 31685
Score: 0
---
An image represents a rectangular grid of picture elements (*pixel*). In the [image](https://golang.org/pkg/image/) package, the pixel is represented as one of the color defined in [image/color](https://golang.org/pkg/image/color/) package. The 2-D geometry of the image is represented as [`image.Rectangle`](https://golang.org/pkg/image/#Rectangle), while [`image.Point`](https://golang.org/pkg/image/#Point) denotes a position on the grid.

![Image and 2-D geometry](https://i.stack.imgur.com/PbRoJ.jpg)

The above figure illustrates the basic concepts of an image in the package. An image of size 15x14 pixels has a rectangular *bounds* started at *upper left* corner (e.g. co-ordinate (-3, -4) in the above figure), and its axes increase right and down to *lower right* corner (e.g. co-ordinate (12, 10) in the figure). Note that the bounds **do not necessarily start from or contain point (0,0)**.

## Image related *type* ##

In `Go`, an image always implement the following [`image.Image`](https://golang.org/pkg/image/#Image) interface

```go
type Image interface {
    // ColorModel returns the Image's color model.
    ColorModel() color.Model
    // Bounds returns the domain for which At can return non-zero color.
    // The bounds do not necessarily contain the point (0, 0).
    Bounds() Rectangle
    // At returns the color of the pixel at (x, y).
    // At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
    // At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
    At(x, y int) color.Color
}
```

in which the [`color.Color`](https://golang.org/pkg/image/color/#Color) interface is defined as

```go
type Color interface {
    // RGBA returns the alpha-premultiplied red, green, blue and alpha values
    // for the color. Each value ranges within [0, 0xffff], but is represented
    // by a uint32 so that multiplying by a blend factor up to 0xffff will not
    // overflow.
    //
    // An alpha-premultiplied color component c has been scaled by alpha (a),
    // so has valid values 0 <= c <= a.
    RGBA() (r, g, b, a uint32)
}
```

and [`color.Model`](https://golang.org/pkg/image/color/#Model) is an interface declared as

```go
type Model interface {
    Convert(c Color) Color
}
```

## Accessing image dimension and pixel ##

Suppose we have an image stored as variable `img`, then we can obtain the dimension and image pixel by:

```
// Image bounds and dimension
b := img.Bounds()
width, height := b.Dx(), b.Dy()
// do something with dimension ...

// Corner co-ordinates
top := b.Min.Y
left := b.Min.X
bottom := b.Max.Y
right := b.Max.X

// Accessing pixel. The (x,y) position must be
// started from (left, top) position not (0,0)
for y := top; y < bottom; y++ {
    for x := left; x < right; x++ {
        cl := img.At(x, y)
        r, g, b, a := cl.RGBA()
        // do something with r,g,b,a color component
    }
}
```

Note that in the package, the value of each `R,G,B,A` component is between `0-65535` (`0x0000 - 0xffff`), **not `0-255`**.
