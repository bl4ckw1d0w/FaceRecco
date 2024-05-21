package main

import (
	"github.com/vee2xx/camtron"
)

func main() {
	// Inicia a captura da webcam e o processamento do stream
	camtron.StartStreamToFileConsumer()

	// Inicia o aplicativo Electron que se conecta à webcam e captura o stream
	camtron.StartCam()
}
