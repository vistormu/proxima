# Proxima: A Simple Templating Markup Language

<p align="center">
    <a href="https://github.com/vistormu/adam_simulator">
        <img src="/assets/proxima_logo.svg">
    </a>
</p>

Proxima is a markup language that transpiles Proxima source code into HTML. It offers clearer syntax instead of plain HTML/CSS coding.

> Important: Proxima is still under development, so there might be breaking changes between releases until v1.0.0 is out.

## Syntax

The syntax is simple, consisting only of five special characters: 
- `@` defines a tag.
- `#` is used for comments.
- `{` and `}` encloses tag arguments.
- `\` is used to escape a character.

## Components

The main feature of Proxima is that you can create your own components. By default, Proxima looks for files under a `components` directory in the root of your project.

### HTML Components

You can create HTML components by placing `.html` files anywhere under the components' directory.
```
<!-- ./components/spacer.html -->

<div style="heigh: 20px;"></div>
```

In your Proxima file, you can use the component.
```
# ./index.prox

@spacer
```

If you want to add arguments, every `@` symbol inside the file will be replaced by the arguments in appearance order.
```
<!-- ./components/smol-text.html -->
<div style="font-size: 10px;">@</div>
```

And in your Proxima file:
```
# ./index.prox

@smol-text{so basically im very smol}
```

### Python Components
Components can also be written in Python. The Python file should be named after the component's name, and the function must be called `function`.

```
# ./components/list.py

def function(*items: tuple[str]) -> str:
    value: str = "<ul>"
    for item in items:
        value += f"<li>{item}</li>"
    value += "</ul>"

    return value
```

```
# ./index.prox

@list{
This is the first sentence
}{
This is the second sentence
}{
This is the third sentence
}
```

### Adding elements to the head of the document

Proxima will auto-detect all elements that should be placed in the head of the document if they are placed on top of the Proxima document (via a component or embedded HTML).

## Usage

Proxima has four commands available:

- generate [flags] [arguments]
    - generates the HTML file. The arguments can be either files or directories
    - the `-c` flag lets you specify the path to the components' directory. By default, the directory is set to `./components`
    - the `-r` flag tells the compiler to recursively search for `.prox` files in the specified directory and subdirectories 
- watch [flags] [file]
    - watches the Proxima file for changes and auto-generates the HTML file
    - the `-c` flag lets you specify the path to the components' directory. By default, the directory is set to `./components`
- version
    - prints the current version
- help
    - prints the Proxima CLI documentation

## Installation
Download the Proxima binary for your machine from the [Releases Page](https://github.com/vistormu/proxima/releases).

Then, rename it to `proxima`:

```bash
mv proxima-<ARCH> proxima
```

Give it execution privileges:

```bash
chmod +x proxima
```

And move it to the system's binaries:

```bash
mv proxima /usr/local/bin/
```
