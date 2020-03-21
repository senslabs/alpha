package models

import (
	"time"

	"github.com/senslabs/alpha/sens/datastore"
)

var t time.Time

type Auth struct {
	Id        string               `db:"id"`
	Email     string               `db:"email"`
	Mobile    string               `db:"mobile"`
	Social    string               `db:"social"`
	FirstName datastore.NullString `db:"first_name"`
	LastName  datastore.NullString `db:"last_name"`
	CreatedAt datastore.NullTime   `db:"created_at"`
	UpdatedAt datastore.NullTime   `db:"updated_at"`
}

func GetAuthFieldMap() map[string]string {
	return map[string]string{"CreatedAt": "created_at", "Email": "email", "FirstName": "first_name", "Id": "id", "LastName": "last_name", "Mobile": "mobile", "Social": "social", "UpdatedAt": "updated_at"}
}

type Org struct {
	Id        string               `db:"id"`
	AuthId    datastore.NullString `db:"auth_id"`
	Name      datastore.NullString `db:"name"`
	CreatedAt datastore.NullTime   `db:"created_at"`
	UpdatedAt datastore.NullTime   `db:"updated_at"`
}

func GetOrgFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "CreatedAt": "created_at", "Id": "id", "Name": "name", "UpdatedAt": "updated_at"}
}

type Op struct {
	Id        string               `db:"id"`
	AuthId    datastore.NullString `db:"auth_id"`
	CreatedAt datastore.NullTime   `db:"created_at"`
	UpdatedAt datastore.NullTime   `db:"updated_at"`
	Status    datastore.NullString `db:"status"`
}

func GetOpFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "CreatedAt": "created_at", "Id": "id", "Status": "status", "UpdatedAt": "updated_at"}
}

type User struct {
	Id        string               `db:"id"`
	AuthId    datastore.NullString `db:"auth_id"`
	CreatedAt datastore.NullTime   `db:"created_at"`
	UpdatedAt datastore.NullTime   `db:"updated_at"`
	Status    datastore.NullString `db:"status"`
}

func GetUserFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "CreatedAt": "created_at", "Id": "id", "Status": "status", "UpdatedAt": "updated_at"}
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

type Device struct {
	Id         string               `db:"id"`
	DeviceId   datastore.NullString `db:"device_id"`
	Name       datastore.NullString `db:"name"`
	OrgId      datastore.NullString `db:"org_id"`
	UserId     datastore.NullString `db:"user_id"`
	CreatedAt  datastore.NullTime   `db:"created_at"`
	Status     datastore.NullString `db:"status"`
	Properties datastore.RawMessage `db:"properties"`
}

func GetDeviceFieldMap() map[string]string {
	return map[string]string{"CreatedAt": "created_at", "DeviceId": "device_id", "Id": "id", "Name": "name", "OrgId": "org_id", "Properties": "properties", "Status": "status", "UserId": "user_id"}
}

type OrgUser struct {
	OrgId    string               `db:"org_id"`
	UserId   string               `db:"user_id"`
	Category datastore.NullString `db:"category"`
}

func GetOrgUserFieldMap() map[string]string {
	return map[string]string{"Category": "category", "OrgId": "org_id", "UserId": "user_id"}
}

type OpUserCategorie struct {
	OpId         string `db:"op_id"`
	UserCategory string `db:"user_category"`
}

func GetOpUserCategorieFieldMap() map[string]string {
	return map[string]string{"OpId": "op_id", "UserCategory": "user_category"}
}

type OpUser struct {
	OpId   string `db:"op_id"`
	UserId string `db:"user_id"`
	Access bool   `db:"access"`
}

func GetOpUserFieldMap() map[string]string {
	return map[string]string{"Access": "access", "OpId": "op_id", "UserId": "user_id"}
}

type OrgEndpointCategorie struct {
	OrgId            string `db:"org_id"`
	EndpointCategory string `db:"endpoint_category"`
}

func GetOrgEndpointCategorieFieldMap() map[string]string {
	return map[string]string{"EndpointCategory": "endpoint_category", "OrgId": "org_id"}
}

type OrgEndpoint struct {
	OrgId      string `db:"org_id"`
	EndpointId string `db:"endpoint_id"`
	Access     bool   `db:"access"`
}

func GetOrgEndpointFieldMap() map[string]string {
	return map[string]string{"Access": "access", "EndpointId": "endpoint_id", "OrgId": "org_id"}
}

