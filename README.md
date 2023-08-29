# VISML

Vistor's Markup Language (visml) aims to be a markup language between Markdown and LaTeX. It is also very opinionated, leaving only one way of doing things.

Alignment
```
@center
This text is centered.

@justify
This text is justified

@right
This text is flushed to the right

@left
This text is flushed to the left
```

Headers
```
@h0
This is a chapter

@h1
This is a section

@h2
This is a subsection

@h3
This is a subsubsection

...
```

Text styles
```
@bold{This is bold text}

@bold
This is a bold paragraph

@italic{Same with italic text}

@striketrough{Same same}

@underline{Same}
```

Lists
```
@bulletlist
- This is one item.
- This is another item

@enumeration
- This is the first element
- This is the second element
```

Links
```
@url{https://github.com/vistormu}
```

Images
```
@image{assets/fig.png}
```

Other commands
```
@center
This text is centered
@break
And this one too
```
