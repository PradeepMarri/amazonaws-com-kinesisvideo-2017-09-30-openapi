package main

import (
	"github.com/amazon-kinesis-video-streams/mcp-server/config"
	"github.com/amazon-kinesis-video-streams/mcp-server/models"
	tools_createsignalingchannel "github.com/amazon-kinesis-video-streams/mcp-server/tools/createsignalingchannel"
	tools_deletesignalingchannel "github.com/amazon-kinesis-video-streams/mcp-server/tools/deletesignalingchannel"
	tools_startedgeconfigurationupdate "github.com/amazon-kinesis-video-streams/mcp-server/tools/startedgeconfigurationupdate"
	tools_tagresource "github.com/amazon-kinesis-video-streams/mcp-server/tools/tagresource"
	tools_updatenotificationconfiguration "github.com/amazon-kinesis-video-streams/mcp-server/tools/updatenotificationconfiguration"
	tools_tagstream "github.com/amazon-kinesis-video-streams/mcp-server/tools/tagstream"
	tools_createstream "github.com/amazon-kinesis-video-streams/mcp-server/tools/createstream"
	tools_describeimagegenerationconfiguration "github.com/amazon-kinesis-video-streams/mcp-server/tools/describeimagegenerationconfiguration"
	tools_getsignalingchannelendpoint "github.com/amazon-kinesis-video-streams/mcp-server/tools/getsignalingchannelendpoint"
	tools_untagstream "github.com/amazon-kinesis-video-streams/mcp-server/tools/untagstream"
	tools_describemediastorageconfiguration "github.com/amazon-kinesis-video-streams/mcp-server/tools/describemediastorageconfiguration"
	tools_describeedgeconfiguration "github.com/amazon-kinesis-video-streams/mcp-server/tools/describeedgeconfiguration"
	tools_describesignalingchannel "github.com/amazon-kinesis-video-streams/mcp-server/tools/describesignalingchannel"
	tools_listtagsforstream "github.com/amazon-kinesis-video-streams/mcp-server/tools/listtagsforstream"
	tools_describestream "github.com/amazon-kinesis-video-streams/mcp-server/tools/describestream"
	tools_listtagsforresource "github.com/amazon-kinesis-video-streams/mcp-server/tools/listtagsforresource"
	tools_getdataendpoint "github.com/amazon-kinesis-video-streams/mcp-server/tools/getdataendpoint"
	tools_listsignalingchannels "github.com/amazon-kinesis-video-streams/mcp-server/tools/listsignalingchannels"
	tools_updatemediastorageconfiguration "github.com/amazon-kinesis-video-streams/mcp-server/tools/updatemediastorageconfiguration"
	tools_updatedataretention "github.com/amazon-kinesis-video-streams/mcp-server/tools/updatedataretention"
	tools_liststreams "github.com/amazon-kinesis-video-streams/mcp-server/tools/liststreams"
	tools_updatestream "github.com/amazon-kinesis-video-streams/mcp-server/tools/updatestream"
	tools_deleteedgeconfiguration "github.com/amazon-kinesis-video-streams/mcp-server/tools/deleteedgeconfiguration"
	tools_untagresource "github.com/amazon-kinesis-video-streams/mcp-server/tools/untagresource"
	tools_updateimagegenerationconfiguration "github.com/amazon-kinesis-video-streams/mcp-server/tools/updateimagegenerationconfiguration"
	tools_updatesignalingchannel "github.com/amazon-kinesis-video-streams/mcp-server/tools/updatesignalingchannel"
	tools_listedgeagentconfigurations "github.com/amazon-kinesis-video-streams/mcp-server/tools/listedgeagentconfigurations"
	tools_deletestream "github.com/amazon-kinesis-video-streams/mcp-server/tools/deletestream"
	tools_describenotificationconfiguration "github.com/amazon-kinesis-video-streams/mcp-server/tools/describenotificationconfiguration"
	tools_describemappedresourceconfiguration "github.com/amazon-kinesis-video-streams/mcp-server/tools/describemappedresourceconfiguration"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_createsignalingchannel.CreateCreatesignalingchannelTool(cfg),
		tools_deletesignalingchannel.CreateDeletesignalingchannelTool(cfg),
		tools_startedgeconfigurationupdate.CreateStartedgeconfigurationupdateTool(cfg),
		tools_tagresource.CreateTagresourceTool(cfg),
		tools_updatenotificationconfiguration.CreateUpdatenotificationconfigurationTool(cfg),
		tools_tagstream.CreateTagstreamTool(cfg),
		tools_createstream.CreateCreatestreamTool(cfg),
		tools_describeimagegenerationconfiguration.CreateDescribeimagegenerationconfigurationTool(cfg),
		tools_getsignalingchannelendpoint.CreateGetsignalingchannelendpointTool(cfg),
		tools_untagstream.CreateUntagstreamTool(cfg),
		tools_describemediastorageconfiguration.CreateDescribemediastorageconfigurationTool(cfg),
		tools_describeedgeconfiguration.CreateDescribeedgeconfigurationTool(cfg),
		tools_describesignalingchannel.CreateDescribesignalingchannelTool(cfg),
		tools_listtagsforstream.CreateListtagsforstreamTool(cfg),
		tools_describestream.CreateDescribestreamTool(cfg),
		tools_listtagsforresource.CreateListtagsforresourceTool(cfg),
		tools_getdataendpoint.CreateGetdataendpointTool(cfg),
		tools_listsignalingchannels.CreateListsignalingchannelsTool(cfg),
		tools_updatemediastorageconfiguration.CreateUpdatemediastorageconfigurationTool(cfg),
		tools_updatedataretention.CreateUpdatedataretentionTool(cfg),
		tools_liststreams.CreateListstreamsTool(cfg),
		tools_updatestream.CreateUpdatestreamTool(cfg),
		tools_deleteedgeconfiguration.CreateDeleteedgeconfigurationTool(cfg),
		tools_untagresource.CreateUntagresourceTool(cfg),
		tools_updateimagegenerationconfiguration.CreateUpdateimagegenerationconfigurationTool(cfg),
		tools_updatesignalingchannel.CreateUpdatesignalingchannelTool(cfg),
		tools_listedgeagentconfigurations.CreateListedgeagentconfigurationsTool(cfg),
		tools_deletestream.CreateDeletestreamTool(cfg),
		tools_describenotificationconfiguration.CreateDescribenotificationconfigurationTool(cfg),
		tools_describemappedresourceconfiguration.CreateDescribemappedresourceconfigurationTool(cfg),
	}
}
