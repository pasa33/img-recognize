package imgrecognize

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

func (m *ImgRecognize) RecognFromFile(imgpath string) (map[string]float64, error) {
	imageBytes, err := os.ReadFile(imgpath)
	if err != nil {
		return nil, fmt.Errorf("errore nella lettura dell'immagine: %v", err)
	}

	return m.RecognFromBytes(imageBytes)
}

func (m *ImgRecognize) RecognFromBytes(imgBytes []byte) (map[string]float64, error) {

	m.sem <- struct{}{} //sem for concurrent
	defer func() { <-m.sem }()

	cmd := exec.Command(m.scriptPath)
	cmd.Stdin = bytes.NewReader(imgBytes)
	cmd.Env = append(os.Environ(), "PYTHONIOENCODING=utf-8")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("errore nell'esecuzione dello script Python: %v\nStderr: %s", err, stderr.String())
	}

	var predictions map[string]float64
	err = json.Unmarshal(stdout.Bytes(), &predictions)
	if err != nil {
		return nil, fmt.Errorf("errore nella decodifica dell'output JSON: %v", err)
	}

	return predictions, nil
}
