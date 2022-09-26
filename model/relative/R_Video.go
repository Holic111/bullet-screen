package relative

import "time"

type Info_Video struct {
	Title string `json:"title"`
	Cover string `json:"cover"` // 封面图
	Desc string `json:"desc"` // 视频简介
	Original bool `json:"original"` // 是否原创
	Partition int `json:"视频分区"`
}

type Update_Video struct {
	Vid int `json:"vid"` // 视频ID
	Title string `json:"title"`
	Cover string `json:"cover"` // 封面图
	Desc string `json:"desc"` // 视频简介
	Original bool `json:"original"` // 是否原创
}

type Paratition_Video struct {
	ID int `json:"vid"`
	Title string `json:"title"`
	Cover string `json:"cover"`
	Created_at time.Time `json:"created_at"`
}