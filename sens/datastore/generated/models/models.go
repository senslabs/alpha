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
	OrgName   *string `db:"org_name" json:",omitempty"`
	CreatedAt *int64  `db:"created_at" json:",omitempty"`
	UpdatedAt *int64  `db:"updated_at" json:",omitempty"`
}

func GetOrgFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "CreatedAt": "created_at", "Id": "id", "OrgName": "org_name", "UpdatedAt": "updated_at"}
}
func GetOrgTypeMap() map[string]string {
	return map[string]string{"AuthId": "*string", "CreatedAt": "*int64", "Id": "*string", "OrgName": "*string", "UpdatedAt": "*int64"}
}

type Op struct {
	Id        *string `db:"id" json:",omitempty"`
	AuthId    *string `db:"auth_id" json:",omitempty"`
	OrgId     *string `db:"org_id" json:",omitempty"`
	CreatedAt *int64  `db:"created_at" json:",omitempty"`
	UpdatedAt *int64  `db:"updated_at" json:",omitempty"`
	Status    *string `db:"status" json:",omitempty"`
}

func GetOpFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "CreatedAt": "created_at", "Id": "id", "OrgId": "org_id", "Status": "status", "UpdatedAt": "updated_at"}
}
func GetOpTypeMap() map[string]string {
	return map[string]string{"AuthId": "*string", "CreatedAt": "*int64", "Id": "*string", "OrgId": "*string", "Status": "*string", "UpdatedAt": "*int64"}
}

type User struct {
	Id          *string `db:"id" json:",omitempty"`
	AuthId      *string `db:"auth_id" json:",omitempty"`
	OrgId       *string `db:"org_id" json:",omitempty"`
	AccessGroup *string `db:"access_group" json:",omitempty"`
	CreatedAt   *int64  `db:"created_at" json:",omitempty"`
	UpdatedAt   *int64  `db:"updated_at" json:",omitempty"`
	Age         *int64  `db:"age" json:",omitempty"`
	Status      *string `db:"status" json:",omitempty"`
}

func GetUserFieldMap() map[string]string {
	return map[string]string{"AccessGroup": "access_group", "Age": "age", "AuthId": "auth_id", "CreatedAt": "created_at", "Id": "id", "OrgId": "org_id", "Status": "status", "UpdatedAt": "updated_at"}
}
func GetUserTypeMap() map[string]string {
	return map[string]string{"AccessGroup": "*string", "Age": "*int64", "AuthId": "*string", "CreatedAt": "*int64", "Id": "*string", "OrgId": "*string", "Status": "*string", "UpdatedAt": "*int64"}
}

type OpUserAccessGroup struct {
	OpId        *string `db:"op_id" json:",omitempty"`
	AccessGroup *string `db:"access_group" json:",omitempty"`
}

