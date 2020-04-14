package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/senslabs/alpha/sens/datastore"
)

var t time.Time

type DeviceActivitie struct {
	DeviceId     *uuid.UUID `db:"device_id" json:",omitempty"`
	ActivityType *string    `db:"activity_type" json:",omitempty"`
	ActiveAt     *int64     `db:"active_at" json:",omitempty"`
}

func GetDeviceActivitieFieldMap() map[string]string {
	return map[string]string{"ActiveAt": "active_at", "ActivityType": "activity_type", "DeviceId": "device_id"}
}

func GetDeviceActivitieReverseFieldMap() map[string]string {
	return map[string]string{"active_at": "ActiveAt", "activity_type": "ActivityType", "device_id": "DeviceId"}
}

func GetDeviceActivitieTypeMap() map[string]string {
	return map[string]string{"ActiveAt": "*int64", "ActivityType": "*string", "DeviceId": "*uuid.UUID"}
}

type Auth struct {
	AuthId    *uuid.UUID `db:"auth_id" json:",omitempty"`
	Email     *string    `db:"email" json:",omitempty"`
	Mobile    *string    `db:"mobile" json:",omitempty"`
	Social    *string    `db:"social" json:",omitempty"`
	FirstName *string    `db:"first_name" json:",omitempty"`
	LastName  *string    `db:"last_name" json:",omitempty"`
	CreatedAt *int64     `db:"created_at" json:",omitempty"`
	UpdatedAt *int64     `db:"updated_at" json:",omitempty"`
	IsSens    *bool      `db:"is_sens" json:",omitempty"`
}

func GetAuthFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "CreatedAt": "created_at", "Email": "email", "FirstName": "first_name", "IsSens": "is_sens", "LastName": "last_name", "Mobile": "mobile", "Social": "social", "UpdatedAt": "updated_at"}
}

func GetAuthReverseFieldMap() map[string]string {
	return map[string]string{"auth_id": "AuthId", "created_at": "CreatedAt", "email": "Email", "first_name": "FirstName", "is_sens": "IsSens", "last_name": "LastName", "mobile": "Mobile", "social": "Social", "updated_at": "UpdatedAt"}
}

func GetAuthTypeMap() map[string]string {
	return map[string]string{"AuthId": "*uuid.UUID", "CreatedAt": "*int64", "Email": "*string", "FirstName": "*string", "IsSens": "*bool", "LastName": "*string", "Mobile": "*string", "Social": "*string", "UpdatedAt": "*int64"}
}

type Org struct {
	OrgId     *uuid.UUID `db:"org_id" json:",omitempty"`
	AuthId    *uuid.UUID `db:"auth_id" json:",omitempty"`
	OrgName   *string    `db:"org_name" json:",omitempty"`
	CreatedAt *int64     `db:"created_at" json:",omitempty"`
	UpdatedAt *int64     `db:"updated_at" json:",omitempty"`
}

func GetOrgFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "CreatedAt": "created_at", "OrgId": "org_id", "OrgName": "org_name", "UpdatedAt": "updated_at"}
}

func GetOrgReverseFieldMap() map[string]string {
	return map[string]string{"auth_id": "AuthId", "created_at": "CreatedAt", "org_id": "OrgId", "org_name": "OrgName", "updated_at": "UpdatedAt"}
}

func GetOrgTypeMap() map[string]string {
	return map[string]string{"AuthId": "*uuid.UUID", "CreatedAt": "*int64", "OrgId": "*uuid.UUID", "OrgName": "*string", "UpdatedAt": "*int64"}
}

type Op struct {
	OpId      *uuid.UUID `db:"op_id" json:",omitempty"`
	AuthId    *uuid.UUID `db:"auth_id" json:",omitempty"`
	OrgId     *uuid.UUID `db:"org_id" json:",omitempty"`
	CreatedAt *int64     `db:"created_at" json:",omitempty"`
	UpdatedAt *int64     `db:"updated_at" json:",omitempty"`
	Status    *string    `db:"status" json:",omitempty"`
}

func GetOpFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "CreatedAt": "created_at", "OpId": "op_id", "OrgId": "org_id", "Status": "status", "UpdatedAt": "updated_at"}
}

func GetOpReverseFieldMap() map[string]string {
	return map[string]string{"auth_id": "AuthId", "created_at": "CreatedAt", "op_id": "OpId", "org_id": "OrgId", "status": "Status", "updated_at": "UpdatedAt"}
}

func GetOpTypeMap() map[string]string {
	return map[string]string{"AuthId": "*uuid.UUID", "CreatedAt": "*int64", "OpId": "*uuid.UUID", "OrgId": "*uuid.UUID", "Status": "*string", "UpdatedAt": "*int64"}
}

type User struct {
	UserId      *uuid.UUID `db:"user_id" json:",omitempty"`
	AuthId      *uuid.UUID `db:"auth_id" json:",omitempty"`
	OrgId       *uuid.UUID `db:"org_id" json:",omitempty"`
	AccessGroup *string    `db:"access_group" json:",omitempty"`
	CreatedAt   *int64     `db:"created_at" json:",omitempty"`
	UpdatedAt   *int64     `db:"updated_at" json:",omitempty"`
	Age         *int64     `db:"age" json:",omitempty"`
	Status      *string    `db:"status" json:",omitempty"`
}

