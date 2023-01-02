go-ipldtool
===========

The multipurpose tool for wrangling data.

**IPLD** is a data interchange standard, with emphasis on utility in the decentralized web.
The `ipld` tool is a command for working with any of the data we can wrangle by using the IPLD standards and conventions.

The aim is to have a playful, but useful, gadget.
It should help you explore and manage data,
and it should also help you understand how to use IPLD and how to create and grow decentralized protocols using the leverage that IPLD provides.

The API philosophy of the ipldtool is human-centric:
debug and diagnostic output formats are the defaults.
The tool is also designed to be friendly for "pipeline" usage,
but you'll have to tell it explicitly what codecs to use in that case.

(Generally, the hope is that the ipldtool should end up feeling a bit like the venerable "`jq`".
It's something you _can_ use in production scripts, but it's mostly for you, as a developer, to glue things together, and be able to build demos fast.)


Features
--------

- Read data in a variety of codecs and transform it into other codecs.  (E.g. JSON-to-CBOR, dagpb-to-dagjson, etc, etc!)

- Walk over data while processing it -- use [paths](https://ipld.io/docs/data-model/pathing/) to select specific sections of data.
	- ... or use [Selectors](https://ipld.io/specs/selectors/) to do even more detailed walks that can match multiple regions of data in complex conditions.

- Compute the [CID](https://ipld.io/glossary/#cid) of data, so you can refer to it with immutable [links](https://ipld.io/glossary/#link).

- Add data hunks to local storage using the `ipld put` command, which will make the data available for reference in larger data structures using [links](https://ipld.io/glossary/#link).

- For [IPLD](https://ipld.io/) data that contains [links](https://ipld.io/glossary/#link), pathing and selectors and other forms of data access can freely traverse links, automatically loading data from local storage as needed.

- [IPLD Schemas](https://ipld.io/docs/schemas/) can be compiled and processed with the `ipld schema` subcommands.
	- ... and they can be used by the `ipld read` and other commands as a lens for interpreting data, too.

... and more, coming soon!


Status
------

The ipld tool should currently be considered in an early alpha status.
It's under very active development.
Some features are working, but may not be completely; in general, there is currently no promise of API stability.
Some features planned features are also missing entirely (perhaps, waiting for you to contribute them?).

The best way to increase the stability and completeness of the ipldtool is to start using it, and if you can, contribute!


Comparisons
-----------

Please not that this tool has, strictly speaking, nothing to do with the IPFS APIs.  IPFS offers some commands which also work with IPLD data, but they do not necessarily use the same names, or follow the same rules, as the commands in this ipldtool.
Many IPFS APIs are also philosophically different in that they may try to do networking in order to satisfy your requests; this ipldtool is very explicitly designed *not* to ever initialize new network requests, and works only with local data.

Prerequisit for install
-----------
Before install the go-ipld, make sure your computer have Go installations, you can read the official guide on [Go website](https://go.dev/doc/install). 


Install
-----------

1. clone the repository

```bash
git clone https://github.com/ipld/go-ipldtool.git
```

2. navigate to the project install directory
```bash
cd go-ipldtool/app
```

3. use Go to install the tools
```bash
go install
```

4. check the ipld commandline installation location in your local path
```bash
go env|grep "GOPATH"
```
5. you can add the $GOPATH/bin to your shell bin by following scripts
```bash
PATH=$PATH:~/your/go/path/bin
```

6. if you sucessfully install with go, then you can type ipld in the command line environment
```bash
ipld
```



License
-------

SPDX-License-Identifier: Apache-2.0 OR MIT
