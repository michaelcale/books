Title: Get, Install And Setup Atom & Gulp
Id: 26863
Score: 0
Body:

1. Install Atom. You can get atom from [here][1]
2. Go to Atom settings (ctrl+,). Packages -> Install go-plus package ([go-plus][2])

After Installing go-plus in Atom:
[![atom-setting-img][3]][3]

3. Get these dependencies using go get or another dependency manager: (open a console and run these commands)

> go get -u golang.org/x/tools/cmd/goimports

> go get -u golang.org/x/tools/cmd/gorename 

> go get -u github.com/sqs/goreturns 

> go get -u github.com/nsf/gocode

> go get -u github.com/alecthomas/gometalinter

> go get -u github.com/zmb3/gogetdoc

> go get -u github.com/rogpeppe/godef

> go get -u golang.org/x/tools/cmd/guru

4. Install Gulp ([Gulpjs][4]) using npm or any other package manager ([gulp-getting-started-doc][5]):

> $ npm install --global gulp


  [1]: https://github.com/atom/atom/releases/tag/v1.12.7 "here"
  [2]: https://atom.io/packages/go-plus
  [3]: https://i.stack.imgur.com/HSbug.png
  [4]: http://gulpjs.com/
  [5]: https://github.com/gulpjs/gulp/blob/master/docs/getting-started.md
|======|