func GetUserFieldMap() map[string]string {
	return map[string]string{"AccessGroup": "access_group", "Age": "age", "AuthId": "auth_id", "CreatedAt": "created_at", "OrgId": "org_id", "Status": "status", "UpdatedAt": "updated_at", "UserId": "user_id"}
}

func GetUserReverseFieldMap() map[string]string {
	return map[string]string{"access_group": "AccessGroup", "age": "Age", "auth_id": "AuthId", "created_at": "CreatedAt", "org_id": "OrgId", "status": "Status", "updated_at": "UpdatedAt", "user_id": "UserId"}
}

func GetUserTypeMap() map[string]string {
	return map[string]string{"AccessGroup": "*string", "Age": "*int64", "AuthId": "*uuid.UUID", "CreatedAt": "*int64", "OrgId": "*uuid.UUID", "Status": "*string", "UpdatedAt": "*int64", "UserId": "*uuid.UUID"}
}

type SurveyQuestion struct {
	SurveyQuestionId *uuid.UUID `db:"survey_question_id" json:",omitempty"`
	Question         *string    `db:"question" json:",omitempty"`
}

func GetSurveyQuestionFieldMap() map[string]string {
	return map[string]string{"Question": "question", "SurveyQuestionId": "survey_question_id"}
}

func GetSurveyQuestionReverseFieldMap() map[string]string {
	return map[string]string{"question": "Question", "survey_question_id": "SurveyQuestionId"}
}

func GetSurveyQuestionTypeMap() map[string]string {
	return map[string]string{"Question": "*string", "SurveyQuestionId": "*uuid.UUID"}
}

type SurveyAnswer struct {
	SurveyAnswerId   *uuid.UUID `db:"survey_answer_id" json:",omitempty"`
	SurveyQuestionId *uuid.UUID `db:"survey_question_id" json:",omitempty"`
	Answer           *string    `db:"answer" json:",omitempty"`
}

func GetSurveyAnswerFieldMap() map[string]string {
	return map[string]string{"Answer": "answer", "SurveyAnswerId": "survey_answer_id", "SurveyQuestionId": "survey_question_id"}
}

func GetSurveyAnswerReverseFieldMap() map[string]string {
	return map[string]string{"answer": "Answer", "survey_answer_id": "SurveyAnswerId", "survey_question_id": "SurveyQuestionId"}
}

func GetSurveyAnswerTypeMap() map[string]string {
	return map[string]string{"Answer": "*string", "SurveyAnswerId": "*uuid.UUID", "SurveyQuestionId": "*uuid.UUID"}
}

type UserSetting struct {
	UserSettingId *uuid.UUID `db:"user_setting_id" json:",omitempty"`
	UserId        *uuid.UUID `db:"user_id" json:",omitempty"`
	CreatedAt     *int64     `db:"created_at" json:",omitempty"`
	Key           *string    `db:"key" json:",omitempty"`
	Value         *string    `db:"value" json:",omitempty"`
}

func GetUserSettingFieldMap() map[string]string {
	return map[string]string{"CreatedAt": "created_at", "Key": "key", "UserId": "user_id", "UserSettingId": "user_setting_id", "Value": "value"}
}

func GetUserSettingReverseFieldMap() map[string]string {
	return map[string]string{"created_at": "CreatedAt", "key": "Key", "user_id": "UserId", "user_setting_id": "UserSettingId", "value": "Value"}
}

func GetUserSettingTypeMap() map[string]string {
	return map[string]string{"CreatedAt": "*int64", "Key": "*string", "UserId": "*uuid.UUID", "UserSettingId": "*uuid.UUID", "Value": "*string"}
}

type UserPropertie struct {
	UserId *uuid.UUID `db:"user_id" json:",omitempty"`
	Key    *string    `db:"key" json:",omitempty"`
	Value  *string    `db:"value" json:",omitempty"`
}

func GetUserPropertieFieldMap() map[string]string {
	return map[string]string{"Key": "key", "UserId": "user_id", "Value": "value"}
}

func GetUserPropertieReverseFieldMap() map[string]string {
	return map[string]string{"key": "Key", "user_id": "UserId", "value": "Value"}
}

func GetUserPropertieTypeMap() map[string]string {
	return map[string]string{"Key": "*string", "UserId": "*uuid.UUID", "Value": "*string"}
}

type ApiKey struct {
	ApiKeyId    *uuid.UUID `db:"api_key_id" json:",omitempty"`
	OrgId       *uuid.UUID `db:"org_id" json:",omitempty"`
	KeyName     *string    `db:"key_name" json:",omitempty"`
	Description *string    `db:"description" json:",omitempty"`
	Key         *string    `db:"key" json:",omitempty"`
}

func GetApiKeyFieldMap() map[string]string {
	return map[string]string{"ApiKeyId": "api_key_id", "Description": "description", "Key": "key", "KeyName": "key_name", "OrgId": "org_id"}
}

func GetApiKeyReverseFieldMap() map[string]string {
	return map[string]string{"api_key_id": "ApiKeyId", "description": "Description", "key": "Key", "key_name": "KeyName", "org_id": "OrgId"}
}

func GetApiKeyTypeMap() map[string]string {
	return map[string]string{"ApiKeyId": "*uuid.UUID", "Description": "*string", "Key": "*string", "KeyName": "*string", "OrgId": "*uuid.UUID"}
}

