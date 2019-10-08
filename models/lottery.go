package models

import (
	"fmt"
	"lottery/database"
	"time"
)

type Lottery struct {
	ID				uint `gorm:"primary_key"`
	CreatedAt		time.Time
	Name			string `gorm:"not null VARCHAR(32)"`
	LotteryType		uint
	Schedule		uint64
	OpenDate		string `gorm:"not null VARCHAR(32)"`
	Number			string `gorm:"not null VARCHAR(64)"`
}

type LotteryJson struct {
	Name			string `json:"name" validate:"required"`
	LotteryType		uint   `json:"type" validate:"required"`
	Schedule		uint64 `json:"schedule" validate:"required"`
	OpenDate		string `json:"date" validate:"required"`
	Number			string `json:"number" validate:"required"`
}

//通过彩票种类期号寻找开奖号码
func GetLotteryBySchedule(lotterytype uint,schedule uint64) *Lottery {
	lottery := new(Lottery)
	lottery.LotteryType = lotterytype
	lottery.Schedule = schedule

	if err := database.DB.First(lottery).Error; err != nil {
		fmt.Printf("GetLotteryBySchedule:%s", err)
	}

	return lottery
}

//获取所有开奖记录
func GetAllLotterys(name, orderBy string, offset, limit int) (lottery []*Lottery) {
	if err := database.GetAll(name, orderBy, offset, limit).Find(&lottery).Error; err != nil {
		fmt.Printf("GetAllLotterys:%s", err)
	}
	return
}

//根据彩票种类获取所有期号
func GetScheduleByType(lotterytype uint) (lottery []* Lottery){
	if err := database.DB.Order("schedule desc").Where("lotterytype = ?", lotterytype).Limit(10).Find(&lottery).Error; err != nil {
		fmt.Printf("GetScheduleByType:%s", err)
	}
	return
}

//插入开奖记录
func CreateLottery(aul *LotteryJson) *Lottery {
	lottery := new(Lottery)
	lottery.Name = aul.Name
	lottery.LotteryType = aul.LotteryType
	lottery.Schedule = aul.Schedule
	lottery.OpenDate = aul.OpenDate
	lottery.Number = aul.Number

	if err := database.DB.Create(lottery).Error; err != nil {
		fmt.Printf("CreateLottery:%s", err)
	}

	return lottery
}
