`ipldtool` ROADMAP
==================

Design Focus
------------

The ipldtool is aimed at providing a playful and useful tool for interacting with IPLD data.

The following are some heuristics and boundary conditions about what to expect from the ipldtool now,
and what to expect from it (or not expect from it) as it grows.

### data model first

The ipldtool's command should make sense if you understand the [IPLD Data Model](https://ipld.io/docs/data-model/);
and they should lead you to understand the IPLD Data Model if you don't yet.

That means pathing should work in the usual ways of the data model, etc.

Nothing should be codec-specific.

### human friendly

The ipldtool should be human-friendly by default.

For example, that means commands like `ipld read` will default to emitting a human-readable debug representation of data,
which is rich in data and clearly annotates the [Data Model kinds](https://ipld.io/docs/data-model/kinds/) of the data.

Using the tool in ways that are friendly to scripting composition may require more flags,
or different subcommands (we aim to have multiple levels of possible engagement, much like git's concept of "porcelain" vs "plumbing" commands).

### statefulness

Most of the commands of the ipldtool are meant to be stateless -- working on data that you give to each command.

Some commands can't avoid being a little bit stateful -- when making large graphs of data, it necessarily takes more than one operation,
because we have to calculate the CID of each object, and then use that CID in other objects, resulting the connected graph of data.
Similarly, commands that walk across graphs of data for you need to load data during the walk in order to be useful,
so they're consulting data accumulated by earlier operations.

We try to make it possible to use stateless forms of commands as much as possible.
So, even when a command has a mode which works by accepting a CID as the argument for what data to work on (which implies we're loading data state stored previously),
it should *also* accept a file name or other stream identifier (e.g. many commands take "`-`" to mean "read from stdin").

#### non-singleton

For those parts of the ipldtool's behavior that _are_ stateful,
it should be extremely easy to run more than one instance of the ipldtool with entirely separated state and storage pools.

The ipldtool should not be a daemon by default.

### files

IPLD isn't (just) about files.

But it's useful to have tools for mapping filesystems into IPLD.

The ipldtool will probably grow to contain some commands for packing filesystems into IPLD-based data structures,
and unpacking those specific recognizable kinds of structures back onto regular filesystems.

But: these should be presented and understood primarily as add-ons and task-specific focuses.
IPLD is much more general than just files, and the commands we showcase first should be the ones that work on the data model.
By the time someone encounters the commands for working with files, ideally, they should already have been heavily exposed to commands that work at the data model level.
This is important for didactic purposes.

### networking

The ipldtool intentionally eschews networking.

Transport can be accomplished by commands that read and write data to and from the local storage.
Commands are available that do this both at the individual block level,
or in bulked operations, which usually then use the CAR format.

You can write other tools which do networking, and have those processes
wrap the ipldtool, shuttling data in and out of the ipldtool as necessary.

The intent of this is that when working with the ipldtool, operations are ***always*** local.

(Correspondingly, operations with the ipldtool should ***always*** have a predictable latency.
No command of the ipldtool should ever, ever have a concept of "timeout" which might be triggered by network latency.)

### CLI and HTTP API isomorphisms

The ipldtool supports many operations over an HTTP API as well as the CLI.

Features available over the HTTP API map as directly to the CLI as possible.
For example:

`$ ipld read bafyfrak path/in/data --output=codec:0x20`

should be the same as:

`GET bafyfrak.ipldtool.local/path/in/data?output=codec:0x20`

(This is not yet standardized, but the idea is to make it as mechanical as possible,
so that things are consistent and unsurprising: reading the CLI docs should
also give you the information needed to craft equivalent HTTP API interactions.)

Not all features available in the CLI are always available in the HTTP API.
For example, the HTTP API is read-only by default (unless writable endpoints are enabled by additional flags).



Feature List
------------

(WIP)
