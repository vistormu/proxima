<a name="readme-top"></a>

<div align="center">

<a href="https://github.com/vistormu/proxima" target="_blank" title="go to the repo"><img width="196px" alt="proxima logo" src="/docs/logo.png"></a>


# proxima<br>a markup-language-wrapper markup language!

_proxima_ is a markup language that transpiles to any text-based format you want

creating python components, you can extend the language with your own logic and reuse them in any part of your document

<br>

[![go version][go_version_img]][go_dev_url]
[![License][repo_license_img]][repo_license_url]

<br>

</div>

> [!WARNING]
> this project is functional but still in development, so expect some bugs and missing features

## basic usage

a common _proxima_ project consists of the following structure:

```plaintext
├── components/         # all proxima components are defined here
│   └── example.py      # example component
├── main.prox           # a proxima file
└── proxima.toml        # configuration file for proxima
```

> [!NOTE]
> running `proxima init` initializes a project with this structure.

the components are defined under the `components` directory in the root of your project. 

a component is a python file with a function with the same name as the file. The types of the arguments passed to the function are always "str".

in this case, we have a python component called "example.py":

```python
# components/example.py

def example(arg: str) -> str:
    return f"the passed argument is: {arg}"
```

then, we can use the component in our Proxima file:

```proxima
# main.prox

this is a proxima example, and @example{2}.
```

finally, by running

```bash
proxima make main.prox -o main.txt
```

the output file should look like this:

```txt
# main.txt

this is a proxima example, and the passed argument is: 2.
```

## advanced usage

### arguments

multiple arguments can be passed to the component by enclosing them in curly braces.

```proxima
# main.prox

this is a proxima example, and @component_with_two_args{one}{two}.
```

also, you can pass the arguments by name by using the following syntax:

```proxima
# main.prox

this is a proxima example, and @component_with_named_args{<first> one}{<second> two}.
```

### moduled components

by deafult, _proxima_ looks recursively for components in the `components` directory and loads them by their filename.

if we have the following structure:

```
├── components/
│   └── module/
│       └── component.py
```

the component is called in the following way:

```proxima
# main.prox

this is how the component is called by default: @component{}.
```

however, if the `use_modules` option is set to `true` in the `proxima.toml` file (further explained in the [configuring proxima](#configuring-proxima) section), the components can be called as:

```proxima
# main.prox

this is how the component is called with modules: @module.component{}.
```

### default components

if the `use_modules` option is set to `true`, each module can have a default component that is called when the module is called without specifying the component. The default component must be called as the module itself.

```
├── components/
│   └── module/
│       └── module.py
```

```proxima
# main.prox

this is how the default component is called: @module{}.
```

## configuring proxima
_proxima_ can be configured by using the `proxima.toml` file

check the default configuration [here](/internal/assets/proxima.toml)

## installation

### homebrew

you can install _proxima_ using homebrew by running the following command:

```bash
brew install vistormu/proxima/proxima
```

### from releases

download the _proxima_ binary for your machine from the [releases page](https://github.com/vistormu/proxima/releases).

move it to the system's binaries

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


[go_version_img]: https://img.shields.io/badge/Go-1.24+-00ADD8?style=for-the-badge&logo=go
[go_dev_url]: https://go.dev/
[go_report_img]: https://goreportcard.com/badge/github.com/vistormu/cahier
[go_report_url]: https://goreportcard.com/report/github.com/vistormu/cahier
[repo_license_img]: https://img.shields.io/github/license/vistormu/cahier?style=for-the-badge
[repo_license_url]: https://github.com/vistormu/cahier/blob/main/LICENSE
