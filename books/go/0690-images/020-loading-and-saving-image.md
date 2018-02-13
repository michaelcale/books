Title: Loading and saving image
Id: 31686
Score: 0
Body:
In memory, an image can be seen as a matrix of pixel (color). However, when an image being stored in a permanent storage, it may be stored as is (RAW format), [Bitmap](https://en.wikipedia.org/wiki/Bitmap) or other image formats with particular compression algorithm for saving storage space, e.g. PNG, JPEG, GIF, etc. When loading an image with particular format, the image must be *decoded* to `image.Image` with corresponding algorithm. An [`image.Decode`](https://golang.org/pkg/image/#Decode) function declared as

    func Decode(r io.Reader) (Image, string, error)

is provided for this particular usage. In order to be able to handle various image formats, prior to calling the `image.Decode` function, the decoder must be registered through [`image.RegisterFormat`](https://golang.org/pkg/image/#RegisterFormat) function defined as

```go
func RegisterFormat(name, magic string,
    decode func(io.Reader) (Image, error), decodeConfig func(io.Reader) (Config, error))
```

Currently, the image package supports three file formats: [JPEG](https://golang.org/pkg/image/jpeg/), [GIF](https://golang.org/pkg/image/gif/) and [PNG](https://golang.org/pkg/image/png/). To register a decoder, add the following

    import _ "image/jpeg"    //register JPEG decoder

to the application's `main` package. Somewhere in your code (not necessary in `main` package), to load a JPEG image, use the following snippets:

```go
f, err := os.Open("inputimage.jpg")
if err != nil {
    log.Fatalf("os.Open() failed with %s\n", err)
}
defer f.Close()

img, fmtName, err := image.Decode(f)
if err != nil {
    log.Fatalf("image.Decode() failed with %s\n", err)
}

// `fmtName` contains the name used during format registration
// Work with `img` ...
```

## Save to PNG ##

 To save an image into particular format, the corresponding *encoder* must be imported explicitly, i.e.

    import "image/png"    //needed to use `png` encoder

then an image can be saved with the following snippets:

```go
f, err := os.Create("outimage.png")
if err != nil {
    log.Fatalf("os.Create() failed with %s\n", err)
}
defer f.Close()

// Encode to `PNG` with `DefaultCompression` level
// then save to file
err = png.Encode(f, img)
if err != nil {
    log.Fatalf("png.Encode() failed with %s\n", err)
}
```

If you want to specify compression level other than `DefaultCompression` level, create an *encoder*, e.g.

```go
enc := png.Encoder{
    CompressionLevel: png.BestSpeed,
}
err := enc.Encode(f, img)
```

## Save to JPEG ##

To save to `jpeg` format, use the following:

```go
import "image/jpeg"

// Somewhere in the same package
f, err := os.Create("outimage.jpg")
if err != nil {
    log.Fatalf("os.Create() failed with %s\n", err)
}
defer f.Close()

// Specify the quality, between 0-100
// Higher is better
opt := jpeg.Options{
    Quality: 90,
}
err = jpeg.Encode(f, img, &opt)
if err != nil {
    log.Fatalf("jpeg.Encode() failed with %s\n", err)
}
```

## Save to GIF ##

To save the image to GIF file, use the following snippets.

```go
import "image/gif"

// Samewhere in the same package
f, err := os.Create("outimage.gif")
if err != nil {
    log.Fatalf("os.Create() failed with %s\n", err)
}
defer f.Close()

opt := gif.Options {
    NumColors: 256,
    // Add more parameters as needed
}

err = gif.Encode(f, img, &opt)
if err != nil {
    log.Fatalf("gif.Encode() failed with %s\n", err)
}
```
