package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"

	shell "github.com/ipfs/go-ipfs-api"
)

type Message struct {
	Action string `json:"action"`
	Data   struct {
		Event string `json:"event"`
		Code  string `json:"code"`
	} `json:"data"`
}

func decodeUnwrap(encoded string) string {
	uEnv, _ := b64.URLEncoding.DecodeString(encoded)
	return string(uEnv)
}

func writeCodeToTempFile(code string) {
	path := "/tmp/43556789705456847598/" // build the temp DIR
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}

	// write the function file body to tmp
	err := ioutil.WriteFile("/tmp/43556789705456847598/lambda_function.py", []byte(code), 0644)
	if err != nil {
		panic(err)
	}
}

func executeLambdaDocker(event string) string {
	cmd := "docker"
	args := []string{
		"run", "--rm", "-v", "/tmp/43556789705456847598:/var/task",
		"lambci/lambda:python3.7", "lambda_function.lambda_handler", event,
	}
	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func ListenForExecute() {
	// Where your local node is running on localhost:5001
	sh := shell.NewShell("localhost:5001")
	sub, _ := sh.PubSubSubscribe("test")
	for true {
		r, _ := sub.Next()
		var msg Message
		err := json.Unmarshal(r.Data, &msg)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(msg) // shows what we parsed out
		code := decodeUnwrap(msg.Data.Code)
		event := decodeUnwrap(msg.Data.Event)
		writeCodeToTempFile(code)
		out := executeLambdaDocker(event)
		// fmt.Println(out) // shows Lambda Docker Resp - not logs
		sh.PubSubPublish("test-response", out)
		time.Sleep(time.Second)
	}
}
