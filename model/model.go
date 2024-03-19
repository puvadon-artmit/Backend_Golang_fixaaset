package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserID       string         `gorm:"type:uuid;primaryKey" json:"user_id"`
	Firstname    string         `json:"firstname" validate:"required"`
	Lastname     string         `json:"lastname" validate:"required"`
	Username     string         `json:"username" validate:"required"`
	Password     string         `json:"password" validate:"required"`
	Status       bool           `json:"status" gorm:"default:true"`
	EmployeeCode string         `json:"employee_code"`
	CreatedAt    *time.Time     `json:"created_at"`
	UpdatedAt    *time.Time     `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Roles        []Role         `gorm:"many2many:user_roles;"`
}

// เอาไว้เก็บ UserID และ RoleID เพื่อ ให้ 1 user มีได้หลาย Role
type UserRoles struct {
	gorm.Model
	UserID string `gorm:"type:uuid"`
	RoleID string `gorm:"type:uuid"`
}

// เอาไว้เก็บ Role ตำแหน่งต่างๆ
type Role struct {
	RoleID          string            `gorm:"type:uuid;primaryKey" json:"role_id"`
	RoleName        string            `json:"role_name" validate:"required"`
	RoleDisplayName string            `json:"role_display_name" validate:"required"`
	RoleDescription string            `json:"role_description" validate:"required"`
	RoleCode        string            `json:"role_code"`
	Status          bool              `json:"status" gorm:"default:true"`
	CreatedAt       *time.Time        `json:"created_at"`
	UpdatedAt       *time.Time        `json:"updated_at"`
	Users           []User            `gorm:"many2many:user_roles;"`
	PermissionGroup []PermissionGroup `gorm:"foreignKey:RoleID"`
}

// เอาไว้เก็บ การทำงานหลักๆ ว่า สามารถทำอะไรได้
type Permission struct {
	PermissionID          string                `gorm:"type:uuid;primaryKey" json:"permission_id"`
	PermissionName        string                `json:"permission_name" validate:"required"`
	PermissionDisplayName string                `json:"permission_display_name" validate:"required"`
	PermissionDescription string                `json:"permission_description" validate:"required"`
	Status                bool                  `json:"status" gorm:"default:true"`
	CreatedAt             *time.Time            `json:"created_at"`
	UpdatedAt             *time.Time            `json:"updated_at"`
	PermissionComponent   []PermissionComponent `gorm:"foreignKey:PermissionID"`
	PermissionGroup       []PermissionGroup     `gorm:"foreignKey:PermissionID"`
}

// เอาไว้เก็บ การทำงานแต่ล่ะ Component โดยละเอียดยิบ
type PermissionComponent struct {
	PermissionComponentID          string            `gorm:"type:uuid;primaryKey" json:"permission_component_id"`
	PermissionComponentName        string            `json:"permission_component_name"`
	PermissionComponentDisplayName string            `json:"permission_component_display_name"`
	PermissionComponentDescription string            `json:"permission_component_description"`
	Status                         bool              `json:"status" gorm:"default:true"`
	CreatedAt                      *time.Time        `json:"created_at"`
	UpdatedAt                      *time.Time        `json:"updated_at"`
	PermissionID                   string            `json:"permission_id"`
	Permission                     Permission        `validate:"-"`
	PermissionGroup                []PermissionGroup `gorm:"foreignKey:PermissionComponentID"`
}

// เอาไว้เก็บ การทำงานโดยรวมทุกอย่าง เพื่อส่งไปให้เลือก Role
type PermissionGroup struct {
	PermissionGroupID     string              `gorm:"type:uuid;primaryKey" json:"permission_group_id"`
	Activate              bool                `json:"avtivates"`
	CreatedAt             *time.Time          `json:"created_at"`
	UpdatedAt             *time.Time          `json:"updated_at"`
	RoleID                string              `json:"role_id"`
	PermissionID          string              `json:"permission_id"`
	PermissionComponentID string              `json:"permission_component_id"`
	Role                  Role                `validate:"-"`
	Permission            Permission          `validate:"-"`
	PermissionComponent   PermissionComponent `validate:"-"`
}

type Main_Category struct {
	Main_Category_ID string `gorm:"type:uuid;primaryKey" json:"main_category_id"`
	Code_ID          string `json:"code_id"`
	Main_name        string `json:"main_name"`
}

type Category struct {
	CategoryID       string         `gorm:"type:uuid;primaryKey" json:"category_id"`
	CategoryName     string         `json:"category_name"`
	CategoryImage    string         `json:"category_image" validate:"required"`
	CreatedAt        *time.Time     `json:"created_at"`
	UpdatedAt        *time.Time     `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	UserID           string         `json:"user_id"`
	User             User
	Main_Category_ID string        `json:"main_category_id"`
	Main_Category    Main_Category `validate:"-"`
}

