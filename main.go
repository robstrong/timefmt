package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	unixNs = flag.Bool("unix-ns", false, "Output unix timestamp in nanoseconds")
	unixUs = flag.Bool("unix-us", false, "Output unix timestamp in microseconds")
	unixMs = flag.Bool("unix-ms", false, "Output unix timestamp in milliseconds")
	unixS  = flag.Bool("unix-s", false, "Output unix timestamp in seconds")
)

var formatters = []func(string) (time.Time, error){
	unixToTime,
	nowToUnix,
	StdFormatter("2006-01-02_15:04:05", time.UTC),
	StdFormatter("2006-01-02-150405", time.UTC),
	StdFormatter("2006-01-02 15:04:05 MST", nil),
	StdFormatter("2006-01-02 15:04:05", time.UTC),
	StdFormatter("Mon, 02 Jan 2006 15:04:05 MST", nil),
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("Usage: timefmt <datetime>")
		os.Exit(1)
	}
	input := strings.Join(flag.Args(), " ")

	for _, formatter := range formatters {
		res, err := formatter(input)
		if err == nil {
			printTime(res)
			os.Exit(0)
		}
	}
	fmt.Println("could not parse format: " + input)
	os.Exit(1)
}

func printTime(t time.Time) {
	switch {
	case *unixS:
		fmt.Println(t.Unix())
	case *unixMs:
		fmt.Println(t.UnixNano() / int64(time.Millisecond))
	case *unixUs:
		fmt.Println(t.UnixNano() / int64(time.Microsecond))
	case *unixNs:
		fmt.Println(t.UnixNano())
	default:
		fmt.Println(t.Local().String())
		fmt.Println(t.UTC().String())
		fmt.Printf("%d\t\tseconds\n", t.Unix())
		fmt.Printf("%d\t\tmilliseconds\n", t.UnixNano()/int64(time.Millisecond))
		fmt.Printf("%d\tmicroseconds\n", t.UnixNano()/int64(time.Microsecond))
		fmt.Printf("%d\tnanoseconds\n", t.UnixNano())
	}
}

func nowToUnix(v string) (time.Time, error) {
	if v != "now" {
		return time.Time{}, errors.New("expected now")
	}
	ts := time.Now()
	return ts, nil
}

func unixToTime(v string) (time.Time, error) {
	ts, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	//we take a datetime 20 years in the future and compare it to different precisions
	//if the future date is greater than the timestamp passed in, it's probably in that precision
	twentyYearsFromNow := time.Now().Add(time.Hour * 24 * 365 * 20)
	switch {
	case twentyYearsFromNow.Unix() > ts:
		return time.Unix(ts, 0), nil
	case twentyYearsFromNow.UnixNano()/int64(time.Millisecond) > ts:
		return time.Unix(0, int64(time.Millisecond)*ts), nil
	case twentyYearsFromNow.UnixNano()/int64(time.Microsecond) > ts:
		return time.Unix(0, ts*int64(time.Microsecond)), nil
	default:
		return time.Unix(0, ts), nil
	}

}

func StdFormatter(format string, loc *time.Location) func(s string) (time.Time, error) {
	return func(s string) (time.Time, error) {
		if loc == nil {
			return time.Parse(format, s)
		}
		return time.ParseInLocation(format, s, loc)
	}
}
