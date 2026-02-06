package common

import (
	"fmt"
	"go-captcha/utils"
	"strconv"

	"github.com/rs/zerolog/log"
	"github.com/yitter/idgenerator-go/idgen"
)

func InitSnowflareId() {
	// init idgen
	workerId, _ := strconv.ParseUint(utils.Getenv("WORKER_ID", "0"), 10, 16)
	var options = idgen.NewIdGeneratorOptions(uint16(workerId))
	options.WorkerIdBitLength = 10 // 默认值6，限定 WorkerId 最大值为2^6-1，即默认最多支持64个节点。1024
	options.SeqBitLength = 10      // 默认值6，限制每毫秒生成的ID个数。若生成速度超过5万个/秒，建议加大 SeqBitLength 到 10。
	idgen.SetIdGenerator(options)
	log.Info().Msg(fmt.Sprintf("Init idgen success %d", idgen.NextId()))
}

func GenIdLong() int64 {
	return idgen.NextId()
}

func GenId() string {
	return strconv.FormatInt(idgen.NextId(), 10)
}