type OpUserAccessGroup struct {
	OpId        *uuid.UUID `db:"op_id" json:",omitempty"`
	AccessGroup *string    `db:"access_group" json:",omitempty"`
}

func GetOpUserAccessGroupFieldMap() map[string]string {
	return map[string]string{"AccessGroup": "access_group", "OpId": "op_id"}
}

func GetOpUserAccessGroupReverseFieldMap() map[string]string {
	return map[string]string{"access_group": "AccessGroup", "op_id": "OpId"}
}

func GetOpUserAccessGroupTypeMap() map[string]string {
	return map[string]string{"AccessGroup": "*string", "OpId": "*uuid.UUID"}
}

type OpUser struct {
	OpId   *uuid.UUID `db:"op_id" json:",omitempty"`
	UserId *uuid.UUID `db:"user_id" json:",omitempty"`
	Access *bool      `db:"access" json:",omitempty"`
}

func GetOpUserFieldMap() map[string]string {
	return map[string]string{"Access": "access", "OpId": "op_id", "UserId": "user_id"}
}

func GetOpUserReverseFieldMap() map[string]string {
	return map[string]string{"access": "Access", "op_id": "OpId", "user_id": "UserId"}
}

func GetOpUserTypeMap() map[string]string {
	return map[string]string{"Access": "*bool", "OpId": "*uuid.UUID", "UserId": "*uuid.UUID"}
}

type OrgEndpointAccessGroup struct {
	OrgId       *uuid.UUID `db:"org_id" json:",omitempty"`
	AccessGroup *string    `db:"access_group" json:",omitempty"`
}

func GetOrgEndpointAccessGroupFieldMap() map[string]string {
	return map[string]string{"AccessGroup": "access_group", "OrgId": "org_id"}
}

func GetOrgEndpointAccessGroupReverseFieldMap() map[string]string {
	return map[string]string{"access_group": "AccessGroup", "org_id": "OrgId"}
}

func GetOrgEndpointAccessGroupTypeMap() map[string]string {
	return map[string]string{"AccessGroup": "*string", "OrgId": "*uuid.UUID"}
}

type OrgEndpoint struct {
	OrgId      *uuid.UUID `db:"org_id" json:",omitempty"`
	EndpointId *uuid.UUID `db:"endpoint_id" json:",omitempty"`
	Access     *bool      `db:"access" json:",omitempty"`
}

func GetOrgEndpointFieldMap() map[string]string {
	return map[string]string{"Access": "access", "EndpointId": "endpoint_id", "OrgId": "org_id"}
}

func GetOrgEndpointReverseFieldMap() map[string]string {
	return map[string]string{"access": "Access", "endpoint_id": "EndpointId", "org_id": "OrgId"}
}

func GetOrgEndpointTypeMap() map[string]string {
	return map[string]string{"Access": "*bool", "EndpointId": "*uuid.UUID", "OrgId": "*uuid.UUID"}
}

type Endpoint struct {
	EndpointId  *uuid.UUID `db:"endpoint_id" json:",omitempty"`
	AccessGroup *string    `db:"access_group" json:",omitempty"`
	Path        *string    `db:"path" json:",omitempty"`
	Secure      *bool      `db:"secure" json:",omitempty"`
}

func GetEndpointFieldMap() map[string]string {
	return map[string]string{"AccessGroup": "access_group", "EndpointId": "endpoint_id", "Path": "path", "Secure": "secure"}
}

func GetEndpointReverseFieldMap() map[string]string {
	return map[string]string{"access_group": "AccessGroup", "endpoint_id": "EndpointId", "path": "Path", "secure": "Secure"}
}

func GetEndpointTypeMap() map[string]string {
	return map[string]string{"AccessGroup": "*string", "EndpointId": "*uuid.UUID", "Path": "*string", "Secure": "*bool"}
}

type OpEndpointAccessGroup struct {
	OpId        *uuid.UUID `db:"op_id" json:",omitempty"`
	AccessGroup *string    `db:"access_group" json:",omitempty"`
}

func GetOpEndpointAccessGroupFieldMap() map[string]string {
	return map[string]string{"AccessGroup": "access_group", "OpId": "op_id"}
}

func GetOpEndpointAccessGroupReverseFieldMap() map[string]string {
	return map[string]string{"access_group": "AccessGroup", "op_id": "OpId"}
}

func GetOpEndpointAccessGroupTypeMap() map[string]string {
	return map[string]string{"AccessGroup": "*string", "OpId": "*uuid.UUID"}
}

type OpEndpoint struct {
	OpId       *uuid.UUID `db:"op_id" json:",omitempty"`
	EndpointId *uuid.UUID `db:"endpoint_id" json:",omitempty"`
	Access     *bool      `db:"access" json:",omitempty"`
}

func GetOpEndpointFieldMap() map[string]string {
	return map[string]string{"Access": "access", "EndpointId": "endpoint_id", "OpId": "op_id"}
}

func GetOpEndpointReverseFieldMap() map[string]string {
	return map[string]string{"access": "Access", "endpoint_id": "EndpointId", "op_id": "OpId"}
}

func GetOpEndpointTypeMap() map[string]string {
	return map[string]string{"Access": "*bool", "EndpointId": "*uuid.UUID", "OpId": "*uuid.UUID"}
}

