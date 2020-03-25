package models

import (
	"time"

	"github.com/senslabs/alpha/sens/datastore"
)

var t time.Time

type Auth struct {
	Id        string               `db:"id"`
	Email     datastore.NullString `db:"email"`
	Mobile    string               `db:"mobile"`
	Social    datastore.NullString `db:"social"`
	FirstName string               `db:"first_name"`
	LastName  string               `db:"last_name"`
	CreatedAt datastore.NullInt64  `db:"created_at"`
	UpdatedAt datastore.NullInt64  `db:"updated_at"`
}

func GetAuthFieldMap() map[string]string {
	return map[string]string{"CreatedAt": "created_at", "Email": "email", "FirstName": "first_name", "Id": "id", "LastName": "last_name", "Mobile": "mobile", "Social": "social", "UpdatedAt": "updated_at"}
}
func GetAuthTypeMap() map[string]string {
	return map[string]string{"CreatedAt": "datastore.NullInt64", "Email": "datastore.NullString", "FirstName": "string", "Id": "string", "LastName": "string", "Mobile": "string", "Social": "datastore.NullString", "UpdatedAt": "datastore.NullInt64"}
}

type Org struct {
	Id        string               `db:"id"`
	AuthId    datastore.NullString `db:"auth_id"`
	Name      string               `db:"name"`
	CreatedAt datastore.NullInt64  `db:"created_at"`
	UpdatedAt datastore.NullInt64  `db:"updated_at"`
}

func GetOrgFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "CreatedAt": "created_at", "Id": "id", "Name": "name", "UpdatedAt": "updated_at"}
}
func GetOrgTypeMap() map[string]string {
	return map[string]string{"AuthId": "datastore.NullString", "CreatedAt": "datastore.NullInt64", "Id": "string", "Name": "string", "UpdatedAt": "datastore.NullInt64"}
}

type Op struct {
	Id        string               `db:"id"`
	AuthId    datastore.NullString `db:"auth_id"`
	CreatedAt datastore.NullInt64  `db:"created_at"`
	UpdatedAt datastore.NullInt64  `db:"updated_at"`
	Status    datastore.NullString `db:"status"`
}

func GetOpFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "CreatedAt": "created_at", "Id": "id", "Status": "status", "UpdatedAt": "updated_at"}
}
func GetOpTypeMap() map[string]string {
	return map[string]string{"AuthId": "datastore.NullString", "CreatedAt": "datastore.NullInt64", "Id": "string", "Status": "datastore.NullString", "UpdatedAt": "datastore.NullInt64"}
}

type User struct {
	Id        string               `db:"id"`
	AuthId    datastore.NullString `db:"auth_id"`
	CreatedAt datastore.NullInt64  `db:"created_at"`
	UpdatedAt datastore.NullInt64  `db:"updated_at"`
	Status    datastore.NullString `db:"status"`
}

func GetUserFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "CreatedAt": "created_at", "Id": "id", "Status": "status", "UpdatedAt": "updated_at"}
}
func GetUserTypeMap() map[string]string {
	return map[string]string{"AuthId": "datastore.NullString", "CreatedAt": "datastore.NullInt64", "Id": "string", "Status": "datastore.NullString", "UpdatedAt": "datastore.NullInt64"}
}

type Endpoint struct {
	Id       string               `db:"id"`
	Category datastore.NullString `db:"category"`
	Path     datastore.NullString `db:"path"`
	Secure   bool                 `db:"secure"`
}

func GetEndpointFieldMap() map[string]string {
	return map[string]string{"Category": "category", "Id": "id", "Path": "path", "Secure": "secure"}
}
func GetEndpointTypeMap() map[string]string {
	return map[string]string{"Category": "datastore.NullString", "Id": "string", "Path": "datastore.NullString", "Secure": "bool"}
}

type Device struct {
	Id         string               `db:"id"`
	DeviceId   datastore.NullString `db:"device_id"`
	Name       datastore.NullString `db:"name"`
	OrgId      datastore.NullString `db:"org_id"`
	UserId     datastore.NullString `db:"user_id"`
	CreatedAt  datastore.NullInt64  `db:"created_at"`
	Status     datastore.NullString `db:"status"`
	Properties datastore.RawMessage `db:"properties"`
}

func GetDeviceFieldMap() map[string]string {
	return map[string]string{"CreatedAt": "created_at", "DeviceId": "device_id", "Id": "id", "Name": "name", "OrgId": "org_id", "Properties": "properties", "Status": "status", "UserId": "user_id"}
}
func GetDeviceTypeMap() map[string]string {
	return map[string]string{"CreatedAt": "datastore.NullInt64", "DeviceId": "datastore.NullString", "Id": "string", "Name": "datastore.NullString", "OrgId": "datastore.NullString", "Properties": "datastore.RawMessage", "Status": "datastore.NullString", "UserId": "datastore.NullString"}
}

