package main

import (
	"encoding/json"
	"os/exec"
	"bytes"
	"os"
	"fmt"
)

type status struct {
	Status string							`json:"status"`
	Out string								`json:"out"`
	Result map[string]interface{}			`json:"result"`
}

func (s *status) getExec(key string, value string ) error {

	cmdPath := os.Getenv("WORK_DIR")
	cmdExec := cmdPath + "WORK_EXEC"
	cmdArguments := []string{key, value}
	cmd := exec.Command(cmdExec, cmdArguments...)
	cmd.Dir = os.Getenv("WORK_DIR")

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Start()

	if err != nil {
		s.Out = err.Error()
		return err
	}

	err = cmd.Wait()
	if err != nil {
		s.Out = err.Error()
		return err
	}

	s.Out = out.String()
	json.Unmarshal(out.Bytes(), &s.Result)

	return nil
}

func (s *status) postExec(key string, postj map[string]interface{}) error{

	var args = make([]string, 0)

	for _, n := range postj {
		str := fmt.Sprint(n)
		args = append(args, str)
	}

	cmdPath := os.Getenv("WORK_DIR")
	cmdExec := cmdPath + os.Getenv("WORK_EXEC")
	cmd := exec.Command(cmdExec, args...)
	cmd.Dir = os.Getenv("WORK_DIR")

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Start()

	if err != nil {
		s.Out = err.Error()
		return err
	}

	err = cmd.Wait()
	if err != nil {
		s.Out = err.Error()
		return err
	}

	s.Out = out.String()
	json.Unmarshal(out.Bytes(), &s.Result)

	return nil
}