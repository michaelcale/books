## How to contribute


### Making a small change

To make a small change to an article, you can click on `Edit on GitHub` link at the bottom of each page. This opens the file directly in GitHub editor.

You can make changes and submit Pull Request from the browser.

### Suggesting a change

For general ideas on how to improve the books, the project etc., use [issue tracker](https://github.com/essentialbooks/books/issues)

### Making more changes

If you plan to make more than a small change, it's good to do it locally.

Toolchain for building the books (i.e. converting markdown sources and source files into HTML) are written in go, so you'll need to [install Go](http://localhost:8080/essential/go/20381-installing-go-compiler).

For cross-platform portability, helper scripts are written in PowerShell, so if you're not on Windows, you'll have to [install it too](https://github.com/PowerShell/PowerShell). Or you can write equivalent bash or what not script. They are trivial.

You'll also need an editor. I use [Visual Studio Code](https://code.visualstudio.com/) with [Code Runner](https://marketplace.visualstudio.com/items?itemName=formulahendry.code-runner) and [Terminal Here](https://marketplace.visualstudio.com/items?itemName=Tyriar.vscode-terminal-here) extensions.

A crucial tool is `./s/preview.ps1`.

It rebuilds all HTML, starts a web server for local preview of changes.

It also watches the source markdown files for changes and rebuilds HTML when they change. That way you can make a change to .md file, save it and refresh corresponding page in the browser to see a change.

### What to improve?

Some articles have implicit notes about what to improve.

You can find them with [this search](https://github.com/essentialbooks/books/search?utf8=%E2%9C%93&q=TODO%3A&type=).

Each book can be improved by adding more articles.

Areas that are likely to always need improvement:
* examples for common tasks (e.g. parsing json/xml/cvs/markdown files, accessing databases etc.)
* examples for popular third-party libraries
