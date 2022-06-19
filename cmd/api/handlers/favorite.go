package handlers

import (
	"github.com/gin-gonic/gin"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	//token := c.Query("token")

}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	uid, _ := c.Get("uid")
	userId := uid.(int64)
	_ = userId
	/*
		videos, err := service.FavouriteList(userId)
		if err != nil {
			c.JSON(http.StatusOK, proto.DouyinPublishListResponse{
				StatusCode: 1,
				StatusMsg:  "Video loads Failed",
			})
		}
		c.JSON(http.StatusOK, proto.DouyinPublishListResponse{
			StatusCode: 0,
			StatusMsg:  "publishList successfully",
			VideoList:  videos,
		})

	*/
}