type OrgOp struct {
	OrgId string `db:"org_id"`
	OpId  string `db:"op_id"`
}

func GetOrgOpFieldMap() map[string]string {
	return map[string]string{"OpId": "op_id", "OrgId": "org_id"}
}
func GetOrgOpTypeMap() map[string]string {
	return map[string]string{"OpId": "string", "OrgId": "string"}
}

type OrgUser struct {
	OrgId    string               `db:"org_id"`
	UserId   string               `db:"user_id"`
	Category datastore.NullString `db:"category"`
}

func GetOrgUserFieldMap() map[string]string {
	return map[string]string{"Category": "category", "OrgId": "org_id", "UserId": "user_id"}
}
func GetOrgUserTypeMap() map[string]string {
	return map[string]string{"Category": "datastore.NullString", "OrgId": "string", "UserId": "string"}
}

type OpUserCategorie struct {
	OpId         string `db:"op_id"`
	UserCategory string `db:"user_category"`
}

func GetOpUserCategorieFieldMap() map[string]string {
	return map[string]string{"OpId": "op_id", "UserCategory": "user_category"}
}
func GetOpUserCategorieTypeMap() map[string]string {
	return map[string]string{"OpId": "string", "UserCategory": "string"}
}

type OpUser struct {
	OpId   string `db:"op_id"`
	UserId string `db:"user_id"`
	Access bool   `db:"access"`
}

func GetOpUserFieldMap() map[string]string {
	return map[string]string{"Access": "access", "OpId": "op_id", "UserId": "user_id"}
}
func GetOpUserTypeMap() map[string]string {
	return map[string]string{"Access": "bool", "OpId": "string", "UserId": "string"}
}

type OrgEndpointCategorie struct {
	OrgId            string `db:"org_id"`
	EndpointCategory string `db:"endpoint_category"`
}

func GetOrgEndpointCategorieFieldMap() map[string]string {
	return map[string]string{"EndpointCategory": "endpoint_category", "OrgId": "org_id"}
}
func GetOrgEndpointCategorieTypeMap() map[string]string {
	return map[string]string{"EndpointCategory": "string", "OrgId": "string"}
}

type OrgEndpoint struct {
	OrgId      string `db:"org_id"`
	EndpointId string `db:"endpoint_id"`
	Access     bool   `db:"access"`
}

func GetOrgEndpointFieldMap() map[string]string {
	return map[string]string{"Access": "access", "EndpointId": "endpoint_id", "OrgId": "org_id"}
}
func GetOrgEndpointTypeMap() map[string]string {
	return map[string]string{"Access": "bool", "EndpointId": "string", "OrgId": "string"}
}

type OpEndpointCategorie struct {
	OpId             string `db:"op_id"`
	EndpointCategory string `db:"endpoint_category"`
}

func GetOpEndpointCategorieFieldMap() map[string]string {
	return map[string]string{"EndpointCategory": "endpoint_category", "OpId": "op_id"}
}
func GetOpEndpointCategorieTypeMap() map[string]string {
	return map[string]string{"EndpointCategory": "string", "OpId": "string"}
}

type OpEndpoint struct {
	OpId       string `db:"op_id"`
	EndpointId string `db:"endpoint_id"`
	Access     bool   `db:"access"`
}

func GetOpEndpointFieldMap() map[string]string {
	return map[string]string{"Access": "access", "EndpointId": "endpoint_id", "OpId": "op_id"}
}
func GetOpEndpointTypeMap() map[string]string {
	return map[string]string{"Access": "bool", "EndpointId": "string", "OpId": "string"}
}

type UserEndpointCategorie struct {
	UserId           string `db:"user_id"`
	EndpointCategory string `db:"endpoint_category"`
}

func GetUserEndpointCategorieFieldMap() map[string]string {
	return map[string]string{"EndpointCategory": "endpoint_category", "UserId": "user_id"}
}
func GetUserEndpointCategorieTypeMap() map[string]string {
	return map[string]string{"EndpointCategory": "string", "UserId": "string"}
}

type UserEndpoint struct {
	UserId     string `db:"user_id"`
	EndpointId string `db:"endpoint_id"`
	Access     bool   `db:"access"`
}

func GetUserEndpointFieldMap() map[string]string {
	return map[string]string{"Access": "access", "EndpointId": "endpoint_id", "UserId": "user_id"}
}
func GetUserEndpointTypeMap() map[string]string {
	return map[string]string{"Access": "bool", "EndpointId": "string", "UserId": "string"}
}

