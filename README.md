<p align="center">
    <a href="https://github.com/vistormu/adam_simulator">
        <img src="/assets/proxima.svg">
    </a>
</p>

Proxima is a markup language written in pure Go that aims to offer a more beautiful and clearer syntax than HTML. It is intended to be used for creating simple webpages with vanilla HTML, CSS and JS; for blog posts, simple documents... 

Proxima has several advantages compared to writing plain HTML:
- It offers a compilation step for syntax checking, so it is not possible to build an HTML with syntax errors.
- You can build HTML components for faster typing.
- You can embed inline HTML syntax

## Syntax

The syntax is very simple, as there is only three special characters: `#` for comments, `\` for escaping a character, and `@` to define a tag. Moreover, there are three times of tags:

- A self-closing tag can have no arguments to offer a specific functionality:
```
@<tag>
```

- A bracketed tag can have as many arguments as needed:
```
@<tag>{<arg1>}{<arg2>}
```

- A wrapping tag wraps the text below until a double line break is encountered:
```
@<tag>
<text>
```

## Installation

The installation process is a bit quirky as there are still no official releases.

First, [install Go](https://go.dev/dl/) if you already don't have it.

Secondly, clone the repository:
```
git clone github.com/vistormu/proxima.git
```
Lastly, build the project with:
```
go build build/main.go
```

## Usage

### Basic Usage
The following `test.prox` file shows some basic usage.
```
# test.prox

@h1
This is a section title!

This is a new paragraph!

This is a paragraph with @bold{some bold text}!
```

Then, execute the binary code with the file as the first argument:
```bash
/path/to/main <filename>.prox
```

As a temporal solution for a terminal command, you can write the following function on your `~/.bashrc` or `~/.zshrc`:
```bash
proxima() {
    if [ "$1" != "generate" ]; then
        echo "Usage: proxima generate"
        return 1
    fi
    ~/path/to/proxima/build/main $1
}
```

In this way, now the usage would be
```bash
proxima generate <filename>.prox
```

If you have multiple `.prox` files in one dorectory, you can use:
```bash
/path/to/main all
```

or
```bash
proxima generate all
```

### Adding a style, script or title
if you want to add a style, script or title, place the respective tags at the beginning of the file. If you want a script to be added at the end of the body, place the tag at the end of the document.

### Components Usage
In order to create your own components, you must create a `components` dir. Inside that directory, there can only `.html` files with an especific name:
- `<tag>-s.html` for self-closing tags.
- `<tag>-b.html` for bracketed tags.
- `<tag>-w.html` for wrapping tags.

If this rules don't apply, the program will not compile.

Inside a `<tag>-b.html` or `<tag>-w.html` every `@` inside the file will be replaced by the number of arguments in order.

For example, if you create this component,
```
# ./components/smalltext-b.html
<div style="font-size: 10px;">@</div>
```

you can use it in your proxima file like:
```
@smalltext{so basically im very smol}
```

## Implemented tags
- Headings: `@h1`, `@h2`, `@h3`
- Text styles: `@bold{<text>}`, `@italic{<text>}`, `@uline{<text>}`, `@strike{<text>}`. `@mark{<text>}`
- Links: `@url{<url>}{<alt text>}`, `@email{<text>}{<alt text>}`
- Images: `@image{<src>}{<width ratio>}`
- External files: `@style{path/to/style.css}`, `@script{path/to/script.js}`
- Other: `@line`, `@break`, `@title{<text>}`
