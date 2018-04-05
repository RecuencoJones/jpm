# JPM

I want to learn Go.

## Requirements

- (npm >= 5 && node >= 8) || nvm
- go

## Setup

### Go CLI

We need to link `jpm-cli` to `GOPATH` so it can be installed and used globally

```ln -s /path/to/jpm/jpm-cli `echo $GOPATH`/src/github.com/RecuencoJones/jpm```

Then, we can cd into the linked path and install

```
cd $GOPATH/src/github.com/RecuencoJones/jpm
go install
```


## CLI Usage

```
jpm install
jpm add <dependency>
jpm publish
```
