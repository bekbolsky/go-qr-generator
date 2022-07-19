# go-qr-generator

![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/bekbolsky/go-qr-generator?filename=go.mod)
![GitHub](https://img.shields.io/github/license/bekbolsky/go-qr-generator)

A simple CLI application to generate a QR code for URL links in a given text file.

# Usage

Make sure you have Go installed and have the `go` command available in your PATH.

First, clone the repository:

```bash
git clone https://github.com/bekbolsky/go-qr-generator.git
```

Then, go to the directory:

```bash
cd go-qr-generator
```

Build the application:

```bash
go build
```

Then, you can use the application:

```bash
go-qr-generator -input=path-to-your-file.txt -output=your-folder-or-path
```

The application will generate a QR code for each line in the input file.

Additionally, you can use the `-help` flag to see the usage:

```bash
go-qr-generator -help
```

There are also some other flags:

- `-transparent`: enables the transparent background of the QR code. Default is `false`.
- `-fgcolor`: sets the foreground color of the QR code. Default is `#000000`.

# License

[MIT](https://choosealicense.com/licenses/mit/)
