---
Title: Convert color image to grayscale
Id: 31693
Score: 0
---
Some digital image processing algorithm such as edge detection, information carried by the image intensity (i.e. grayscale value) is sufficient. Using color information (`R, G, B` channel) may provides slightly better result, but the algorithm complexity will be increased. Thus, in this case, we need to convert the color image to grayscale image prior to applying such algorithm.

The following code is an example of converting arbitrary image to 8-bit grayscale image. The image is retrieved from remote location using `net/http` package, converted to grayscale, and finally saved as PNG image.

```go
package main

import (
    "image"
    "log"
    "net/http"
    "os"

    _ "image/jpeg"
    "image/png"
)

func main() {
    // Load image from remote through http
    // The Go gopher was designed by Renee French. (http://reneefrench.blogspot.com/)
    // Images are available under the Creative Commons 3.0 Attributions license.
    resp, err := http.Get("http://golang.org/doc/gopher/fiveyears.jpg")
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    // Decode image to JPEG
    img, _, err := image.Decode(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Image type: %T", img)

    // Converting image to grayscale
    grayImg := image.NewGray(img.Bounds())
    for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
        for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
            grayImg.Set(x, y, img.At(x, y))
        }
    }

    // Working with grayscale image, e.g. convert to png
    f, err := os.Create("fiveyears_gray.png")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    if err := png.Encode(f, grayImg); err != nil {
        log.Fatal(err)
    }
}
```

 Color conversion occurs when assigning pixel through `Set(x, y int, c color.Color)` which is implemented in [`image.go`](https://golang.org/src/image/image.go?s=19292:19335#L691) as

```go
func (p *Gray) Set(x, y int, c color.Color) {
    if !(Point{x, y}.In(p.Rect)) {
        return
    }

    i := p.PixOffset(x, y)
    p.Pix[i] = color.GrayModel.Convert(c).(color.Gray).Y
}
```

in which, `color.GrayModel` is defined in [`color.go`](https://golang.org/src/image/color/color.go?s=2699:2728#L110) as

```go
func grayModel(c Color) Color {
    if _, ok := c.(Gray); ok {
        return c
    }
    r, g, b, _ := c.RGBA()

    // These coefficients (the fractions 0.299, 0.587 and 0.114) are the same
    // as those given by the JFIF specification and used by func RGBToYCbCr in
    // ycbcr.go.
    //
    // Note that 19595 + 38470 + 7471 equals 65536.
    //
    // The 24 is 16 + 8. The 16 is the same as used in RGBToYCbCr. The 8 is
    // because the return value is 8 bit color, not 16 bit color.
    y := (19595*r + 38470*g + 7471*b + 1<<15) >> 24

    return Gray{uint8(y)}
}
```

Based on the above facts, the intensity `Y` is calculated with the following formula:

> Luminance: Y = 0.299**R** + 0.587**G** + 0.114**B**

If we want to apply different [formula/algorithms](http://journals.plos.org/plosone/article?id=10.1371/journal.pone.0029740) to convert a color into an intesity, e.g.

> Mean: Y = (**R** + **G** + **B**) / 3
> Luma: Y = 0.2126**R** + 0.7152**G** + 0.0722**B**
> Luster: Y = (min(**R**, **G**, **B**) + max(**R**, **G**, **B**))/2

then, the following snippets can be used.

```go
// Converting image to grayscale
grayImg := image.NewGray(img.Bounds())
for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
    for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
        R, G, B, _ := img.At(x, y).RGBA()
        //Luma: Y = 0.2126*R + 0.7152*G + 0.0722*B
        Y := (0.2126*float64(R) + 0.7152*float64(G) + 0.0722*float64(B)) * (255.0 / 65535)
        grayPix := color.Gray{uint8(Y)}
        grayImg.Set(x, y, grayPix)
    }
}
```

The above calculation is done by floating point multiplication, and certainly is not the most efficient one, but it's enough for demonstrating the idea. The other point is, when calling `Set(x, y int, c color.Color)` with `color.Gray` as third argument, the color model will not perform color conversion as can be seen in the previous `grayModel` function.
