### HTML2TXT

`html2txt` is a simple app for converting HTML files to TXT format.
It's using the https://github.com/jaytaylor/html2text under the hood.

To install:
```sh
go get github.com/cycloidio/html2txt
```

Example usage:
```sh
cat {{path_to_html_file}} | html2txt
```

Currently we're only supporting reading input from stdin and stdout for providing output.
All errors are being written to stderr.

For tests and more examples please visit https://github.com/jaytaylor/html2text.

For contributions, please introduce yourself to contribution guide.