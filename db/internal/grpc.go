package internal

import (
	"context"
	"github.com/Xanonymous-GitHub/sxcctw/db/pkg/orm"
	"github.com/Xanonymous-GitHub/sxcctw/db/pkg/proto/pb"
)

type RecordService struct {
	pb.UnimplementedRecordServiceServer
}

func (r *RecordService) CreateRecord(_ context.Context, req *pb.CreateRecordRequest) (*pb.CreateRecordResponse, error) {
	// TODO(TU): do sth filter or verifications here.
	return SaveRecord(orm.DB, req)
}

func (r *RecordService) GetOriginUrl(_ context.Context, req *pb.GetOriginUrlRequest) (*pb.GetOriginUrlResponse, error) {
	return LoadRecord(orm.DB, req)
}
