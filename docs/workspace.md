`workspace` subcommands
=======================

The `ipld workspace` subcommands are for creating, configuring, and asking questions about
the local "workspace".

A "workspace" is the on-disk area where data is stored.

Not all `ipld` commands require a workspace,
but most of the stateful ones do,
as do any commands which read data that spans multiple blocks.


Docs
----

[testmark]:# (docs/script)
```
ipld workspace --help
```

[testmark]:# (docs/output)
```text
NAME:
   ipld workspace - Create, configure, or interogate a workspace for the ipldtool.  (You'll need a workspace for any of the stateful commands.)

USAGE:
   ipld workspace command [command options] [arguments...]

COMMANDS:
   new      Creates the local filesystem markers for a new workspace.
   find     Tells you what the current workspace is.
   help, h  Shows a list of commands or help for one command

OPTIONS:
   --help, -h  show help (default: false)
   
```

Examples
--------

[testmark]:# (new-workspace/script)
```
ipld workspace new
find .
```

[testmark]:# (new-workspace/output)
```text
.
./.ipld
```