func GetOpUserAccessGroupFieldMap() map[string]string {
	return map[string]string{"AccessGroup": "access_group", "OpId": "op_id"}
}
func GetOpUserAccessGroupTypeMap() map[string]string {
	return map[string]string{"AccessGroup": "*string", "OpId": "*string"}
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

type Endpoint struct {
	Id          *string `db:"id" json:",omitempty"`
	AccessGroup *string `db:"access_group" json:",omitempty"`
	Path        *string `db:"path" json:",omitempty"`
	Secure      *bool   `db:"secure" json:",omitempty"`
}

func GetEndpointFieldMap() map[string]string {
	return map[string]string{"AccessGroup": "access_group", "Id": "id", "Path": "path", "Secure": "secure"}
}
func GetEndpointTypeMap() map[string]string {
	return map[string]string{"AccessGroup": "*string", "Id": "*string", "Path": "*string", "Secure": "*bool"}
}

type OrgEndpointAccessGroup struct {
	OrgId       *string `db:"org_id" json:",omitempty"`
	AccessGroup *string `db:"access_group" json:",omitempty"`
}

func GetOrgEndpointAccessGroupFieldMap() map[string]string {
	return map[string]string{"AccessGroup": "access_group", "OrgId": "org_id"}
}
func GetOrgEndpointAccessGroupTypeMap() map[string]string {
	return map[string]string{"AccessGroup": "*string", "OrgId": "*string"}
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

type OpEndpointAccessGroup struct {
	OpId        *string `db:"op_id" json:",omitempty"`
	AccessGroup *string `db:"access_group" json:",omitempty"`
}

func GetOpEndpointAccessGroupFieldMap() map[string]string {
	return map[string]string{"AccessGroup": "access_group", "OpId": "op_id"}
}
func GetOpEndpointAccessGroupTypeMap() map[string]string {
	return map[string]string{"AccessGroup": "*string", "OpId": "*string"}
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

type UserEndpointAccessGroup struct {
	UserId           *string `db:"user_id" json:",omitempty"`
	EndpointCategory *string `db:"endpoint_category" json:",omitempty"`
}

func GetUserEndpointAccessGroupFieldMap() map[string]string {
	return map[string]string{"EndpointCategory": "endpoint_category", "UserId": "user_id"}
}
func GetUserEndpointAccessGroupTypeMap() map[string]string {
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

type Device struct {
	Id        *string `db:"id" json:",omitempty"`
	DeviceId  *string `db:"device_id" json:",omitempty"`
	Name      *string `db:"name" json:",omitempty"`
	OrgId     *string `db:"org_id" json:",omitempty"`
	UserId    *string `db:"user_id" json:",omitempty"`
	CreatedAt *int64  `db:"created_at" json:",omitempty"`
	Status    *string `db:"status" json:",omitempty"`
}

func GetDeviceFieldMap() map[string]string {
	return map[string]string{"CreatedAt": "created_at", "DeviceId": "device_id", "Id": "id", "Name": "name", "OrgId": "org_id", "Status": "status", "UserId": "user_id"}
}
func GetDeviceTypeMap() map[string]string {
	return map[string]string{"CreatedAt": "*int64", "DeviceId": "*string", "Id": "*string", "Name": "*string", "OrgId": "*string", "Status": "*string", "UserId": "*string"}
}

type DeviceActivitie struct {
	DeviceId *string `db:"device_id" json:",omitempty"`
	ActiveAt *int64  `db:"active_at" json:",omitempty"`
	Type     *string `db:"type" json:",omitempty"`
}

func GetDeviceActivitieFieldMap() map[string]string {
	return map[string]string{"ActiveAt": "active_at", "DeviceId": "device_id", "Type": "type"}
}
func GetDeviceActivitieTypeMap() map[string]string {
	return map[string]string{"ActiveAt": "*int64", "DeviceId": "*string", "Type": "*string"}
}

type Alert struct {
	Id        *string `db:"id" json:",omitempty"`
	UserId    *string `db:"user_id" json:",omitempty"`
	CreatedAt *int64  `db:"created_at" json:",omitempty"`
	AlertName *string `db:"alert_name" json:",omitempty"`
	Status    *string `db:"status" json:",omitempty"`
	Remarks   *string `db:"remarks" json:",omitempty"`
}

func GetAlertFieldMap() map[string]string {
	return map[string]string{"AlertName": "alert_name", "CreatedAt": "created_at", "Id": "id", "Remarks": "remarks", "Status": "status", "UserId": "user_id"}
}
func GetAlertTypeMap() map[string]string {
	return map[string]string{"AlertName": "*string", "CreatedAt": "*int64", "Id": "*string", "Remarks": "*string", "Status": "*string", "UserId": "*string"}
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
	OrgId     *string `db:"org_id" json:",omitempty"`
}

func GetOpDetailViewFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "Email": "email", "FirstName": "first_name", "Id": "id", "LastName": "last_name", "Mobile": "mobile", "OrgId": "org_id", "Social": "social"}
}
func GetOpDetailViewTypeMap() map[string]string {
	return map[string]string{"AuthId": "*string", "Email": "*string", "FirstName": "*string", "Id": "*string", "LastName": "*string", "Mobile": "*string", "OrgId": "*string", "Social": "*string"}
}

type UserDetailView struct {
	AuthId    *string `db:"auth_id" json:",omitempty"`
	Email     *string `db:"email" json:",omitempty"`
	Mobile    *string `db:"mobile" json:",omitempty"`
	Social    *string `db:"social" json:",omitempty"`
	FirstName *string `db:"first_name" json:",omitempty"`
	LastName  *string `db:"last_name" json:",omitempty"`
	Id        *string `db:"id" json:",omitempty"`
	OrgId     *string `db:"org_id" json:",omitempty"`
}

func GetUserDetailViewFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "Email": "email", "FirstName": "first_name", "Id": "id", "LastName": "last_name", "Mobile": "mobile", "OrgId": "org_id", "Social": "social"}
}
func GetUserDetailViewTypeMap() map[string]string {
	return map[string]string{"AuthId": "*string", "Email": "*string", "FirstName": "*string", "Id": "*string", "LastName": "*string", "Mobile": "*string", "OrgId": "*string", "Social": "*string"}
}

