package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// FormatConfig represents the FormatConfig schema from the OpenAPI specification
type FormatConfig struct {
}

// DescribeSignalingChannelInput represents the DescribeSignalingChannelInput schema from the OpenAPI specification
type DescribeSignalingChannelInput struct {
	Channelarn interface{} `json:"ChannelARN,omitempty"`
	Channelname interface{} `json:"ChannelName,omitempty"`
}

// MediaStorageConfiguration represents the MediaStorageConfiguration schema from the OpenAPI specification
type MediaStorageConfiguration struct {
	Streamarn interface{} `json:"StreamARN,omitempty"`
	Status interface{} `json:"Status"`
}

// UpdateSignalingChannelInput represents the UpdateSignalingChannelInput schema from the OpenAPI specification
type UpdateSignalingChannelInput struct {
	Channelarn interface{} `json:"ChannelARN"`
	Currentversion interface{} `json:"CurrentVersion"`
	Singlemasterconfiguration interface{} `json:"SingleMasterConfiguration,omitempty"`
}

// UpdateStreamInput represents the UpdateStreamInput schema from the OpenAPI specification
type UpdateStreamInput struct {
	Streamname interface{} `json:"StreamName,omitempty"`
	Currentversion interface{} `json:"CurrentVersion"`
	Devicename interface{} `json:"DeviceName,omitempty"`
	Mediatype interface{} `json:"MediaType,omitempty"`
	Streamarn interface{} `json:"StreamARN,omitempty"`
}

// DescribeEdgeConfigurationOutput represents the DescribeEdgeConfigurationOutput schema from the OpenAPI specification
type DescribeEdgeConfigurationOutput struct {
	Lastupdatedtime interface{} `json:"LastUpdatedTime,omitempty"`
	Streamarn interface{} `json:"StreamARN,omitempty"`
	Streamname interface{} `json:"StreamName,omitempty"`
	Syncstatus interface{} `json:"SyncStatus,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Edgeagentstatus interface{} `json:"EdgeAgentStatus,omitempty"`
	Edgeconfig interface{} `json:"EdgeConfig,omitempty"`
	Failedstatusdetails interface{} `json:"FailedStatusDetails,omitempty"`
}

// DescribeNotificationConfigurationInput represents the DescribeNotificationConfigurationInput schema from the OpenAPI specification
type DescribeNotificationConfigurationInput struct {
	Streamarn interface{} `json:"StreamARN,omitempty"`
	Streamname interface{} `json:"StreamName,omitempty"`
}

// DescribeSignalingChannelOutput represents the DescribeSignalingChannelOutput schema from the OpenAPI specification
type DescribeSignalingChannelOutput struct {
	Channelinfo interface{} `json:"ChannelInfo,omitempty"`
}

// DeleteStreamInput represents the DeleteStreamInput schema from the OpenAPI specification
type DeleteStreamInput struct {
	Currentversion interface{} `json:"CurrentVersion,omitempty"`
	Streamarn interface{} `json:"StreamARN"`
}

// GetDataEndpointOutput represents the GetDataEndpointOutput schema from the OpenAPI specification
type GetDataEndpointOutput struct {
	Dataendpoint interface{} `json:"DataEndpoint,omitempty"`
}

// ChannelNameCondition represents the ChannelNameCondition schema from the OpenAPI specification
type ChannelNameCondition struct {
	Comparisonoperator interface{} `json:"ComparisonOperator,omitempty"`
	Comparisonvalue interface{} `json:"ComparisonValue,omitempty"`
}

// DescribeMediaStorageConfigurationInput represents the DescribeMediaStorageConfigurationInput schema from the OpenAPI specification
type DescribeMediaStorageConfigurationInput struct {
	Channelarn interface{} `json:"ChannelARN,omitempty"`
	Channelname interface{} `json:"ChannelName,omitempty"`
}

// DeleteEdgeConfigurationOutput represents the DeleteEdgeConfigurationOutput schema from the OpenAPI specification
type DeleteEdgeConfigurationOutput struct {
}

// UntagStreamOutput represents the UntagStreamOutput schema from the OpenAPI specification
type UntagStreamOutput struct {
}

// UpdateSignalingChannelOutput represents the UpdateSignalingChannelOutput schema from the OpenAPI specification
type UpdateSignalingChannelOutput struct {
}

// UpdateMediaStorageConfigurationInput represents the UpdateMediaStorageConfigurationInput schema from the OpenAPI specification
type UpdateMediaStorageConfigurationInput struct {
	Channelarn interface{} `json:"ChannelARN"`
	Mediastorageconfiguration interface{} `json:"MediaStorageConfiguration"`
}

// UpdateNotificationConfigurationOutput represents the UpdateNotificationConfigurationOutput schema from the OpenAPI specification
type UpdateNotificationConfigurationOutput struct {
}

// LastUploaderStatus represents the LastUploaderStatus schema from the OpenAPI specification
type LastUploaderStatus struct {
	Jobstatusdetails interface{} `json:"JobStatusDetails,omitempty"`
	Lastcollectedtime interface{} `json:"LastCollectedTime,omitempty"`
	Lastupdatedtime interface{} `json:"LastUpdatedTime,omitempty"`
	Uploaderstatus interface{} `json:"UploaderStatus,omitempty"`
}

// DescribeNotificationConfigurationOutput represents the DescribeNotificationConfigurationOutput schema from the OpenAPI specification
type DescribeNotificationConfigurationOutput struct {
	Notificationconfiguration interface{} `json:"NotificationConfiguration,omitempty"`
}

// StreamInfo represents the StreamInfo schema from the OpenAPI specification
type StreamInfo struct {
	Dataretentioninhours interface{} `json:"DataRetentionInHours,omitempty"`
	Streamarn interface{} `json:"StreamARN,omitempty"`
	Mediatype interface{} `json:"MediaType,omitempty"`
	Version interface{} `json:"Version,omitempty"`
	Streamname interface{} `json:"StreamName,omitempty"`
	Devicename interface{} `json:"DeviceName,omitempty"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
}

