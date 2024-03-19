package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	SetupAsset_checkRoutes "github.com/puvadon-artmit/gofiber-template/api/asset_check"
	SetupAsset_RoundRoutes "github.com/puvadon-artmit/gofiber-template/api/asset_round"
	AssetsRoutes "github.com/puvadon-artmit/gofiber-template/api/assets"
	Assets_StoryRoutes "github.com/puvadon-artmit/gofiber-template/api/assets_story"
	authRoutes "github.com/puvadon-artmit/gofiber-template/api/auth"
	SetupAutoclik_dataRoutes "github.com/puvadon-artmit/gofiber-template/api/autoclik_data"
	branchRoutes "github.com/puvadon-artmit/gofiber-template/api/branch"
	SetupBranch_AutoclikRoutes "github.com/puvadon-artmit/gofiber-template/api/branch_autoclik"
	SetupBranch_StoryRoutes "github.com/puvadon-artmit/gofiber-template/api/branch_story"
	categoryRoutes "github.com/puvadon-artmit/gofiber-template/api/category"
	SetupCategory_StoryRoutes "github.com/puvadon-artmit/gofiber-template/api/category_story"
	SetupCount_AutoclikRoutes "github.com/puvadon-artmit/gofiber-template/api/count_autoclik"
	Setupcount_categoryRoutes "github.com/puvadon-artmit/gofiber-template/api/count_category"
	Setupcount_main_categoryRoutes "github.com/puvadon-artmit/gofiber-template/api/count_main_category"
	groundRouter "github.com/puvadon-artmit/gofiber-template/api/ground"
	Ground_storyRouter "github.com/puvadon-artmit/gofiber-template/api/ground_story"
	groupRoutes "github.com/puvadon-artmit/gofiber-template/api/group"
	SetupGroup_StoryRoutes "github.com/puvadon-artmit/gofiber-template/api/group_story"
	asset_countRoutes "github.com/puvadon-artmit/gofiber-template/api/inspection-assets"
	Inspection_storyRoutes "github.com/puvadon-artmit/gofiber-template/api/inspection_story"
	item_modelRoutes "github.com/puvadon-artmit/gofiber-template/api/item_model"
	SetupLocationRoutes "github.com/puvadon-artmit/gofiber-template/api/location"
	SetupLocation_storyRoutes "github.com/puvadon-artmit/gofiber-template/api/location_story"
	SetupMain_branchRoutes "github.com/puvadon-artmit/gofiber-template/api/main-branch"
	SetupMain_Branch_StoryRoutes "github.com/puvadon-artmit/gofiber-template/api/main_branch_story"
	SetupMain_CategoryRoutes "github.com/puvadon-artmit/gofiber-template/api/main_category"
	SetupMaliwan_dataRoutes "github.com/puvadon-artmit/gofiber-template/api/maliwan_data"
	manufacturerRoutes "github.com/puvadon-artmit/gofiber-template/api/manufacturer"
	permissionRoutes "github.com/puvadon-artmit/gofiber-template/api/permission"
	SetupPermission_componentyRoutes "github.com/puvadon-artmit/gofiber-template/api/permissioncomponent"
	SetupPermissionGroupRoutes "github.com/puvadon-artmit/gofiber-template/api/permissiongroup"
	SetupPhotos_checkRoutes "github.com/puvadon-artmit/gofiber-template/api/photos_check"
	responsibleRoutes "github.com/puvadon-artmit/gofiber-template/api/responsible"
	SetupResponsible_StoryRoutes "github.com/puvadon-artmit/gofiber-template/api/responsible-story"
	roleRoutes "github.com/puvadon-artmit/gofiber-template/api/role"
	SetupRound_CountRoutes "github.com/puvadon-artmit/gofiber-template/api/round_count"
	SetupRound_Count_StoryRoutes "github.com/puvadon-artmit/gofiber-template/api/round_count_story"
	SetupSignatureRoutes "github.com/puvadon-artmit/gofiber-template/api/sign-confirmation"
	statusRoutes "github.com/puvadon-artmit/gofiber-template/api/status"
	storyRoutes "github.com/puvadon-artmit/gofiber-template/api/story"
	typeRoutes "github.com/puvadon-artmit/gofiber-template/api/type_things"
	typeplanRoutes "github.com/puvadon-artmit/gofiber-template/api/typeplan"
	typeplan_storyRoutes "github.com/puvadon-artmit/gofiber-template/api/typeplan_story"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("api", logger.New())

	// authRoutes.SetupAuthRoutes(api)
	authRoutes.SetupAuthRoutes(api)
	roleRoutes.SetupRoleRoutes(api)
	permissionRoutes.SetupPermissionRoutes(api)
	categoryRoutes.SetupCategoryRoutes(api)
	branchRoutes.SetupBranchRoutes(api)
	groupRoutes.SetupGroupRoutes(api)
	statusRoutes.SetupStatusRoutes(api)
	typeRoutes.SetupTypeRoutes(api)
	manufacturerRoutes.SetupManufacturerRoutes(api)
	item_modelRoutes.SetupItem_modelRoutes(api)
	storyRoutes.SetupStoryRoutes(api)
	groundRouter.SetupGroundRoutes(api)
	Ground_storyRouter.SetupGround_storyRoutes(api)
	responsibleRoutes.SetupResponsibleRoutes(api)
	AssetsRoutes.SetupAssetsRoutes(api)
	Assets_StoryRoutes.SetupAssets_StoryRoutes(api)
	typeplanRoutes.SetupTypeplanRoutes(api)
	typeplan_storyRoutes.SetupTypeplan_StoryRoutes(api)
	asset_countRoutes.SetupAsset_countRoutes(api)
	Inspection_storyRoutes.SetupInspection_StoryRoutes(api)
	SetupAsset_checkRoutes.SetupAsset_checkRoutes(api)
	SetupLocationRoutes.SetupLocationRoutes(api)
	SetupLocation_storyRoutes.SetupLocation_storyRoutes(api)
	SetupSignatureRoutes.SetupSignatureRoutes(api)
	SetupMain_branchRoutes.SetupMain_branchRoutes(api)
	SetupBranch_StoryRoutes.SetupBranch_StoryRoutes(api)
	SetupMain_Branch_StoryRoutes.SetupMain_Branch_StoryRoutes(api)
	SetupGroup_StoryRoutes.SetupGroup_StoryRoutes(api)
	SetupResponsible_StoryRoutes.SetupResponsible_StoryRoutes(api)
	SetupAutoclik_dataRoutes.SetupAutoclik_dataRoutes(api)
	SetupMaliwan_dataRoutes.SetupMaliwan_dataRoutes(api)
	SetupRound_CountRoutes.SetupRound_CountRoutes(api)
	SetupRound_Count_StoryRoutes.SetupRound_Count_StoryRoutes(api)
	SetupAsset_RoundRoutes.SetupAsset_RoundRoutes(api)
	SetupMain_CategoryRoutes.SetupMain_CategoryRoutes(api)
	SetupBranch_AutoclikRoutes.SetupBranch_AutoclikRoutes(api)
	SetupCount_AutoclikRoutes.SetupCount_AutoclikRoutes(api)
	SetupPhotos_checkRoutes.SetupPhotos_checkRoutes(api)
	Setupcount_categoryRoutes.Setupcount_categoryRoutes(api)
	Setupcount_main_categoryRoutes.Setupcount_main_categoryRoutes(api)
	Setupcount_categoryRoutes.Setupcount_categoryRoutes(api)
	Setupcount_main_categoryRoutes.Setupcount_main_categoryRoutes(api)
	SetupPermission_componentyRoutes.SetupPermission_componentyRoutes(api)
	SetupPermissionGroupRoutes.SetupPermissionGroupRoutes(api)
	SetupCategory_StoryRoutes.SetupCategory_StoryRoutes(api)

}
