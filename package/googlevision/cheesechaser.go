package vision

import (
	"bytes"
	"context"
	"image"
	"image/jpeg"

	vision "cloud.google.com/go/vision/apiv1"
	// Replacing the deprecated package
	visionpb "cloud.google.com/go/vision/v2/apiv1/visionpb"
	log "github.com/sirupsen/logrus"
)

// CheeseChase detecta faces em uma imagem usando a API do Google Cloud Vision e retorna as anotações das faces.
func CheeseChase(faceImage image.Image) ([]*visionpb.FaceAnnotation, error) {
	// Cria um novo cliente da API do Google Cloud Vision
	context := context.Background()
	client, err := vision.NewImageAnnotatorClient(context)
	if err != nil {
		log.Errorf("failed to create client: %s", err)
	}
	defer client.Close()

	// Converte a imagem para um formato que a API do Google Cloud Vision pode entender
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, faceImage, nil); err != nil {
		log.Errorf("failed to encode image: %s", err)
	}
	image, err := vision.NewImageFromReader(buffer)
	if err != nil {
		log.Errorf("failed to create image: %s", err)
	}

	// Detecta faces na imagem
	faces, err := client.DetectFaces(context, image, nil, 10)
	if err != nil {
		log.Errorf("failed to detect faces: %s", err)
	}

	return faces, nil
}