// DeleteSignalingChannelOutput represents the DeleteSignalingChannelOutput schema from the OpenAPI specification
type DeleteSignalingChannelOutput struct {
}

// ChannelInfo represents the ChannelInfo schema from the OpenAPI specification
type ChannelInfo struct {
	Channeltype interface{} `json:"ChannelType,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Singlemasterconfiguration interface{} `json:"SingleMasterConfiguration,omitempty"`
	Version interface{} `json:"Version,omitempty"`
	Channelarn interface{} `json:"ChannelARN,omitempty"`
	Channelname interface{} `json:"ChannelName,omitempty"`
	Channelstatus interface{} `json:"ChannelStatus,omitempty"`
}

// EdgeAgentStatus represents the EdgeAgentStatus schema from the OpenAPI specification
type EdgeAgentStatus struct {
	Lastuploaderstatus interface{} `json:"LastUploaderStatus,omitempty"`
	Lastrecorderstatus interface{} `json:"LastRecorderStatus,omitempty"`
}

// DescribeMappedResourceConfigurationInput represents the DescribeMappedResourceConfigurationInput schema from the OpenAPI specification
type DescribeMappedResourceConfigurationInput struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Streamarn interface{} `json:"StreamARN,omitempty"`
	Streamname interface{} `json:"StreamName,omitempty"`
}

// DescribeStreamInput represents the DescribeStreamInput schema from the OpenAPI specification
type DescribeStreamInput struct {
	Streamarn interface{} `json:"StreamARN,omitempty"`
	Streamname interface{} `json:"StreamName,omitempty"`
}

// TagStreamInput represents the TagStreamInput schema from the OpenAPI specification
type TagStreamInput struct {
	Streamarn interface{} `json:"StreamARN,omitempty"`
	Streamname interface{} `json:"StreamName,omitempty"`
	Tags interface{} `json:"Tags"`
}

// ListTagsForStreamInput represents the ListTagsForStreamInput schema from the OpenAPI specification
type ListTagsForStreamInput struct {
	Streamarn interface{} `json:"StreamARN,omitempty"`
	Streamname interface{} `json:"StreamName,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DeleteSignalingChannelInput represents the DeleteSignalingChannelInput schema from the OpenAPI specification
type DeleteSignalingChannelInput struct {
	Channelarn interface{} `json:"ChannelARN"`
	Currentversion interface{} `json:"CurrentVersion,omitempty"`
}

// UpdateDataRetentionInput represents the UpdateDataRetentionInput schema from the OpenAPI specification
type UpdateDataRetentionInput struct {
	Dataretentionchangeinhours interface{} `json:"DataRetentionChangeInHours"`
	Operation interface{} `json:"Operation"`
	Streamarn interface{} `json:"StreamARN,omitempty"`
	Streamname interface{} `json:"StreamName,omitempty"`
	Currentversion interface{} `json:"CurrentVersion"`
}

// DescribeEdgeConfigurationInput represents the DescribeEdgeConfigurationInput schema from the OpenAPI specification
type DescribeEdgeConfigurationInput struct {
	Streamarn interface{} `json:"StreamARN,omitempty"`
	Streamname interface{} `json:"StreamName,omitempty"`
}

// CreateStreamOutput represents the CreateStreamOutput schema from the OpenAPI specification
type CreateStreamOutput struct {
	Streamarn interface{} `json:"StreamARN,omitempty"`
}

// StartEdgeConfigurationUpdateOutput represents the StartEdgeConfigurationUpdateOutput schema from the OpenAPI specification
type StartEdgeConfigurationUpdateOutput struct {
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Edgeconfig interface{} `json:"EdgeConfig,omitempty"`
	Failedstatusdetails interface{} `json:"FailedStatusDetails,omitempty"`
	Lastupdatedtime interface{} `json:"LastUpdatedTime,omitempty"`
	Streamarn interface{} `json:"StreamARN,omitempty"`
	Streamname interface{} `json:"StreamName,omitempty"`
	Syncstatus interface{} `json:"SyncStatus,omitempty"`
}

// DescribeImageGenerationConfigurationInput represents the DescribeImageGenerationConfigurationInput schema from the OpenAPI specification
type DescribeImageGenerationConfigurationInput struct {
	Streamarn interface{} `json:"StreamARN,omitempty"`
	Streamname interface{} `json:"StreamName,omitempty"`
}

// NotificationConfiguration represents the NotificationConfiguration schema from the OpenAPI specification
type NotificationConfiguration struct {
	Destinationconfig interface{} `json:"DestinationConfig"`
	Status interface{} `json:"Status"`
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Key interface{} `json:"Key"`
	Value interface{} `json:"Value"`
}

// GetSignalingChannelEndpointInput represents the GetSignalingChannelEndpointInput schema from the OpenAPI specification
type GetSignalingChannelEndpointInput struct {
	Channelarn interface{} `json:"ChannelARN"`
	Singlemasterchannelendpointconfiguration interface{} `json:"SingleMasterChannelEndpointConfiguration,omitempty"`
}

// CreateSignalingChannelInput represents the CreateSignalingChannelInput schema from the OpenAPI specification
type CreateSignalingChannelInput struct {
	Channelname interface{} `json:"ChannelName"`
	Channeltype interface{} `json:"ChannelType,omitempty"`
	Singlemasterconfiguration interface{} `json:"SingleMasterConfiguration,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// MediaSourceConfig represents the MediaSourceConfig schema from the OpenAPI specification
type MediaSourceConfig struct {
	Mediaurisecretarn interface{} `json:"MediaUriSecretArn"`
	Mediauritype interface{} `json:"MediaUriType"`
}

// ListEdgeAgentConfigurationsInput represents the ListEdgeAgentConfigurationsInput schema from the OpenAPI specification
type ListEdgeAgentConfigurationsInput struct {
	Hubdevicearn interface{} `json:"HubDeviceArn"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// UpdateDataRetentionOutput represents the UpdateDataRetentionOutput schema from the OpenAPI specification
type UpdateDataRetentionOutput struct {
}

// ScheduleConfig represents the ScheduleConfig schema from the OpenAPI specification
type ScheduleConfig struct {
	Durationinseconds interface{} `json:"DurationInSeconds"`
	Scheduleexpression interface{} `json:"ScheduleExpression"`
}

// ListStreamsInput represents the ListStreamsInput schema from the OpenAPI specification
type ListStreamsInput struct {
	Streamnamecondition interface{} `json:"StreamNameCondition,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// StreamNameCondition represents the StreamNameCondition schema from the OpenAPI specification
type StreamNameCondition struct {
	Comparisonoperator interface{} `json:"ComparisonOperator,omitempty"`
	Comparisonvalue interface{} `json:"ComparisonValue,omitempty"`
}

// ListTagsForStreamOutput represents the ListTagsForStreamOutput schema from the OpenAPI specification
type ListTagsForStreamOutput struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// ListStreamsOutput represents the ListStreamsOutput schema from the OpenAPI specification
type ListStreamsOutput struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Streaminfolist interface{} `json:"StreamInfoList,omitempty"`
}

// UpdateStreamOutput represents the UpdateStreamOutput schema from the OpenAPI specification
type UpdateStreamOutput struct {
}

// SingleMasterChannelEndpointConfiguration represents the SingleMasterChannelEndpointConfiguration schema from the OpenAPI specification
type SingleMasterChannelEndpointConfiguration struct {
	Protocols interface{} `json:"Protocols,omitempty"`
	Role interface{} `json:"Role,omitempty"`
}

// UntagResourceOutput represents the UntagResourceOutput schema from the OpenAPI specification
type UntagResourceOutput struct {
}

// TagResourceOutput represents the TagResourceOutput schema from the OpenAPI specification
type TagResourceOutput struct {
}

// ListTagsForResourceOutput represents the ListTagsForResourceOutput schema from the OpenAPI specification
type ListTagsForResourceOutput struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// ListTagsForResourceInput represents the ListTagsForResourceInput schema from the OpenAPI specification
type ListTagsForResourceInput struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Resourcearn interface{} `json:"ResourceARN"`
}

// ImageGenerationConfiguration represents the ImageGenerationConfiguration schema from the OpenAPI specification
type ImageGenerationConfiguration struct {
	Widthpixels interface{} `json:"WidthPixels,omitempty"`
	Destinationconfig interface{} `json:"DestinationConfig"`
	Format interface{} `json:"Format"`
	Formatconfig interface{} `json:"FormatConfig,omitempty"`
	Heightpixels interface{} `json:"HeightPixels,omitempty"`
	Imageselectortype interface{} `json:"ImageSelectorType"`
	Samplinginterval interface{} `json:"SamplingInterval"`
	Status interface{} `json:"Status"`
}

// LastRecorderStatus represents the LastRecorderStatus schema from the OpenAPI specification
type LastRecorderStatus struct {
	Lastcollectedtime interface{} `json:"LastCollectedTime,omitempty"`
	Lastupdatedtime interface{} `json:"LastUpdatedTime,omitempty"`
	Recorderstatus interface{} `json:"RecorderStatus,omitempty"`
	Jobstatusdetails interface{} `json:"JobStatusDetails,omitempty"`
}

// DescribeStreamOutput represents the DescribeStreamOutput schema from the OpenAPI specification
type DescribeStreamOutput struct {
	Streaminfo interface{} `json:"StreamInfo,omitempty"`
}

// UpdateImageGenerationConfigurationInput represents the UpdateImageGenerationConfigurationInput schema from the OpenAPI specification
type UpdateImageGenerationConfigurationInput struct {
	Imagegenerationconfiguration interface{} `json:"ImageGenerationConfiguration,omitempty"`
	Streamarn interface{} `json:"StreamARN,omitempty"`
	Streamname interface{} `json:"StreamName,omitempty"`
}

// DescribeImageGenerationConfigurationOutput represents the DescribeImageGenerationConfigurationOutput schema from the OpenAPI specification
type DescribeImageGenerationConfigurationOutput struct {
	Imagegenerationconfiguration interface{} `json:"ImageGenerationConfiguration,omitempty"`
}

// DescribeMappedResourceConfigurationOutput represents the DescribeMappedResourceConfigurationOutput schema from the OpenAPI specification
type DescribeMappedResourceConfigurationOutput struct {
	Mappedresourceconfigurationlist interface{} `json:"MappedResourceConfigurationList,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ResourceEndpointListItem represents the ResourceEndpointListItem schema from the OpenAPI specification
type ResourceEndpointListItem struct {
	Protocol interface{} `json:"Protocol,omitempty"`
	Resourceendpoint interface{} `json:"ResourceEndpoint,omitempty"`
}

// MappedResourceConfigurationListItem represents the MappedResourceConfigurationListItem schema from the OpenAPI specification
type MappedResourceConfigurationListItem struct {
	Arn interface{} `json:"ARN,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// ImageGenerationDestinationConfig represents the ImageGenerationDestinationConfig schema from the OpenAPI specification
type ImageGenerationDestinationConfig struct {
	Uri interface{} `json:"Uri"`
	Destinationregion interface{} `json:"DestinationRegion"`
}

// CreateStreamInput represents the CreateStreamInput schema from the OpenAPI specification
type CreateStreamInput struct {
	Streamname interface{} `json:"StreamName"`
	Tags interface{} `json:"Tags,omitempty"`
	Dataretentioninhours interface{} `json:"DataRetentionInHours,omitempty"`
	Devicename interface{} `json:"DeviceName,omitempty"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	Mediatype interface{} `json:"MediaType,omitempty"`
}

// LocalSizeConfig represents the LocalSizeConfig schema from the OpenAPI specification
type LocalSizeConfig struct {
	Maxlocalmediasizeinmb interface{} `json:"MaxLocalMediaSizeInMB,omitempty"`
	Strategyonfullsize interface{} `json:"StrategyOnFullSize,omitempty"`
}

// DescribeMediaStorageConfigurationOutput represents the DescribeMediaStorageConfigurationOutput schema from the OpenAPI specification
type DescribeMediaStorageConfigurationOutput struct {
	Mediastorageconfiguration interface{} `json:"MediaStorageConfiguration,omitempty"`
}

// GetDataEndpointInput represents the GetDataEndpointInput schema from the OpenAPI specification
type GetDataEndpointInput struct {
	Apiname interface{} `json:"APIName"`
	Streamarn interface{} `json:"StreamARN,omitempty"`
	Streamname interface{} `json:"StreamName,omitempty"`
}

// DeletionConfig represents the DeletionConfig schema from the OpenAPI specification
type DeletionConfig struct {
	Localsizeconfig interface{} `json:"LocalSizeConfig,omitempty"`
	Deleteafterupload interface{} `json:"DeleteAfterUpload,omitempty"`
	Edgeretentioninhours interface{} `json:"EdgeRetentionInHours,omitempty"`
}

// UntagStreamInput represents the UntagStreamInput schema from the OpenAPI specification
type UntagStreamInput struct {
	Streamarn interface{} `json:"StreamARN,omitempty"`
	Streamname interface{} `json:"StreamName,omitempty"`
	Tagkeylist interface{} `json:"TagKeyList"`
}

// ListEdgeAgentConfigurationsEdgeConfig represents the ListEdgeAgentConfigurationsEdgeConfig schema from the OpenAPI specification
type ListEdgeAgentConfigurationsEdgeConfig struct {
	Syncstatus interface{} `json:"SyncStatus,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Edgeconfig EdgeConfig `json:"EdgeConfig,omitempty"` // A description of the stream's edge configuration that will be used to sync with the Edge Agent IoT Greengrass component. The Edge Agent component will run on an IoT Hub Device setup at your premise.
	Failedstatusdetails interface{} `json:"FailedStatusDetails,omitempty"`
	Lastupdatedtime interface{} `json:"LastUpdatedTime,omitempty"`
	Streamarn interface{} `json:"StreamARN,omitempty"`
	Streamname interface{} `json:"StreamName,omitempty"`
}

// NotificationDestinationConfig represents the NotificationDestinationConfig schema from the OpenAPI specification
type NotificationDestinationConfig struct {
	Uri interface{} `json:"Uri"`
}

// EdgeConfig represents the EdgeConfig schema from the OpenAPI specification
type EdgeConfig struct {
	Deletionconfig interface{} `json:"DeletionConfig,omitempty"`
	Hubdevicearn interface{} `json:"HubDeviceArn"`
	Recorderconfig interface{} `json:"RecorderConfig"`
	Uploaderconfig interface{} `json:"UploaderConfig,omitempty"`
}

// DeleteStreamOutput represents the DeleteStreamOutput schema from the OpenAPI specification
type DeleteStreamOutput struct {
}

// ListSignalingChannelsInput represents the ListSignalingChannelsInput schema from the OpenAPI specification
type ListSignalingChannelsInput struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Channelnamecondition interface{} `json:"ChannelNameCondition,omitempty"`
}

// ResourceTags represents the ResourceTags schema from the OpenAPI specification
type ResourceTags struct {
}

// UntagResourceInput represents the UntagResourceInput schema from the OpenAPI specification
type UntagResourceInput struct {
	Tagkeylist interface{} `json:"TagKeyList"`
	Resourcearn interface{} `json:"ResourceARN"`
}

// CreateSignalingChannelOutput represents the CreateSignalingChannelOutput schema from the OpenAPI specification
type CreateSignalingChannelOutput struct {
	Channelarn interface{} `json:"ChannelARN,omitempty"`
}

// GetSignalingChannelEndpointOutput represents the GetSignalingChannelEndpointOutput schema from the OpenAPI specification
type GetSignalingChannelEndpointOutput struct {
	Resourceendpointlist interface{} `json:"ResourceEndpointList,omitempty"`
}

// TagResourceInput represents the TagResourceInput schema from the OpenAPI specification
type TagResourceInput struct {
	Resourcearn interface{} `json:"ResourceARN"`
	Tags interface{} `json:"Tags"`
}

// UpdateImageGenerationConfigurationOutput represents the UpdateImageGenerationConfigurationOutput schema from the OpenAPI specification
type UpdateImageGenerationConfigurationOutput struct {
}

// ListEdgeAgentConfigurationsOutput represents the ListEdgeAgentConfigurationsOutput schema from the OpenAPI specification
type ListEdgeAgentConfigurationsOutput struct {
	Edgeconfigs interface{} `json:"EdgeConfigs,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// RecorderConfig represents the RecorderConfig schema from the OpenAPI specification
type RecorderConfig struct {
	Mediasourceconfig interface{} `json:"MediaSourceConfig"`
	Scheduleconfig interface{} `json:"ScheduleConfig,omitempty"`
}

// UploaderConfig represents the UploaderConfig schema from the OpenAPI specification
type UploaderConfig struct {
	Scheduleconfig interface{} `json:"ScheduleConfig"`
}

// StartEdgeConfigurationUpdateInput represents the StartEdgeConfigurationUpdateInput schema from the OpenAPI specification
type StartEdgeConfigurationUpdateInput struct {
	Streamname interface{} `json:"StreamName,omitempty"`
	Edgeconfig interface{} `json:"EdgeConfig"`
	Streamarn interface{} `json:"StreamARN,omitempty"`
}

// TagStreamOutput represents the TagStreamOutput schema from the OpenAPI specification
type TagStreamOutput struct {
}

// DeleteEdgeConfigurationInput represents the DeleteEdgeConfigurationInput schema from the OpenAPI specification
type DeleteEdgeConfigurationInput struct {
	Streamarn interface{} `json:"StreamARN,omitempty"`
	Streamname interface{} `json:"StreamName,omitempty"`
}

// SingleMasterConfiguration represents the SingleMasterConfiguration schema from the OpenAPI specification
type SingleMasterConfiguration struct {
	Messagettlseconds interface{} `json:"MessageTtlSeconds,omitempty"`
}

// UpdateMediaStorageConfigurationOutput represents the UpdateMediaStorageConfigurationOutput schema from the OpenAPI specification
type UpdateMediaStorageConfigurationOutput struct {
}

// UpdateNotificationConfigurationInput represents the UpdateNotificationConfigurationInput schema from the OpenAPI specification
type UpdateNotificationConfigurationInput struct {
	Streamarn interface{} `json:"StreamARN,omitempty"`
	Streamname interface{} `json:"StreamName,omitempty"`
	Notificationconfiguration interface{} `json:"NotificationConfiguration,omitempty"`
}

// ListSignalingChannelsOutput represents the ListSignalingChannelsOutput schema from the OpenAPI specification
type ListSignalingChannelsOutput struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Channelinfolist interface{} `json:"ChannelInfoList,omitempty"`
}