type Session struct {
	Id        string               `db:"id"`
	UserId    datastore.NullString `db:"user_id"`
	Name      datastore.NullString `db:"name"`
	Type      datastore.NullString `db:"type"`
	StartedAt datastore.NullInt64  `db:"started_at"`
	EndedAt   datastore.NullInt64  `db:"ended_at"`
}

func GetSessionFieldMap() map[string]string {
	return map[string]string{"EndedAt": "ended_at", "Id": "id", "Name": "name", "StartedAt": "started_at", "Type": "type", "UserId": "user_id"}
}
func GetSessionTypeMap() map[string]string {
	return map[string]string{"EndedAt": "datastore.NullInt64", "Id": "string", "Name": "datastore.NullString", "StartedAt": "datastore.NullInt64", "Type": "datastore.NullString", "UserId": "datastore.NullString"}
}

type SessionEvent struct {
	UserId     string               `db:"user_id"`
	Name       string               `db:"name"`
	StartedAt  int64                `db:"started_at"`
	EndedAt    datastore.NullInt64  `db:"ended_at"`
	Properties datastore.RawMessage `db:"properties"`
}

func GetSessionEventFieldMap() map[string]string {
	return map[string]string{"EndedAt": "ended_at", "Name": "name", "Properties": "properties", "StartedAt": "started_at", "UserId": "user_id"}
}
func GetSessionEventTypeMap() map[string]string {
	return map[string]string{"EndedAt": "datastore.NullInt64", "Name": "string", "Properties": "datastore.RawMessage", "StartedAt": "int64", "UserId": "string"}
}

type SessionRecord struct {
	UserId     string               `db:"user_id"`
	Name       string               `db:"name"`
	Timestamp  int64                `db:"timestamp"`
	Value      float64              `db:"value"`
	Properties datastore.RawMessage `db:"properties"`
}

func GetSessionRecordFieldMap() map[string]string {
	return map[string]string{"Name": "name", "Properties": "properties", "Timestamp": "timestamp", "UserId": "user_id", "Value": "value"}
}
func GetSessionRecordTypeMap() map[string]string {
	return map[string]string{"Name": "string", "Properties": "datastore.RawMessage", "Timestamp": "int64", "UserId": "string", "Value": "float64"}
}

type SessionPropertie struct {
	SessionId datastore.NullString `db:"session_id"`
	Name      datastore.NullString `db:"name"`
	Value     datastore.NullString `db:"value"`
	Rowid     int64                `db:"rowid"`
}

func GetSessionPropertieFieldMap() map[string]string {
	return map[string]string{"Name": "name", "Rowid": "rowid", "SessionId": "session_id", "Value": "value"}
}
func GetSessionPropertieTypeMap() map[string]string {
	return map[string]string{"Name": "datastore.NullString", "Rowid": "int64", "SessionId": "datastore.NullString", "Value": "datastore.NullString"}
}

type OrgDetailView struct {
	AuthId    string `db:"auth_id"`
	Email     string `db:"email"`
	Mobile    string `db:"mobile"`
	Social    string `db:"social"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Id        string `db:"id"`
	OrgName   string `db:"org_name"`
}

func GetOrgDetailViewFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "Email": "email", "FirstName": "first_name", "Id": "id", "LastName": "last_name", "Mobile": "mobile", "OrgName": "org_name", "Social": "social"}
}
func GetOrgDetailViewTypeMap() map[string]string {
	return map[string]string{"AuthId": "string", "Email": "string", "FirstName": "string", "Id": "string", "LastName": "string", "Mobile": "string", "OrgName": "string", "Social": "string"}
}

type OpDetailView struct {
	AuthId    string `db:"auth_id"`
	Email     string `db:"email"`
	Mobile    string `db:"mobile"`
	Social    string `db:"social"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Id        string `db:"id"`
}

func GetOpDetailViewFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "Email": "email", "FirstName": "first_name", "Id": "id", "LastName": "last_name", "Mobile": "mobile", "Social": "social"}
}
func GetOpDetailViewTypeMap() map[string]string {
	return map[string]string{"AuthId": "string", "Email": "string", "FirstName": "string", "Id": "string", "LastName": "string", "Mobile": "string", "Social": "string"}
}

type UserDetailView struct {
	AuthId    string `db:"auth_id"`
	Email     string `db:"email"`
	Mobile    string `db:"mobile"`
	Social    string `db:"social"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Id        string `db:"id"`
}

func GetUserDetailViewFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "Email": "email", "FirstName": "first_name", "Id": "id", "LastName": "last_name", "Mobile": "mobile", "Social": "social"}
}
func GetUserDetailViewTypeMap() map[string]string {
	return map[string]string{"AuthId": "string", "Email": "string", "FirstName": "string", "Id": "string", "LastName": "string", "Mobile": "string", "Social": "string"}
}
