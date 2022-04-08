package main

import (
	"encoding/json"
	"testing"
	"time"
)

func TestSwtich(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log("s", i)
		switch i {
		default:
			continue

		case 1:
			t.Log(i)
		}
		t.Log("nc")
	}
}

func TestTodayDuration(t *testing.T) {
	y, m, d := time.Now().Date()
	startTime := time.Date(y, m, d, 0, 0, 0, 0, time.Local).Unix()
	endTime1 := time.Date(y, m, d, 23, 59, 59, 0, time.Local).Unix()
	endTime := time.Date(y, m, d, 23, 59, 59, 0, time.Local).Add(time.Hour * 24).Unix()

	t.Log(startTime, endTime1, endTime)

	jsonStr := `{

		"banner_list": [
		],
		"platform_seckill": {
			"show": true,
			"title": "天天半价",
			"sub_title": "天天半价副标题PPE"
		},
		"good_seckill": {
			"show": true,
			"title": "好物秒杀",
			"sub_title": "好物秒杀副标题PPE"
		},
		"more_seckill": {
			"show": true,
			"title": "更多秒杀",
			"sub_title": "更多秒杀副标题PPE"
		},
		"activity": {
			"show": true,
			"name": "chilltest",
			"start_time": "2020-12-03 00:00:00",
			"end_time": "2021-01-30 00:00:00"
		},
		"super_seckill": {
			"show": true,
			"title": "超级秒杀直播间",
			"sub_title": "天天有惊喜PPE",
			"list": [
				{
					"show": true,
					"cover": "https://xxx",
					"topic": "诺亚美肤润唇膏",
					"user_id":75743141851,
					"start_time": "2021-02-04 15:12:00",
					"end_time": "2021-02-07 23:59:59"
				}
			]
		},
		"top_user": [
			{
				"user_id": 75743141851,
				"start_time": "2021-02-04 15:12:00",
				"end_time": "2021-02-07 23:59:59"
			},{
				"user_id": 60739805715,
				"start_time": "2021-01-03 00:00:00",
				"end_time": "2021-01-05 17:10:00"
			}
		]
	}`

	var res CMSSeckillChannelConfig
	json.Unmarshal([]byte(jsonStr), &res)
	t.Log(res)
}

type CMSSeckillChannelResponse struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

type SeckillChannelModelConfStruct struct {
	Show     bool   `json:"show"`
	Title    string `json:"title"`
	SubTitle string `json:"sub_title"`
}

type SeckillChannelActivityConfStruct struct {
	Show      bool   `json:"show"`
	Name      string `json:"name"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type SeckillBannerStruct struct {
	BannerImage string `json:"banner_image"`
	LinkUrl     string `json:"link_url"`
}
type SeckillChannelTopUser struct {
	UserId    int64  `json:"user_id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}
type SeckillChannelSuperSeckillConfStruct struct {
	Show     bool   `json:"show"`
	Title    string `json:"title"`
	SubTitle string `json:"sub_title"`
	List     []struct {
		Show      bool   `json:"show"`
		Cover     string `json:"cover"`
		Topic     string `json:"topic"`
		UserID    int64  `json:"user_id"`
		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`
	} `json:"list"`
}
type CMSSeckillChannelConfig struct {
	BannerList      []SeckillBannerStruct                `json:"banner_list"`
	PlatformSeckill SeckillChannelModelConfStruct        `json:"platform_seckill"`
	SuperSeckill    SeckillChannelSuperSeckillConfStruct `json:"super_seckill"`
	GoodSeckill     SeckillChannelModelConfStruct        `json:"good_seckill"`
	MoreSeckill     SeckillChannelModelConfStruct        `json:"more_seckill"`
	Activity        SeckillChannelActivityConfStruct     `json:"activity"`
	TopUser         []SeckillChannelTopUser              `json:"top_user"`
	RoomList        []int64                              // 强插的UserList对应的开播直播间
	SuperRoomList   []int64                              // SuperSeckill对应的开播直播间
}
