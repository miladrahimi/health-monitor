package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

var targets []string
var ctx = context.Background()
var rdb *redis.Client

func main() {
	setupLogger()
	setupRedis()
	configTargets()
	startWorker()

	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.HandleFunc("/data", handleData)

	err := http.ListenAndServe(":7575", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func setupLogger() {
	n := fmt.Sprintf("./logs/%s.log", time.Now().Format("2006-01-02"))
	f, err := os.OpenFile(n, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(f)
	log.Println("START")
}

func setupRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "",
		DB:       0,
	})
}

func configTargets() {
	targets = strings.Split(os.Getenv("TARGETS"), ",")
	if len(targets) > 5 {
		log.Fatalln("Cannot handle more than 5 targets.")
	}

	log.Println(targets)
}

func handleData(w http.ResponseWriter, _ *http.Request) {
	w.Header().Add("Content-type", "application/json")

	keys, err := rdb.Keys(ctx, "*").Result()
	if err != nil {
		publishError(w, err)
	}

	values := make(map[string]map[string]int64)
	for _, k := range keys {
		kp := strings.Split(k, "=")

		cmd := rdb.Get(ctx, k)
		if cmd.Err() != nil {
			publishError(w, cmd.Err())
		}

		v, err := strconv.Atoi(cmd.Val())
		if err != nil {
			publishError(w, err)
		}

		if values[kp[0]] == nil {
			values[kp[0]] = make(map[string]int64)
		}

		values[kp[0]][kp[1]] = int64(v)
	}

	j, err := json.Marshal(values)
	if err != nil {
		publishError(w, err)
	}

	if _, err := fmt.Fprintf(w, string(j)); err != nil {
		log.Fatal(err)
	}
}

func publishError(w http.ResponseWriter, e error) {
	w.WriteHeader(500)

	if _, err := fmt.Fprint(w, "{\"error\":\"Internal Error\"}"); err != nil {
		log.Println(e)
		log.Fatal(err)
	}

	log.Fatal(e)
}

func startWorker() {
	go func() {
		for range time.Tick(20 * time.Second) {
			for _, t := range targets {
				go func(s string) {
					err := call(s)
					if err != nil {
						log.Println(err)
					}
				}(t)
			}
		}
	}()
}

func call(target string) error {
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	start := time.Now()
	response, err := client.Get(target)
	if err != nil {
		return err
	}
	latency := time.Now().Sub(start)

	l := int64(0)
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		l = latency.Nanoseconds()
	}

	if err = persist(target, l); err != nil {
		return err
	}

	return nil
}

func persist(target string, latency int64) error {
	loc, err := time.LoadLocation(os.Getenv("TIMEZONE"))
	if err != nil {
		return err
	}

	u, err := url.Parse(target)
	if err != nil {
		return err
	}

	t := time.Now().In(loc)
	key := fmt.Sprintf("%s=%s", u.Host, t.Format("2006/01/02 15:04"))
	if err := rdb.Set(ctx, key, latency, 24*time.Hour).Err(); err != nil {
		return err
	}

	return nil
}