type Category_Story struct {
	Category_Story_ID    string   `gorm:"type:uuid;primaryKey" json:"category_story_id"`
	Category_Story_Name  string   `json:"category_story_name"`
	Category_Story_Dtail string   `json:"category_story_detail"`
	UserID               string   `json:"user_id"`
	User                 User     `validate:"-"`
	CategoryID           string   `json:"category_id"`
	Category             Category `validate:"-"`
}

type Branch struct {
	BranchID     string         `gorm:"type:uuid;primaryKey" json:"id_branch"`
	ZipCode      *string        `json:"zip_code"`
	County       *string        `json:"county"`
	Province     *string        `json:"province"`
	Comment      *string        `json:"comment"`
	Building     *string        `json:"building"`
	Address      *string        `json:"address"`
	Town         *string        `json:"town"`
	RoomNumber   *string        `json:"room_number"`
	CreatedAt    *time.Time     `json:"created_at"`
	UpdatedAt    *time.Time     `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	UserID       string         `json:"user_id"`
	GroupID      string         `json:"group_id"`
	Group        Group          `validate:"-"`
	MainBranchID string         `json:"main_branch_id"`
	Main_branch  Main_branch    `validate:"-"`
}

type Branch_Story struct {
	BranchStoryID      string         `gorm:"type:uuid;primaryKey" json:"branch_story_id"`
	BranchStoryName    string         `json:"branch_story_name"`
	BranchStoryDetails string         `json:"branch_story_details"`
	CreatedAt          *time.Time     `json:"created_at"`
	UpdatedAt          *time.Time     `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"index"`
	UserID             string         `json:"user_id"`
	User               User           `validate:"-"`
	BranchID           string         `json:"branch_id"`
	Branch             Branch         `validate:"-"`
}

type Main_branch struct {
	MainBranchID                string  `gorm:"type:uuid;primaryKey" json:"main_branch_id"`
	MainBranchName              *string `json:"main_branch_name"`
	Company_name                *string `json:"company_name"`
	Company_branch_name         *string `json:"company_branch_name"`
	Company_branch_name_en      *string `json:"company_branch_name_en"`
	Company_branch_no           *string `json:"company_branch_no"`
	Company_address             *string `json:"company_address"`
	Company_taxid               *string `json:"company_taxid"`
	Company_header              *string `json:"company_header"`
	BranchRef1                  *string `json:"branchRef1"`
	BranchRef2                  *string `json:"branchRef2"`
	TokenLineCreditNotification *string `json:"tokenLineCreditNotification"`
	Navurl                      *string `json:"navurl"`
	Parent                      int64   `json:"parent"`
}

type Main_Branch_Story struct {
	MainBranchStoryID      string         `gorm:"type:uuid;primaryKey" json:"main_branch_story_id"`
	MainBranchStoryName    string         `json:"main_branch_story_name"`
	MainBranchStoryDetails string         `json:"main_branch_story_details"`
	CreatedAt              *time.Time     `json:"created_at"`
	UpdatedAt              *time.Time     `json:"updated_at"`
	DeletedAt              gorm.DeletedAt `gorm:"index"`
	UserID                 string         `json:"user_id"`
	User                   User           `validate:"-"`
	MainBranchID           string         `json:"main_branch_id"`
	MainBranch             Main_branch    `validate:"-"`
}

type Group struct {
	GroupID   string         `gorm:"type:uuid;primaryKey" json:"group_id"`
	GroupName *string        `json:"name_group" validate:"required"`
	Comment   *string        `json:"comment"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserID    string         `json:"user_id"`
}