type UserEndpointAccessGroup struct {
	UserId           *uuid.UUID `db:"user_id" json:",omitempty"`
	EndpointCategory *string    `db:"endpoint_category" json:",omitempty"`
}

func GetUserEndpointAccessGroupFieldMap() map[string]string {
	return map[string]string{"EndpointCategory": "endpoint_category", "UserId": "user_id"}
}

func GetUserEndpointAccessGroupReverseFieldMap() map[string]string {
	return map[string]string{"endpoint_category": "EndpointCategory", "user_id": "UserId"}
}

func GetUserEndpointAccessGroupTypeMap() map[string]string {
	return map[string]string{"EndpointCategory": "*string", "UserId": "*uuid.UUID"}
}

type UserEndpoint struct {
	UserId     *uuid.UUID `db:"user_id" json:",omitempty"`
	EndpointId *uuid.UUID `db:"endpoint_id" json:",omitempty"`
	Access     *bool      `db:"access" json:",omitempty"`
}

func GetUserEndpointFieldMap() map[string]string {
	return map[string]string{"Access": "access", "EndpointId": "endpoint_id", "UserId": "user_id"}
}

func GetUserEndpointReverseFieldMap() map[string]string {
	return map[string]string{"access": "Access", "endpoint_id": "EndpointId", "user_id": "UserId"}
}

func GetUserEndpointTypeMap() map[string]string {
	return map[string]string{"Access": "*bool", "EndpointId": "*uuid.UUID", "UserId": "*uuid.UUID"}
}

type Device struct {
	DeviceId   *uuid.UUID `db:"device_id" json:",omitempty"`
	CreatedAt  *int64     `db:"created_at" json:",omitempty"`
	DeviceName *string    `db:"device_name" json:",omitempty"`
	OrgId      *uuid.UUID `db:"org_id" json:",omitempty"`
	UserId     *uuid.UUID `db:"user_id" json:",omitempty"`
	Status     *string    `db:"status" json:",omitempty"`
}

func GetDeviceFieldMap() map[string]string {
	return map[string]string{"CreatedAt": "created_at", "DeviceId": "device_id", "DeviceName": "device_name", "OrgId": "org_id", "Status": "status", "UserId": "user_id"}
}

func GetDeviceReverseFieldMap() map[string]string {
	return map[string]string{"created_at": "CreatedAt", "device_id": "DeviceId", "device_name": "DeviceName", "org_id": "OrgId", "status": "Status", "user_id": "UserId"}
}

func GetDeviceTypeMap() map[string]string {
	return map[string]string{"CreatedAt": "*int64", "DeviceId": "*uuid.UUID", "DeviceName": "*string", "OrgId": "*uuid.UUID", "Status": "*string", "UserId": "*uuid.UUID"}
}

type Alert struct {
	AlertId   *uuid.UUID `db:"alert_id" json:",omitempty"`
	UserId    *uuid.UUID `db:"user_id" json:",omitempty"`
	CreatedAt *int64     `db:"created_at" json:",omitempty"`
	AlertName *string    `db:"alert_name" json:",omitempty"`
	Status    *string    `db:"status" json:",omitempty"`
	Remarks   *string    `db:"remarks" json:",omitempty"`
}

func GetAlertFieldMap() map[string]string {
	return map[string]string{"AlertId": "alert_id", "AlertName": "alert_name", "CreatedAt": "created_at", "Remarks": "remarks", "Status": "status", "UserId": "user_id"}
}

func GetAlertReverseFieldMap() map[string]string {
	return map[string]string{"alert_id": "AlertId", "alert_name": "AlertName", "created_at": "CreatedAt", "remarks": "Remarks", "status": "Status", "user_id": "UserId"}
}

func GetAlertTypeMap() map[string]string {
	return map[string]string{"AlertId": "*uuid.UUID", "AlertName": "*string", "CreatedAt": "*int64", "Remarks": "*string", "Status": "*string", "UserId": "*uuid.UUID"}
}

type Session struct {
	SessionId   *uuid.UUID `db:"session_id" json:",omitempty"`
	UserId      *uuid.UUID `db:"user_id" json:",omitempty"`
	SessionName *string    `db:"session_name" json:",omitempty"`
	SessionType *string    `db:"session_type" json:",omitempty"`
	StartedAt   *int64     `db:"started_at" json:",omitempty"`
	EndedAt     *int64     `db:"ended_at" json:",omitempty"`
}

func GetSessionFieldMap() map[string]string {
	return map[string]string{"EndedAt": "ended_at", "SessionId": "session_id", "SessionName": "session_name", "SessionType": "session_type", "StartedAt": "started_at", "UserId": "user_id"}
}

func GetSessionReverseFieldMap() map[string]string {
	return map[string]string{"ended_at": "EndedAt", "session_id": "SessionId", "session_name": "SessionName", "session_type": "SessionType", "started_at": "StartedAt", "user_id": "UserId"}
}

func GetSessionTypeMap() map[string]string {
	return map[string]string{"EndedAt": "*int64", "SessionId": "*uuid.UUID", "SessionName": "*string", "SessionType": "*string", "StartedAt": "*int64", "UserId": "*uuid.UUID"}
}

