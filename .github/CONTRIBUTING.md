# Contributing to Sonobabble

:tada::grinning: First off, thanks for checking out Sonobabble! :tada::grinning: Before you do anything, please see
[the readme][readme] for more general details about the project.

-----------------------------------------------------------------------------------------------------------------------

## Getting Started

 1. Install [Golang][go website]! Sonobabble is powered by Google’s free and [open source][go github] language Go.
    Although at the time of writing Go is version go1.8, Sonobabble is currently using go1.7.4. Hopefully, Sonobabble
    will be updated to go1.8 soon.

 2. Install [Gorilla Mux][gorilla mux]! Gorilla Mux is a URL router for Golang, which is required by Sonobabble.

 3. 	go get github.com/skunkmb/sonobabble

 4. 	cd $GOPATH/src/github.com/skunkmb/sonobabble

 5. 	go run start.go

 6. Open up [localhost on port 8080][localhost port 8080].

 7. Enjoy!

## Committing

When committing to a fork of Sonobabble, preferably follow some simple and generic guidelines:

### 1. Use imperative tense

Do

> Add foobar because baz

instead of

> Added foobar because baz

### 2. Use proper styling in the title

Start a title with a capital letter, and don’t put a period at the end. Also, in the title, because it is not rendered
with Markdown, _don’t_ do any special notation for variable or function names, such as backticks, in order to more
closely align to [GoDoc][godoc]’s standards.

Do

> Change foobar

instead of

