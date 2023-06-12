package subscriber

import (
	"context"
	"g05-fooddelivery/common"
	"g05-fooddelivery/component/appctx"
	asyncjob "g05-fooddelivery/component/async_job"
	"g05-fooddelivery/pubsub"
	"log"
)

type consumerJob struct {
	Title string
	Hld   func(ctx context.Context, message *pubsub.Message) error
}
type consumerEngine struct {
	appCtx appctx.AppContext
	//rtEngine skio.RealtimeEngine
}

func NewEngine(appContext appctx.AppContext) *consumerEngine {
	return &consumerEngine{appCtx: appContext}
}
func (engine *consumerEngine) Start() error {
	engine.startSubTopic(common.TopicUserLikeRestaurant, true, IncreaseLikeCountAfterUserLikeRestaurant(engine.appCtx))
	engine.startSubTopic(common.TopicUserUnLikeRestaurant, true, DereaseLikeCountAfterUserLikeRestaurant(engine.appCtx))
	return nil
}

type GroupJob interface {
	Run(ctx context.Context) error
}

func (engine *consumerEngine) startSubTopic(topic pubsub.Topic, isConcurrent bool, consumerJobs ...consumerJob) error {
	c, _ := engine.appCtx.GetPubSub().Subscribe(context.Background(), topic)
	for _, item := range consumerJobs {
		log.Println("Setup consumer for:", item.Title)
	}
	getJobHandler := func(job *consumerJob, message *pubsub.Message) asyncjob.JobHandler {
		return func(ctx context.Context) error {
			log.Println("running job for", job.Title, ". Value:", message.Data())
			return job.Hld(ctx, message)
		}
	}
	go func() {
		for {
			msg := <-c
			jobHdlArr := make([]asyncjob.Job, len(consumerJobs))
			for i := range consumerJobs {
				jobHdl := getJobHandler(&consumerJobs[i], msg)
				jobHdlArr[i] = asyncjob.NewJob(jobHdl)
			}
			group := asyncjob.NewGroup(isConcurrent, jobHdlArr...)
			if err := group.Run(context.Background()); err != nil {
				log.Println(err)
			}
		}
	}()
	return nil
}
