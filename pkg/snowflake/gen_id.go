package snowflake

import (
	"fmt"
	"github.com/sony/sonyflake"
	"time"
)

var (
	// 实例
	sonyFlake *sonyflake.Sonyflake
	// 机器 ID
	sonyMachineID uint16
)

// 返回全局定义的机器 ID
func getMachineID() (uint16, error) {
	return sonyMachineID, nil
}

// Init 需传入当前的机器ID
func Init(machineId uint16) (err error) {
	sonyMachineID = machineId
	// 初始化一个起始时间
	t, _ := time.Parse("2006-01-02", "2022-02-09")
	// 生成全局配置
	settings := sonyflake.Settings{
		StartTime: t,
		MachineID: getMachineID,
	}
	// 用配置生成 sonyflake 节点
	sonyFlake = sonyflake.NewSonyflake(settings)
	return
}

// GetID 返回生成的id值
func GetID() (id uint64, err error) {
	if sonyFlake == nil {
		err = fmt.Errorf("snoy flake not inited")
		return
	}
	// 拿到 sonyFlake 节点生成 id 值
	id, err = sonyFlake.NextID()
	return
}
