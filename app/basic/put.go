package basic

import (
	"context"
	"encoding/base32"

	flatfs "github.com/ipfs/go-ds-flatfs"
	"github.com/urfave/cli/v2"

	"github.com/ipld/go-ipld-prime/storage/dsadapter"
)

var Cmd_Put = &cli.Command{
	Name:     "put",
	Category: "Basic",
	Usage:    "Put a single block of data into storage.",
	Action:   Action_Put,
}

func Action_Put(args *cli.Context) error {
	shardFn, err := flatfs.ParseShardFunc("/repo/flatfs/shard/v1/next-to-last/3")
	if err != nil {
		return err
	}
	ds, err := flatfs.CreateOrOpen("/tmp/foobar", shardFn, false)
	if err != nil {
		return err
	}
	defer ds.Close()

	store := dsadapter.Adapter{
		Wrapped: ds,
		EscapingFunc: func(raw string) string {
			return base32.StdEncoding.EncodeToString([]byte(raw))
		},
	}

	if err := store.Put(context.Background(), "some key", []byte("zonk")); err != nil {
		return err
	}

	if err := store.Put(context.Background(), "another key", []byte("hork")); err != nil {
		return err
	}

	return nil
}

// Storage configuration: still in debate, but I think I'd like storage config to look roughly like this:
/*

type StorageConfig [ModedStorageSpec]

type ModedStorageSpec union {
	| "rw:" StorageModeReadWrite
	| "ro:" StorageModeReadOnly
	| "wb:" StorageModeWriteBack
} representation stringprefix

type StorageModeReadWrite StorageSpec
type StorageModeReadOnly StorageSpec
type StorageModeWriteBack StorageSpec

type StorageSpec struct {
	engine string
	param string
} representation stringjoin (":")

*/
// And so when taking a real system we know and looking at how that would be stated as a CLI flag:
// `--storage=rw:flatfs:/path:/repo/flatfs/shard/v1/next-to-last/3`
// (I can't say I care for half (or even two thirds) of the redundancy in that last hunk, but that's what flatfs does right now.  (Maybe we should tersen it up in our user-facing porcelain.))
// That flag would also be allowed repeatedly.

// Is this sufficiently ghastly that we should just force you to have a config file, and maybe make subcommands to help you set it up?
//  Well, it's dang close.
//  But I still want to be able to choose the modes per-command.  And that's easiest if it's controlled by flags or env vars.

// Perhaps a mid-way approach:
// You can have *no* config options in the CLI -- and make those a stateful part of the storage.
// The only parameter left is the storage root path.
// There are then also commands that can pre-establish a storage and configure it with options.  But this is optional.
// And that's it.
// This seems probably good.
