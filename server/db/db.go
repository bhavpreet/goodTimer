package db

import (
	"context"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

type DB interface {
	Initialize(context.Context) (func() error, error)
	GetCollection(ctx context.Context, name string) (Collection, error)
	ListCollections(ctx context.Context) ([]string, error)
	Plug(context.Context, func(collection string, key []byte, value []byte)) error
}

type Collection interface {
	Write(ctx context.Context, key []byte, value []byte) error
	Read(ctx context.Context, key []byte) ([]byte, error)
	DeleteCollection(ctx context.Context) error
	List(ctx context.Context) ([][]byte, error)
}

type boltDB struct {
	*bolt.DB
	plugs []func(collection string, key, value []byte)
}

type collection struct {
	b    *boltDB
	name []byte
}

func NewDefaultDB() DB {
	return new(boltDB)
}

func (b *boltDB) Initialize(ctx context.Context) (func() error, error) {
	var err error
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	b.DB, err = bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Printf("Error occurred while opening db: %v", err)
	}

	return b.Close, err
}

func (b *boltDB) GetCollection(ctx context.Context, name string) (Collection, error) {
	c := new(collection)
	c.b = b
	var err error
	_, err = b.getBucket(ctx, name)
	if err != nil {
		return nil, err
	}
	c.name = []byte(name)
	return c, nil
}

func (b *boltDB) ListCollections(ctx context.Context) ([]string, error) {
	c := []string{}
	err := b.View(func(tx *bolt.Tx) error {
		return tx.ForEach(func(name []byte, _ *bolt.Bucket) error {
			c = append(c, string(name))
			return nil
		})
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return c, nil
}

func (b *boltDB) getBucket(ctx context.Context, name string) (*bolt.Bucket, error) {
	var bkt *bolt.Bucket
	// Create default bucket if not exists
	err := b.Update(func(tx *bolt.Tx) error {
		var err error
		if bkt = tx.Bucket([]byte(name)); bkt == nil {
			bkt, err = tx.CreateBucket([]byte(name))
			if err != nil {
				return fmt.Errorf("create bucket: %s", err)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return bkt, nil
}

func (c *collection) Write(ctx context.Context, key, value []byte) error {
	var err error
	defer func() {
		if err == nil {
			c.b.callbacks(ctx, string(c.name), key, value)
		}
	}()

	err = c.b.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket(c.name).Put(key, value)
		return err
	})
	return nil
}

func (c *collection) Read(ctx context.Context, key []byte) ([]byte, error) {
	// var err error
	// defer func() {
	// 	if err == nil {
	// 		b.callbacks(ctx, key, value)
	// 	}
	// }()

	var value []byte

	if err := c.b.View(func(tx *bolt.Tx) error {
		value = tx.Bucket(c.name).Get(key)
		return nil
	}); err != nil {
		return nil, err
	}

	return value, nil
}

func (c *collection) DeleteCollection(ctx context.Context) error {
	err := c.b.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket(c.name)
	})
	return err
}

func (c *collection) List(ctx context.Context) ([][]byte, error) {
	ret := [][]byte{}
	err := c.b.View(func(tx *bolt.Tx) error {
		tx.Bucket(c.name).ForEach(func(k, v []byte) error {
			ret = append(ret, v)
			return nil
		})
		return nil
	})
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (b *boltDB) Plug(ctx context.Context, fn func(collection string, key, value []byte)) error {
	b.plugs = append(b.plugs, fn)
	return nil
}

func (b *boltDB) callbacks(ctx context.Context, collection string, key, value []byte) {
	for _, fn := range b.plugs {
		fn(collection, key, value)
	}
}
