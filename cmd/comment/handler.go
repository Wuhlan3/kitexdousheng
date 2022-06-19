package main

import (
	"context"
	"kitexdousheng/cmd/comment/service"
	"kitexdousheng/kitex_gen/comment"
	"kitexdousheng/pkg/errno"
)

// CommentSrvImpl implements the last service interface defined in the IDL.
type CommentSrvImpl struct{}

// CommentAction implements the CommentSrvImpl interface.
func (s *CommentSrvImpl) CommentAction(ctx context.Context, req *comment.DouyinCommentActionRequest) (resp *comment.DouyinCommentActionResponse, err error) {
	// TODO: Your code here...
	err = service.CommentAction(req.UserId, req.VideoId, *req.CommentText)
	if err != nil {
		return &comment.DouyinCommentActionResponse{
			StatusCode: errno.ServiceErr.ErrCode,
			StatusMsg:  &errno.ServiceErr.ErrMsg,
		}, nil
	}
	return &comment.DouyinCommentActionResponse{
		StatusCode: errno.Success.ErrCode,
		StatusMsg:  &errno.Success.ErrMsg,
		Comment: &comment.Comment{
			Id:      0,
			User:    nil,
			Content: *req.CommentText,
		},
	}, nil
}

// CommentList implements the CommentSrvImpl interface.
func (s *CommentSrvImpl) CommentList(ctx context.Context, req *comment.DouyinCommentListRequest) (resp *comment.DouyinCommentListResponse, err error) {
	// TODO: Your code here...
	comments, err := service.CommentList(req.VideoId)
	if err != nil {
		return &comment.DouyinCommentListResponse{
			StatusCode:  errno.ServiceErr.ErrCode,
			StatusMsg:   &errno.ServiceErr.ErrMsg,
			CommentList: comments,
		}, nil
	}
	return &comment.DouyinCommentListResponse{
		StatusCode:  errno.Success.ErrCode,
		StatusMsg:   &errno.Success.ErrMsg,
		CommentList: comments,
	}, nil
}
