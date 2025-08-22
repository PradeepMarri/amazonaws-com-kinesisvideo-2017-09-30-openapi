package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/amazon-kinesis-video-streams/mcp-server/config"
	"github.com/amazon-kinesis-video-streams/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func UpdatestreamHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody map[string]interface{}
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/updateStream", cfg.BaseURL)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.UpdateStreamOutput
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateUpdatestreamTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_updateStream",
		mcp.WithDescription("<p>Updates stream metadata, such as the device name and media type.</p> <p>You must provide the stream name or the Amazon Resource Name (ARN) of the stream.</p> <p>To make sure that you have the latest version of the stream before updating it, you can specify the stream version. Kinesis Video Streams assigns a version to each stream. When you update a stream, Kinesis Video Streams assigns a new version number. To get the latest stream version, use the <code>DescribeStream</code> API. </p> <p> <code>UpdateStream</code> is an asynchronous operation, and takes time to complete.</p>"),
		mcp.WithString("StreamName", mcp.Description("Input parameter: <p>The name of the stream whose metadata you want to update.</p> <p>The stream name is an identifier for the stream, and must be unique for each account and region.</p>")),
		mcp.WithString("CurrentVersion", mcp.Required(), mcp.Description("Input parameter: The version of the stream whose metadata you want to update.")),
		mcp.WithString("DeviceName", mcp.Description("Input parameter: <p>The name of the device that is writing to the stream. </p> <note> <p> In the current implementation, Kinesis Video Streams does not use this name. </p> </note>")),
		mcp.WithString("MediaType", mcp.Description("Input parameter: <p>The stream's media type. Use <code>MediaType</code> to specify the type of content that the stream contains to the consumers of the stream. For more information about media types, see <a href=\"http://www.iana.org/assignments/media-types/media-types.xhtml\">Media Types</a>. If you choose to specify the <code>MediaType</code>, see <a href=\"https://tools.ietf.org/html/rfc6838#section-4.2\">Naming Requirements</a>.</p> <p>To play video on the console, you must specify the correct video type. For example, if the video in the stream is H.264, specify <code>video/h264</code> as the <code>MediaType</code>.</p>")),
		mcp.WithString("StreamARN", mcp.Description("Input parameter: The ARN of the stream whose metadata you want to update.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    UpdatestreamHandler(cfg),
	}
}