type SessionSetting struct {
	SessionSettingId *uuid.UUID `db:"session_setting_id" json:",omitempty"`
	UserId           *uuid.UUID `db:"user_id" json:",omitempty"`
	Key              *string    `db:"key" json:",omitempty"`
	CreatedAt        *int64     `db:"created_at" json:",omitempty"`
	SessionType      *string    `db:"session_type" json:",omitempty"`
	Value            *string    `db:"value" json:",omitempty"`
}

func GetSessionSettingFieldMap() map[string]string {
	return map[string]string{"CreatedAt": "created_at", "Key": "key", "SessionSettingId": "session_setting_id", "SessionType": "session_type", "UserId": "user_id", "Value": "value"}
}

func GetSessionSettingReverseFieldMap() map[string]string {
	return map[string]string{"created_at": "CreatedAt", "key": "Key", "session_setting_id": "SessionSettingId", "session_type": "SessionType", "user_id": "UserId", "value": "Value"}
}

func GetSessionSettingTypeMap() map[string]string {
	return map[string]string{"CreatedAt": "*int64", "Key": "*string", "SessionSettingId": "*uuid.UUID", "SessionType": "*string", "UserId": "*uuid.UUID", "Value": "*string"}
}

type VitalBaseline struct {
	VitalBaselineId *uuid.UUID `db:"vital_baseline_id" json:",omitempty"`
	UserId          *uuid.UUID `db:"user_id" json:",omitempty"`
	Key             *string    `db:"key" json:",omitempty"`
	CreatedAt       *int64     `db:"created_at" json:",omitempty"`
	LowerLimit      *int64     `db:"lower_limit" json:",omitempty"`
	UpperLimit      *int64     `db:"upper_limit" json:",omitempty"`
}

func GetVitalBaselineFieldMap() map[string]string {
	return map[string]string{"CreatedAt": "created_at", "Key": "key", "LowerLimit": "lower_limit", "UpperLimit": "upper_limit", "UserId": "user_id", "VitalBaselineId": "vital_baseline_id"}
}

func GetVitalBaselineReverseFieldMap() map[string]string {
	return map[string]string{"created_at": "CreatedAt", "key": "Key", "lower_limit": "LowerLimit", "upper_limit": "UpperLimit", "user_id": "UserId", "vital_baseline_id": "VitalBaselineId"}
}

func GetVitalBaselineTypeMap() map[string]string {
	return map[string]string{"CreatedAt": "*int64", "Key": "*string", "LowerLimit": "*int64", "UpperLimit": "*int64", "UserId": "*uuid.UUID", "VitalBaselineId": "*uuid.UUID"}
}

type SessionEvent struct {
	UserId     *uuid.UUID            `db:"user_id" json:",omitempty"`
	Key        *string               `db:"key" json:",omitempty"`
	StartedAt  *int64                `db:"started_at" json:",omitempty"`
	EndedAt    *int64                `db:"ended_at" json:",omitempty"`
	Properties *datastore.RawMessage `db:"properties" json:",omitempty"`
}

func GetSessionEventFieldMap() map[string]string {
	return map[string]string{"EndedAt": "ended_at", "Key": "key", "Properties": "properties", "StartedAt": "started_at", "UserId": "user_id"}
}

func GetSessionEventReverseFieldMap() map[string]string {
	return map[string]string{"ended_at": "EndedAt", "key": "Key", "properties": "Properties", "started_at": "StartedAt", "user_id": "UserId"}
}

func GetSessionEventTypeMap() map[string]string {
	return map[string]string{"EndedAt": "*int64", "Key": "*string", "Properties": "*datastore.RawMessage", "StartedAt": "*int64", "UserId": "*uuid.UUID"}
}

type SessionRecord struct {
	UserId     *uuid.UUID            `db:"user_id" json:",omitempty"`
	Key        *string               `db:"key" json:",omitempty"`
	Timestamp  *int64                `db:"timestamp" json:",omitempty"`
	Value      *float64              `db:"value" json:",omitempty"`
	Properties *datastore.RawMessage `db:"properties" json:",omitempty"`
}

func GetSessionRecordFieldMap() map[string]string {
	return map[string]string{"Key": "key", "Properties": "properties", "Timestamp": "timestamp", "UserId": "user_id", "Value": "value"}
}

func GetSessionRecordReverseFieldMap() map[string]string {
	return map[string]string{"key": "Key", "properties": "Properties", "timestamp": "Timestamp", "user_id": "UserId", "value": "Value"}
}

func GetSessionRecordTypeMap() map[string]string {
	return map[string]string{"Key": "*string", "Properties": "*datastore.RawMessage", "Timestamp": "*int64", "UserId": "*uuid.UUID", "Value": "*float64"}
}

type SessionPropertie struct {
	SessionId *uuid.UUID `db:"session_id" json:",omitempty"`
	Key       *string    `db:"key" json:",omitempty"`
	Value     *string    `db:"value" json:",omitempty"`
}

func GetSessionPropertieFieldMap() map[string]string {
	return map[string]string{"Key": "key", "SessionId": "session_id", "Value": "value"}
}

func GetSessionPropertieReverseFieldMap() map[string]string {
	return map[string]string{"key": "Key", "session_id": "SessionId", "value": "Value"}
}

func GetSessionPropertieTypeMap() map[string]string {
	return map[string]string{"Key": "*string", "SessionId": "*uuid.UUID", "Value": "*string"}
}