type Group_Story struct {
	GroupStoryID      string         `gorm:"type:uuid;primaryKey" json:"group_story_id"`
	GroupStoryName    string         `json:"group_story_name"`
	GroupStoryDetails *string        `json:"group_story_details"`
	CreatedAt         *time.Time     `json:"created_at"`
	UpdatedAt         *time.Time     `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index"`
	UserID            string         `json:"user_id"`
	User              User           `validate:"-"`
	GroupID           string         `json:"group_id"`
	Group             Group          `validate:"-"`
}

type Status struct {
	StatusID   string         `gorm:"type:uuid;primaryKey" json:"status_id"`
	StatusName *string        `json:"status_name" validate:"required"`
	Comment    *string        `json:"comment"`
	CreatedAt  *time.Time     `json:"created_at"`
	UpdatedAt  *time.Time     `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	UserID     string         `json:"user_id"`
}

type Type_things struct {
	TypeID     string         `gorm:"type:uuid;primaryKey" json:"type_id"`
	TypeName   *string        `json:"type_name" validate:"required"`
	Comment    *string        `json:"comment"`
	CreatedAt  *time.Time     `json:"created_at"`
	UpdatedAt  *time.Time     `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	UserID     string         `json:"user_id"`
	CategoryID string         `json:"category_id"`
	Category   Category       `validate:"-"`
}

type Manufacturer struct {
	ManufacturerID   string         `gorm:"type:uuid;primaryKey" json:"manufacturer_id"`
	ManufacturerName *string        `json:"manufacturer_name" validate:"required"`
	Comment          *string        `json:"comment"`
	CreatedAt        *time.Time     `json:"created_at"`
	UpdatedAt        *time.Time     `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	UserID           string         `json:"user_id"`
}

type Item_model struct {
	ItemModelID   string         `gorm:"type:uuid;primaryKey" json:"item_model_id"`
	ItemModelName *string        `json:"item_model_name" validate:"required"`
	Comment       *string        `json:"comment"`
	ProductNumber *int           `json:"product_number"`
	Weight        *string        `json:"weight"`
	RequiredUnits *int           `json:"required_units"`
	Frontpicture  *string        `json:"frontpicture"`
	CreatedAt     *time.Time     `json:"created_at"`
	UpdatedAt     *time.Time     `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	UserID        string         `json:"user_id"`
	TypeID        string         `json:"type_id"`
	Type          Type_things    `validate:"-"`
}

type Story struct {
	StoryID     string         `gorm:"type:uuid;primaryKey" json:"story_id"`
	StoryName   *string        `json:"story_name" validate:"required"`
	Username    *string        `json:"username"`
	Details     *string        `json:"details"`
	CreatedAt   *time.Time     `json:"created_at"`
	UpdatedAt   *time.Time     `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	UserID      string         `json:"user_id"`
	ItemModelID string         `json:"item_model_id"`
	Item_model  Item_model     `validate:"-"`
}

const thaiDateFormat = "2006-01-02 15:04:05"

func (s *Story) FormatThaiTime() {
	if s.CreatedAt != nil {
		thaiTime := s.CreatedAt.Format(thaiDateFormat)
		createdAt, _ := time.Parse(thaiDateFormat, thaiTime)
		s.CreatedAt = &createdAt
	}
}

type Ground struct {
	GroundID      string         `gorm:"type:uuid;primaryKey" json:"ground_id"`
	GroundName    *string        `json:"ground_name"`
	Location_code *string        `json:"location_code"`
	Building      *string        `json:"building"`
	Floor         *string        `json:"floor"`
	Room          *string        `json:"room"`
	GroundImage   *string        `json:"ground_image"`
	Comment       *string        `json:"comment"`
	CreatedAt     *time.Time     `json:"created_at"`
	UpdatedAt     *time.Time     `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	UserID        string         `json:"user_id"`
}

