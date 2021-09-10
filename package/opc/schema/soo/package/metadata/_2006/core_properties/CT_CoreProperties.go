package core_properties

import "time"

type CT_CoreProperties struct {
	Category       *string
	ContentStatus  *string
	Created        *schema.XSDAny
	Creator        *schema.XSDAny
	Description    *schema.XSDAny
	Identifier     *schema.XSDAny
	Keywords       *CT_Keywords
	Language       *schema.XSDAny
	LastModifiedBy *string
	LastPrinted    *time.Time
	Modified       *schema.XSDAny
	Revision       *string
	Subject        *schema.XSDAny
	Title          *schema.XSDAny
	Version        *string
}