type DeviceView struct {
	DeviceId  *string `db:"device_id" json:",omitempty"`
	Name      *string `db:"name" json:",omitempty"`
	OrgId     *string `db:"org_id" json:",omitempty"`
	UserId    *string `db:"user_id" json:",omitempty"`
	CreatedAt *int64  `db:"created_at" json:",omitempty"`
	Status    *string `db:"status" json:",omitempty"`
}

func GetDeviceViewFieldMap() map[string]string {
	return map[string]string{"CreatedAt": "created_at", "DeviceId": "device_id", "Name": "name", "OrgId": "org_id", "Status": "status", "UserId": "user_id"}
}
func GetDeviceViewTypeMap() map[string]string {
	return map[string]string{"CreatedAt": "*int64", "DeviceId": "*string", "Name": "*string", "OrgId": "*string", "Status": "*string", "UserId": "*string"}
}

type UserAlertView struct {
	UserId    *string `db:"user_id" json:",omitempty"`
	OrgId     *string `db:"org_id" json:",omitempty"`
	FirstName *string `db:"first_name" json:",omitempty"`
	LastName  *string `db:"last_name" json:",omitempty"`
	CreatedAt *int64  `db:"created_at" json:",omitempty"`
	AlertName *string `db:"alert_name" json:",omitempty"`
	Status    *string `db:"status" json:",omitempty"`
	Remarks   *string `db:"remarks" json:",omitempty"`
}

func GetUserAlertViewFieldMap() map[string]string {
	return map[string]string{"AlertName": "alert_name", "CreatedAt": "created_at", "FirstName": "first_name", "LastName": "last_name", "OrgId": "org_id", "Remarks": "remarks", "Status": "status", "UserId": "user_id"}
}
func GetUserAlertViewTypeMap() map[string]string {
	return map[string]string{"AlertName": "*string", "CreatedAt": "*int64", "FirstName": "*string", "LastName": "*string", "OrgId": "*string", "Remarks": "*string", "Status": "*string", "UserId": "*string"}
}

type SleepView struct {
	Id        *string `db:"id" json:",omitempty"`
	UserId    *string `db:"user_id" json:",omitempty"`
	Name      *string `db:"name" json:",omitempty"`
	Type      *string `db:"type" json:",omitempty"`
	StartedAt *int64  `db:"started_at" json:",omitempty"`
	EndedAt   *int64  `db:"ended_at" json:",omitempty"`
}

func GetSleepViewFieldMap() map[string]string {
	return map[string]string{"EndedAt": "ended_at", "Id": "id", "Name": "name", "StartedAt": "started_at", "Type": "type", "UserId": "user_id"}
}
func GetSleepViewTypeMap() map[string]string {
	return map[string]string{"EndedAt": "*int64", "Id": "*string", "Name": "*string", "StartedAt": "*int64", "Type": "*string", "UserId": "*string"}
}

type MeditationView struct {
	Id        *string `db:"id" json:",omitempty"`
	UserId    *string `db:"user_id" json:",omitempty"`
	Name      *string `db:"name" json:",omitempty"`
	Type      *string `db:"type" json:",omitempty"`
	StartedAt *int64  `db:"started_at" json:",omitempty"`
	EndedAt   *int64  `db:"ended_at" json:",omitempty"`
}

func GetMeditationViewFieldMap() map[string]string {
	return map[string]string{"EndedAt": "ended_at", "Id": "id", "Name": "name", "StartedAt": "started_at", "Type": "type", "UserId": "user_id"}
}
func GetMeditationViewTypeMap() map[string]string {
	return map[string]string{"EndedAt": "*int64", "Id": "*string", "Name": "*string", "StartedAt": "*int64", "Type": "*string", "UserId": "*string"}
}

