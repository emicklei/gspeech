# gspeech

Simple tool that uses the Google Text-to-Speech API

## requirements

- GCP project with Text-to-Speech API enabled
- Create a Serviceaccount with Owner permissions :-(
- Generate and download an API key for this service account (call it service-account-key.json)

## install

This step requires the Go SDK.

    go install

## usage

    Usage of gspeech:
    -g string
            male|female|neutral (default "female")
    -i string
            text or ssml file name
    -l string
            English (Austrialia) (default "en-AU")
    -p	also play the result next to generating the MP3
    -v string
            pick a voice from https://cloud.google.com/text-to-speech/docs/voices (default "en-AU-Wavenet-C")

## run

    GOOGLE_APPLICATION_CREDENTIALS=service-account-key.json gspeech -p -i welcome.ssml
