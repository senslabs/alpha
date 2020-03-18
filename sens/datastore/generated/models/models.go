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
	Name      datastore.NullString `db:"name"`
	CreatedAt datastore.NullTime   `db:"created_at"`
	UpdatedAt datastore.NullTime   `db:"updated_at"`
}

func GetOrgFieldMap() map[string]string {
	return map[string]string{"CreatedAt": "created_at", "Id": "id", "Name": "name", "UpdatedAt": "updated_at"}
}

type Op struct {
	Id        string               `db:"id"`
	CreatedAt datastore.NullTime   `db:"created_at"`
	UpdatedAt datastore.NullTime   `db:"updated_at"`
	Status    datastore.NullString `db:"status"`
}

func GetOpFieldMap() map[string]string {
	return map[string]string{"CreatedAt": "created_at", "Id": "id", "Status": "status", "UpdatedAt": "updated_at"}
}

type User struct {
	Id        string               `db:"id"`
	CreatedAt datastore.NullTime   `db:"created_at"`
	UpdatedAt datastore.NullTime   `db:"updated_at"`
	Status    datastore.NullString `db:"status"`
}

func GetUserFieldMap() map[string]string {
	return map[string]string{"CreatedAt": "created_at", "Id": "id", "Status": "status", "UpdatedAt": "updated_at"}
}

type Endpoint struct {
	Id           string               `db:"id"`
	Category     datastore.NullString `db:"category"`
	Path         datastore.NullString `db:"path"`
	Secure       bool                 `db:"secure"`
	NextEndpoint datastore.NullString `db:"next_endpoint"`
}

func GetEndpointFieldMap() map[string]string {
	return map[string]string{"Category": "category", "Id": "id", "NextEndpoint": "next_endpoint", "Path": "path", "Secure": "secure"}
}

type Device struct {
	Id             string               `db:"id"`
	OrgId          datastore.NullString `db:"org_id"`
	UserId         datastore.NullString `db:"user_id"`
	RegisteredAt   datastore.NullTime   `db:"registered_at"`
	UnregisteredAt datastore.NullTime   `db:"unregistered_at"`
	PairedAt       datastore.NullTime   `db:"paired_at"`
	UnpairedAt     datastore.NullTime   `db:"unpaired_at"`
	Tags           datastore.RawMessage `db:"tags"`
	Status         datastore.NullString `db:"status"`
}

func GetDeviceFieldMap() map[string]string {
	return map[string]string{"Id": "id", "OrgId": "org_id", "PairedAt": "paired_at", "RegisteredAt": "registered_at", "Status": "status", "Tags": "tags", "UnpairedAt": "unpaired_at", "UnregisteredAt": "unregistered_at", "UserId": "user_id"}
}

type OrgAuth struct {
	OrgId  string `db:"org_id"`
	AuthId string `db:"auth_id"`
}

func GetOrgAuthFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "OrgId": "org_id"}
}

type OpAuth struct {
	OpId   string `db:"op_id"`
	AuthId string `db:"auth_id"`
}

func GetOpAuthFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "OpId": "op_id"}
}

type UserAuth struct {
	UserId string `db:"user_id"`
	AuthId string `db:"auth_id"`
}

func GetUserAuthFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "UserId": "user_id"}
}

type OrgOp struct {
	OrgId string `db:"org_id"`
	OpId  string `db:"op_id"`
}

func GetOrgOpFieldMap() map[string]string {
	return map[string]string{"OpId": "op_id", "OrgId": "org_id"}
}

type OrgUser struct {
	OrgId  string `db:"org_id"`
	UserId string `db:"user_id"`
}

func GetOrgUserFieldMap() map[string]string {
	return map[string]string{"OrgId": "org_id", "UserId": "user_id"}
}

type OpUser struct {
	OpId   string `db:"op_id"`
	UserId string `db:"user_id"`
}

func GetOpUserFieldMap() map[string]string {
	return map[string]string{"OpId": "op_id", "UserId": "user_id"}
}

type OrgEndpoint struct {
	OrgId      string `db:"org_id"`
	EndpointId string `db:"endpoint_id"`
}

func GetOrgEndpointFieldMap() map[string]string {
	return map[string]string{"EndpointId": "endpoint_id", "OrgId": "org_id"}
}

type OpEndpoint struct {
	OpId       string `db:"op_id"`
	EndpointId string `db:"endpoint_id"`
}

func GetOpEndpointFieldMap() map[string]string {
	return map[string]string{"EndpointId": "endpoint_id", "OpId": "op_id"}
}

type UserEndpoint struct {
	UserId     string `db:"user_id"`
	EndpointId string `db:"endpoint_id"`
}

func GetUserEndpointFieldMap() map[string]string {
	return map[string]string{"EndpointId": "endpoint_id", "UserId": "user_id"}
}

type OrgDetail struct {
	AuthId    string `db:"auth_id"`
	Email     string `db:"email"`
	Mobile    string `db:"mobile"`
	Social    string `db:"social"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	OrgId     string `db:"org_id"`
}

func GetOrgDetailFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "Email": "email", "FirstName": "first_name", "LastName": "last_name", "Mobile": "mobile", "OrgId": "org_id", "Social": "social"}
}

type OpDetail struct {
	AuthId    string `db:"auth_id"`
	Email     string `db:"email"`
	Mobile    string `db:"mobile"`
	Social    string `db:"social"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	OpId      string `db:"op_id"`
}

func GetOpDetailFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "Email": "email", "FirstName": "first_name", "LastName": "last_name", "Mobile": "mobile", "OpId": "op_id", "Social": "social"}
}

type UserDetail struct {
	AuthId    string `db:"auth_id"`
	Email     string `db:"email"`
	Mobile    string `db:"mobile"`
	Social    string `db:"social"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	UserId    string `db:"user_id"`
}

func GetUserDetailFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "Email": "email", "FirstName": "first_name", "LastName": "last_name", "Mobile": "mobile", "Social": "social", "UserId": "user_id"}
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
