package internal

import (
	"context"
	"github.com/Xanonymous-GitHub/sxcctw/db/pkg/orm"
	"github.com/Xanonymous-GitHub/sxcctw/db/pkg/proto"
)

type RecordService struct {
	proto.UnimplementedRecordServiceServer
}

func (r *RecordService) CreateRecord(_ context.Context, req *proto.CreateRecordRequest) (*proto.CreateRecordResponse, error) {
	// TODO(TU): do sth filter or verifications here.
	return SaveRecord(orm.DB, req)
}

func (r *RecordService) GetOriginUrl(_ context.Context, req *proto.GetOriginUrlRequest) (*proto.GetOriginUrlResponse, error) {
	return LoadRecord(orm.DB, req)
}
