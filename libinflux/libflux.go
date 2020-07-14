package libinflux

import (
	"errors"
	"log"
	"time"

	"github.com/influxdata/influxdb/client/v2"
)

//Pointer --
type Pointer interface {
	MakePoint(chan *client.Point)
}

//InfluxLib --
type InfluxLib struct {
	points      chan *client.Point
	cl          client.Client
	waitTime    time.Duration
	maxPointNum int32
	dbServer    string
	name        string
	password    string
	database    string
}

//FluxConfig --
type FluxConfig struct {
	WaitPushTime int
	MaxPointNum  int32
	DbServer     string
	Name         string
	Password     string
	Database     string
}

//NewInfluxdata --
func NewInfluxdata(config FluxConfig) (nifd *InfluxLib, err error) {
	nifd = &InfluxLib{}
	nifd.points = make(chan *client.Point, 10000)
	nifd.waitTime = time.Duration(config.WaitPushTime) * time.Second
	nifd.maxPointNum = config.MaxPointNum
	nifd.dbServer = config.DbServer
	nifd.name = config.Name
	nifd.password = config.Password
	nifd.database = config.Database
	nifd.cl, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:     nifd.dbServer,
		Username: nifd.name,
		Password: nifd.password,
	})
	if err != nil {
		return nil, err
	}
	return nifd, nil
}

//RegisterPointer --
func (ifd *InfluxLib) RegisterPointer(pers []Pointer) (err error) {
	if len(pers) < 1 {
		return errors.New("Pers is empty")
	}
	for _, v := range pers {
		go v.MakePoint(ifd.points)
	}
	ifd.readPoint()
	return nil
}

//GetPointChan --
func (ifd *InfluxLib) GetPointChan() (chp chan *client.Point) {
	return ifd.points
}

//Start --
func (ifd *InfluxLib) Start() {
	ifd.readPoint()
}

func (ifd *InfluxLib) readPoint() {
	var count int32
	timer := time.NewTimer(ifd.waitTime)
	bp := ifd.newBP()
	startTime := time.Now()
	for {
		if count == ifd.maxPointNum {
			go ifd.writeToDb(bp)
			bp = ifd.newBP()
			count = 0
			startTime = time.Now()
		}
		if time.Now().Sub(startTime) >= ifd.waitTime {
			if count > 0 {
				go ifd.writeToDb(bp)
				bp = ifd.newBP()
				count = 0
				startTime = time.Now()
			}
		}
		timer.Reset(ifd.waitTime)
		select {
		case p := <-ifd.points:
			bp.AddPoint(p)
			count++
		case <-timer.C:
			if count < 1 {
				break
			}
			go ifd.writeToDb(bp)
			bp = ifd.newBP()
			count = 0
			startTime = time.Now()
		}

	}

}

func (ifd *InfluxLib) newBP() (bp client.BatchPoints) {
	var err error
	bp, err = client.NewBatchPoints(client.BatchPointsConfig{Database: ifd.database})
	if err != nil || bp == nil {
		log.Println(err)
		panic(err)
	}
	return bp
}

func (ifd *InfluxLib) writeToDb(oldBp client.BatchPoints) {
	err := ifd.cl.Write(oldBp)
	if err != nil {
		log.Println(err)
	}
}

//QueryDB --
func (ifd *InfluxLib) QueryDB(cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: ifd.database,
	}
	if response, err := ifd.cl.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}

//FirstTime --
func (ifd *InfluxLib) FirstTime(cmd string) (fst int64, err error) {
	req, err := ifd.QueryDB(cmd)
	if err != nil {
		return 0, err
	}
	var colID int
	colID = -1
	if len(req) > 0 && len(req[0].Series) > 0 {
		for k, v := range req[0].Series[0].Columns {
			if v == "time" {
				colID = k
				break
			}
		}
	}
	if colID >= 0 && len(req[0].Series[0].Values) > 0 && len(req[0].Series[0].Values[0]) > 0 {
		vst := req[0].Series[0].Values[0][colID].(string)
		vt, err := time.Parse("2006-01-02T15:04:05.99999999Z", vst)
		if err != nil {
			return 0, err
		}
		fst = vt.UnixNano()
	}
	if fst <= 0 {
		return 0, nil
	}
	return fst, nil
}

//
