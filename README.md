[![license](http://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/arthurkay/dload/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/arthurkay/gopherchain)](https://goreportcard.com/report/github.com/arthurkay/dload)

# Dload

A golang CLI module for file download. You only need to use one method that initiates the download and sets the download directory.

## Installation

The module can be added to your module by using:

```bash
go get github.com/arthurkay/dload
```

To use this module in your applicaton, first import it, then you can use it by using the `Download(url, dest string)`. e.g

```go
import (
    ...
    "github.com/arthurkay/dload"
    ...
)
```

Then somewhere in your function call the Download method passing the url to download from and the destination to save the file to, like below:

```go

...
dload.Download("example.com/download/something.mp4", "./")
...
```


&copy; Arthur Kalikiti 2021
