package vm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

type InstrumentRequest struct {
	Name         string   `json:"name"`
	Input        string   `json:"input"`
	Instructions []uint64 `json:"instructions"`
}

type ExecutionRegistry interface {
	Register(pc uint64)
	SendRegistriesToFuzzer()
}

type executionRegistry struct {
	name         string
	input        string
	instructions []uint64
}

var lock = &sync.Mutex{}
var instance *executionRegistry

func GetRegistryInstance(contractName string, input string) *executionRegistry {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()

		if instance == nil {
			instance = &executionRegistry{
				name:         contractName,
				input:        input,
				instructions: make([]uint64, 0, 3),
			}
		}
	}

	return instance
}

func (r *executionRegistry) Register(pc uint64) {
	r.instructions = append(r.instructions, pc)
}

func (r executionRegistry) SendRegistriesToFuzzer() {
	fuzzerHost := os.Getenv("FUZZER_HOST")
	if fuzzerHost == "" {
		fuzzerHost = "localhost"
	}
	url := fmt.Sprintf("http://%s:8888/instrument", fuzzerHost)

	request := InstrumentRequest{
		Name:         r.name,
		Input:        r.input,
		Instructions: r.instructions,
	}
	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("Error Occurred. %+v", err)
	}
	res, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Printf("Error Occurred. %+v", err)
	}
	log.Printf("Sending execution log: %s", res.Status)
}