type Ground_Story struct {
	Ground_storyID   string         `gorm:"type:uuid;primaryKey" json:"ground_story_id"`
	Ground_StoryName *string        `json:"groundstory_name"`
	Ground_Username  *string        `json:"ground_username"`
	Ground_Details   *string        `json:"ground_details"`
	CreatedAt        *time.Time     `json:"created_at"`
	UpdatedAt        *time.Time     `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	UserID           string         `json:"user_id"`
	GroundID         string         `json:"ground_id"`
	Ground           Ground         `validate:"-"`
}

type Assets_Story struct {
	Assets_storyID   string         `gorm:"type:uuid;primaryKey" json:"assets_story_id"`
	Assets_StoryName *string        `json:"assetsstory_name"`
	Assets_Username  *string        `json:"assets_username"`
	Assets_Details   *string        `json:"assets_details"`
	CreatedAt        *time.Time     `json:"created_at"`
	UpdatedAt        *time.Time     `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	UserID           string         `json:"user_id"`
	AssetsID         string         `json:"assets_id"`
	Assets           Assets         `validate:"-"`
}

type Responsible struct {
	ResponsibleID   string         `gorm:"type:uuid;primaryKey" json:"responsible_id"`
	ResponsibleName *string        `json:"responsible_name"`
	EmployeeCode    *string        `json:"employee_code"`
	Comment         *string        `json:"comment"`
	CreatedAt       *time.Time     `json:"created_at"`
	UpdatedAt       *time.Time     `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	UserID          string         `json:"user_id"`
	GroupID         string         `json:"group_id"`
	Group           Group          `validate:"-"`
}

type Responsible_Story struct {
	ResponsibleStoryID       string         `gorm:"type:uuid;primaryKey" json:"responsible_story_id"`
	Responsible_StoryName    *string        `json:"responsible_story_name"`
	Responsible_StoryDetails *string        `json:"responsible_story_details"`
	CreatedAt                *time.Time     `json:"created_at"`
	UpdatedAt                *time.Time     `json:"updated_at"`
	DeletedAt                gorm.DeletedAt `gorm:"index"`
	UserID                   string         `json:"user_id"`
	User                     User           `validate:"-"`
	ResponsibleID            string         `json:"responsible_id"`
	Responsible              Responsible    `validate:"-"`
}

type Assets struct {
	AssetsID       string         `gorm:"type:uuid;primaryKey" json:"assets_id"`
	Model_name     string         `json:"model_name"`
	Manufacturer   string         `json:"manufacturer"`
	Serial_Code    string         `json:"serial_code"`
	Type           string         `json:"type"`
	Model          string         `json:"model"`
	Branch         string         `json:"branch"`
	Username       string         `json:"username"`
	Property_code  string         `json:"property_code"`
	Status         string         `json:"status"`
	Group_hardware string         `json:"group_hardware"`
	Group          string         `json:"group"`
	User_hardware  string         `json:"user_hardware"`
	Phone_number   string         `json:"phone_number"`
	Posting_group  string         `json:"posting_group"`
	Latest_time    time.Time      `json:"latest_time"`
	ResponsibleID  string         `json:"responsible_id"`
	Responsible    Responsible    `validate:"-"`
	Comment1       string         `json:"comment1"`
	Comment2       string         `json:"comment2"`
	Comment3       string         `json:"comment3"`
	CreatedAt      *time.Time     `json:"created_at"`
	UpdatedAt      *time.Time     `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	UserID         string         `json:"user_id"`
	User           User           `validate:"-"`
	CategoryID     string         `json:"category_id"`
	Category       Category       `validate:"-"`
	ItemModelID    string         `json:"item_model_id"`
	Item_model     Item_model     `validate:"-"`
	GroundID       string         `json:"ground_id"`
	Ground         Ground         `validate:"-"`
}

