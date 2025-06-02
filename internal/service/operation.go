package service

import (
	"context"

	pb "review-o/api/operation/v1"
	"review-o/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type OperationService struct {
	pb.UnimplementedOperationServer

	uc  *biz.OperationUsecase
	log *log.Helper
}

func NewOperationService(uc *biz.OperationUsecase, logger log.Logger) *OperationService {
	return &OperationService{uc: uc, log: log.NewHelper(logger)}
}

func (s *OperationService) AppealReview(ctx context.Context, req *pb.AppealReviewRequest) (*pb.AppealReviewReply, error) {
	s.log.WithContext(ctx).Infof("AppealReview: %+v", req)
	res, err := s.uc.AppealReview(ctx, req)
	if err != nil {
		s.log.WithContext(ctx).Errorf("AppealReview: %+v", err)
		return nil, err
	}
	s.log.WithContext(ctx).Infof("AppealReview: %+v", res)
	return res, nil
}
