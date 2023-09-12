# proxima

<p align="center">
    <a href="https://github.com/vistormu/adam_simulator">
        <img src="/assets/proxima.svg">
    </a>
</p>

Proxima is a markup language written in pure Go that offers the simplicity of Markdown with the power of LaTeX. It transpiles into HTML, so it is also suitable for teh browser.

> It is currently under development, so actually it is not even as powerful as Markdown ;).

Its syntax is very simple:

- A tag wraps the selected text with some functionality:
```
@<tag>{<text>}
```

- A tag can also wrap the text below until a double line break is encountered.
```
@<tag>
<text>
```

Proxima also supports comments with the `#` character.

## Installation

First, [install Go](https://go.dev/dl/) if you already don't have it.

Proxima generates the PDF using `wkhtmltopdf`, so download it [here](https://wkhtmltopdf.org/index.html).

Currently, the only way of installing is through source code.
- Clone the repository:
  ```
  git clone github.com/vistormu/proxima.git
  ```
- Build the project with:
  ```
  go build build/main.go
  ```

## Usage
The following `test.prox` file shows some basic usage.
```
# test.prox

@h1
This is a section title!

@center
This is centered text!

This is a new paragraph!

@right
This is a paragraph flushed to the right and with @bold{some bold text}!
```

Then, execute the binary code with the file as the first argument:
```
/path/to/main <filename>.prox
```

Also, you can generate the HTML file too with:
```
/path/to/main <filename>.prox --html
```

## Full syntax
- Alignment: `@center`, `@right`
- Headings: `@h1`, `@h2`, `@h3`
- Text styles: `@bold`, `@italic`, `@uline`, `@striket`
- Links: `@url`
- TBI: `@ulist`, `@nlist`, `@email`, `@image`

## TODOs
- Implement more features
- Implement some formatting algorithm
- Change the default style via CSS
