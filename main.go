package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/sashabaranov/go-openai"
)

var modelName string
var instruction string
var inputPath string
var outputPath string
var responseTimeoutSeconds int

var openAiKey string = os.Getenv("OPENAI_KEY")

func init() {
	flag.StringVar(&modelName, "model", "gpt-3.5-turbo", "OpenAI model to use")
	flag.StringVar(&instruction, "instruction", "", "prompt to use, will be prepended to input")
	flag.StringVar(&inputPath, "input", "", "where to get the input from (- means stdin)")
	flag.StringVar(&outputPath, "output", "", "where to put the output (- means stdout)")
	flag.IntVar(&responseTimeoutSeconds, "timeout", 600, "timeout for model response, in seconds")
	flag.Parse()
}

func main() {
	openAiClient := openai.NewClient(openAiKey)

	var inFile io.Reader
	if inputPath == "" || inputPath == "-" {
		inFile = os.Stdin
	} else {
		inFileCloser, err := os.Open(inputPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: Failed to open input file %s, error %v", inputPath, err)
		}
		inFile = inFileCloser
		defer inFileCloser.Close()
	}

	var outFile io.Writer
	if outputPath == "" || outputPath == "-" {
		outFile = os.Stdout
	} else {
		outFileCloser, err := os.Create(outputPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: Failed to open input file %s, error %v\n", outputPath, err)
			os.Exit(1)
		}
		outFile = outFileCloser
		defer outFileCloser.Close()
	}

	inBytes, err := io.ReadAll(inFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to read input file %s, error %v\n", inputPath, err)
		os.Exit(1)
	}
	inString := string(inBytes)

	ctx, ctxCancel := context.WithTimeout(context.Background(), time.Second*time.Duration(responseTimeoutSeconds))
	defer ctxCancel()
	openAiResp, err := openAiClient.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: modelName,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: modelPreprompt,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: instruction + "\n\n" + inString,
			},
		},
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to get response from OpenAI API, error %v\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(outFile, "%s\n", openAiResp.Choices[0].Message.Content)
}
