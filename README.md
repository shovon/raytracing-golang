# Ray Tracing in One Weekend in Go

![Output from ray tracing](/output.png)

This is a ray tracing implementation in Go.

All implementation was translated from the minibook _[Ray Tracing In One Weekend](https://www.amazon.ca/Ray-Tracing-Weekend-Minibooks-Book-ebook/dp/B01B5AODD8)_. The book opted for C++, but I figured a good way to both practice Go as well as learn ray tracing would be to go through the minibook, without copying and pasting. Writing in Go is a good way to motivate me to read through the book.

## Running the ray tracer

Be sure to have [Go](https://golang.org/) installed. Then, from this directory, just run:

```
go run . > image.ppm
```

The program outputs portable pixmap format to the console/stdout, and so, the `>` symbol writes the output to a file.

Rendering a whole 1024 by 576 image takes around 10 minutes my 2019 M1 MacBook Pro. If you want something rendered quicker, change the `nx` and `ny` variables in `main()`, in `main.go` to something smaller than 1024 by 576 (perhaps 200 and 100, respectively).
