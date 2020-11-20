package userful

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"regexp"
	"time"

	"github.com/go-redis/redis/v8"

	ccb "gitlab.zaojiu.com/ziyuan/common/pkg/cb"
	"gitlab.zaojiu.com/ziyuan/common/pkg/log"
	"gitlab.zaojiu.com/ziyuan/common/pkg/orm"
	"gitlab.zaojiu.com/ziyuan/common/pkg/store"
)

type Job interface {
	Do() bool
}

type Worker struct {
	JobQueue chan Job
}

func NewWorker() Worker {
	return Worker{
		JobQueue: make(chan Job),
	}
}

func (w Worker) Start(ctx context.Context, wp chan chan Job, resCh chan bool) {
	go func() {
		for {
			wp <- w.JobQueue
			select {
			case <-ctx.Done():
				return
			case job := <-w.JobQueue:
				res := job.Do()
				resCh <- res
			}
		}
	}()
}

type WorkerPool struct {
	Workerlen   int
	JobQueue    chan Job
	WorkerQueue chan chan Job
}

func NewWorkerPool(workerlen int) *WorkerPool {
	return &WorkerPool{
		Workerlen:   workerlen,
		JobQueue:    make(chan Job),
		WorkerQueue: make(chan chan Job, workerlen),
	}
}

func (w *WorkerPool) Start(ctx context.Context, resCh chan bool) {
	// workerpool init
	for i := 0; i < w.Workerlen; i++ {
		worker := NewWorker()
		worker.Start(ctx, w.WorkerQueue, resCh)
	}
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			case job := <-w.JobQueue:
				select {
				case <-ctx.Done():
					return
				case worker := <-w.WorkerQueue:
					select {
					case <-ctx.Done():
						return
					case worker <- job:
					}
				}
			}
		}
	}(ctx)
}

type Task struct {
	Content string // 需要检查的内容
	Text    string // 违禁词字段
}

func (t *Task) Do() bool {
	match, _ := regexp.MatchString(t.Text, t.Content)
	return match
}

func GenTasks(ctx context.Context, wp *WorkerPool, content string, bannedWrodsList *[]string) {
	go func() {
		for _, text := range *bannedWrodsList {
			task := &Task{Content: content, Text: text}
			select {
			case <-ctx.Done():
				return
			case wp.JobQueue <- task:
			}
		}
	}()
}

func BannedWordTest1(content *string, bannedWordsList *[]string) bool {
	status := true
	for i, word := range *bannedWordsList {
		match, _ := regexp.MatchString(word, *content)
		fmt.Println(i, word)
		if match {
			status = false
			break
		}
	}
	return status
}

func BannedWordChunk(ctx context.Context, resCh chan bool, num int, content *string, words *[]string) {
	go func() {
		for _, word := range *words {
			match, _ := regexp.MatchString(word, *content)
			select {
			case <-ctx.Done():
				return
			case resCh <- match:
			}
		}
	}()
}

func BannedWordTest2(content *string, bannedWordsList *[]string) bool {
	status := true
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	resCh := make(chan bool, 100)
	length := 20
	listLength := len(*bannedWordsList)
	chuck := int(math.Floor(float64(listLength) / float64(length)))
	for i := 0; i < length; i++ {
		start := i * chuck
		end := (i + 1) * chuck
		if i == length-1 {
			end = listLength
		}
		newList := (*bannedWordsList)[start:end]
		BannedWordChunk(ctx, resCh, i, content, &newList)
	}
	for i := 0; i < listLength; i++ {
		res := <-resCh
		if res {
			status = false
			break
		}
	}
	return status
}

func CheckBannedWords(content *string, bannedWordsList *[]string) bool {
	status := true
	// 执行结果通道
	resCh := make(chan bool, 100)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// workerlen并发个数
	workerlen := 50
	listLength := len(*bannedWordsList)
	if listLength < workerlen {
		workerlen = listLength
	}
	pool := NewWorkerPool(workerlen)
	pool.Start(ctx, resCh)
	// 将任务放入jobqueue中
	GenTasks(ctx, pool, *content, bannedWordsList)
	for i := 0; i < listLength; i++ {
		res := <-resCh
		if res {
			status = false
			break
		}
	}
	return status
}

// 检查用户所输入内容是否属于违禁词
// 返回 true 时表示可以正常发送
func IsBannedWord(text *string, db *store.DB, rdb *store.RDB) bool {
	status := true
	bannedWordsList := GetBannedWordsList(db, rdb)
	start := time.Now().UnixNano()
	if len(bannedWordsList) == 0 {
		return status
	}
	BannedWordTest1(text, &bannedWordsList)
	// BannedWordTest2(text, &bannedWordsList)
	// CheckBannedWords(text, &bannedWordsList)
	end := time.Now().UnixNano()
	fmt.Println("=========== total:", end-start)
	return status
}

func GetBannedWordsList(db *store.DB, rdb *store.RDB) []string {
	ctx := context.Background()
	rdbName, err := rdb.Get(ctx, ccb.BannedWordsKey).Result()
	flag := true // 是否查询数据库
	res := []string{}
	if err == redis.Nil || err == nil {
		err := json.Unmarshal([]byte(rdbName), &res)
		if err != nil {
			log.Error("banned word unmarshal", log.ZError(err))
			rdb.Del(ctx, ccb.BannedWordsKey)
		} else {
			flag = false
		}
	} else if err != nil {
		log.Error("rdb get", log.ZError(err))
	}
	if flag {
		var words []orm.BannedWord
		err = db.Find(&words).Error
		if err != nil {
			log.Error("banned word find", log.ZError(err))
			return res
		}
		for _, word := range words {
			res = append(res, word.Name)
		}
		rdbContents, err := json.Marshal(res)
		if err != nil {
			log.Error("contents marshal", log.ZError(err))
		} else {
			rdb.Set(ctx, ccb.BannedWordsKey, string(rdbContents), 0)
		}
	}
	return res
}
