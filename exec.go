package main

import (
	"encoding/json"
	"os/exec"
	"os"
	"fmt"
)

type status struct {
	Status string							`json:"status"`
	Out string								`json:"stdout"`
	Result map[string]interface{}			`json:"jsonst"`
}

func (s *status) getExec(key string, value string ) error {

	cmdPath := os.Getenv("WORK_DIR")
	cmdExec := cmdPath + "WORK_EXEC"
	cmdArguments := []string{key, value}
	cmd := exec.Command(cmdExec, cmdArguments...)
	cmd.Dir = os.Getenv("WORK_DIR")
	out, err := cmd.CombinedOutput()

	if err != nil {
		s.Out = err.Error()
		return err
	}

	s.Out = string(out)
	json.Unmarshal(out, &s.Result)

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
	out, err := cmd.CombinedOutput()

	if err != nil {
		s.Out = err.Error()
		return err
	}

	s.Out = string(out)
	json.Unmarshal(out, &s.Result)

	return nil
}

func (s *status) putExec(key string, p string) error{

	cmdPath := os.Getenv("WORK_DIR")
	cmdExec := cmdPath + os.Getenv("WORK_EXEC")
	cmd := exec.Command(cmdExec, p)
	cmd.Dir = os.Getenv("WORK_DIR")
	out, err := cmd.CombinedOutput()

	if err != nil {
		s.Out = err.Error()
		return err
	}

	s.Out = string(out)
	json.Unmarshal(out, &s.Result)

	return nil
}