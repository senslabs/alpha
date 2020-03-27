package models

import (
	"time"

	"github.com/senslabs/alpha/sens/datastore"
)

var t time.Time

type Auth struct {
	Id        *string `db:"id" json:",omitempty"`
	Email     *string `db:"email" json:",omitempty"`
	Mobile    *string `db:"mobile" json:",omitempty"`
	Social    *string `db:"social" json:",omitempty"`
	FirstName *string `db:"first_name" json:",omitempty"`
	LastName  *string `db:"last_name" json:",omitempty"`
	CreatedAt *int64  `db:"created_at" json:",omitempty"`
	UpdatedAt *int64  `db:"updated_at" json:",omitempty"`
}

func GetAuthFieldMap() map[string]string {
	return map[string]string{"CreatedAt": "created_at", "Email": "email", "FirstName": "first_name", "Id": "id", "LastName": "last_name", "Mobile": "mobile", "Social": "social", "UpdatedAt": "updated_at"}
}
func GetAuthTypeMap() map[string]string {
	return map[string]string{"CreatedAt": "*int64", "Email": "*string", "FirstName": "*string", "Id": "*string", "LastName": "*string", "Mobile": "*string", "Social": "*string", "UpdatedAt": "*int64"}
}

type Org struct {
	Id        *string `db:"id" json:",omitempty"`
	AuthId    *string `db:"auth_id" json:",omitempty"`
	Name      *string `db:"name" json:",omitempty"`
	CreatedAt *int64  `db:"created_at" json:",omitempty"`
	UpdatedAt *int64  `db:"updated_at" json:",omitempty"`
}

func GetOrgFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "CreatedAt": "created_at", "Id": "id", "Name": "name", "UpdatedAt": "updated_at"}
}
func GetOrgTypeMap() map[string]string {
	return map[string]string{"AuthId": "*string", "CreatedAt": "*int64", "Id": "*string", "Name": "*string", "UpdatedAt": "*int64"}
}

type Op struct {
	Id        *string `db:"id" json:",omitempty"`
	AuthId    *string `db:"auth_id" json:",omitempty"`
	CreatedAt *int64  `db:"created_at" json:",omitempty"`
	UpdatedAt *int64  `db:"updated_at" json:",omitempty"`
	Status    *string `db:"status" json:",omitempty"`
}

func GetOpFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "CreatedAt": "created_at", "Id": "id", "Status": "status", "UpdatedAt": "updated_at"}
}
func GetOpTypeMap() map[string]string {
	return map[string]string{"AuthId": "*string", "CreatedAt": "*int64", "Id": "*string", "Status": "*string", "UpdatedAt": "*int64"}
}

type User struct {
	Id        *string `db:"id" json:",omitempty"`
	AuthId    *string `db:"auth_id" json:",omitempty"`
	CreatedAt *int64  `db:"created_at" json:",omitempty"`
	UpdatedAt *int64  `db:"updated_at" json:",omitempty"`
	Status    *string `db:"status" json:",omitempty"`
}

func GetUserFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "CreatedAt": "created_at", "Id": "id", "Status": "status", "UpdatedAt": "updated_at"}
}
func GetUserTypeMap() map[string]string {
	return map[string]string{"AuthId": "*string", "CreatedAt": "*int64", "Id": "*string", "Status": "*string", "UpdatedAt": "*int64"}
}

type Endpoint struct {
	Id       *string `db:"id" json:",omitempty"`
	Category *string `db:"category" json:",omitempty"`
	Path     *string `db:"path" json:",omitempty"`
	Secure   *bool   `db:"secure" json:",omitempty"`
}

func GetEndpointFieldMap() map[string]string {
	return map[string]string{"Category": "category", "Id": "id", "Path": "path", "Secure": "secure"}
}
func GetEndpointTypeMap() map[string]string {
	return map[string]string{"Category": "*string", "Id": "*string", "Path": "*string", "Secure": "*bool"}
}

