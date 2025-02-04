package client

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/carlzhao/seata-golang/v2/pkg/apis"
	"github.com/carlzhao/seata-golang/v2/pkg/client/config"
	"github.com/carlzhao/seata-golang/v2/pkg/client/rm"
	"github.com/carlzhao/seata-golang/v2/pkg/client/tcc"
	"github.com/carlzhao/seata-golang/v2/pkg/client/tm"
)

// Init init resource manager，init transaction manager, expose a port to listen tc
// call back request.
func Init(config *config.Configuration) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	conn, err := grpc.DialContext(ctx, config.ServerAddressing,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithKeepaliveParams(config.GetClientParameters()))
	if err != nil {
		panic(errors.New(fmt.Sprintf("did not connect: %v", err)))
	}

	resourceManagerClient := apis.NewResourceManagerServiceClient(conn)
	transactionManagerClient := apis.NewTransactionManagerServiceClient(conn)

	rm.InitResourceManager(config.Addressing, resourceManagerClient)
	tm.InitTransactionManager(config.Addressing, transactionManagerClient)
	rm.RegisterTransactionServiceServer(tcc.GetTCCResourceManager())
}
