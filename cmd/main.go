package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gammazero/workerpool"

	"github.com/kevinsudut/nsq-publisher-using-workerpool/config"
	"github.com/kevinsudut/nsq-publisher-using-workerpool/lib/csv"
	"github.com/kevinsudut/nsq-publisher-using-workerpool/lib/log"
	"github.com/kevinsudut/nsq-publisher-using-workerpool/lib/mq"
	"github.com/kevinsudut/nsq-publisher-using-workerpool/util"
)

var (
	logObj log.Logger
)

func errMessage(message string, err error) string {
	return fmt.Sprintf("%s: %s", message, err.Error())
}

type Data struct {
	UserID  string `json:"userid"`
	ShopID  string `json:"shopid"`
	QuestID string `json:"questid"`
}

type NSQPayload struct {
	UserID           int64  `json:"user_id"`
	ShopID           int64  `json:"shop_id"`
	UniqueIdentifier string `json:"unique_identifier"`
	Timestamp        int64  `json:"timestamp"`
}

func readCSVData(cfg config.CSVConfig) ([]Data, error) {
	var res []Data

	csvContent, err := csv.ReadFile(fmt.Sprintf("./%s/%s", cfg.Path, cfg.FileName))
	if err != nil {
		return nil, err
	}

	err = csvContent.ParseToStruct(&res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func publishData(cfg config.NSQConfig, rawData []Data) error {
	publisher, err := mq.InitProducer(cfg.MessageQueue)
	if err != nil {
		return err
	}

	data, err := util.ChunkSliceOfInterface(rawData, len(rawData)/cfg.NumPayloadMultiPublish)
	if err != nil {
		return err
	}

	wp := workerpool.New(cfg.NumWorker)

	for _, datum := range data {
		var temp []Data

		err = util.ConvertSliceOfInterfaceToStruct(datum, &temp)
		if err != nil {
			logObj.Errorln(fmt.Sprintf("[Error:%s][Payload:%+v]", err.Error(), datum))
			continue
		}

		payload := buildNSQPayload(temp)

		if len(payload) == 0 {
			continue
		}

		wp.Submit(func() {
			err = publisher.MultiPublishJSON(cfg.TopicName, payload)
			if err != nil {
				logObj.Errorln(fmt.Sprintf("[Error:%s][Payload:%+v]", err.Error(), payload))
				return
			}

			logObj.Infoln(fmt.Sprintf("[Success][Payload:%+v]", payload))
		})
	}

	wp.StopWait()

	return err
}

func Atoi64(str string, defaultValue int64) int64 {
	i64, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return defaultValue
	}

	return i64
}

func buildNSQPayload(data []Data) []NSQPayload {
	res := []NSQPayload{}

	for i := 0; i < len(data); i++ {
		res = append(res, NSQPayload{
			UserID:           Atoi64(data[i].UserID, 0),
			ShopID:           Atoi64(data[i].ShopID, 0),
			UniqueIdentifier: fmt.Sprintf("%s:%s", data[i].ShopID, data[i].QuestID),
			Timestamp:        time.Now().Unix(),
		})
	}

	return res
}

func main() {
	cfg, err := config.ReadConfig("./config.yaml")
	if err != nil {
		panic(errMessage("ReadConfig", err))
	}

	fmt.Println("Successfully read config file")

	timeNow := time.Now().UnixNano()
	logObj, err = log.InitLog(log.LogConfig{
		LogLevel:      cfg.Log.LogLevel,
		Path:          cfg.Log.Path,
		ErrorFileName: fmt.Sprintf("%d_error", timeNow),
		InfoFileName:  fmt.Sprintf("%d_info", timeNow),
	})
	if err != nil {
		panic(errMessage("InitLog", err))
	}

	fmt.Println("Successfully init log")

	csvData, err := readCSVData(cfg.CSV)
	if err != nil {
		panic(errMessage("ReadCSVData", err))
	}

	fmt.Println("Successfully read csv file")

	err = publishData(cfg.NSQ, csvData)
	if err != nil {
		panic(errMessage("PublishData", err))
	}

	fmt.Println("Successfully publish all data")

	time.Sleep(time.Second * 5)
}
