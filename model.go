package imgrecognize

type ImgRecognize struct {
	scriptPath string
	sem        chan struct{}
}

func NewImgRecognize(maxConcurrent int) *ImgRecognize {
	return &ImgRecognize{
		sem:        make(chan struct{}, maxConcurrent),
		scriptPath: "./image_recognition.exe",
	}
}

func (m *ImgRecognize) SetScriptPath(scriptPath string) {
	m.scriptPath = scriptPath
}
