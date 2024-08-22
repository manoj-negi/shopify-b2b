package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type ProductStatus string
type MetafieldType string
type VariantInventoryPolicy string

const (
	ProductStatusActive   ProductStatus = "active"
	ProductStatusArchived ProductStatus = "archived"
	ProductStatusDraft    ProductStatus = "draft"
)

type Image struct {
	Id                uint64     `json:"id,omitempty"`
	ProductId         uint64     `json:"product_id,omitempty"`
	Position          int        `json:"position,omitempty"`
	CreatedAt         *time.Time `json:"created_at,omitempty"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty"`
	Width             int        `json:"width,omitempty"`
	Height            int        `json:"height,omitempty"`
	Src               string     `json:"src,omitempty"`
	Attachment        string     `json:"attachment,omitempty"`
	Filename          string     `json:"filename,omitempty"`
	Alt               string     `json:"alt,omitempty"`
	VariantIds        []uint64   `json:"variant_ids,omitempty"`
	AdminGraphqlApiId string     `json:"admin_graphql_api_id,omitempty"`
}

type Product struct {
	Id                             uint64          `json:"id,omitempty"`
	Title                          string          `json:"title,omitempty"`
	BodyHTML                       string          `json:"body_html,omitempty"`
	Vendor                         string          `json:"vendor,omitempty"`
	ProductType                    string          `json:"product_type,omitempty"`
	Handle                         string          `json:"handle,omitempty"`
	CreatedAt                      *time.Time      `json:"created_at,omitempty"`
	UpdatedAt                      *time.Time      `json:"updated_at,omitempty"`
	PublishedAt                    *time.Time      `json:"published_at,omitempty"`
	PublishedScope                 string          `json:"published_scope,omitempty"`
	Tags                           string          `json:"tags,omitempty"`
	Status                         ProductStatus   `json:"status,omitempty"`
	Options                        []ProductOption `json:"options,omitempty"`
	Variants                       []Variant       `json:"variants,omitempty"`
	Image                          Image           `json:"image,omitempty"`
	Images                         []Image         `json:"images,omitempty"`
	TemplateSuffix                 string          `json:"template_suffix,omitempty"`
	MetafieldsGlobalTitleTag       string          `json:"metafields_global_title_tag,omitempty"`
	MetafieldsGlobalDescriptionTag string          `json:"metafields_global_description_tag,omitempty"`
	Metafields                     []Metafield     `json:"metafields,omitempty"`
	AdminGraphqlApiId              string          `json:"admin_graphql_api_id,omitempty"`
}

type Metafield struct {
	CreatedAt         *time.Time    `json:"created_at,omitempty"`
	Description       string        `json:"description,omitempty"`    // Description of the metafield.
	Id                uint64        `json:"id,omitempty"`             // Assigned by Shopify, used for updating a metafield.
	Key               string        `json:"key,omitempty"`            // The unique identifier for a metafield within its namespace, 3-64 characters long.
	Namespace         string        `json:"namespace,omitempty"`      // The container for a group of metafields, 3-255 characters long.
	OwnerId           uint64        `json:"owner_id,omitempty"`       // The unique Id of the resource the metafield is for, i.e.: an Order Id.
	OwnerResource     string        `json:"owner_resource,omitempty"` // The type of reserouce the metafield is for, i.e.: and Order.
	UpdatedAt         *time.Time    `json:"updated_at,omitempty"`     //
	Value             interface{}   `json:"value,omitempty"`          // The data stored in the metafield. Always stored as a string, use Type field for actual data type.
	Type              MetafieldType `json:"type,omitempty"`           // One of Shopify's defined types, see MetafieldType.
	AdminGraphqlApiId string        `json:"admin_graphql_api_id,omitempty"`
}

