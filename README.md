# theRemix/go90s

[![CircleCI](https://circleci.com/gh/theRemix/go90s.svg?style=svg)](https://circleci.com/gh/theRemix/go90s)

Go lib for Nineties Namer

Outputs a random name. Good for when you need to generate random nostalgic names.

`GetName(input)` will hash your input so you'll always get the same output... until the `names` list gets updated.

## 90s.fun Implementation

[https://90s.fun](https://90s.fun)


## Usage

```go
import "github.com/theRemix/go90s"


fmt.Println( go90s.GetRandomName() )
// powerpuff furby

fmt.Println( go90s.GetRandomName() )
// full crystal pepsi

fmt.Println( go90s.GetName("1.2.1") )
// dawsons hercules

fmt.Println( go90s.GetName("v1.2.0-rc8") )
// powerpuff vanilla coke

fmt.Println( go90s.GetName("v1.2.0-rc8") )
// powerpuff vanilla coke

```
## Tests

```sh
go test
```

**Note** the tests will fail for a while, until we get a larger name list, see [#Contributions](#Contributions)

## More Sample Outputs

```
mc hammer
house of animaniacs
wu tang buffy
saved by the pager
trapper slinky
rangers meets world
vanilla dolls
beastie martin
doug meets world
missy rugrats
are you afraid of the furby
cypress nintendo
```

## Contributions

Calling out all 90s kids! Accepting PRs for [lib/names.go](./lib/names.go)

**Adjectives** are the first set of words, they should make sense pairing with any **Names**

**Names** are next set of words, usually the last bit. Cartoon titles and famous memorable names are good here.

**Verbs** are rarely appended to a **Name** for extra flavor.



