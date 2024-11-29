# proxima: a markup-language-wrapper markup language!

<p align="center">
    <a href="https://github.com/vistormu/proxima">
        <img src="/assets/proxima_logo.svg">
    </a>
</p>

_proxima_ is a markup language that transpiles to any text-based format you want. Creating Python components, you can extend the language with your own logic and reuse them in any part of your document.

> [!WARNING]
> Proxima is still under development, so there might be breaking changes between releases until v1.0.0 is out.

## basic usage

A common _proxima_ project consists of the following structure:

```plaintext
├── components/         # All proxima components are defined here
│   └── example.py      # Example component
├── main.prox           # A proxima file
└── proxima.toml        # Configuration file for proxima
```
> [!NOTE]
> running `proxima init` initializes a project with this structure.

The components are defined under the `components` directory in the root of your project. 

A component is a python file with a function with the same name as the file. The types of the arguments passed to the function are always "str".

In this case, we have a python component called "example.py":

```python
# components/example.py

def example(arg: str) -> str:
    return f"the passed argument is: {arg}"
```

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

## advanced usage

### arguments

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

### moduled components

By deafult, _proxima_ looks recursively for components in the `components` directory and loads them by their filename.

If we have the following structure:

```
├── components/
│   └── module/
│       └── component.py
```

the component is called in the following way:

```proxima
# main.prox

This is how the component is called by default: @component{}.
```

However, if the `use_modules` option is set to `true` in the `proxima.toml` file (further explained in the [configuring proxima](#configuring-proxima) section), the components can be called as:

```proxima
# main.prox

This is how the component is called with modules: @module.component{}.
```

### default components

If the `use_modules` option is set to `true`, each module can have a default component that is called when the module is called without specifying the component. The default component must be called as the module itself.

```
├── components/
│   └── module/
│       └── module.py
```

```proxima
# main.prox

This is how the default component is called: @module{}.
```

## configuring proxima
_proxima_ can be configured by using the `proxima.toml` file. Here is the default configuration:

```toml
[parser]
line_break_value = "\n"          # The value that will be used as a line break
double_line_break_value = "\n\n" # The value that will be used as a double line break

[evaluator]
python = "python3"               # The python command to evaluate the components
begin_with = ""                  # The string that will be added at the beginning of the evaluated file
end_with = ""                    # The string that will be added at the end of the evaluated file
# text_replacement = [           # Global text replacements
#    {from = "_", to = "\\_"},
# ]

[components]
path = "./components/"           # The directory where the components are stored
use_modules = false              # If true, the components will be called as @module.component
exclude = []                     # Components to exclude from being loaded
```

## installation
Download the _proxima_ binary for your machine from the [Releases Page](https://github.com/vistormu/proxima/releases).

Move it to the system's binaries

```bash
mv proxima-<ARCH> /usr/local/bin/proxima
```

and give it execution privileges:

```bash
chmod +x /usr/local/bin/proxima
```

## tooling
_proxima_ syntax highlighting is available [here](https://github.com/vistormu/tree-sitter-proxima.git).

## future plans
The main goal of _proxima_ is to be a simple, high-configurable, and extensible. The following features are planned for the future:

- [ ] LSP
- [ ] More configuration options
- [ ] Documentation


_proxima_ is a project I made for fun and my personal use. I do not plan to add support for other languages, as I mostly use Python and I would not be able to test them as good as I would like. However, I am open to pull requests and issues, so feel free to contribute!
