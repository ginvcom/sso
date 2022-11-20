package config

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func GetWhiteList(hosts []string) (err error) {

	client, err := clientv3.New(clientv3.Config{
		Endpoints:   hosts,
		DialTimeout: time.Duration(5) * time.Second,
	})

	if err != nil {
		err = errors.Errorf("connect etcd failed, err: %w", err)
		return
	}

	key := "/sso/api/whitelist"
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	resp, err := client.Get(ctx, key)
	cancel()
	if err != nil {
		err = errors.Errorf("get etcd whitelist failed, err: %v", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s, %s\n", ev.Key, ev.Value)
	}

	go func() {
		rch := client.Watch(context.Background(), key, clientv3.WithPrefix())
		for wresp := range rch {
			for _, ev := range wresp.Events {
				fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
		}
	}()
	return
}
