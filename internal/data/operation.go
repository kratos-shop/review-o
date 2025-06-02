package data

import (
	"context"

	pb "review-o/api/operation/v1"
	reviewPb "review-o/api/review/v1"

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
	r.log.Infof("AppealReview req %+v", req)
	reviewReq := &reviewPb.OperationAppealReviewRequest{
		AppealId:     req.AppealId,
		ReviewId:     req.ReviewId,
		StoreId:      req.StoreId,
		Status:       req.Status,
		AppealReason: req.AppealReason,
	}
	reply, err := r.data.grpcClient.OperationAppealReview(ctx, reviewReq)
	if err != nil {
		r.log.Errorf("OperationAppealReview err %+v", err)
		return nil, err
	}
	r.log.Infof("OperationAppealReview reply %+v", reply)
	return &pb.AppealReviewReply{AppealId: reply.Id}, nil
}
