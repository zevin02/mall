package model

// 秒杀的商品
type SkillGoods struct {
	Id         uint `gorm:"primarykey"` //在秒杀中的秒杀id
	ProductId  uint `gorm:"not null"`   //秒杀活动中的商品原id
	BossId     uint `gorm:"not null"`
	Title      string
	Money      float64
	Num        int `gorm:"not null"`
	CustomId   uint
	CustomName string
}

// 当前这个是需要在mq中的消费者中进行消费的
type SkillGood2MQ struct {
	SkillGoodId uint    `json:"skill_good_id"`
	ProductId   uint    `json:"product_id"`
	BossId      uint    `json:"boss_id"`
	UserId      uint    `json:"user_id"`
	Money       float64 `json:"money"`
	AddressId   uint    `json:"address_id"`
	Key         string  `json:"key"`
}
