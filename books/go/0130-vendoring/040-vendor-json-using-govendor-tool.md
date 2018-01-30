Title: vendor.json using Govendor tool
Id: 30033
Score: 2
Body:
  
    # It creates vendor folder and vendor.json inside it
    govendor init

    # Add dependencies in vendor.json
    govendor fetch <dependency>

    # Usage on new repository
    # fetch depenencies in vendor.json
    govendor sync


Example vendor.json

    {

    "comment": "",
    "ignore": "test",
    "package": [
        {
            "checksumSHA1": "kBeNcaKk56FguvPSUCEaH6AxpRc=",
            "path": "github.com/golang/protobuf/proto",
            "revision": "2bba0603135d7d7f5cb73b2125beeda19c09f4ef",
            "revisionTime": "2017-03-31T03:19:02Z"
        },
        {
            "checksumSHA1": "1DRAxdlWzS4U0xKN/yQ/fdNN7f0=",
            "path": "github.com/syndtr/goleveldb/leveldb/errors",
            "revision": "8c81ea47d4c41a385645e133e15510fc6a2a74b4",
            "revisionTime": "2017-04-09T01:48:31Z"
        }
    ],
    "rootPath": "github.com/sample"

    }

|======|
