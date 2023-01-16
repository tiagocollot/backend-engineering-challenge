package main

import (
	"bufio"
	"fmt"
	"io"
	"time"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// avgState struct to store sum and count
type avgState struct {
	sum float64
	cnt uint
}

// movingAvg struct to store the windowHead, windowTail, end time, byMinute and state
type movingAvg struct {
	windowHead time.Time
	windowTail time.Time
	end        time.Time
	byMinute   map[time.Time]avgState
	state      avgState
}

// calculateAvg function to read and write the average of the data
func calculateAvg(reader io.Reader, writer io.Writer, wsize uint) error {
	ma, err := readAggregated(reader, wsize)
	if err != nil {
		return err
	}

	// iterate over the data and write the average to the writer
	for ma.windowHead.Before(ma.end) {
		ma.writeAvg(writer)
		ma.advanceHead()
		ma.advanceTail()
	}
	return nil
}

// writeAvg function to write the avg to the writer
func (ma *movingAvg) writeAvg(writer io.Writer) error {
	var avg float64
	if ma.state.cnt != 0 {
		avg = ma.state.sum / float64(ma.state.cnt)
	} // set the json data
	data, err := sjson.SetBytes([]byte{}, "date", ma.windowHead.Add(time.Minute).Format("2006-01-02 15:04:00"))
	if err != nil {
		return err
	}
	data, err = sjson.SetBytes(data, "average_delivery_time", avg)
	if err != nil {
		return err
	}
	_, err = writer.Write(data)
	writer.Write([]byte("\n"))
	return err
}

// advanceHead function to move the windowHead
func (ma *movingAvg) advanceHead() {
	ma.windowHead = ma.windowHead.Add(time.Minute)
	if v, ok := ma.byMinute[ma.windowHead]; ok {
		ma.state.sum += v.sum
		ma.state.cnt += v.cnt
	}
}

// advanceTail function to move the windowTail
func (ma *movingAvg) advanceTail() {
	if v, ok := ma.byMinute[ma.windowTail]; ok {
		ma.state.sum -= v.sum
		ma.state.cnt -= v.cnt
	}
	ma.windowTail = ma.windowTail.Add(time.Minute)
}

// readAggregated function to read and aggregate the data
func readAggregated(reader io.Reader, wsize uint) (*movingAvg, error) {
	ma := movingAvg{
		byMinute: make(map[time.Time]avgState),
	}
	sc := bufio.NewScanner(reader)
	for sc.Scan() {
		data := sc.Bytes()
		ts, err := getTimestamp(data)
		if err != nil {
			return nil, err
		}
		if ma.windowHead.IsZero() {
			ma.windowHead = ts.Add(-time.Minute)
			ma.windowTail = ts.Add(-time.Duration(wsize) * time.Minute)
		}
		ma.end = ts
		dur, err := getDuration(data)
		if err != nil {
			return nil, err
		}
		v := ma.byMinute[ts]
		v.sum += dur
		v.cnt += 1
		ma.byMinute[ts] = v
	}
	ma.end = ma.end.Add(time.Minute)
	return &ma, sc.Err()
}

// getTimestamp function to get the timestamp from the data
func getTimestamp(data []byte) (time.Time, error) {
	tsRes := gjson.GetBytes(data, "timestamp")
	if !tsRes.Exists() {
		return time.Time{}, fmt.Errorf("no timestamp in %s", data)
	}
	ts, err := time.Parse("2006-01-02 15:04:05.000000", tsRes.String())
	if err != nil {
		return time.Time{}, err
	}
	ts = ts.Round(time.Minute)
	return ts, nil
}

// getDuration function to get the duration from the data
func getDuration(data []byte) (float64, error) {
	durRes := gjson.GetBytes(data, "duration")
	if !durRes.Exists() {
		return 0, fmt.Errorf("no duration in %s", data)
	}
	return durRes.Float(), nil
}
