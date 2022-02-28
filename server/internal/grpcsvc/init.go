package grpcsvc

import (
	"github.com/Xanonymous-GitHub/sxcctw/db/pkg/proto/pb"
	"github.com/Xanonymous-GitHub/sxcctw/server/pkg/vp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"strconv"
)

func init() {
	dbSvcAddr := ":" + strconv.Itoa(vp.Cvp.GetInt("dbServerPort"))

	// TODO(TU): remove insecure credentials when in production mode.
	devCredentials := insecure.NewCredentials()

	dial, err := grpc.Dial(dbSvcAddr, grpc.WithTransportCredentials(devCredentials))

	if err != nil {
		log.Fatalf("can not init grpc client!\n%v", err)
	}

	// Initialize service clients
	RecordSvcClient = pb.NewRecordServiceClient(dial)
}
