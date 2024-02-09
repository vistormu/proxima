# Proxima: A Simple Templating Markup Language

<p align="center">
    <a href="https://github.com/vistormu/adam_simulator">
        <img src="/assets/proxima_logo.svg">
    </a>
</p>

Proxima is a markup language that transpiles Proxima source code into HTML. It offers clearer syntax instead of plain HTML/CSS coding.

> Important: Proxima is still under development, so there might be breaking changes between releases until v1.0.0 is out.

## Syntax

The syntax is very simple, consisting only of five special characters: 
- `@` defines a tag.
- `#` is used for comments.
- `{` and `}` encloses tag arguments.
- `\` is used to escape a character.


## Basic Usage

### Components

A proxima file can be written as an HTML file. However, you can create HTML components placing `.html` files under a `components` directory in the root of your project. For example:
```
<!-- ./components/spacer.html -->

<div style="heigh: 20px;"></div>
```

So now in your proxima file you can use
```
# ./index.prox

@spacer
```

If you want to add arguments, every `@` symbol inside of the file will be replaced by the arguments in appearance order.
```
<!-- ./components/smol-text.html -->
<div style="font-size: 10px;">@</div>
```

And in your proxima file:
```
# ./index.prox

@smol-text{so basically im very smol}
```

If the tag has only one argument, it can wrap the content until a double linebreak is encountered.
```
# ./index.prox

@smol-text
so basically
im very smol
```

### HTML file generation

For creating the HTML file, use the command:
```
proxima <filename>.prox
```

For generating all proxima files recursively, use:
```
proxima all
```

### Adding a style, script or title
Proxima offers builtin tags such as `@style`, `@script`, and `@title`. These can be placed at the beginning of a `.prox` file to include them into the `<header>` section of the HTML. Moreover, all embedded HTML code that begins with `<link` will also be placed in the `<header>` section.

## Installation
Download the install script from [Releases Page](https://github.com/vistormu/proxima/releases) and execute it:

```
sh install-proxima.sh
```