type Device struct {
	Id         *string               `db:"id" json:",omitempty"`
	DeviceId   *string               `db:"device_id" json:",omitempty"`
	Name       *string               `db:"name" json:",omitempty"`
	OrgId      *string               `db:"org_id" json:",omitempty"`
	UserId     *string               `db:"user_id" json:",omitempty"`
	CreatedAt  *int64                `db:"created_at" json:",omitempty"`
	Status     *string               `db:"status" json:",omitempty"`
	Properties *datastore.RawMessage `db:"properties" json:",omitempty"`
}

func GetDeviceFieldMap() map[string]string {
	return map[string]string{"CreatedAt": "created_at", "DeviceId": "device_id", "Id": "id", "Name": "name", "OrgId": "org_id", "Properties": "properties", "Status": "status", "UserId": "user_id"}
}
func GetDeviceTypeMap() map[string]string {
	return map[string]string{"CreatedAt": "*int64", "DeviceId": "*string", "Id": "*string", "Name": "*string", "OrgId": "*string", "Properties": "*datastore.RawMessage", "Status": "*string", "UserId": "*string"}
}

type OrgOp struct {
	OrgId *string `db:"org_id" json:",omitempty"`
	OpId  *string `db:"op_id" json:",omitempty"`
}

func GetOrgOpFieldMap() map[string]string {
	return map[string]string{"OpId": "op_id", "OrgId": "org_id"}
}
func GetOrgOpTypeMap() map[string]string {
	return map[string]string{"OpId": "*string", "OrgId": "*string"}
}

type OrgUser struct {
	OrgId    *string `db:"org_id" json:",omitempty"`
	UserId   *string `db:"user_id" json:",omitempty"`
	Category *string `db:"category" json:",omitempty"`
}

func GetOrgUserFieldMap() map[string]string {
	return map[string]string{"Category": "category", "OrgId": "org_id", "UserId": "user_id"}
}
func GetOrgUserTypeMap() map[string]string {
	return map[string]string{"Category": "*string", "OrgId": "*string", "UserId": "*string"}
}

type OpUserCategorie struct {
	OpId         *string `db:"op_id" json:",omitempty"`
	UserCategory *string `db:"user_category" json:",omitempty"`
}

func GetOpUserCategorieFieldMap() map[string]string {
	return map[string]string{"OpId": "op_id", "UserCategory": "user_category"}
}
func GetOpUserCategorieTypeMap() map[string]string {
	return map[string]string{"OpId": "*string", "UserCategory": "*string"}
}

type OpUser struct {
	OpId   *string `db:"op_id" json:",omitempty"`
	UserId *string `db:"user_id" json:",omitempty"`
	Access *bool   `db:"access" json:",omitempty"`
}

func GetOpUserFieldMap() map[string]string {
	return map[string]string{"Access": "access", "OpId": "op_id", "UserId": "user_id"}
}
func GetOpUserTypeMap() map[string]string {
	return map[string]string{"Access": "*bool", "OpId": "*string", "UserId": "*string"}
}

type OrgEndpointCategorie struct {
	OrgId            *string `db:"org_id" json:",omitempty"`
	EndpointCategory *string `db:"endpoint_category" json:",omitempty"`
}

func GetOrgEndpointCategorieFieldMap() map[string]string {
	return map[string]string{"EndpointCategory": "endpoint_category", "OrgId": "org_id"}
}
func GetOrgEndpointCategorieTypeMap() map[string]string {
	return map[string]string{"EndpointCategory": "*string", "OrgId": "*string"}
}

type OrgEndpoint struct {
	OrgId      *string `db:"org_id" json:",omitempty"`
	EndpointId *string `db:"endpoint_id" json:",omitempty"`
	Access     *bool   `db:"access" json:",omitempty"`
}

func GetOrgEndpointFieldMap() map[string]string {
	return map[string]string{"Access": "access", "EndpointId": "endpoint_id", "OrgId": "org_id"}
}
func GetOrgEndpointTypeMap() map[string]string {
	return map[string]string{"Access": "*bool", "EndpointId": "*string", "OrgId": "*string"}
}

