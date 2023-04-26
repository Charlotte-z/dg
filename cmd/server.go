package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"

	"google.golang.org/grpc"
)

var (
	dgraph = flag.String("d", "127.0.0.1:8088", "Dgraph Alpha address")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*dgraph, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	dg := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	resp, err := dg.NewTxn().Query(context.Background(), `{
  me(func: eq(name@en, "Steven Spielberg")) @filter(has(director.film)) {
    name@en
    director.film @filter(allofterms(name@en, "jones indiana"))  {
      name@en
    }
  }
}`)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Response: %s\n", resp.Json)
}