type Asset_check struct {
	Asset_checkID    string         `gorm:"type:uuid;primaryKey" json:"asset_check_id"`
	Additional_notes *string        `json:"additional_notes"`
	CreatedAt        *time.Time     `json:"created_at"`
	UpdatedAt        *time.Time     `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	Property_code    *string        `json:"property_code"`
	StatusAsset      *string        `json:"statusasset"`
	Round_CountID    uuid.UUID      `gorm:"type:uuid" json:"round_count_id"`
	Round_Count      Round_Count    `validate:"-"`
	Photos_check     []Photos_check `gorm:"many2many:photocheck;" json:"photos_check"`
	UserID           string         `json:"user_id"`
	User             User           `validate:"-"`
}

type PhotoCheck struct {
	gorm.Model
	Asset_checkID  string `gorm:"type:uuid"`
	Photos_checkID string `gorm:"type:uuid"`
}

type Photos_check struct {
	Photos_checkID string `gorm:"type:uuid;primaryKey" json:"photos_check_id"`
	Photos         string `json:"photos"`
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
	DeletedAt      gorm.DeletedAt
}

type Asset_count struct {
	Asset_countID string         `gorm:"type:uuid;primaryKey" json:"asset_count_id"`
	Plan_Code     string         `json:"plan_code"`
	Plan_Name     *string        `json:"plan_name"`
	TypeplanName  *string        `json:"typeplan_name"`
	Project_name  *string        `json:"project_name"`
	Comment       *string        `json:"comment"`
	Plan_start    *string        `json:"plan_start"`
	Plan_end      *string        `json:"plan_end"`
	Status        *string        `json:"status"`
	CreatedAt     *time.Time     `json:"created_at"`
	UpdatedAt     *time.Time     `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	UserID        string         `json:"user_id"`
	User          User           `validate:"-"`
	Group         string         `json:"group"`
	MainBranchID  string         `json:"main_branch_id"`
	Main_branch   Main_branch    `validate:"-"`
	// Main_Category_ID *string        `json:"main_category_id"`
	// CategoryID       *string        `json:"category_id"`
}

type Asset_count_Main_Category struct {
	Asset_count_Main_CategoryID string        `gorm:"type:uuid;primaryKey" json:"asset_count_main_category_id"`
	Main_Category_ID            *string       `json:"main_category_id"`
	Main_Category               Main_Category `validate:"-"`
	Asset_countID               string        `json:"asset_count_id"`
	Asset_count                 Asset_count   `validate:"-"`
}

type Asset_count_Category struct {
	Asset_count_CategoryID string      `gorm:"type:uuid;primaryKey" json:"asset_count_category_id"`
	Category_ID            *string     `json:"category_id"`
	Category               Category    `validate:"-"`
	Asset_countID          string      `json:"asset_count_id"`
	Asset_count            Asset_count `validate:"-"`
}

type Inspection_story struct {
	Inspection_storyID  string         `gorm:"type:uuid;primaryKey" json:"inspection_story_id"`
	Inspection_name     *string        `json:"inspection_name"`
	Inspection_Username *string        `json:"inspection_username"`
	Inspection_Details  *string        `json:"inspection_details"`
	CreatedAt           *time.Time     `json:"created_at"`
	UpdatedAt           *time.Time     `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"index"`
	Asset_countID       string         `json:"asset_count_id"`
	Asset_count         Asset_count    `validate:"-"`
}