type AuthDetailView struct {
	AuthId    *uuid.UUID `db:"auth_id" json:",omitempty"`
	Email     *string    `db:"email" json:",omitempty"`
	Mobile    *string    `db:"mobile" json:",omitempty"`
	Social    *string    `db:"social" json:",omitempty"`
	FirstName *string    `db:"first_name" json:",omitempty"`
	LastName  *string    `db:"last_name" json:",omitempty"`
	IsSens    *bool      `db:"is_sens" json:",omitempty"`
}

func GetAuthDetailViewFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "Email": "email", "FirstName": "first_name", "IsSens": "is_sens", "LastName": "last_name", "Mobile": "mobile", "Social": "social"}
}

func GetAuthDetailViewReverseFieldMap() map[string]string {
	return map[string]string{"auth_id": "AuthId", "email": "Email", "first_name": "FirstName", "is_sens": "IsSens", "last_name": "LastName", "mobile": "Mobile", "social": "Social"}
}

func GetAuthDetailViewTypeMap() map[string]string {
	return map[string]string{"AuthId": "*uuid.UUID", "Email": "*string", "FirstName": "*string", "IsSens": "*bool", "LastName": "*string", "Mobile": "*string", "Social": "*string"}
}

type OrgDetailView struct {
	AuthId    *uuid.UUID `db:"auth_id" json:",omitempty"`
	Email     *string    `db:"email" json:",omitempty"`
	Mobile    *string    `db:"mobile" json:",omitempty"`
	Social    *string    `db:"social" json:",omitempty"`
	FirstName *string    `db:"first_name" json:",omitempty"`
	LastName  *string    `db:"last_name" json:",omitempty"`
	OrgId     *uuid.UUID `db:"org_id" json:",omitempty"`
	OrgName   *string    `db:"org_name" json:",omitempty"`
}

func GetOrgDetailViewFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "Email": "email", "FirstName": "first_name", "LastName": "last_name", "Mobile": "mobile", "OrgId": "org_id", "OrgName": "org_name", "Social": "social"}
}

func GetOrgDetailViewReverseFieldMap() map[string]string {
	return map[string]string{"auth_id": "AuthId", "email": "Email", "first_name": "FirstName", "last_name": "LastName", "mobile": "Mobile", "org_id": "OrgId", "org_name": "OrgName", "social": "Social"}
}

func GetOrgDetailViewTypeMap() map[string]string {
	return map[string]string{"AuthId": "*uuid.UUID", "Email": "*string", "FirstName": "*string", "LastName": "*string", "Mobile": "*string", "OrgId": "*uuid.UUID", "OrgName": "*string", "Social": "*string"}
}

type OpDetailView struct {
	AuthId    *uuid.UUID `db:"auth_id" json:",omitempty"`
	Email     *string    `db:"email" json:",omitempty"`
	Mobile    *string    `db:"mobile" json:",omitempty"`
	Social    *string    `db:"social" json:",omitempty"`
	FirstName *string    `db:"first_name" json:",omitempty"`
	LastName  *string    `db:"last_name" json:",omitempty"`
	OpId      *uuid.UUID `db:"op_id" json:",omitempty"`
	OrgId     *uuid.UUID `db:"org_id" json:",omitempty"`
}

func GetOpDetailViewFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "Email": "email", "FirstName": "first_name", "LastName": "last_name", "Mobile": "mobile", "OpId": "op_id", "OrgId": "org_id", "Social": "social"}
}

func GetOpDetailViewReverseFieldMap() map[string]string {
	return map[string]string{"auth_id": "AuthId", "email": "Email", "first_name": "FirstName", "last_name": "LastName", "mobile": "Mobile", "op_id": "OpId", "org_id": "OrgId", "social": "Social"}
}

func GetOpDetailViewTypeMap() map[string]string {
	return map[string]string{"AuthId": "*uuid.UUID", "Email": "*string", "FirstName": "*string", "LastName": "*string", "Mobile": "*string", "OpId": "*uuid.UUID", "OrgId": "*uuid.UUID", "Social": "*string"}
}

type UserDetailView struct {
	AuthId    *uuid.UUID `db:"auth_id" json:",omitempty"`
	Email     *string    `db:"email" json:",omitempty"`
	Mobile    *string    `db:"mobile" json:",omitempty"`
	Social    *string    `db:"social" json:",omitempty"`
	FirstName *string    `db:"first_name" json:",omitempty"`
	LastName  *string    `db:"last_name" json:",omitempty"`
	UserId    *uuid.UUID `db:"user_id" json:",omitempty"`
	OrgId     *uuid.UUID `db:"org_id" json:",omitempty"`
}

func GetUserDetailViewFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "Email": "email", "FirstName": "first_name", "LastName": "last_name", "Mobile": "mobile", "OrgId": "org_id", "Social": "social", "UserId": "user_id"}
}

func GetUserDetailViewReverseFieldMap() map[string]string {
	return map[string]string{"auth_id": "AuthId", "email": "Email", "first_name": "FirstName", "last_name": "LastName", "mobile": "Mobile", "org_id": "OrgId", "social": "Social", "user_id": "UserId"}
}

func GetUserDetailViewTypeMap() map[string]string {
	return map[string]string{"AuthId": "*uuid.UUID", "Email": "*string", "FirstName": "*string", "LastName": "*string", "Mobile": "*string", "OrgId": "*uuid.UUID", "Social": "*string", "UserId": "*uuid.UUID"}
}

