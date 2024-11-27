# proxima: a markup-language-wrapper markup language!

<p align="center">
    <a href="https://github.com/vistormu/proxima">
        <img src="/assets/proxima_logo.svg">
    </a>
</p>

_proxima_ is a markup language that transpiles to any text-based format you want. Creating Python components, you can extend the language with your own logic and reuse them in any part of your document.

> [!WARNING]
> Proxima is still under development, so there might be breaking changes between releases until v1.0.0 is out.

## Usage

A common proxima project consists of the following structure:

```plaintext
├── components/         # All proxima components are defined here
│   └── example.py      # Example component
├── main.prox           # A proxima file
└── proxima.toml        # Configuration file for proxima
```
> [!NOTE]
> running `proxima init` initializes a project with this structure.

The components are defined under the `components` directory in the root of your project. In this case, we have a python component called "example.py"

```python
# components/example.py

def example(arg: str) -> str:
    return f"the passed argument is: {arg}"
```

Proxima will execute the function with the same name as the file. This function will always have string arguments and will return a string.

Then, we can use the component in our Proxima file:

```proxima
# main.prox

This is a proxima example, and @example{2}.
```

Finally, by running

```bash
proxima make main.prox -o main.txt
```

the output file should look like this:

```txt
# main.txt

This is a proxima example, and the passed argument is: 2.
```

Multiple arguments can be passed to the component by enclosing them in curly braces.

```proxima
# main.prox

This is a proxima example, and @component_with_two_args{one}{two}.
```

Also, you can pass the arguments by name by using the following syntax:

```proxima
# main.prox

This is a proxima example, and @component_with_named_args{<first> one}{<second> two}.
```

## Configuring proxima
_proxima_ can be configured by using the `proxima.toml` file. Here is the default configuration:

```toml
[parser]
line_break_value = "\n" # The value that will be used as a line break
double_line_break_value = "\n\n" # The value that will be used as a double line break

[evaluator]
python = "python3" # The python command to evaluate the components
# text_replacement = [ # Global text replacements
#    {from = "_", to = "\\_"},
# ]

[components]
path = "./components/" # The directory where the components are stored
use_modules = false # If true, the components will be called as @module.component
exclude = [] # Components to exclude from being loaded
```

## Installation
Download the _proxima_ binary for your machine from the [Releases Page](https://github.com/vistormu/proxima/releases).

Move it to the system's binaries

```bash
mv proxima-<ARCH> /usr/local/bin/proxima
```

and give it execution privileges:

```bash
chmod +x /usr/local/bin/proxima
```

## Tooling
_proxima_ syntax highlighting is available [here](https://github.com/vistormu/tree-sitter-proxima.git).

## Future plans
The main goal of _proxima_ is to be a simple, high-configurable, and extensible markup language that can transpile to any other markup language. The following features are planned for the future:

- [ ] LSP
- [ ] More configuration options
- [ ] Documentation
- [ ] Add support for other dynamic languages