type UserSessionView struct {
	Type      *string `db:"type" json:",omitempty"`
	Timestamp *int64  `db:"timestamp" json:",omitempty"`
	UserId    *string `db:"user_id" json:",omitempty"`
}

func GetUserSessionViewFieldMap() map[string]string {
	return map[string]string{"Timestamp": "timestamp", "Type": "type", "UserId": "user_id"}
}
func GetUserSessionViewTypeMap() map[string]string {
	return map[string]string{"Timestamp": "*int64", "Type": "*string", "UserId": "*string"}
}

type SessionPropertie struct {
	SessionId *string `db:"session_id" json:",omitempty"`
	Name      *string `db:"name" json:",omitempty"`
	Value     *string `db:"value" json:",omitempty"`
}

func GetSessionPropertieFieldMap() map[string]string {
	return map[string]string{"Name": "name", "SessionId": "session_id", "Value": "value"}
}
func GetSessionPropertieTypeMap() map[string]string {
	return map[string]string{"Name": "*string", "SessionId": "*string", "Value": "*string"}
}

type SleepSummarie struct {
	UserId     *string               `db:"user_id" json:",omitempty"`
	Duration   *int64                `db:"duration" json:",omitempty"`
	Properties *datastore.RawMessage `db:"properties" json:",omitempty"`
	SessionId  *string               `db:"session_id" json:",omitempty"`
}

func GetSleepSummarieFieldMap() map[string]string {
	return map[string]string{"Duration": "duration", "Properties": "properties", "SessionId": "session_id", "UserId": "user_id"}
}
func GetSleepSummarieTypeMap() map[string]string {
	return map[string]string{"Duration": "*int64", "Properties": "*datastore.RawMessage", "SessionId": "*string", "UserId": "*string"}
}

type MeditationSummarie struct {
	UserId     *string               `db:"user_id" json:",omitempty"`
	Duration   *int64                `db:"duration" json:",omitempty"`
	Properties *datastore.RawMessage `db:"properties" json:",omitempty"`
	SessionId  *string               `db:"session_id" json:",omitempty"`
}

func GetMeditationSummarieFieldMap() map[string]string {
	return map[string]string{"Duration": "duration", "Properties": "properties", "SessionId": "session_id", "UserId": "user_id"}
}
func GetMeditationSummarieTypeMap() map[string]string {
	return map[string]string{"Duration": "*int64", "Properties": "*datastore.RawMessage", "SessionId": "*string", "UserId": "*string"}
}

type UserSummaryView struct {
	Count  *int64  `db:"count" json:",omitempty"`
	Type   *string `db:"type" json:",omitempty"`
	UserId *string `db:"user_id" json:",omitempty"`
}

func GetUserSummaryViewFieldMap() map[string]string {
	return map[string]string{"Count": "count", "Type": "type", "UserId": "user_id"}
}
func GetUserSummaryViewTypeMap() map[string]string {
	return map[string]string{"Count": "*int64", "Type": "*string", "UserId": "*string"}
}

type UserSleepView struct {
	UserId     *string               `db:"user_id" json:",omitempty"`
	Duration   *int64                `db:"duration" json:",omitempty"`
	Properties *datastore.RawMessage `db:"properties" json:",omitempty"`
	SessionId  *string               `db:"session_id" json:",omitempty"`
}

func GetUserSleepViewFieldMap() map[string]string {
	return map[string]string{"Duration": "duration", "Properties": "properties", "SessionId": "session_id", "UserId": "user_id"}
}
func GetUserSleepViewTypeMap() map[string]string {
	return map[string]string{"Duration": "*int64", "Properties": "*datastore.RawMessage", "SessionId": "*string", "UserId": "*string"}
}

type UserMeditationView struct {
	UserId     *string               `db:"user_id" json:",omitempty"`
	Duration   *int64                `db:"duration" json:",omitempty"`
	Properties *datastore.RawMessage `db:"properties" json:",omitempty"`
	SessionId  *string               `db:"session_id" json:",omitempty"`
}

func GetUserMeditationViewFieldMap() map[string]string {
	return map[string]string{"Duration": "duration", "Properties": "properties", "SessionId": "session_id", "UserId": "user_id"}
}
func GetUserMeditationViewTypeMap() map[string]string {
	return map[string]string{"Duration": "*int64", "Properties": "*datastore.RawMessage", "SessionId": "*string", "UserId": "*string"}
}
