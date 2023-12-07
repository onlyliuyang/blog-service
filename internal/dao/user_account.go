package dao

import (
	"fmt"
	_const "github.com/blog-service/const"
	"github.com/blog-service/global"
	"github.com/blog-service/internal/model"
	"github.com/blog-service/pkg/convert"
	"github.com/blog-service/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/goccy/go-json"
	"strconv"
	"time"
)

const USER_ACCOUNT_KEY = "user_account_%d"
const USER_ACCOUNT_ZSET = "user_account_zset"

func (d *Dao) RegisterUser(ctx *gin.Context, userAccount *model.UserAccount) error {
	var err error
	err = userAccount.Register(ctx, d.engine)
	if err != nil {
		return err
	}

	cacheKey := fmt.Sprintf(USER_ACCOUNT_KEY, userAccount.Uid)
	cacheFields := make(map[string]interface{})
	bytes, _ := json.Marshal(userAccount)
	err = json.Unmarshal(bytes, &cacheFields)
	if err != nil {
		return err
	}

	pipe := global.RedisDB.TxPipeline()
	pipe.HMSet(cacheKey, cacheFields)
	pipe.Expire(cacheKey, time.Second*3600)
	if _, err = pipe.Exec(); err != nil {
		return err
	}
	return nil
}

func (d *Dao) GetUserInfo(ctx *gin.Context, userId int) (userInfo *model.UserAccount, err error) {
	cacheKey := fmt.Sprintf(USER_ACCOUNT_KEY, userId)
	userInfoFromCache, err := global.RedisDB.HGetAll(cacheKey).Result()
	if err != nil {
		return nil, err
	}

	//从缓存中获取数据
	if len(userInfoFromCache) > 0 {
		mapData := make(map[string]interface{})
		for key, val := range userInfoFromCache {
			if key == "created_at" || key == "updated_at" {
				parseTime, _ := time.Parse(time.RFC3339, val)
				mapData[key] = parseTime.Format(_const.DATE_LAYOUT)
			} else {
				mapData[key] = val
			}
		}
		userInfo = &model.UserAccount{}
		err = convert.MapToStruct(mapData, userInfo)
		if err != nil {
			return nil, err
		}
		return userInfo, nil
	}

	//如果缓存没有，则从DB中获取数据
	var userAccount model.UserAccount
	userInfoFromDB, err := userAccount.GetUserInfo(ctx, d.engine, userId)

	if err != nil {
		return nil, err
	}

	//回写缓存
	cacheFields := make(map[string]interface{})
	bytes, _ := json.Marshal(userInfoFromDB)
	err = json.Unmarshal(bytes, &cacheFields)
	if err != nil {
		return userInfo, err
	}
	global.RedisDB.HMSet(cacheKey, cacheFields)

	return &userInfoFromDB, nil
}

func (d *Dao) ImportUserList(ctx *gin.Context) {
	var userAccount model.UserAccount
	offset, limit := 0, 100
	for {
		userList, err := userAccount.GetListByPage(ctx, d.engine, limit, offset)
		if err != nil {
			global.Logger.WithFields(logger.Fields{
				"offset": offset,
				"limit":  limit,
			}).Errorof(ctx, "获取用户列表失败:"+err.Error())
			break
		}

		if len(userList) <= 0 {
			global.Logger.Infof(ctx, "用户列表数据完成")
			return
		}

		memberList := make([]redis.Z, len(userList))
		for idx, user := range userList {
			memberList[idx] = redis.Z{
				Score:  float64(user.CreatedAt.Unix()),
				Member: user.Uid,
			}
		}
		_, err = global.RedisDB.ZAdd(USER_ACCOUNT_ZSET, memberList...).Result()
		if err != nil {
			global.Logger.WithFields(logger.Fields{
				"offset": offset,
				"limit":  limit,
			}).Errorof(ctx, "用户列表写入失败:"+err.Error())
			break
		}
		offset += limit
	}
}

func (d *Dao) GetUserList(ctx *gin.Context, cursor string, limit int64) (userList []*model.UserAccount, nextCursor string, err error) {
	var maxCursor = cursor
	if cursor == "" {
		maxCursor = "+inf"
	}

	data, err := global.RedisDB.WithContext(ctx).ZRevRangeByScoreWithScores(USER_ACCOUNT_ZSET, redis.ZRangeBy{
		Min:    "-inf",
		Max:    maxCursor,
		Offset: 1,
		Count:  limit,
	}).Result()

	if err != nil {
		global.Logger.WithFields(logger.Fields{
			"cursor": cursor,
			"limit":  limit,
		}).Errorof(ctx, "获取用户列表失败: ", err.Error())
		return
	}

	for idx, value := range data {
		userId := value.Member
		score := value.Score
		userInfo, _ := d.GetUserInfo(ctx, convert.StrTo(userId.(string)).MustInt())
		userList = append(userList, userInfo)
		if idx == len(data)-1 {
			nextCursor = strconv.FormatFloat(score, 'f', 0, 64)
		}
	}
	return
}
