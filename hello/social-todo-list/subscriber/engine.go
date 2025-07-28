package subscriber

import (
	"context"
	"fmt"
	"log"
	"main/common"
	"main/common/asyncjob"
	"main/pubsub"

	goservice "github.com/200Lab-Education/go-sdk"
)

type subJob struct {
	Title string
	Hld   func(ctx context.Context, message *pubsub.Message) error
}

type pbEngine struct {
	serviceCtx goservice.ServiceContext // Assuming appctx.AppContext is defined elsewhere
}

func NewEngine(serviceCtx goservice.ServiceContext) *pbEngine {
	return &pbEngine{serviceCtx: serviceCtx}
}

func (engine *pbEngine) Start() error {
	engine.startSubTopic(common.TopicUserLikedItem, true, IncreaseLikeCountAfterUserLikeItem(engine.serviceCtx), PushNotificationAfterUserLikeItem(engine.serviceCtx))

	engine.startSubTopic(common.TopicUserUnLikedItem, true, DecreaseLikeCountAfterUserLikeItem(engine.serviceCtx))
	return nil
}

func (engine *pbEngine) startSubTopic(topic pubsub.Topic, isConcurrent bool, jobs ...subJob) error {
	ps := engine.serviceCtx.MustGet(common.PluginPubSub).(pubsub.PubSub)

	c, _ := ps.Subscribe(context.Background(), topic)

	for _, item := range jobs {
		fmt.Printf("v...: \"Setup subscriber for:\", \"%s\"\n", item.Title)
	}

	getJobHandler := func(job *subJob, message *pubsub.Message) asyncjob.JobHandler {
		return func(ctx context.Context) error {
			fmt.Printf("v...: \"running job for\", \"%s\", \". Value:\", \"%s\"\n", job.Title, message.Data())
			return job.Hld(ctx, message)
		}
	}

	go func() {
		for msg := range c {
			jobHdlArr := make([]asyncjob.Job, len(jobs))

			for i := range jobs {
				jobHdl := getJobHandler(&jobs[i], msg)
				jobHdlArr[i] = asyncjob.NewJob(jobHdl, asyncjob.WithName(jobs[i].Title))
			}

			group := asyncjob.NewGroup(isConcurrent, jobHdlArr...)

			if err := group.Run(context.Background()); err != nil {
				log.Println(err)
			}

		}
	}()

	return nil
}
