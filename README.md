# Proxima: A Markup Language Wrapper Markup Language

<p align="center">
    <a href="https://github.com/vistormu/adam_simulator">
        <img src="/assets/proxima_logo.svg">
    </a>
</p>

Proxima is a markup language that wraps your favorite markup language. The main feature is that you can create your own components in your favorite dynamic language: Python, JavaScript, Ruby, and Lua. In this way, you can create reusable components that can be used in any Proxima file.

> Important: Proxima is still under development, so there might be breaking changes between releases until v1.0.0 is out.

## Syntax

The syntax is designed to be minimal and simple, consisting only of five special characters: 
- `@` defines a tag.
- `#` is used for comments.
- `{` and `}` encloses tag arguments.
- `\` is used to escape a character.

## Components

The components are defined under the `components` directory in the root of your project. The components can be written in HTML, Python, JavaScript, Ruby, and Lua.

In the following example, we can create a simple HTML list component in Python:

```
# ./components/list.py

def function(*items: str) -> str:
    value: str = "<ul>"
    for item in items:
        value += f"<li>{item}</li>"
    value += "</ul>"

    return value
```

Then, we can use the component in our Proxima file:

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
