# proxima: A Markup-Language-Wrapper Markup Language

<p align="center">
    <a href="https://github.com/vistormu/adam_simulator">
        <img src="/assets/proxima_logo.svg">
    </a>
</p>

_proxima_ is a markup language that wraps your favorite markup language with additional logic! You can create your own components in your favorite dynamic language: Python, JavaScript, Ruby, and Lua. These components can be then reused in any part of your document!

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

## Configuring proxima
_proxima_ can be configured by using the `proxima.toml` file. Here, you can change the line break values, the language runtimes, specify the components directory, and wether the components should use modules (e.g. if true, a component defined under "components/module/component.py" is then called in a .prox file as "@module.component").

```toml
[parser]
line_break_value = "\n"
double_line_break_value = "\n\n"

[evaluator]
python_cmd = "python3 -c"
javascript_cmd = "node -e"
lua_cmd = "lua -e"
ruby_cmd = "ruby -e"

[components]
components_dir = "./components/"
use_modules = false
exclude = []
```

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

## Tooling
Proxima syntax highlighting is available [here](https://github.com/vistormu/tree-sitter-proxima.git), and the LSP is under development.
