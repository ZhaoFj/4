package rocketmq

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"go.uber.org/zap"
)

func main() {
	mqAddr := fmt.Sprintf("%s:%d", "ddhanta.cn", 9876)
	p, err := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{mqAddr})),
		producer.WithRetry(2),
	)
	if err != nil {
		panic(err)
	}
	err = p.Start()
	if err != nil {
		zap.S().Error("生产者错误：" + err.Error())
		os.Exit(1)
	}
	topic := "hantamall"
	for i := 0; i < 10; i++ {
		msg := &primitive.Message{
			Topic: topic,
			Body:  []byte("hello hanta mall" + strconv.Itoa(i)),
		}
		res, err := p.SendSync(context.Background(), msg)
		if err != nil {
			zap.S().Error("发送消息错误：" + err.Error())
			os.Exit(1)
		} else {
			zap.S().Error("发送消息成功" + res.String() + "-" + res.MsgID)
		}
	}
	err = p.Shutdown()
	if err != nil {
		zap.S().Error("生产者shutdown错误：" + err.Error())
		os.Exit(1)
	}
}
