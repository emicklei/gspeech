// Command quickstart generates an audio file with the content "Hello, World!".
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
)

var oPlay = flag.Bool("p", false, "also play the result next to generating the MP3")
var oInput = flag.String("i", "", "text or ssml file name")
var oGender = flag.String("g", "female", "male|female|neutral")
var oVoice = flag.String("v", "en-AU-Wavenet-C", "pick a voice from https://cloud.google.com/text-to-speech/docs/voices")
var oLang = flag.String("l", "en-AU", "English (Austrialia)")

func main() {
	flag.Parse()
	if len(*oInput) == 0 {
		flag.Usage()
		os.Exit(0)
	}
	inputName := *oInput
	content, err := os.ReadFile(inputName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Text content read from file: %v\n", inputName)

	outputName := strings.ReplaceAll(inputName, ".txt", ".mp3")
	isSSML := strings.HasSuffix(inputName, ".ssml")
	if isSSML {
		outputName = strings.ReplaceAll(inputName, ".ssml", ".mp3")
	}

	// Instantiates a client.
	ctx := context.Background()

	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	var input *texttospeechpb.SynthesisInput
	if isSSML {
		input = &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Ssml{Ssml: string(content)},
		}
	} else {
		input = &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: string(content)},
		}
	}

	var ssmlGender texttospeechpb.SsmlVoiceGender
	switch *oGender {
	case "male":
		ssmlGender = texttospeechpb.SsmlVoiceGender_MALE
	case "neutral":
		ssmlGender = texttospeechpb.SsmlVoiceGender_NEUTRAL
	default:
		ssmlGender = texttospeechpb.SsmlVoiceGender_FEMALE
	}

	// Perform the text-to-speech request on the text input with the selected
	// voice parameters and audio file type.
	req := texttospeechpb.SynthesizeSpeechRequest{
		Input: input,
		Voice: &texttospeechpb.VoiceSelectionParams{
			Name:         *oVoice,
			LanguageCode: *oLang,
			SsmlGender:   ssmlGender,
		},
		// Select the type of audio file you want returned.
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
		},
	}

	resp, err := client.SynthesizeSpeech(ctx, &req)
	if err != nil {
		log.Fatal(err)
	}

	// The resp's AudioContent is binary.
	err = os.WriteFile(outputName, resp.AudioContent, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Audio content written to file: %v\n", outputName)

	if *oPlay {
		if err := exec.Command("open", outputName).Run(); err != nil {
			log.Fatal(err)
		}
	}
}