> change \`foobar\`.

### 3. Limit the title to 50 characters or less

Be descriptive, but not over-the-top.

Do

> Change foobar to baz instead of qux

instead of

> Change foobar to do baz instead of qux when inside of foobaz because barbaz

or

> Change variable

### 4. Inside of the body, _do not_ use Markdown

Because commits are rendered (usually) as plaintext, not with Markdown, do not do any Markdown-specific formatting
inside of a commit title or message that will look ugly without Markdown. For example, do not use backticks, in order
to more closely align to [GoDoc][godoc]’s standards. Bullet points made from hyphens, or numeric lists, are fine.

Do

> Change foobar to baz instead of qux
>
> Call foobar with baz set to true, so that in the future qux, barbaz, and  
> foobaz won’t happen. Do this for three main reasons:
>
>  &nbsp;1\. Lorem ipsum.
>
>  &nbsp;2\. Dolor sit amet.
>
>  &nbsp;3\. Consectetur adipiscing elit.
>
> Also do it for two other reasons:
>
>  &nbsp;\- Quisque euismod
>
>  &nbsp;\- Lorem et lobortis

instead of

> Change \`foobar\` to \`baz\` instead of \`qux\`
>
> \# Call \`foobar\` with \`baz\` set to true, so that in the future \`qux\`,  
> \`barbaz\`, and \`foobaz\` won’t happen. Do this for  
> \[three main reasons](https://foobarqux.net):
>
>  &nbsp;1\. Lorem ipsum.
>
>  &nbsp;2\. Dolor sit amet.
>
>  &nbsp;3\. Consectetur adipiscing elit.
>
> _Also do it for two other reasons_:
>
>  &nbsp;\- Quisque euismod
>
>  &nbsp;\- Lorem et lobortis

### 5. Wrap the body at 72 characters

Wrapping the body at 72 characters is common for a Git commit, because commit messages are often seen as inside of an
80-character-wide a terminal.

### 6. Explain what and why, not how

The specific changes of a pull request can always be seen by clicking on “Files changed” in GitHub, so a pull request
with a body that only states specific code changes isn’t very helpful.

Do

> Call foobar with baz set to true, so that in the future qux won’t happen.

instead of

> Instead of calling foobar(true) on line 100, call foobar(false).

## Making a pull request

When making a pull request, the rules are very similar to those when making a commit. However, there are a few
exceptions:

### 1. Inside of the body, _do_ use Markdown

Although commit messages should not have Markdown because they are seen as plaintext, Markdown can and should be used
in a pull requests’s body to make them more readable.

Note that Markdown like backticks _still should not_ be put in a pull request’s title, as it is _not_ rendered with
Markdown.

Do
> Change foobar to baz instead of qux
>
> ### Call `foobar` with `baz` set to true, so that in the future `qux`, `barbaz`, and `foobaz` won’t happen.
>
> Do this for [three main reasons](https://foobarqux.net):
>
>  1. Lorem ipsum.
>
>  2. Dolor sit amet.
>
>  3. Consectetur adipiscing elit.
>
> _Also do it for two other reasons_:
>
>  - Quisque euismod
>
>  - Lorem et lobortis

instead of

> Change `foobar` to `baz` instead of `qux`
>
> Call foobar with baz set to true, so that in the future qux, barbaz, and foobaz won’t happen. Do this for three
> main reasons: lorem ipsum, dolor sit amet, and consectetur adipiscing elit.
>
> Also do it for two other reasons: quisque euismod, and lorem et lobortis.

### 2. _Do not_ wrap the body at 72 characters

Wrapping the body at 72 characters is common for a Git commit, but it does not have to be done in a pull request’s
body, because it is seen in a web browser, not a 80-character-wide terminal.

Do

> Change foobar to baz instead of qux
>
> ### Call `foobar` with `baz` set to true, so that in the future `qux`, `barbaz`, and `foobaz` won’t happen.
>
> Do this for [three main reasons](https://foobarqux.net):
>
>  1. Lorem ipsum.
>
>  2. Dolor sit amet.
>
>  3. Consectetur adipiscing elit.
>
> _Also do it for two other reasons_:
>
>  - Quisque euismod
>
>  - Lorem et lobortis

instead of

> Change foobar to baz instead of qux
>
> ### Call `foobar` with `baz` set to true, so that in the future  
> `qux`, `barbaz`, and `foobaz` won’t happen.
>
> Do this for  
> [three main reasons](https://foobarqux.net):
>
>  1. Lorem ipsum.
>
>  2. Dolor sit amet.
>
>  3. Consectetur adipiscing elit.
>
> _Also do it for two other reasons_:
>
>  - Quisque euismod
>
>  - Lorem et lobortis

## Making an issue

Generally, there are not really any specific “rules” for when making issues. Just use common sense and good formatting,
just like when making a pull request.

# Style

Code styling is fairly complicated. For the most part, when submitting code to Sonobabble, try to follow
[CodeReviewComments][codereviewcomments] and [Effective Go][effective go], and _always_ use [`go fmt`][go fmt].

However, there are few things to do differently than just following those guides:

## 1. Code Wrapping

Wrap all lines so that they do not go over 119 characters. As mentioned in [`c0f58b9`][c0f58b9], a line-width of the
commonly-used 79 can stifle descriptive variable names. Also, for those who want to set their tab-width to 8 spaces, 79
can seem very small. Since Effective Go [specifically says][effective go line-width] that

> Go has no line length limit. Don't worry about overflowing a punched card.

use a line width of 119, not 79.

## 2. Variables

### a. Variable name length

CodeReviewComments [says][codereviewcomments variable names] that

> Variable names in Go should be short rather than long. This is especially true for local variables with limited
> scope. Prefer `c` to `lineCount`. Prefer `i` to `sliceIndex`.

However, although short variable names may have some slight advantages for things like coding speed, they actually can
hinder code readability and simplicity in the long run. Go _against_ CodeReviewComments and instead feel free to use
long variable names.

### b. Error names

When contributing to Sonobabble, do not name `error` type variables `err` or `error`. Instead, name them something
relevant to the specific task that the error problem happened on. For example, if checking for an error on `os.Stat`,
call the `error` `statError`.

### c. Variable declarations

Although this may seem very inconvenient, _always_ use explicit variable types. _Never_ say something like

```golang
foo := bar()
```

because at first glance, there is no way to know what `bar` returns. Instead, say

```golang
var foo string = bar()
```

Also, for functions that return multiple values, use the `var()` syntax. For example,

```golang
var (
	foo      string
	fooError error
)
foo, fooError = bar()
```

## 3. Commenting

When commenting, only use two slashes when the entire comment fits on one line. In any other case, use a slash and an
asterisk, so that comments can be easily edited later.

Do

```golang
// fooBar lorem ipsum dolor sit amet.
```

or

```golang
/*
	fooBar Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut
	labore et dolore magna aliqua. Ut enim ad minim veniam.
*/
```

but not

```golang
/* fooBar lorem ipsum dolor sit amet. */
```

or

```golang
/*
	fooBar lorem ipsum dolor sit amet.
*/
```

or

```golang
// fooBar Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut
// labore et dolore magna aliqua. Ut enim ad minim veniam.
```

Also, _never_ use backticks or other Markdown in comments. Follow [GoDoc][godoc]’s guide on documentation comments
(start with the function or variable’s name, end with a period, et cetera).

Comments should be full sentences, starting with a capital letter (unless its the name of a function or variable), and
ending with a period.

Comments should not reference specific parameter names or return types. One can always find the type or name of a
parameter or return value. Instead, they should be in “human-speak.”

Do

```golang
// foo returns a number that is equal to a given number plus 1.
func foo(uint8 numberToReturn) uint8 {
	return numberToReturn + 1
}
```

instead of

```golang
// foo returns uint8 numberToReturn + 1.
func foo(uint8 numberToReturn) uint8 {
	return numberToReturn
}
```

Lastly, use comments when needed, but not crazily. Comments should refer to the code that is under the comment without
any newlines between them, or the entirety of a statement it is directly under the comment. For example,

```golang
// This comment applies to foo and bar, but not baz or the if statement.
foo()
bar()

