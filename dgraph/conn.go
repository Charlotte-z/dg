package dgraph

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
	dgraph = flag.String("d", "127.0.0.1:8089", "Dgraph Alpha address")
)

func connectDgraphClient() {
	flag.Parse()
	conn, err := grpc.Dial(*dgraph, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	dg := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	resp, err := dg.NewTxn().Query(context.Background(), `{
  me(func: allofterms(name@en, "jones indiana")) {
    name@en
    genre {
      name@en
    }
  }
}`)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Response: %s\n", resp.Json)
}
