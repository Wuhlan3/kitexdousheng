package service

import (
	"errors"
	"fmt"
	"kitexdousheng/cmd/repository/db"
	"kitexdousheng/pkg/constants"
	"strconv"

	"github.com/go-redis/redis/v8"
)

func AddVIdListToRedis(video *db.Video) (err error) {
	//判断feed是否存在，如果不存在，要重新创建
	err = updateFeed()
	if err != nil {
		return err
	}

	//pipe.ZAdd(constants.CONTEXT, "feed", &redis.Z{Score: float64(video.CreatedAt.UnixMilli()) / 1000, Member: videoIDStr})
	videoIDStr := strconv.FormatInt(video.Id, 10)
	//添加新的视频序列号
	if err := constants.REDIS.ZAdd(constants.CONTEXT, "feed", &redis.Z{Score: float64(video.CreateTime.UnixMilli()) / 1000, Member: videoIDStr}).Err(); err != nil {
		return err
	}
	return nil
}

func updateFeed() error {
	n, err := constants.REDIS.Exists(constants.CONTEXT, "feed").Result()
	if err != nil {
		return err
	}
	// "feed"不存在
	if n <= 0 {
		videosList, err := db.NewVideoDaoInstance().QueryVideoList(100)
		fmt.Println(videosList)
		if err != nil {
			return err
		}
		if len(*videosList) == 0 {
			return errors.New("not found video")
		}

		var listZ = make([]*redis.Z, 0, len(*videosList))
		for _, video := range *videosList {
			listZ = append(listZ, &redis.Z{Score: float64(video.CreateTime.UnixMilli()) / 1000, Member: video.Id})
		}

		if err := constants.REDIS.ZAdd(constants.CONTEXT, "feed", listZ...).Err(); err != nil {
			return err
		}
	}
	return nil
}
