package models

import (
	"time"
)

var t time.Time

type Org struct {
	Id        string     `db:"id"`
	Name      NullString `db:"name"`
	CreatedAt NullTime   `db:"created_at"`
	UpdatedAt NullTime   `db:"updated_at"`
}

func GetOrgFieldMap() map[string]string {
	return map[string]string{"CreatedAt": "created_at", "Id": "id", "Name": "name", "UpdatedAt": "updated_at"}
}

type Op struct {
	Id        string   `db:"id"`
	CreatedAt NullTime `db:"created_at"`
	UpdatedAt NullTime `db:"updated_at"`
}

func GetOpFieldMap() map[string]string {
	return map[string]string{"CreatedAt": "created_at", "Id": "id", "UpdatedAt": "updated_at"}
}

type User struct {
	Id        string   `db:"id"`
	CreatedAt NullTime `db:"created_at"`
	UpdatedAt NullTime `db:"updated_at"`
}

func GetUserFieldMap() map[string]string {
	return map[string]string{"CreatedAt": "created_at", "Id": "id", "UpdatedAt": "updated_at"}
}

type Endpoint struct {
	Id   string     `db:"id"`
	Path NullString `db:"path"`
}

func GetEndpointFieldMap() map[string]string {
	return map[string]string{"Id": "id", "Path": "path"}
}

type Group struct {
	Id   string     `db:"id"`
	Name NullString `db:"name"`
}

func GetGroupFieldMap() map[string]string {
	return map[string]string{"Id": "id", "Name": "name"}
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

type EndpointGroup struct {
	EndpointId string `db:"endpoint_id"`
	GroupId    string `db:"group_id"`
}

func GetEndpointGroupFieldMap() map[string]string {
	return map[string]string{"EndpointId": "endpoint_id", "GroupId": "group_id"}
}

type Auth struct {
	Id        string     `db:"id"`
	Email     string     `db:"email"`
	Mobile    string     `db:"mobile"`
	Social    string     `db:"social"`
	FirstName NullString `db:"first_name"`
	LastName  NullString `db:"last_name"`
	CreatedAt NullTime   `db:"created_at"`
	UpdatedAt NullTime   `db:"updated_at"`
}

func GetAuthFieldMap() map[string]string {
	return map[string]string{"CreatedAt": "created_at", "Email": "email", "FirstName": "first_name", "Id": "id", "LastName": "last_name", "Mobile": "mobile", "Social": "social", "UpdatedAt": "updated_at"}
}
