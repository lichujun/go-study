package demo

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

const (
	timestampBits    = 41
	dataCenterIdBits = 5
	machineIdBits    = 5
	sequenceBits     = 12

	maxTimestamp    = -1 ^ (-1 << timestampBits)
	maxDataCenterId = -1 ^ (-1 << dataCenterIdBits)
	maxMachineId    = -1 ^ (-1 << machineIdBits)
	maxSequence     = -1 ^ (-1 << sequenceBits)

	timestampBitLeft    = dataCenterIdBits + machineIdBits + sequenceBits
	dataCenterIdBitLeft = machineIdBits + sequenceBits
	machineIdBitLeft    = sequenceBits

	startTimestamp = 1612772988000
)

/**
41位时间戳-5位数据中心id-5位机器id-12位sequence
*/
type SnowWorker struct {
	mu           sync.Mutex
	dataCenterId int64
	machineId    int64
	timestamp    int64
	sequence     int64
}

func NewSnowWorker(dataCenterId, machineId int) (*SnowWorker, error) {
	if dataCenterId > maxDataCenterId {
		return nil, errors.New("dataCenterId invalid")
	}
	if machineId > maxMachineId {
		return nil, errors.New("machineId invalid")
	}
	return &SnowWorker{
		dataCenterId: int64(dataCenterId),
		machineId:    int64(machineId),
		timestamp:    -1,
		sequence:     0,
	}, nil
}

func (worker *SnowWorker) NextId() (int64, error) {
	// 上锁，防止并发问题
	worker.mu.Lock()
	defer worker.mu.Unlock()
	currentTimeStamp := getCurrentTimestamp()
	// 1、超过最大时间
	// 2、防止时间回调的问题
	if worker.timestamp > maxTimestamp || worker.timestamp > currentTimeStamp {
		return 0, errors.New("can't grater than currentTimeStamp")
	}
	if worker.timestamp == currentTimeStamp {
		worker.sequence = (worker.sequence + 1) & maxSequence
		if worker.sequence == 0 {
			// 等待下一毫秒
			for worker.timestamp >= currentTimeStamp {
				currentTimeStamp = getCurrentTimestamp()
			}
		}
	} else {
		worker.sequence = 0
	}
	worker.timestamp = currentTimeStamp
	return ((worker.timestamp - startTimestamp) << timestampBitLeft) | (worker.dataCenterId << dataCenterIdBitLeft) |
		(worker.machineId << machineIdBitLeft) | worker.sequence, nil
}

func getCurrentTimestamp() int64 {
	return time.Now().UnixNano() / 1e6
}

func TestSnowWorker(t *testing.T) {
	worker, err := NewSnowWorker(1, 2)
	if err != nil {
		log.Fatal(err)
	}

	i := 0
	r := make([]int64, 0)
	for {
		id, _ := worker.NextId()
		r = append(r, id)
		i++
		if i > 10000000 {
			break
		}
	}
	j := 0
	for _, v := range r {
		if v > 1 {
		}
		j++
	}
	fmt.Println(j)
	fmt.Println(len(unique(r)))
	fmt.Println(worker)

}
func unique(m []int64) []int64 {
	s := make([]int64, 0)
	smap := make(map[int64]int64)
	for _, value := range m {
		//计算map长度
		length := len(smap)
		smap[value] = 1
		//比较map长度, 如果map长度不相等， 说明key不存在
		if len(smap) != length {
			s = append(s, value)
		}
	}

	return s

}
