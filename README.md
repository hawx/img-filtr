# img-filtr

Port of [straup/filtr][filtr] to use [hawx/img][img] instead of
ImageMagick/GraphicMagick. This means it is slow, and produces different
results, if that sounds like the kind of thing you like install:

``` bash
$ go get hawx.me/code/img # if not already
$ go get hawx.me/code/img-filtr
$ img filtr dthr < in.jpg > out.png
$ img filtr --help
...
$ img filtr drdl --help
...
```

- [x] brdl
- [x] dazd
- [x] dthr
- [x] dthrpxl (equivalent to `img pxl | img filtr dthr`)
- [x] edwn
- [x] heathr
- [x] postcrd
- [x] postr
- [x] pxl (in img)
- [x] pxldthr (equivalent to `img filtr dthr | img pxl`)
- [x] rockstr


[img]:   http://hawx.me/code/img
[filtr]: http://straup.github.com/filtr/
