package biz

import (
	"context"

	pb "review-o/api/operation/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type OperationRepo interface {
	AppealReview(ctx context.Context, req *pb.AppealReviewRequest) (*pb.AppealReviewReply, error)
}

type OperationUsecase struct {
	repo OperationRepo
	log  *log.Helper
}

func NewOperationUsecase(repo OperationRepo, logger log.Logger) *OperationUsecase {
	return &OperationUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *OperationUsecase) AppealReview(ctx context.Context, req *pb.AppealReviewRequest) (*pb.AppealReviewReply, error) {
	uc.log.WithContext(ctx).Infof("AppealReview: %+v", req)
	res, err := uc.repo.AppealReview(ctx, req)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("AppealReview: %+v", err)
		return nil, err
	}
	uc.log.WithContext(ctx).Infof("AppealReview: %+v", res)
	return res, nil
}
