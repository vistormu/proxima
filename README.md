# Proxima: A Simple Templating Markup Language

<p align="center">
    <a href="https://github.com/vistormu/adam_simulator">
        <img src="/assets/proxima_logo.svg">
    </a>
</p>

Proxima is a markup language written in pure Go that aims to offer a more beautiful and clearer syntax than HTML. It is intended to be used for creating simple webpages with vanilla HTML, CSS and JS; for blog posts, simple documents... 

Proxima has several advantages compared to writing plain HTML:
- It offers a compilation step for syntax checking, so it is not possible to build an HTML with syntax errors.
- You can build HTML components for faster typing.
- You can embed inline HTML syntax

> Important: Proxima is still under development, so there might be breaking changes between releases until v1.0.0 is out.

## Syntax

The syntax is very simple, as there is only five special characters: 
- `@` is used to define a tag.
- `#` is used for comments.
- `{` and '}' are used for enclosing arguments.
- `\` is used to escape a character.

A tag can be used in three ways:

- A tag with no arguments
```
@<tag>
```

- A tag can have as many arguments as needed:
```
@<tag>{<arg1>}{<arg2>}
```

- If the tag has only one argument, it can wrap the content until a double linebreak is encountered.
```
@<tag>
<text>
```

## Installation
Download the binary for your operating system in the [Releases Page](https://github.com/vistormu/proxima/releases).

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
    shift ~/proxima-linux-arm64 $@
}
```

In this way, now the usage would be
```bash
proxima generate <filename>.prox
```

If you have multiple `.prox` files in one directory, you can use:
```bash
proxima generate all
```

### Adding a style, script or title
If you want to add a style, script or title, place the respective tags at the beginning of the file. If you want a script to be added at the end of the body, place the tag at the end of the document.

### Components Usage
In order to create your own components, you must create a `components` directory in the root of your project. Inside that directory, you can define your new components by creating a `.html` file with the name of your component. For example:
```
<!-- ./components/spacer.html -->

<div style="heigh: 20px;"></div>
```

So now in your proxima file you can use
```
# ./index.prox

@spacer
```

If you want to add arguments, every `@` symbol inside of the file will be replaces by the arguments in appearance order.
```
<!-- ./components/smol-text.html -->
<div style="font-size: 10px;">@</div>
```

And in your proxima file:
```
@smol-text{so basically im very smol}
```

## Pre-implemented tags
The basic HTML tags are already pre-implemented.

- Headings: `@h1`, `@h2`, `@h3`
- Text styles: `@bold{<text>}`, `@italic{<text>}`, `@uline{<text>}`, `@strike{<text>}`. `@mark{<text>}`
- Links: `@url{<url>}{<alt text>}`, `@email{<text>}{<alt text>}`
- Images: `@image{<src>}{<width ratio>}`
- External files: `@style{path/to/style.css}`, `@script{path/to/script.js}`
- Other: `@line`, `@break`, `@title{<text>}`