type DeviceView struct {
	DeviceId   *uuid.UUID `db:"device_id" json:",omitempty"`
	DeviceName *string    `db:"device_name" json:",omitempty"`
	OrgId      *uuid.UUID `db:"org_id" json:",omitempty"`
	UserId     *uuid.UUID `db:"user_id" json:",omitempty"`
	CreatedAt  *int64     `db:"created_at" json:",omitempty"`
	Status     *string    `db:"status" json:",omitempty"`
}

func GetDeviceViewFieldMap() map[string]string {
	return map[string]string{"CreatedAt": "created_at", "DeviceId": "device_id", "DeviceName": "device_name", "OrgId": "org_id", "Status": "status", "UserId": "user_id"}
}

func GetDeviceViewReverseFieldMap() map[string]string {
	return map[string]string{"created_at": "CreatedAt", "device_id": "DeviceId", "device_name": "DeviceName", "org_id": "OrgId", "status": "Status", "user_id": "UserId"}
}

func GetDeviceViewTypeMap() map[string]string {
	return map[string]string{"CreatedAt": "*int64", "DeviceId": "*uuid.UUID", "DeviceName": "*string", "OrgId": "*uuid.UUID", "Status": "*string", "UserId": "*uuid.UUID"}
}

type UserSessionView struct {
	ActivityType *string    `db:"activity_type" json:",omitempty"`
	Timestamp    *int64     `db:"timestamp" json:",omitempty"`
	UserId       *uuid.UUID `db:"user_id" json:",omitempty"`
}

func GetUserSessionViewFieldMap() map[string]string {
	return map[string]string{"ActivityType": "activity_type", "Timestamp": "timestamp", "UserId": "user_id"}
}

func GetUserSessionViewReverseFieldMap() map[string]string {
	return map[string]string{"activity_type": "ActivityType", "timestamp": "Timestamp", "user_id": "UserId"}
}

func GetUserSessionViewTypeMap() map[string]string {
	return map[string]string{"ActivityType": "*string", "Timestamp": "*int64", "UserId": "*uuid.UUID"}
}

type UserAlertView struct {
	UserId    *uuid.UUID `db:"user_id" json:",omitempty"`
	OrgId     *uuid.UUID `db:"org_id" json:",omitempty"`
	FirstName *string    `db:"first_name" json:",omitempty"`
	LastName  *string    `db:"last_name" json:",omitempty"`
	CreatedAt *int64     `db:"created_at" json:",omitempty"`
	AlertName *string    `db:"alert_name" json:",omitempty"`
	Status    *string    `db:"status" json:",omitempty"`
	Remarks   *string    `db:"remarks" json:",omitempty"`
}

func GetUserAlertViewFieldMap() map[string]string {
	return map[string]string{"AlertName": "alert_name", "CreatedAt": "created_at", "FirstName": "first_name", "LastName": "last_name", "OrgId": "org_id", "Remarks": "remarks", "Status": "status", "UserId": "user_id"}
}

func GetUserAlertViewReverseFieldMap() map[string]string {
	return map[string]string{"alert_name": "AlertName", "created_at": "CreatedAt", "first_name": "FirstName", "last_name": "LastName", "org_id": "OrgId", "remarks": "Remarks", "status": "Status", "user_id": "UserId"}
}

func GetUserAlertViewTypeMap() map[string]string {
	return map[string]string{"AlertName": "*string", "CreatedAt": "*int64", "FirstName": "*string", "LastName": "*string", "OrgId": "*uuid.UUID", "Remarks": "*string", "Status": "*string", "UserId": "*uuid.UUID"}
}

type SleepView struct {
	SessionId   *uuid.UUID `db:"session_id" json:",omitempty"`
	UserId      *uuid.UUID `db:"user_id" json:",omitempty"`
	SessionName *string    `db:"session_name" json:",omitempty"`
	SessionType *string    `db:"session_type" json:",omitempty"`
	StartedAt   *int64     `db:"started_at" json:",omitempty"`
	EndedAt     *int64     `db:"ended_at" json:",omitempty"`
}

func GetSleepViewFieldMap() map[string]string {
	return map[string]string{"EndedAt": "ended_at", "SessionId": "session_id", "SessionName": "session_name", "SessionType": "session_type", "StartedAt": "started_at", "UserId": "user_id"}
}

func GetSleepViewReverseFieldMap() map[string]string {
	return map[string]string{"ended_at": "EndedAt", "session_id": "SessionId", "session_name": "SessionName", "session_type": "SessionType", "started_at": "StartedAt", "user_id": "UserId"}
}

func GetSleepViewTypeMap() map[string]string {
	return map[string]string{"EndedAt": "*int64", "SessionId": "*uuid.UUID", "SessionName": "*string", "SessionType": "*string", "StartedAt": "*int64", "UserId": "*uuid.UUID"}
}

type MeditationView struct {
	SessionId   *uuid.UUID `db:"session_id" json:",omitempty"`
	UserId      *uuid.UUID `db:"user_id" json:",omitempty"`
	SessionName *string    `db:"session_name" json:",omitempty"`
	SessionType *string    `db:"session_type" json:",omitempty"`
	StartedAt   *int64     `db:"started_at" json:",omitempty"`
	EndedAt     *int64     `db:"ended_at" json:",omitempty"`
}