type OpEndpointCategorie struct {
	OpId             string `db:"op_id"`
	EndpointCategory string `db:"endpoint_category"`
}

func GetOpEndpointCategorieFieldMap() map[string]string {
	return map[string]string{"EndpointCategory": "endpoint_category", "OpId": "op_id"}
}

type OpEndpoint struct {
	OpId       string `db:"op_id"`
	EndpointId string `db:"endpoint_id"`
	Access     bool   `db:"access"`
}

func GetOpEndpointFieldMap() map[string]string {
	return map[string]string{"Access": "access", "EndpointId": "endpoint_id", "OpId": "op_id"}
}

type UserEndpointCategorie struct {
	UserId           string `db:"user_id"`
	EndpointCategory string `db:"endpoint_category"`
}

func GetUserEndpointCategorieFieldMap() map[string]string {
	return map[string]string{"EndpointCategory": "endpoint_category", "UserId": "user_id"}
}

type UserEndpoint struct {
	UserId     string `db:"user_id"`
	EndpointId string `db:"endpoint_id"`
	Access     bool   `db:"access"`
}

func GetUserEndpointFieldMap() map[string]string {
	return map[string]string{"Access": "access", "EndpointId": "endpoint_id", "UserId": "user_id"}
}

type UserDetailView struct {
	AuthId    string `db:"auth_id"`
	Email     string `db:"email"`
	Mobile    string `db:"mobile"`
	Social    string `db:"social"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	UserId    string `db:"user_id"`
}

func GetUserDetailViewFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "Email": "email", "FirstName": "first_name", "LastName": "last_name", "Mobile": "mobile", "Social": "social", "UserId": "user_id"}
}

type OpDetailView struct {
	AuthId    string `db:"auth_id"`
	Email     string `db:"email"`
	Mobile    string `db:"mobile"`
	Social    string `db:"social"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	OpId      string `db:"op_id"`
}

func GetOpDetailViewFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "Email": "email", "FirstName": "first_name", "LastName": "last_name", "Mobile": "mobile", "OpId": "op_id", "Social": "social"}
}

type OrgDetailView struct {
	AuthId    string `db:"auth_id"`
	Email     string `db:"email"`
	Mobile    string `db:"mobile"`
	Social    string `db:"social"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	OrgId     string `db:"org_id"`
	OrgName   string `db:"org_name"`
}

func GetOrgDetailViewFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "Email": "email", "FirstName": "first_name", "LastName": "last_name", "Mobile": "mobile", "OrgId": "org_id", "OrgName": "org_name", "Social": "social"}
}

type Session struct {
	Id        string               `db:"id"`
	UserId    datastore.NullString `db:"user_id"`
	Name      datastore.NullString `db:"name"`
	Type      datastore.NullString `db:"type"`
	StartedAt datastore.NullTime   `db:"started_at"`
	EndedAt   datastore.NullTime   `db:"ended_at"`
}

func GetSessionFieldMap() map[string]string {
	return map[string]string{"EndedAt": "ended_at", "Id": "id", "Name": "name", "StartedAt": "started_at", "Type": "type", "UserId": "user_id"}
}

type SessionEvent struct {
	UserId     string               `db:"user_id"`
	Name       string               `db:"name"`
	StartedAt  time.Time            `db:"started_at"`
	EndedAt    datastore.NullTime   `db:"ended_at"`
	Properties datastore.RawMessage `db:"properties"`
}

func GetSessionEventFieldMap() map[string]string {
	return map[string]string{"EndedAt": "ended_at", "Name": "name", "Properties": "properties", "StartedAt": "started_at", "UserId": "user_id"}
}

type SessionRecord struct {
	UserId     string               `db:"user_id"`
	Name       string               `db:"name"`
	Timestamp  time.Time            `db:"timestamp"`
	Value      datastore.RawMessage `db:"value"`
	Properties datastore.RawMessage `db:"properties"`
}

func GetSessionRecordFieldMap() map[string]string {
	return map[string]string{"Name": "name", "Properties": "properties", "Timestamp": "timestamp", "UserId": "user_id", "Value": "value"}
}

type SessionPropertie struct {
	SessionId datastore.NullString `db:"session_id"`
	Name      datastore.NullString `db:"name"`
	Value     datastore.NullString `db:"value"`
	Rowid     datastore.RawMessage `db:"rowid"`
}

func GetSessionPropertieFieldMap() map[string]string {
	return map[string]string{"Name": "name", "Rowid": "rowid", "SessionId": "session_id", "Value": "value"}
}