type Variant struct {
	Id                   uint64                 `json:"id,omitempty"`
	ProductId            uint64                 `json:"product_id,omitempty"`
	Title                string                 `json:"title,omitempty"`
	Sku                  string                 `json:"sku,omitempty"`
	Position             int                    `json:"position,omitempty"`
	Grams                int                    `json:"grams,omitempty"`
	InventoryPolicy      VariantInventoryPolicy `json:"inventory_policy,omitempty"`
	Price                *decimal.Decimal       `json:"price,omitempty"`
	CompareAtPrice       *decimal.Decimal       `json:"compare_at_price,omitempty"`
	FulfillmentService   string                 `json:"fulfillment_service,omitempty"`
	InventoryManagement  string                 `json:"inventory_management,omitempty"`
	InventoryItemId      uint64                 `json:"inventory_item_id,omitempty"`
	Option1              string                 `json:"option1,omitempty"`
	Option2              string                 `json:"option2,omitempty"`
	Option3              string                 `json:"option3,omitempty"`
	CreatedAt            *time.Time             `json:"created_at,omitempty"`
	UpdatedAt            *time.Time             `json:"updated_at,omitempty"`
	Taxable              bool                   `json:"taxable,omitempty"`
	TaxCode              string                 `json:"tax_code,omitempty"`
	Barcode              string                 `json:"barcode,omitempty"`
	ImageId              uint64                 `json:"image_id,omitempty"`
	InventoryQuantity    int                    `json:"inventory_quantity,omitempty"`
	Weight               *decimal.Decimal       `json:"weight,omitempty"`
	WeightUnit           string                 `json:"weight_unit,omitempty"`
	OldInventoryQuantity int                    `json:"old_inventory_quantity,omitempty"`
	RequireShipping      bool                   `json:"requires_shipping"`
	AdminGraphqlApiId    string                 `json:"admin_graphql_api_id,omitempty"`
	Metafields           []Metafield            `json:"metafields,omitempty"`
	PresentmentPrices    []presentmentPrices    `json:"presentment_prices,omitempty"`
}
type AmountSetEntry struct {
	Amount       *decimal.Decimal `json:"amount,omitempty"`
	CurrencyCode string           `json:"currency_code,omitempty"`
}

type presentmentPrices struct {
	Price          *AmountSetEntry `json:"price,omitempty"`
	CompareAtPrice *AmountSetEntry `json:"compare_at_price,omitempty"`
}

type ProductOption struct {
	Id        uint64   `json:"id,omitempty"`
	ProductId uint64   `json:"product_id,omitempty"`
	Name      string   `json:"name,omitempty"`
	Position  int      `json:"position,omitempty"`
	Values    []string `json:"values,omitempty"`
}

type ProductListOptions struct {
	ListOptions
	CollectionId          uint64          `url:"collection_id,omitempty"`
	ProductType           string          `url:"product_type,omitempty"`
	Vendor                string          `url:"vendor,omitempty"`
	Handle                string          `url:"handle,omitempty"`
	PublishedAtMin        time.Time       `url:"published_at_min,omitempty"`
	PublishedAtMax        time.Time       `url:"published_at_max,omitempty"`
	PublishedStatus       string          `url:"published_status,omitempty"`
	PresentmentCurrencies string          `url:"presentment_currencies,omitempty"`
	Status                []ProductStatus `url:"status,omitempty,comma"`
	Title                 string          `url:"title,omitempty"`
}

type ListOptions struct {
	// PageInfo is used with new pagination search.
	PageInfo string `url:"page_info,omitempty"`

	// Page is used to specify a specific page to load.
	// It is the deprecated way to do pagination.
	Page         int       `url:"page,omitempty"`
	Limit        int       `url:"limit,omitempty"`
	SinceId      *uint64   `url:"since_id,omitempty"`
	CreatedAtMin time.Time `url:"created_at_min,omitempty"`
	CreatedAtMax time.Time `url:"created_at_max,omitempty"`
	UpdatedAtMin time.Time `url:"updated_at_min,omitempty"`
	UpdatedAtMax time.Time `url:"updated_at_max,omitempty"`
	Order        string    `url:"order,omitempty"`
	Fields       string    `url:"fields,omitempty"`
	Vendor       string    `url:"vendor,omitempty"`
	Ids          []uint64  `url:"ids,omitempty,comma"`
}
