package builder

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm"
)

func CreateMaliwan_data(db *gorm.DB, Odata_etag string, No string, Description_2 string, Description_ENG string,
	No_2 string, Description string, Maximum_Inventory int32, Safety_Stock_Quantity int32, Product_Group_Code string, Created_From_Nonstock_Item bool,
	Line_Type string, Type_Insurance bool, Type_Registration_fee bool, Type_Accessory bool, Type_Deposit bool,
	Stockkeeping_Unit_Exists bool, Assembly_BOM bool, Shelf_No string,
	Gen_Prod_Posting_Group string, Inventory_Posting_Group string, Inventory float64, Base_Unit_of_Measure string,
	Item_Category_Code string, VAT_Prod_Posting_Group string, Item_Disc_Group string, Vendor_No string, Vendor_Item_No string,
	Tariff_No string, Search_Description string, Overhead_Rate int32, Blocked bool, Last_Date_Modified string, Sales_Unit_of_Measure string,
	Replenishment_System string, Purch_Unit_of_Measure string, Lead_Time_Calculation string, Manufacturing_Policy string, Flushing_Method string,
	Assembly_Policy string, Item_Tracking_Code string, Promotion_Model_Code string, Default_Location_Code string, Qty_on_Purch_Order int32, Qty_on_Service_Order float64,
	Qty_on_Sales_Order float64, Unit_Price_Price float64, Brand string, Costing_Method string, Stockout_Warning string, Global_Dimension_1_Filter string,
	Global_Dimension_2_Filter string, Location_Filter string, Drop_Shipment_Filter string, Variant_Filter string, Lot_No_Filter string, Serial_No_Filter string, Date_Filter string) error {
	return db.Create(&model.Maliwan_data{
		Maliwan_dataID:             uuid.New().String(),
		Odata_etag:                 &Odata_etag,
		No:                         &No,
		No_2:                       &No_2,
		Description:                &Description,
		Description_2:              &Description_2,
		Description_ENG:            &Description_ENG,
		Inventory:                  &Inventory,
		Maximum_Inventory:          &Maximum_Inventory,
		Safety_Stock_Quantity:      &Safety_Stock_Quantity,
		Product_Group_Code:         &Product_Group_Code,
		Created_From_Nonstock_Item: &Created_From_Nonstock_Item,
		Line_Type:                  &Line_Type,
		Type_Insurance:             &Type_Insurance,
		Type_Registration_fee:      &Type_Registration_fee,
		Type_Accessory:             &Type_Accessory,
		Type_Deposit:               &Type_Deposit,
		Stockkeeping_Unit_Exists:   &Stockkeeping_Unit_Exists,
		Assembly_BOM:               &Assembly_BOM,
		Base_Unit_of_Measure:       &Base_Unit_of_Measure,
		Shelf_No:                   &Shelf_No,
		Gen_Prod_Posting_Group:     &Gen_Prod_Posting_Group,
		Inventory_Posting_Group:    &Inventory_Posting_Group,
		Item_Category_Code:         &Item_Category_Code,
		VAT_Prod_Posting_Group:     &VAT_Prod_Posting_Group,
		Item_Disc_Group:            &Item_Disc_Group,
		Vendor_No:                  &Vendor_No,
		Vendor_Item_No:             &Vendor_Item_No,
		Tariff_No:                  &Tariff_No,
		Search_Description:         &Search_Description,
		Overhead_Rate:              &Overhead_Rate,
		Blocked:                    &Blocked,
		Last_Date_Modified:         &Last_Date_Modified,
		Sales_Unit_of_Measure:      &Sales_Unit_of_Measure,
		Replenishment_System:       &Replenishment_System,
		Purch_Unit_of_Measure:      &Purch_Unit_of_Measure,
		Lead_Time_Calculation:      &Lead_Time_Calculation,
		Manufacturing_Policy:       &Manufacturing_Policy,
		Flushing_Method:            &Flushing_Method,
		Assembly_Policy:            &Assembly_Policy,
		Item_Tracking_Code:         &Item_Tracking_Code,
		Promotion_Model_Code:       &Promotion_Model_Code,
		Default_Location_Code:      &Default_Location_Code,
		Qty_on_Purch_Order:         &Qty_on_Purch_Order,
		Qty_on_Service_Order:       &Qty_on_Service_Order,
		Qty_on_Sales_Order:         &Qty_on_Sales_Order,
		Unit_Price_Price:           &Unit_Price_Price,
		Brand:                      &Brand,
		Costing_Method:             &Costing_Method,
		Stockout_Warning:           &Stockout_Warning,
		Global_Dimension_1_Filter:  &Global_Dimension_1_Filter,
		Global_Dimension_2_Filter:  &Global_Dimension_2_Filter,
		Location_Filter:            &Location_Filter,
		Drop_Shipment_Filter:       &Drop_Shipment_Filter,
		Variant_Filter:             &Variant_Filter,
		Lot_No_Filter:              &Lot_No_Filter,
		Serial_No_Filter:           &Serial_No_Filter,
		Date_Filter:                &Date_Filter,
	}).Error
}
