
## How to install

```
go get github.com/naotookuda/go-template-cli
```

## Usage

Create `hello.template` file like this :
```
Hello {{.firstName}} {{.lastName}} !
```

and `person.yaml` YAML file like this :

```
firstName: Naoto
lastName: Okuda
```

You can cmpile the template like this :
```
gt hello.template < person.yaml
```

