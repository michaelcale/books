---
Title: Cropping image
Id: 31687
Score: 0
---
Most of image type in [image](https://golang.org/pkg/image/) package having `SubImage(r Rectangle) Image` method, except `image.Uniform`. Based on this fact, We can implement a function to crop an arbitrary image as follows

```go
func CropImage(img image.Image, cropRect image.Rectangle) (cropImg image.Image, newImg bool) {
    //Interface for asserting whether `img`
    //implements SubImage or not.
    //This can be defined globally.
    type CropableImage interface {
        image.Image
        SubImage(r image.Rectangle) image.Image
    }

    if p, ok := img.(CropableImage); ok {
        // Call SubImage. This should be fast,
        // since SubImage (usually) shares underlying pixel.
        cropImg = p.SubImage(cropRect)
    } else if cropRect = cropRect.Intersect(img.Bounds()); !cropRect.Empty() {
        // If `img` does not implement `SubImage`,
        // copy (and silently convert) the image portion to RGBA image.
        rgbaImg := image.NewRGBA(cropRect)
        for y := cropRect.Min.Y; y < cropRect.Max.Y; y++ {
            for x := cropRect.Min.X; x < cropRect.Max.X; x++ {
                rgbaImg.Set(x, y, img.At(x, y))
            }
        }
        cropImg = rgbaImg
        newImg = true
    } else {
        // Return an empty RGBA image
        cropImg = &image.RGBA{}
        newImg = true
    }

    return cropImg, newImg
}
```

Note that the cropped image may shared its underlying pixels with the original image. If this is the case, any modification to the cropped image will affect the original image.
