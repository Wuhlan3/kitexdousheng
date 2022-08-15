// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package constants

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	REDIS   *redis.Client          // Redis 缓存接口
	CONTEXT = context.Background() // 上下文信息
)

const (
	ApiServiceName              = "api"
	FavoriteServiceName         = "favorite"
	CommentServiceName          = "comment"
	PublishServiceName          = "Publish"
	FeedServiceName             = "feed"
	UserServiceName             = "user"
	RelationServiceName         = "relation"
	MySQLDefaultDSN             = "root:123456@tcp(localhost:3306)/kitexdousheng?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress                 = "127.0.0.1:2379"
	CPURateLimit        float64 = 80.0
)

// 过期时间
var (
	FAVORITE_EXPIRE       = 10 * time.Minute
	VIDEO_COMMENTS_EXPIRE = 10 * time.Minute
	COMMENT_EXPIRE        = 10 * time.Minute
	FOLLOW_EXPIRE         = 10 * time.Minute
	USER_INFO_EXPIRE      = 10 * time.Minute
	VIDEO_EXPIRE          = 10 * time.Minute
	PUBLISH_EXPIRE        = 10 * time.Minute
	EMPTY_EXPIRE          = 10 * time.Minute
	EXPIRE_TIME_JITTER    = 10 * time.Minute
)