type Round_Count struct {
	Round_CountID string         `gorm:"type:uuid;primaryKey" json:"round_count_id"`
	Round         *string        `json:"round"`
	Status        *string        `json:"status"`
	CreatedAt     *time.Time     `json:"created_at"`
	UpdatedAt     *time.Time     `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Asset_countID string         `json:"asset_count_id"`
	Asset_count   Asset_count    `validate:"-"`
}

type Asset_Round struct {
	Asset_RoundID string      `gorm:"type:uuid;primaryKey" json:"asset_round_id"`
	AssetsID      string      `json:"assets_id"`
	Assets        Assets      `validate:"-"`
	Asset_countID string      `json:"asset_count_id"`
	Asset_count   Asset_count `validate:"-"`
	Round_CountID string      `json:"round_count_id"`
	Round_Count   Round_Count `validate:"-"`
}

type Round_Count_Story struct {
	Round_Count_StoryID string         `gorm:"type:uuid;primaryKey" json:"round_count_story_id"`
	Round_Name_Strory   *string        `json:"round_name_strory"`
	Round_Name          *string        `json:"round_name"`
	Round_Strory_Detail *string        `json:"round_strory_detail"`
	CreatedAt           *time.Time     `json:"created_at"`
	UpdatedAt           *time.Time     `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"index"`
	Round_CountID       string         `json:"round_count_id"`
	Round_Count         Round_Count    `validate:"-"`
	UserID              string         `json:"user_id"`
	User                User           `validate:"-"`
}

type Typeplan struct {
	TypeplanID   string         `gorm:"type:uuid;primaryKey" json:"typeplan_id"`
	TypeplanName *string        `json:"typeplan_name"`
	Comment      *string        `json:"comment"`
	CreatedAt    *time.Time     `json:"created_at"`
	UpdatedAt    *time.Time     `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type Typeplan_story struct {
	Typeplan_storyID   string         `gorm:"type:uuid;primaryKey" json:"typeplan_story_id"`
	Typeplan_StoryName *string        `json:"typeplan_story_name"`
	Typeplan_Username  *string        `json:"typeplan_username"`
	Typeplan_Details   *string        `json:"typeplan_details"`
	CreatedAt          *time.Time     `json:"created_at"`
	UpdatedAt          *time.Time     `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"index"`
	UserID             string         `json:"user_id"`
	TypeplanID         string         `json:"typeplan_id"`
	Typeplan           Typeplan       `validate:"-"`
}

type Location struct {
	LocationID    string `gorm:"type:uuid;primaryKey" json:"location_id"`
	Coordinates_x string `json:"coordinates_x"`
	Coordinates_y string `json:"coordinates_y"`
	AssetsID      string `json:"assets_id"`
	Assets        Assets `validate:"-"`
	GroundID      string `json:"ground_id"`
	Ground        Ground `validate:"-"`
}

type Location_story struct {
	Location_storyID   string         `gorm:"type:uuid;primaryKey" json:"location_story_id"`
	Location_StoryName *string        `json:"location_story_name"`
	Location_Username  *string        `json:"location_username"`
	Location_Details   *string        `json:"location_details"`
	CreatedAt          *time.Time     `json:"created_at"`
	UpdatedAt          *time.Time     `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"index"`
	LocationID         string         `json:"location_id"`
	Location           Location       `validate:"-"`
}

type Signature struct {
	SignatureID   string         `gorm:"type:uuid;primaryKey" json:"signature_id"`
	Signature     *string        `json:"signature"`
	UserID        string         `json:"user_id"`
	User          User           `validate:"-"`
	CreatedAt     *time.Time     `json:"created_at"`
	UpdatedAt     *time.Time     `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Asset_countID string         `json:"asset_count_id"`
	Asset_count   Asset_count    `validate:"-"`
}

type Branch_Autoclik struct {
	BranchAutoclik_ID string         `gorm:"type:uuid;primaryKey" json:"branchautoclik_id"`
	Branch_Name       string         `json:"branch_name"`
	Branch_Code       string         `json:"branch_code"`
	CreatedAt         *time.Time     `json:"created_at"`
	UpdatedAt         *time.Time     `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}

type Count_Autoclik struct {
	Autoclik_countID   string          `gorm:"type:uuid;primaryKey" json:"autoclik_count_id"`
	Plan_Code          string          `json:"plan_code"`
	Plan_Name          *string         `json:"plan_name"`
	TypeplanName       *string         `json:"typeplan_name"`
	Project_name       *string         `json:"project_name"`
	Comment            *string         `json:"comment"`
	Plan_start         *string         `json:"plan_start"`
	Plan_end           *string         `json:"plan_end"`
	Status             *string         `json:"status"`
	Product_Group_Code *string         `json:"Product_Group_Code"`
	CreatedAt          *time.Time      `json:"created_at"`
	UpdatedAt          *time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt  `gorm:"index"`
	UserID             string          `json:"user_id"`
	User               User            `validate:"-"`
	Group              string          `json:"group"`
	BranchAutoclik_ID  string          `json:"branch_autoclik_id"`
	Branch_Autoclik    Branch_Autoclik `validate:"-"`
}