baz()

// This comment applies to the if statement and all of its contents (qux, bazbar, and the return statement).
if true {
	// This comment only applies to qux.
	qux()

	// This comment only applies to bazbar and the return.
	bazbar()
	return true
}
```

## 4. Error messages

Returned error messages should start with a _lowercase_ letter and _not_ end with a period. They should also _not_
start with the name of the function. Instead, the name should be added by the parent function. For example,

```golang
func getErrorMessage() error {
	return errors.New("received a reason to return an error")
}

func main() {
	var getError error = getErrorMessage()

	if getError != nil {
		// Add the name of the function to the error message after-the-fact.
		panic("getErrorMessage: %s", getError)
	}
}
```

## Statements

Any statement with curly brackets (function declarations, if statements, loops, et cetera) should have a newline above
it, _unless_ it is the first or statement in a parent statement.

Do

```golang
func main() {
	if foo {
		bar()
	}

	if baz {
		qux()
	}
}
```

instead of

```golang
func main() {

	if foo {
		bar()
	}
	if baz {
		qux()
	}

}
```

# Conclusion

:+1: If you’ve stuck it out this far, thanks! :+1:

Please note that these rules are not set in stone, and are very likely to change.

[readme]: ../README.md
[go website]: https://golang.org
[go github]: https://github.com/golang/go
[gorilla mux]: https://github.com/gorilla/mux
[localhost port 8080]: http://localhost:8080
[pull request]: https://github.com/skunkmb/sonobabble/compare
[issue]: https://github.com/skunkmb/sonobabble/issues/new
[godoc]: https://blog.golang.org/godoc-documenting-go-code
[codereviewcomments]: https://github.com/golang/go/wiki/CodeReviewComments
[effective go]: https://golang.org/doc/effective_go.html
[go fmt]: https://blog.golang.org/go-fmt-your-code
[c0f58b9]: https://github.com/skunkmb/sonobabble/commit/c0f58b96d81f3c78b52d49e556f155ed44efacfa
[effective go line-width]: https://golang.org/doc/effective_go.html#formatting
[codereviewcomments variable names]: https://github.com/golang/go/wiki/CodeReviewComments#variable-names
