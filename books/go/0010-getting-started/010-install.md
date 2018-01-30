Title: Installing
Id: 20381
Score: 1
Body:

# Installing on Windows

Download `.msi` installer from [https://golang.org/dl/](https://golang.org/dl/) and run it.

# Installing on Mac OS

## Using official binaries

Download `.pkg` installer from [https://golang.org/dl/](https://golang.org/dl/) and run it.

## Using Homebrew

* if you don't have Homebrew installed, install it following [instructions](https://brew.sh/)
* `brew install go`

# On Ubuntu

```
$ sudo apt-get update
$ sudo apt-get install -y build-essential git curl wget
$ wget https://storage.googleapis.com/golang/go<versions>.gz
```
You can find the version lists [here][1].
```
# To install go1.7 use
$ wget https://storage.googleapis.com/golang/go1.7.linux-amd64.tar.gz


# Untar the file
$ sudo tar -C /usr/local -xzf go1.7.linux-amd64.tar.gz
$ sudo chown -R $USER:$USER /usr/local/go
$ rm go1.5.4.linux-amd64.tar.gz
```
Update `$GOPATH`

    $ mkdir $HOME/go

Add following two lines at the end of the ~/.bashrc file
```
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:/usr/local/go/bin:$PATH
```
```
$ nano ~/.bashrc
  export GOPATH=$HOME/go
  export PATH=$GOPATH/bin:/usr/local/go/bin:$PATH

$ source ~/.bashrc
```
Now are set to go, test your go version using:
```
$ go version
go version go<version> linux/amd64
```

# Other OSes

Follow [official instructions](https://golang.org/doc/install).

|======|