type OpEndpointCategorie struct {
	OpId             *string `db:"op_id" json:",omitempty"`
	EndpointCategory *string `db:"endpoint_category" json:",omitempty"`
}

func GetOpEndpointCategorieFieldMap() map[string]string {
	return map[string]string{"EndpointCategory": "endpoint_category", "OpId": "op_id"}
}
func GetOpEndpointCategorieTypeMap() map[string]string {
	return map[string]string{"EndpointCategory": "*string", "OpId": "*string"}
}

type OpEndpoint struct {
	OpId       *string `db:"op_id" json:",omitempty"`
	EndpointId *string `db:"endpoint_id" json:",omitempty"`
	Access     *bool   `db:"access" json:",omitempty"`
}

func GetOpEndpointFieldMap() map[string]string {
	return map[string]string{"Access": "access", "EndpointId": "endpoint_id", "OpId": "op_id"}
}
func GetOpEndpointTypeMap() map[string]string {
	return map[string]string{"Access": "*bool", "EndpointId": "*string", "OpId": "*string"}
}

type UserEndpointCategorie struct {
	UserId           *string `db:"user_id" json:",omitempty"`
	EndpointCategory *string `db:"endpoint_category" json:",omitempty"`
}

func GetUserEndpointCategorieFieldMap() map[string]string {
	return map[string]string{"EndpointCategory": "endpoint_category", "UserId": "user_id"}
}
func GetUserEndpointCategorieTypeMap() map[string]string {
	return map[string]string{"EndpointCategory": "*string", "UserId": "*string"}
}

type UserEndpoint struct {
	UserId     *string `db:"user_id" json:",omitempty"`
	EndpointId *string `db:"endpoint_id" json:",omitempty"`
	Access     *bool   `db:"access" json:",omitempty"`
}

func GetUserEndpointFieldMap() map[string]string {
	return map[string]string{"Access": "access", "EndpointId": "endpoint_id", "UserId": "user_id"}
}
func GetUserEndpointTypeMap() map[string]string {
	return map[string]string{"Access": "*bool", "EndpointId": "*string", "UserId": "*string"}
}

type Session struct {
	Id        *string `db:"id" json:",omitempty"`
	UserId    *string `db:"user_id" json:",omitempty"`
	Name      *string `db:"name" json:",omitempty"`
	Type      *string `db:"type" json:",omitempty"`
	StartedAt *int64  `db:"started_at" json:",omitempty"`
	EndedAt   *int64  `db:"ended_at" json:",omitempty"`
}

func GetSessionFieldMap() map[string]string {
	return map[string]string{"EndedAt": "ended_at", "Id": "id", "Name": "name", "StartedAt": "started_at", "Type": "type", "UserId": "user_id"}
}
func GetSessionTypeMap() map[string]string {
	return map[string]string{"EndedAt": "*int64", "Id": "*string", "Name": "*string", "StartedAt": "*int64", "Type": "*string", "UserId": "*string"}
}

type SessionEvent struct {
	UserId     *string               `db:"user_id" json:",omitempty"`
	Name       *string               `db:"name" json:",omitempty"`
	StartedAt  *int64                `db:"started_at" json:",omitempty"`
	EndedAt    *int64                `db:"ended_at" json:",omitempty"`
	Properties *datastore.RawMessage `db:"properties" json:",omitempty"`
}

func GetSessionEventFieldMap() map[string]string {
	return map[string]string{"EndedAt": "ended_at", "Name": "name", "Properties": "properties", "StartedAt": "started_at", "UserId": "user_id"}
}
func GetSessionEventTypeMap() map[string]string {
	return map[string]string{"EndedAt": "*int64", "Name": "*string", "Properties": "*datastore.RawMessage", "StartedAt": "*int64", "UserId": "*string"}
}

type SessionRecord struct {
	UserId     *string               `db:"user_id" json:",omitempty"`
	Name       *string               `db:"name" json:",omitempty"`
	Timestamp  *int64                `db:"timestamp" json:",omitempty"`
	Value      *float64              `db:"value" json:",omitempty"`
	Properties *datastore.RawMessage `db:"properties" json:",omitempty"`
}