func GetMeditationViewFieldMap() map[string]string {
	return map[string]string{"EndedAt": "ended_at", "SessionId": "session_id", "SessionName": "session_name", "SessionType": "session_type", "StartedAt": "started_at", "UserId": "user_id"}
}

func GetMeditationViewReverseFieldMap() map[string]string {
	return map[string]string{"ended_at": "EndedAt", "session_id": "SessionId", "session_name": "SessionName", "session_type": "SessionType", "started_at": "StartedAt", "user_id": "UserId"}
}

func GetMeditationViewTypeMap() map[string]string {
	return map[string]string{"EndedAt": "*int64", "SessionId": "*uuid.UUID", "SessionName": "*string", "SessionType": "*string", "StartedAt": "*int64", "UserId": "*uuid.UUID"}
}

type SleepSummarie struct {
	UserId     *uuid.UUID            `db:"user_id" json:",omitempty"`
	Duration   *int64                `db:"duration" json:",omitempty"`
	Properties *datastore.RawMessage `db:"properties" json:",omitempty"`
	SessionId  *uuid.UUID            `db:"session_id" json:",omitempty"`
}

func GetSleepSummarieFieldMap() map[string]string {
	return map[string]string{"Duration": "duration", "Properties": "properties", "SessionId": "session_id", "UserId": "user_id"}
}

func GetSleepSummarieReverseFieldMap() map[string]string {
	return map[string]string{"duration": "Duration", "properties": "Properties", "session_id": "SessionId", "user_id": "UserId"}
}

func GetSleepSummarieTypeMap() map[string]string {
	return map[string]string{"Duration": "*int64", "Properties": "*datastore.RawMessage", "SessionId": "*uuid.UUID", "UserId": "*uuid.UUID"}
}

type MeditationSummarie struct {
	UserId     *uuid.UUID            `db:"user_id" json:",omitempty"`
	Duration   *int64                `db:"duration" json:",omitempty"`
	Properties *datastore.RawMessage `db:"properties" json:",omitempty"`
	SessionId  *uuid.UUID            `db:"session_id" json:",omitempty"`
}

func GetMeditationSummarieFieldMap() map[string]string {
	return map[string]string{"Duration": "duration", "Properties": "properties", "SessionId": "session_id", "UserId": "user_id"}
}

func GetMeditationSummarieReverseFieldMap() map[string]string {
	return map[string]string{"duration": "Duration", "properties": "Properties", "session_id": "SessionId", "user_id": "UserId"}
}

func GetMeditationSummarieTypeMap() map[string]string {
	return map[string]string{"Duration": "*int64", "Properties": "*datastore.RawMessage", "SessionId": "*uuid.UUID", "UserId": "*uuid.UUID"}
}

type UserSummaryView struct {
	Count        *int64     `db:"count" json:",omitempty"`
	ActivityType *string    `db:"activity_type" json:",omitempty"`
	UserId       *uuid.UUID `db:"user_id" json:",omitempty"`
}

func GetUserSummaryViewFieldMap() map[string]string {
	return map[string]string{"ActivityType": "activity_type", "Count": "count", "UserId": "user_id"}
}

func GetUserSummaryViewReverseFieldMap() map[string]string {
	return map[string]string{"activity_type": "ActivityType", "count": "Count", "user_id": "UserId"}
}

func GetUserSummaryViewTypeMap() map[string]string {
	return map[string]string{"ActivityType": "*string", "Count": "*int64", "UserId": "*uuid.UUID"}
}

type UserSleepView struct {
	UserId     *uuid.UUID            `db:"user_id" json:",omitempty"`
	Duration   *int64                `db:"duration" json:",omitempty"`
	Properties *datastore.RawMessage `db:"properties" json:",omitempty"`
	SessionId  *uuid.UUID            `db:"session_id" json:",omitempty"`
}

func GetUserSleepViewFieldMap() map[string]string {
	return map[string]string{"Duration": "duration", "Properties": "properties", "SessionId": "session_id", "UserId": "user_id"}
}

func GetUserSleepViewReverseFieldMap() map[string]string {
	return map[string]string{"duration": "Duration", "properties": "Properties", "session_id": "SessionId", "user_id": "UserId"}
}

func GetUserSleepViewTypeMap() map[string]string {
	return map[string]string{"Duration": "*int64", "Properties": "*datastore.RawMessage", "SessionId": "*uuid.UUID", "UserId": "*uuid.UUID"}
}

type UserMeditationView struct {
	UserId     *uuid.UUID            `db:"user_id" json:",omitempty"`
	Duration   *int64                `db:"duration" json:",omitempty"`
	Properties *datastore.RawMessage `db:"properties" json:",omitempty"`
	SessionId  *uuid.UUID            `db:"session_id" json:",omitempty"`
}

func GetUserMeditationViewFieldMap() map[string]string {
	return map[string]string{"Duration": "duration", "Properties": "properties", "SessionId": "session_id", "UserId": "user_id"}
}

func GetUserMeditationViewReverseFieldMap() map[string]string {
	return map[string]string{"duration": "Duration", "properties": "Properties", "session_id": "SessionId", "user_id": "UserId"}
}

func GetUserMeditationViewTypeMap() map[string]string {
	return map[string]string{"Duration": "*int64", "Properties": "*datastore.RawMessage", "SessionId": "*uuid.UUID", "UserId": "*uuid.UUID"}
}
