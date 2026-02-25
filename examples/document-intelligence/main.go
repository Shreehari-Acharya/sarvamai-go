package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Shreehari-Acharya/sarvam-go-sdk"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/docintel"
)

func main() {
	ctx := context.Background()

	apiKey := os.Getenv("SARVAM_API_KEY")
	if apiKey == "" {
		log.Fatal("SARVAM_API_KEY environment variable not set")
	}

	client, err := sarvamai.NewClient(sarvamai.Config{
		APIKey: apiKey,
	})
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Step 1: Initialize a new document intelligence job
	fmt.Println("Creating document intelligence job...")
	authToken := "my-secret-token"
	resp, err := client.DocumentIntelligence.Initialize(
		ctx,
		docintel.WithLanguage("hi-IN"),
		docintel.WithOutputFormat(docintel.OutputFormatMD),
		docintel.WithCallback("https://example.com/webhook", &authToken),
	)
	if err != nil {
		log.Fatalf("Initialize failed: %v", err)
	}

	fmt.Printf("Job ID: %s\n", resp.JobID)
	fmt.Printf("Job State: %s\n", resp.JobState)
	fmt.Printf("Storage Container: %s\n", resp.StorageContainerType)

	jobID := resp.JobID

	// Step 2: Get upload URL for the file
	fmt.Println("\nGetting upload URL...")
	uploadResp, err := client.DocumentIntelligence.GetUploadLinks(ctx, jobID, "document.pdf")
	if err != nil {
		log.Fatalf("GetUploadLinks failed: %v", err)
	}

	fmt.Printf("Upload URL: %s\n", uploadResp.UploadUrls["document.pdf"].FileURL)

	// Step 3: Start processing the job
	fmt.Println("\nStarting job processing...")
	startResp, err := client.DocumentIntelligence.Start(ctx, jobID)
	if err != nil {
		log.Fatalf("Start failed: %v", err)
	}

	fmt.Printf("Job State: %s\n", startResp.JobState)

	// Step 4: Poll for status (in real usage, you'd use webhook callbacks)
	fmt.Println("\nPolling for job completion...")
	for {
		status, err := client.DocumentIntelligence.GetStatus(ctx, jobID)
		if err != nil {
			log.Fatalf("GetStatus failed: %v", err)
		}

		fmt.Printf("Job State: %s\n", status.JobState)

		if status.JobState == docintel.JobStateCompleted ||
			status.JobState == docintel.JobStatePartiallyCompleted ||
			status.JobState == docintel.JobStateFailed {
			break
		}

		time.Sleep(5 * time.Second)
	}

	// Step 5: Get download URLs
	fmt.Println("\nGetting download URLs...")
	downloadResp, err := client.DocumentIntelligence.GetDownloadLinks(ctx, jobID)
	if err != nil {
		log.Fatalf("GetDownloadLinks failed: %v", err)
	}

	for filename, details := range downloadResp.DownloadURLs {
		fmt.Printf("Download URL for %s: %s\n", filename, details.FileURL)
	}

	fmt.Println("\nDone!")
}