func GetSessionRecordFieldMap() map[string]string {
	return map[string]string{"Name": "name", "Properties": "properties", "Timestamp": "timestamp", "UserId": "user_id", "Value": "value"}
}
func GetSessionRecordTypeMap() map[string]string {
	return map[string]string{"Name": "*string", "Properties": "*datastore.RawMessage", "Timestamp": "*int64", "UserId": "*string", "Value": "*float64"}
}

type SessionPropertie struct {
	SessionId *string `db:"session_id" json:",omitempty"`
	Name      *string `db:"name" json:",omitempty"`
	Value     *string `db:"value" json:",omitempty"`
	Rowid     *int64  `db:"rowid" json:",omitempty"`
}

func GetSessionPropertieFieldMap() map[string]string {
	return map[string]string{"Name": "name", "Rowid": "rowid", "SessionId": "session_id", "Value": "value"}
}
func GetSessionPropertieTypeMap() map[string]string {
	return map[string]string{"Name": "*string", "Rowid": "*int64", "SessionId": "*string", "Value": "*string"}
}

type OrgDetailView struct {
	AuthId    *string `db:"auth_id" json:",omitempty"`
	Email     *string `db:"email" json:",omitempty"`
	Mobile    *string `db:"mobile" json:",omitempty"`
	Social    *string `db:"social" json:",omitempty"`
	FirstName *string `db:"first_name" json:",omitempty"`
	LastName  *string `db:"last_name" json:",omitempty"`
	Id        *string `db:"id" json:",omitempty"`
	OrgName   *string `db:"org_name" json:",omitempty"`
}

func GetOrgDetailViewFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "Email": "email", "FirstName": "first_name", "Id": "id", "LastName": "last_name", "Mobile": "mobile", "OrgName": "org_name", "Social": "social"}
}
func GetOrgDetailViewTypeMap() map[string]string {
	return map[string]string{"AuthId": "*string", "Email": "*string", "FirstName": "*string", "Id": "*string", "LastName": "*string", "Mobile": "*string", "OrgName": "*string", "Social": "*string"}
}

type OpDetailView struct {
	AuthId    *string `db:"auth_id" json:",omitempty"`
	Email     *string `db:"email" json:",omitempty"`
	Mobile    *string `db:"mobile" json:",omitempty"`
	Social    *string `db:"social" json:",omitempty"`
	FirstName *string `db:"first_name" json:",omitempty"`
	LastName  *string `db:"last_name" json:",omitempty"`
	Id        *string `db:"id" json:",omitempty"`
}

func GetOpDetailViewFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "Email": "email", "FirstName": "first_name", "Id": "id", "LastName": "last_name", "Mobile": "mobile", "Social": "social"}
}
func GetOpDetailViewTypeMap() map[string]string {
	return map[string]string{"AuthId": "*string", "Email": "*string", "FirstName": "*string", "Id": "*string", "LastName": "*string", "Mobile": "*string", "Social": "*string"}
}

type UserDetailView struct {
	AuthId    *string `db:"auth_id" json:",omitempty"`
	Email     *string `db:"email" json:",omitempty"`
	Mobile    *string `db:"mobile" json:",omitempty"`
	Social    *string `db:"social" json:",omitempty"`
	FirstName *string `db:"first_name" json:",omitempty"`
	LastName  *string `db:"last_name" json:",omitempty"`
	Id        *string `db:"id" json:",omitempty"`
}

func GetUserDetailViewFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "Email": "email", "FirstName": "first_name", "Id": "id", "LastName": "last_name", "Mobile": "mobile", "Social": "social"}
}
func GetUserDetailViewTypeMap() map[string]string {
	return map[string]string{"AuthId": "*string", "Email": "*string", "FirstName": "*string", "Id": "*string", "LastName": "*string", "Mobile": "*string", "Social": "*string"}
}
