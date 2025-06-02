package data

import (
	"context"

	pb "review-o/api/operation/v1"

	"review-o/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type operationRepo struct {
	data *Data
	log  *log.Helper
}

func NewOperationRepo(data *Data, logger log.Logger) biz.OperationRepo {
	return &operationRepo{data: data, log: log.NewHelper(logger)}
}

func (r *operationRepo) AppealReview(ctx context.Context, req *pb.AppealReviewRequest) (*pb.AppealReviewReply, error) {
	// r.data.grpcClient.AppealReview(ctx, req)
	return nil, nil
}
