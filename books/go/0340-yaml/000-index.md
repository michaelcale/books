---
Title: YAML
Id: 193
SOId: 2503
---

[YAML](http://yaml.org/) is a popular format for serializing data in a human friendly format. Think JSON but easier to read.

Thanks to its expressivness and readability, YAML is popular as a format for configuration files.

It's also used in more complex scenarios like driving [Ansible](https://www.ansible.com/) server automation.

There is no package in standard library for hanlding YAML format but there are community libraries including [gopkg.in/yaml.v2](http://gopkg.in/yaml.v2).

## Reading YAML file into a Go struct

@file data.yml noplayground

@file yaml_deserialize.go output noplayground

YAML decoding is very similar to [JSON decoding](182).

If you know the structure of YAML file, you can define structs that mirror this structure and pass a pointer to a struct describing top-level structure to `yaml.Decoder.Decode()` function (or `yaml.Unmarshal()` if decoding from `[]byte` slice).

YAML decoder does intelligent mapping between struct field names and names in YAML file so that e.g. `name` value in YAML is decoded into field `Name` in a struct.

It's best to create explicit mappings using `yaml` struct tags. I only omitted them from the example to illustrate the behavior when they are not specified.

## Writing Go struct to YAML file

@file yaml_serialize.go output sha1:c69c304984394d3765680a32c6573bc4e1faaeed goplayground:dGHkgP4rNQ4

`yaml.Marshal` takes interface{} as an argument. You can pass any Go value, it’ll be wrapped into interface{} with their type.

Marshaller will use reflection to inspect passed value and encode it as YAML strings.

When serializing structs, only exported fields (whose names start with capital letter) are serialized / deserialized.

In our example, fullName is not serialized.

Structs are serialized as YAML dictionaries. By default dictionary keys are the same as struct field names.

Struct field Name is serialized under dictionary key Name.

We can provide custom mappings with struct tags.

We can attach arbitrary struct tags string to struct fields.

`yaml:"age"` instructs YAML encoder / decoder to use name `age` for dictionary key representing field `Age`.

When serializing structs, passing the value and a pointer to it generates the same result.

Passing a pointer is more efficient becase passing by value creates unnecessary copy.
