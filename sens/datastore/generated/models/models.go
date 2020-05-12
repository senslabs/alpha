package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/senslabs/alpha/sens/datastore"
)

var t time.Time

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

type Alert struct {
	AlertId     *uuid.UUID `db:"alert_id" json:",omitempty"`
	UserId      *uuid.UUID `db:"user_id" json:",omitempty"`
	CreatedAt   *int64     `db:"created_at" json:",omitempty"`
	AlertName   *string    `db:"alert_name" json:",omitempty"`
	Status      *string    `db:"status" json:",omitempty"`
	Remarks     *string    `db:"remarks" json:",omitempty"`
	Valid       *bool      `db:"valid" json:",omitempty"`
	AlertRuleId *uuid.UUID `db:"alert_rule_id" json:",omitempty"`
	UpdatedAt   *int64     `db:"updated_at" json:",omitempty"`
}

func GetAlertFieldMap() map[string]string {
	return map[string]string{"AlertId": "alert_id", "AlertName": "alert_name", "AlertRuleId": "alert_rule_id", "CreatedAt": "created_at", "Remarks": "remarks", "Status": "status", "UpdatedAt": "updated_at", "UserId": "user_id", "Valid": "valid"}
}

func GetAlertReverseFieldMap() map[string]string {
	return map[string]string{"alert_id": "AlertId", "alert_name": "AlertName", "alert_rule_id": "AlertRuleId", "created_at": "CreatedAt", "remarks": "Remarks", "status": "Status", "updated_at": "UpdatedAt", "user_id": "UserId", "valid": "Valid"}
}

func GetAlertTypeMap() map[string]string {
	return map[string]string{"AlertId": "*uuid.UUID", "AlertName": "*string", "AlertRuleId": "*uuid.UUID", "CreatedAt": "*int64", "Remarks": "*string", "Status": "*string", "UpdatedAt": "*int64", "UserId": "*uuid.UUID", "Valid": "*bool"}
}

type Session struct {
	SessionId   *uuid.UUID `db:"session_id" json:",omitempty"`
	UserId      *uuid.UUID `db:"user_id" json:",omitempty"`
	SessionName *string    `db:"session_name" json:",omitempty"`
	SessionType *string    `db:"session_type" json:",omitempty"`
	StartedAt   *int64     `db:"started_at" json:",omitempty"`
	EndedAt     *int64     `db:"ended_at" json:",omitempty"`
	State       *string    `db:"state" json:",omitempty"`
}

func GetSessionFieldMap() map[string]string {
	return map[string]string{"EndedAt": "ended_at", "SessionId": "session_id", "SessionName": "session_name", "SessionType": "session_type", "StartedAt": "started_at", "State": "state", "UserId": "user_id"}
}

func GetSessionReverseFieldMap() map[string]string {
	return map[string]string{"ended_at": "EndedAt", "session_id": "SessionId", "session_name": "SessionName", "session_type": "SessionType", "started_at": "StartedAt", "state": "State", "user_id": "UserId"}
}

func GetSessionTypeMap() map[string]string {
	return map[string]string{"EndedAt": "*int64", "SessionId": "*uuid.UUID", "SessionName": "*string", "SessionType": "*string", "StartedAt": "*int64", "State": "*string", "UserId": "*uuid.UUID"}
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

type Baseline struct {
	BaselineId *uuid.UUID `db:"baseline_id" json:",omitempty"`
	UserId     *uuid.UUID `db:"user_id" json:",omitempty"`
	Key        *string    `db:"key" json:",omitempty"`
	CreatedAt  *int64     `db:"created_at" json:",omitempty"`
	LowerLimit *int64     `db:"lower_limit" json:",omitempty"`
	UpperLimit *int64     `db:"upper_limit" json:",omitempty"`
}

func GetBaselineFieldMap() map[string]string {
	return map[string]string{"BaselineId": "baseline_id", "CreatedAt": "created_at", "Key": "key", "LowerLimit": "lower_limit", "UpperLimit": "upper_limit", "UserId": "user_id"}
}

func GetBaselineReverseFieldMap() map[string]string {
	return map[string]string{"baseline_id": "BaselineId", "created_at": "CreatedAt", "key": "Key", "lower_limit": "LowerLimit", "upper_limit": "UpperLimit", "user_id": "UserId"}
}

func GetBaselineTypeMap() map[string]string {
	return map[string]string{"BaselineId": "*uuid.UUID", "CreatedAt": "*int64", "Key": "*string", "LowerLimit": "*int64", "UpperLimit": "*int64", "UserId": "*uuid.UUID"}
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

type AuthView struct {
	AuthId    *uuid.UUID `db:"auth_id" json:",omitempty"`
	Email     *string    `db:"email" json:",omitempty"`
	Mobile    *string    `db:"mobile" json:",omitempty"`
	Social    *string    `db:"social" json:",omitempty"`
	FirstName *string    `db:"first_name" json:",omitempty"`
	LastName  *string    `db:"last_name" json:",omitempty"`
	IsSens    *bool      `db:"is_sens" json:",omitempty"`
}

func GetAuthViewFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "Email": "email", "FirstName": "first_name", "IsSens": "is_sens", "LastName": "last_name", "Mobile": "mobile", "Social": "social"}
}

func GetAuthViewReverseFieldMap() map[string]string {
	return map[string]string{"auth_id": "AuthId", "email": "Email", "first_name": "FirstName", "is_sens": "IsSens", "last_name": "LastName", "mobile": "Mobile", "social": "Social"}
}

func GetAuthViewTypeMap() map[string]string {
	return map[string]string{"AuthId": "*uuid.UUID", "Email": "*string", "FirstName": "*string", "IsSens": "*bool", "LastName": "*string", "Mobile": "*string", "Social": "*string"}
}

type OrgView struct {
	AuthId    *uuid.UUID `db:"auth_id" json:",omitempty"`
	Email     *string    `db:"email" json:",omitempty"`
	Mobile    *string    `db:"mobile" json:",omitempty"`
	Social    *string    `db:"social" json:",omitempty"`
	FirstName *string    `db:"first_name" json:",omitempty"`
	LastName  *string    `db:"last_name" json:",omitempty"`
	OrgId     *uuid.UUID `db:"org_id" json:",omitempty"`
	OrgName   *string    `db:"org_name" json:",omitempty"`
}

func GetOrgViewFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "Email": "email", "FirstName": "first_name", "LastName": "last_name", "Mobile": "mobile", "OrgId": "org_id", "OrgName": "org_name", "Social": "social"}
}

func GetOrgViewReverseFieldMap() map[string]string {
	return map[string]string{"auth_id": "AuthId", "email": "Email", "first_name": "FirstName", "last_name": "LastName", "mobile": "Mobile", "org_id": "OrgId", "org_name": "OrgName", "social": "Social"}
}

func GetOrgViewTypeMap() map[string]string {
	return map[string]string{"AuthId": "*uuid.UUID", "Email": "*string", "FirstName": "*string", "LastName": "*string", "Mobile": "*string", "OrgId": "*uuid.UUID", "OrgName": "*string", "Social": "*string"}
}

type OpView struct {
	AuthId    *uuid.UUID `db:"auth_id" json:",omitempty"`
	Email     *string    `db:"email" json:",omitempty"`
	Mobile    *string    `db:"mobile" json:",omitempty"`
	Social    *string    `db:"social" json:",omitempty"`
	FirstName *string    `db:"first_name" json:",omitempty"`
	LastName  *string    `db:"last_name" json:",omitempty"`
	OpId      *uuid.UUID `db:"op_id" json:",omitempty"`
	OrgId     *uuid.UUID `db:"org_id" json:",omitempty"`
}

func GetOpViewFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "Email": "email", "FirstName": "first_name", "LastName": "last_name", "Mobile": "mobile", "OpId": "op_id", "OrgId": "org_id", "Social": "social"}
}

func GetOpViewReverseFieldMap() map[string]string {
	return map[string]string{"auth_id": "AuthId", "email": "Email", "first_name": "FirstName", "last_name": "LastName", "mobile": "Mobile", "op_id": "OpId", "org_id": "OrgId", "social": "Social"}
}

func GetOpViewTypeMap() map[string]string {
	return map[string]string{"AuthId": "*uuid.UUID", "Email": "*string", "FirstName": "*string", "LastName": "*string", "Mobile": "*string", "OpId": "*uuid.UUID", "OrgId": "*uuid.UUID", "Social": "*string"}
}

type UserView struct {
	AuthId    *uuid.UUID `db:"auth_id" json:",omitempty"`
	Email     *string    `db:"email" json:",omitempty"`
	Mobile    *string    `db:"mobile" json:",omitempty"`
	Social    *string    `db:"social" json:",omitempty"`
	FirstName *string    `db:"first_name" json:",omitempty"`
	LastName  *string    `db:"last_name" json:",omitempty"`
	UserId    *uuid.UUID `db:"user_id" json:",omitempty"`
	OrgId     *uuid.UUID `db:"org_id" json:",omitempty"`
}

func GetUserViewFieldMap() map[string]string {
	return map[string]string{"AuthId": "auth_id", "Email": "email", "FirstName": "first_name", "LastName": "last_name", "Mobile": "mobile", "OrgId": "org_id", "Social": "social", "UserId": "user_id"}
}

func GetUserViewReverseFieldMap() map[string]string {
	return map[string]string{"auth_id": "AuthId", "email": "Email", "first_name": "FirstName", "last_name": "LastName", "mobile": "Mobile", "org_id": "OrgId", "social": "Social", "user_id": "UserId"}
}

func GetUserViewTypeMap() map[string]string {
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

type OrgActivityView struct {
	ActivityType *string    `db:"activity_type" json:",omitempty"`
	Timestamp    *int64     `db:"timestamp" json:",omitempty"`
	UserId       *uuid.UUID `db:"user_id" json:",omitempty"`
	OrgId        *uuid.UUID `db:"org_id" json:",omitempty"`
}

func GetOrgActivityViewFieldMap() map[string]string {
	return map[string]string{"ActivityType": "activity_type", "OrgId": "org_id", "Timestamp": "timestamp", "UserId": "user_id"}
}

func GetOrgActivityViewReverseFieldMap() map[string]string {
	return map[string]string{"activity_type": "ActivityType", "org_id": "OrgId", "timestamp": "Timestamp", "user_id": "UserId"}
}

func GetOrgActivityViewTypeMap() map[string]string {
	return map[string]string{"ActivityType": "*string", "OrgId": "*uuid.UUID", "Timestamp": "*int64", "UserId": "*uuid.UUID"}
}

type OrgActivitySummaryView struct {
	Count        *int64     `db:"count" json:",omitempty"`
	ActivityType *string    `db:"activity_type" json:",omitempty"`
	UserId       *uuid.UUID `db:"user_id" json:",omitempty"`
	OrgId        *uuid.UUID `db:"org_id" json:",omitempty"`
}

func GetOrgActivitySummaryViewFieldMap() map[string]string {
	return map[string]string{"ActivityType": "activity_type", "Count": "count", "OrgId": "org_id", "UserId": "user_id"}
}

func GetOrgActivitySummaryViewReverseFieldMap() map[string]string {
	return map[string]string{"activity_type": "ActivityType", "count": "Count", "org_id": "OrgId", "user_id": "UserId"}
}

func GetOrgActivitySummaryViewTypeMap() map[string]string {
	return map[string]string{"ActivityType": "*string", "Count": "*int64", "OrgId": "*uuid.UUID", "UserId": "*uuid.UUID"}
}

type OrgQuarterUsageView struct {
	Count        *int64     `db:"count" json:",omitempty"`
	ActivityType *string    `db:"activity_type" json:",omitempty"`
	OrgId        *uuid.UUID `db:"org_id" json:",omitempty"`
	Date         []byte     `db:"date" json:",omitempty"`
}

func GetOrgQuarterUsageViewFieldMap() map[string]string {
	return map[string]string{"ActivityType": "activity_type", "Count": "count", "Date": "date", "OrgId": "org_id"}
}

func GetOrgQuarterUsageViewReverseFieldMap() map[string]string {
	return map[string]string{"activity_type": "ActivityType", "count": "Count", "date": "Date", "org_id": "OrgId"}
}

func GetOrgQuarterUsageViewTypeMap() map[string]string {
	return map[string]string{"ActivityType": "*string", "Count": "*int64", "Date": "[]byte", "OrgId": "*uuid.UUID"}
}

type OrgSessionRecordView struct {
	UserId      *uuid.UUID            `db:"user_id" json:",omitempty"`
	OrgId       *uuid.UUID            `db:"org_id" json:",omitempty"`
	SessionId   *uuid.UUID            `db:"session_id" json:",omitempty"`
	SessionType *string               `db:"session_type" json:",omitempty"`
	StartedAt   *int64                `db:"started_at" json:",omitempty"`
	EndedAt     *int64                `db:"ended_at" json:",omitempty"`
	Key         *string               `db:"key" json:",omitempty"`
	Timestamp   *int64                `db:"timestamp" json:",omitempty"`
	Value       *float64              `db:"value" json:",omitempty"`
	Properties  *datastore.RawMessage `db:"properties" json:",omitempty"`
}

func GetOrgSessionRecordViewFieldMap() map[string]string {
	return map[string]string{"EndedAt": "ended_at", "Key": "key", "OrgId": "org_id", "Properties": "properties", "SessionId": "session_id", "SessionType": "session_type", "StartedAt": "started_at", "Timestamp": "timestamp", "UserId": "user_id", "Value": "value"}
}

func GetOrgSessionRecordViewReverseFieldMap() map[string]string {
	return map[string]string{"ended_at": "EndedAt", "key": "Key", "org_id": "OrgId", "properties": "Properties", "session_id": "SessionId", "session_type": "SessionType", "started_at": "StartedAt", "timestamp": "Timestamp", "user_id": "UserId", "value": "Value"}
}

func GetOrgSessionRecordViewTypeMap() map[string]string {
	return map[string]string{"EndedAt": "*int64", "Key": "*string", "OrgId": "*uuid.UUID", "Properties": "*datastore.RawMessage", "SessionId": "*uuid.UUID", "SessionType": "*string", "StartedAt": "*int64", "Timestamp": "*int64", "UserId": "*uuid.UUID", "Value": "*float64"}
}

type DevicePropertie struct {
	DeviceId *uuid.UUID `db:"device_id" json:",omitempty"`
	Key      *string    `db:"key" json:",omitempty"`
	Value    *string    `db:"value" json:",omitempty"`
}

func GetDevicePropertieFieldMap() map[string]string {
	return map[string]string{"DeviceId": "device_id", "Key": "key", "Value": "value"}
}

func GetDevicePropertieReverseFieldMap() map[string]string {
	return map[string]string{"device_id": "DeviceId", "key": "Key", "value": "Value"}
}

func GetDevicePropertieTypeMap() map[string]string {
	return map[string]string{"DeviceId": "*uuid.UUID", "Key": "*string", "Value": "*string"}
}

type OrgSessionEventView struct {
	UserId         *uuid.UUID            `db:"user_id" json:",omitempty"`
	OrgId          *uuid.UUID            `db:"org_id" json:",omitempty"`
	SessionId      *uuid.UUID            `db:"session_id" json:",omitempty"`
	SessionType    *string               `db:"session_type" json:",omitempty"`
	Key            *string               `db:"key" json:",omitempty"`
	EventStartedAt *int64                `db:"event_started_at" json:",omitempty"`
	EventEndedAt   *int64                `db:"event_ended_at" json:",omitempty"`
	Properties     *datastore.RawMessage `db:"properties" json:",omitempty"`
}

func GetOrgSessionEventViewFieldMap() map[string]string {
	return map[string]string{"EventEndedAt": "event_ended_at", "EventStartedAt": "event_started_at", "Key": "key", "OrgId": "org_id", "Properties": "properties", "SessionId": "session_id", "SessionType": "session_type", "UserId": "user_id"}
}

func GetOrgSessionEventViewReverseFieldMap() map[string]string {
	return map[string]string{"event_ended_at": "EventEndedAt", "event_started_at": "EventStartedAt", "key": "Key", "org_id": "OrgId", "properties": "Properties", "session_id": "SessionId", "session_type": "SessionType", "user_id": "UserId"}
}

func GetOrgSessionEventViewTypeMap() map[string]string {
	return map[string]string{"EventEndedAt": "*int64", "EventStartedAt": "*int64", "Key": "*string", "OrgId": "*uuid.UUID", "Properties": "*datastore.RawMessage", "SessionId": "*uuid.UUID", "SessionType": "*string", "UserId": "*uuid.UUID"}
}

type OrgSessionEventDetailView struct {
	UserId         *uuid.UUID            `db:"user_id" json:",omitempty"`
	OrgId          *uuid.UUID            `db:"org_id" json:",omitempty"`
	SessionId      *uuid.UUID            `db:"session_id" json:",omitempty"`
	SessionType    *string               `db:"session_type" json:",omitempty"`
	EventStartedAt *datastore.RawMessage `db:"event_started_at" json:",omitempty"`
	EventEndedAt   *datastore.RawMessage `db:"event_ended_at" json:",omitempty"`
	Key            *string               `db:"key" json:",omitempty"`
}

func GetOrgSessionEventDetailViewFieldMap() map[string]string {
	return map[string]string{"EventEndedAt": "event_ended_at", "EventStartedAt": "event_started_at", "Key": "key", "OrgId": "org_id", "SessionId": "session_id", "SessionType": "session_type", "UserId": "user_id"}
}

func GetOrgSessionEventDetailViewReverseFieldMap() map[string]string {
	return map[string]string{"event_ended_at": "EventEndedAt", "event_started_at": "EventStartedAt", "key": "Key", "org_id": "OrgId", "session_id": "SessionId", "session_type": "SessionType", "user_id": "UserId"}
}

func GetOrgSessionEventDetailViewTypeMap() map[string]string {
	return map[string]string{"EventEndedAt": "*datastore.RawMessage", "EventStartedAt": "*datastore.RawMessage", "Key": "*string", "OrgId": "*uuid.UUID", "SessionId": "*uuid.UUID", "SessionType": "*string", "UserId": "*uuid.UUID"}
}

type OrgSessionDetailView struct {
	UserId      *uuid.UUID            `db:"user_id" json:",omitempty"`
	OrgId       *uuid.UUID            `db:"org_id" json:",omitempty"`
	SessionId   *uuid.UUID            `db:"session_id" json:",omitempty"`
	SessionType *string               `db:"session_type" json:",omitempty"`
	StartedAt   *int64                `db:"started_at" json:",omitempty"`
	EndedAt     *int64                `db:"ended_at" json:",omitempty"`
	Key         *string               `db:"key" json:",omitempty"`
	Timestamps  *datastore.RawMessage `db:"timestamps" json:",omitempty"`
	Values      *datastore.RawMessage `db:"values" json:",omitempty"`
	Min         *float64              `db:"min" json:",omitempty"`
	Max         *float64              `db:"max" json:",omitempty"`
	Avg         *float64              `db:"avg" json:",omitempty"`
}

func GetOrgSessionDetailViewFieldMap() map[string]string {
	return map[string]string{"Avg": "avg", "EndedAt": "ended_at", "Key": "key", "Max": "max", "Min": "min", "OrgId": "org_id", "SessionId": "session_id", "SessionType": "session_type", "StartedAt": "started_at", "Timestamps": "timestamps", "UserId": "user_id", "Values": "values"}
}

func GetOrgSessionDetailViewReverseFieldMap() map[string]string {
	return map[string]string{"avg": "Avg", "ended_at": "EndedAt", "key": "Key", "max": "Max", "min": "Min", "org_id": "OrgId", "session_id": "SessionId", "session_type": "SessionType", "started_at": "StartedAt", "timestamps": "Timestamps", "user_id": "UserId", "values": "Values"}
}

func GetOrgSessionDetailViewTypeMap() map[string]string {
	return map[string]string{"Avg": "*float64", "EndedAt": "*int64", "Key": "*string", "Max": "*float64", "Min": "*float64", "OrgId": "*uuid.UUID", "SessionId": "*uuid.UUID", "SessionType": "*string", "StartedAt": "*int64", "Timestamps": "*datastore.RawMessage", "UserId": "*uuid.UUID", "Values": "*datastore.RawMessage"}
}

type OrgPropertie struct {
	OrgId *uuid.UUID `db:"org_id" json:",omitempty"`
	Key   *string    `db:"key" json:",omitempty"`
	Value *string    `db:"value" json:",omitempty"`
}

func GetOrgPropertieFieldMap() map[string]string {
	return map[string]string{"Key": "key", "OrgId": "org_id", "Value": "value"}
}

func GetOrgPropertieReverseFieldMap() map[string]string {
	return map[string]string{"key": "Key", "org_id": "OrgId", "value": "Value"}
}

func GetOrgPropertieTypeMap() map[string]string {
	return map[string]string{"Key": "*string", "OrgId": "*uuid.UUID", "Value": "*string"}
}

type OpSetting struct {
	OpSettingId *uuid.UUID `db:"op_setting_id" json:",omitempty"`
	OpId        *uuid.UUID `db:"op_id" json:",omitempty"`
	CreatedAt   *int64     `db:"created_at" json:",omitempty"`
	Key         *string    `db:"key" json:",omitempty"`
	Value       *string    `db:"value" json:",omitempty"`
}

func GetOpSettingFieldMap() map[string]string {
	return map[string]string{"CreatedAt": "created_at", "Key": "key", "OpId": "op_id", "OpSettingId": "op_setting_id", "Value": "value"}
}

func GetOpSettingReverseFieldMap() map[string]string {
	return map[string]string{"created_at": "CreatedAt", "key": "Key", "op_id": "OpId", "op_setting_id": "OpSettingId", "value": "Value"}
}

func GetOpSettingTypeMap() map[string]string {
	return map[string]string{"CreatedAt": "*int64", "Key": "*string", "OpId": "*uuid.UUID", "OpSettingId": "*uuid.UUID", "Value": "*string"}
}

type OpPropertie struct {
	OpId  *uuid.UUID `db:"op_id" json:",omitempty"`
	Key   *string    `db:"key" json:",omitempty"`
	Value *string    `db:"value" json:",omitempty"`
}

func GetOpPropertieFieldMap() map[string]string {
	return map[string]string{"Key": "key", "OpId": "op_id", "Value": "value"}
}

func GetOpPropertieReverseFieldMap() map[string]string {
	return map[string]string{"key": "Key", "op_id": "OpId", "value": "Value"}
}

func GetOpPropertieTypeMap() map[string]string {
	return map[string]string{"Key": "*string", "OpId": "*uuid.UUID", "Value": "*string"}
}

type OrgAlertView struct {
	UserId    *uuid.UUID `db:"user_id" json:",omitempty"`
	OrgId     *uuid.UUID `db:"org_id" json:",omitempty"`
	AlertId   *uuid.UUID `db:"alert_id" json:",omitempty"`
	FirstName *string    `db:"first_name" json:",omitempty"`
	LastName  *string    `db:"last_name" json:",omitempty"`
	CreatedAt *int64     `db:"created_at" json:",omitempty"`
	AlertName *string    `db:"alert_name" json:",omitempty"`
	Status    *string    `db:"status" json:",omitempty"`
	Remarks   *string    `db:"remarks" json:",omitempty"`
}

func GetOrgAlertViewFieldMap() map[string]string {
	return map[string]string{"AlertId": "alert_id", "AlertName": "alert_name", "CreatedAt": "created_at", "FirstName": "first_name", "LastName": "last_name", "OrgId": "org_id", "Remarks": "remarks", "Status": "status", "UserId": "user_id"}
}

func GetOrgAlertViewReverseFieldMap() map[string]string {
	return map[string]string{"alert_id": "AlertId", "alert_name": "AlertName", "created_at": "CreatedAt", "first_name": "FirstName", "last_name": "LastName", "org_id": "OrgId", "remarks": "Remarks", "status": "Status", "user_id": "UserId"}
}

func GetOrgAlertViewTypeMap() map[string]string {
	return map[string]string{"AlertId": "*uuid.UUID", "AlertName": "*string", "CreatedAt": "*int64", "FirstName": "*string", "LastName": "*string", "OrgId": "*uuid.UUID", "Remarks": "*string", "Status": "*string", "UserId": "*uuid.UUID"}
}

type OrgLatestAlertView struct {
	UserId    *uuid.UUID            `db:"user_id" json:",omitempty"`
	OrgId     *uuid.UUID            `db:"org_id" json:",omitempty"`
	FirstName *string               `db:"first_name" json:",omitempty"`
	LastName  *string               `db:"last_name" json:",omitempty"`
	Timestamp *int64                `db:"timestamp" json:",omitempty"`
	Alerts    *datastore.RawMessage `db:"alerts" json:",omitempty"`
}

func GetOrgLatestAlertViewFieldMap() map[string]string {
	return map[string]string{"Alerts": "alerts", "FirstName": "first_name", "LastName": "last_name", "OrgId": "org_id", "Timestamp": "timestamp", "UserId": "user_id"}
}

func GetOrgLatestAlertViewReverseFieldMap() map[string]string {
	return map[string]string{"alerts": "Alerts", "first_name": "FirstName", "last_name": "LastName", "org_id": "OrgId", "timestamp": "Timestamp", "user_id": "UserId"}
}

func GetOrgLatestAlertViewTypeMap() map[string]string {
	return map[string]string{"Alerts": "*datastore.RawMessage", "FirstName": "*string", "LastName": "*string", "OrgId": "*uuid.UUID", "Timestamp": "*int64", "UserId": "*uuid.UUID"}
}

type OrgSetting struct {
	OrgSettingId *uuid.UUID `db:"org_setting_id" json:",omitempty"`
	OrgId        *uuid.UUID `db:"org_id" json:",omitempty"`
	CreatedAt    *int64     `db:"created_at" json:",omitempty"`
	Key          *string    `db:"key" json:",omitempty"`
	Value        *string    `db:"value" json:",omitempty"`
}

func GetOrgSettingFieldMap() map[string]string {
	return map[string]string{"CreatedAt": "created_at", "Key": "key", "OrgId": "org_id", "OrgSettingId": "org_setting_id", "Value": "value"}
}

func GetOrgSettingReverseFieldMap() map[string]string {
	return map[string]string{"created_at": "CreatedAt", "key": "Key", "org_id": "OrgId", "org_setting_id": "OrgSettingId", "value": "Value"}
}

func GetOrgSettingTypeMap() map[string]string {
	return map[string]string{"CreatedAt": "*int64", "Key": "*string", "OrgId": "*uuid.UUID", "OrgSettingId": "*uuid.UUID", "Value": "*string"}
}

type AlertEscalation struct {
	AlertEscalationId *uuid.UUID `db:"alert_escalation_id" json:",omitempty"`
	AlertRuleId       *uuid.UUID `db:"alert_rule_id" json:",omitempty"`
	EscalationGroup   *string    `db:"escalation_group" json:",omitempty"`
	EscalationLevel   *int64     `db:"escalation_level" json:",omitempty"`
	Snooze            *int64     `db:"snooze" json:",omitempty"`
	Medium            *string    `db:"medium" json:",omitempty"`
	MediumValue       *string    `db:"medium_value" json:",omitempty"`
	CreatedAt         *int64     `db:"created_at" json:",omitempty"`
	Timeout           *int64     `db:"timeout" json:",omitempty"`
	Status            *string    `db:"status" json:",omitempty"`
}

func GetAlertEscalationFieldMap() map[string]string {
	return map[string]string{"AlertEscalationId": "alert_escalation_id", "AlertRuleId": "alert_rule_id", "CreatedAt": "created_at", "EscalationGroup": "escalation_group", "EscalationLevel": "escalation_level", "Medium": "medium", "MediumValue": "medium_value", "Snooze": "snooze", "Status": "status", "Timeout": "timeout"}
}

func GetAlertEscalationReverseFieldMap() map[string]string {
	return map[string]string{"alert_escalation_id": "AlertEscalationId", "alert_rule_id": "AlertRuleId", "created_at": "CreatedAt", "escalation_group": "EscalationGroup", "escalation_level": "EscalationLevel", "medium": "Medium", "medium_value": "MediumValue", "snooze": "Snooze", "status": "Status", "timeout": "Timeout"}
}

func GetAlertEscalationTypeMap() map[string]string {
	return map[string]string{"AlertEscalationId": "*uuid.UUID", "AlertRuleId": "*uuid.UUID", "CreatedAt": "*int64", "EscalationGroup": "*string", "EscalationLevel": "*int64", "Medium": "*string", "MediumValue": "*string", "Snooze": "*int64", "Status": "*string", "Timeout": "*int64"}
}

type UserSettingView struct {
	UserId    *uuid.UUID `db:"user_id" json:",omitempty"`
	CreatedAt *int64     `db:"created_at" json:",omitempty"`
	Key       *string    `db:"key" json:",omitempty"`
	Value     *string    `db:"value" json:",omitempty"`
}

func GetUserSettingViewFieldMap() map[string]string {
	return map[string]string{"CreatedAt": "created_at", "Key": "key", "UserId": "user_id", "Value": "value"}
}

func GetUserSettingViewReverseFieldMap() map[string]string {
	return map[string]string{"created_at": "CreatedAt", "key": "Key", "user_id": "UserId", "value": "Value"}
}

func GetUserSettingViewTypeMap() map[string]string {
	return map[string]string{"CreatedAt": "*int64", "Key": "*string", "UserId": "*uuid.UUID", "Value": "*string"}
}

type BaselineView struct {
	UserId     *uuid.UUID `db:"user_id" json:",omitempty"`
	CreatedAt  *int64     `db:"created_at" json:",omitempty"`
	Key        *string    `db:"key" json:",omitempty"`
	LowerLimit *int64     `db:"lower_limit" json:",omitempty"`
	UpperLimit *int64     `db:"upper_limit" json:",omitempty"`
}

func GetBaselineViewFieldMap() map[string]string {
	return map[string]string{"CreatedAt": "created_at", "Key": "key", "LowerLimit": "lower_limit", "UpperLimit": "upper_limit", "UserId": "user_id"}
}

func GetBaselineViewReverseFieldMap() map[string]string {
	return map[string]string{"created_at": "CreatedAt", "key": "Key", "lower_limit": "LowerLimit", "upper_limit": "UpperLimit", "user_id": "UserId"}
}

func GetBaselineViewTypeMap() map[string]string {
	return map[string]string{"CreatedAt": "*int64", "Key": "*string", "LowerLimit": "*int64", "UpperLimit": "*int64", "UserId": "*uuid.UUID"}
}

type OrgSessionView struct {
	SessionId   *uuid.UUID `db:"session_id" json:",omitempty"`
	UserId      *uuid.UUID `db:"user_id" json:",omitempty"`
	OrgId       *uuid.UUID `db:"org_id" json:",omitempty"`
	SessionName *string    `db:"session_name" json:",omitempty"`
	SessionType *string    `db:"session_type" json:",omitempty"`
	StartedAt   *int64     `db:"started_at" json:",omitempty"`
	EndedAt     *int64     `db:"ended_at" json:",omitempty"`
}

func GetOrgSessionViewFieldMap() map[string]string {
	return map[string]string{"EndedAt": "ended_at", "OrgId": "org_id", "SessionId": "session_id", "SessionName": "session_name", "SessionType": "session_type", "StartedAt": "started_at", "UserId": "user_id"}
}

func GetOrgSessionViewReverseFieldMap() map[string]string {
	return map[string]string{"ended_at": "EndedAt", "org_id": "OrgId", "session_id": "SessionId", "session_name": "SessionName", "session_type": "SessionType", "started_at": "StartedAt", "user_id": "UserId"}
}

func GetOrgSessionViewTypeMap() map[string]string {
	return map[string]string{"EndedAt": "*int64", "OrgId": "*uuid.UUID", "SessionId": "*uuid.UUID", "SessionName": "*string", "SessionType": "*string", "StartedAt": "*int64", "UserId": "*uuid.UUID"}
}

type UserSessionCountView struct {
	UserId *uuid.UUID `db:"user_id" json:",omitempty"`
	OrgId  *uuid.UUID `db:"org_id" json:",omitempty"`
	Count  *int64     `db:"count" json:",omitempty"`
}

func GetUserSessionCountViewFieldMap() map[string]string {
	return map[string]string{"Count": "count", "OrgId": "org_id", "UserId": "user_id"}
}

func GetUserSessionCountViewReverseFieldMap() map[string]string {
	return map[string]string{"count": "Count", "org_id": "OrgId", "user_id": "UserId"}
}

func GetUserSessionCountViewTypeMap() map[string]string {
	return map[string]string{"Count": "*int64", "OrgId": "*uuid.UUID", "UserId": "*uuid.UUID"}
}

type SessionDurationView struct {
	UserId      *uuid.UUID            `db:"user_id" json:",omitempty"`
	OrgId       *uuid.UUID            `db:"org_id" json:",omitempty"`
	SessionId   *uuid.UUID            `db:"session_id" json:",omitempty"`
	StageEpochs *datastore.RawMessage `db:"stage_epochs" json:",omitempty"`
	Epochs      *int64                `db:"epochs" json:",omitempty"`
}

func GetSessionDurationViewFieldMap() map[string]string {
	return map[string]string{"Epochs": "epochs", "OrgId": "org_id", "SessionId": "session_id", "StageEpochs": "stage_epochs", "UserId": "user_id"}
}

func GetSessionDurationViewReverseFieldMap() map[string]string {
	return map[string]string{"epochs": "Epochs", "org_id": "OrgId", "session_id": "SessionId", "stage_epochs": "StageEpochs", "user_id": "UserId"}
}

func GetSessionDurationViewTypeMap() map[string]string {
	return map[string]string{"Epochs": "*int64", "OrgId": "*uuid.UUID", "SessionId": "*uuid.UUID", "StageEpochs": "*datastore.RawMessage", "UserId": "*uuid.UUID"}
}

type OrgSessionInfoView struct {
	UserId      *uuid.UUID            `db:"user_id" json:",omitempty"`
	OrgId       *uuid.UUID            `db:"org_id" json:",omitempty"`
	SessionId   *uuid.UUID            `db:"session_id" json:",omitempty"`
	SessionType *string               `db:"session_type" json:",omitempty"`
	SessionName *string               `db:"session_name" json:",omitempty"`
	StartedAt   *int64                `db:"started_at" json:",omitempty"`
	EndedAt     *int64                `db:"ended_at" json:",omitempty"`
	Properties  *datastore.RawMessage `db:"properties" json:",omitempty"`
}

func GetOrgSessionInfoViewFieldMap() map[string]string {
	return map[string]string{"EndedAt": "ended_at", "OrgId": "org_id", "Properties": "properties", "SessionId": "session_id", "SessionName": "session_name", "SessionType": "session_type", "StartedAt": "started_at", "UserId": "user_id"}
}

func GetOrgSessionInfoViewReverseFieldMap() map[string]string {
	return map[string]string{"ended_at": "EndedAt", "org_id": "OrgId", "properties": "Properties", "session_id": "SessionId", "session_name": "SessionName", "session_type": "SessionType", "started_at": "StartedAt", "user_id": "UserId"}
}

func GetOrgSessionInfoViewTypeMap() map[string]string {
	return map[string]string{"EndedAt": "*int64", "OrgId": "*uuid.UUID", "Properties": "*datastore.RawMessage", "SessionId": "*uuid.UUID", "SessionName": "*string", "SessionType": "*string", "StartedAt": "*int64", "UserId": "*uuid.UUID"}
}

type OrgMeditationView struct {
	SessionId   *uuid.UUID            `db:"session_id" json:",omitempty"`
	UserId      *uuid.UUID            `db:"user_id" json:",omitempty"`
	OrgId       *uuid.UUID            `db:"org_id" json:",omitempty"`
	SessionName *string               `db:"session_name" json:",omitempty"`
	SessionType *string               `db:"session_type" json:",omitempty"`
	StartedAt   *int64                `db:"started_at" json:",omitempty"`
	EndedAt     *int64                `db:"ended_at" json:",omitempty"`
	Properties  *datastore.RawMessage `db:"properties" json:",omitempty"`
}

func GetOrgMeditationViewFieldMap() map[string]string {
	return map[string]string{"EndedAt": "ended_at", "OrgId": "org_id", "Properties": "properties", "SessionId": "session_id", "SessionName": "session_name", "SessionType": "session_type", "StartedAt": "started_at", "UserId": "user_id"}
}

func GetOrgMeditationViewReverseFieldMap() map[string]string {
	return map[string]string{"ended_at": "EndedAt", "org_id": "OrgId", "properties": "Properties", "session_id": "SessionId", "session_name": "SessionName", "session_type": "SessionType", "started_at": "StartedAt", "user_id": "UserId"}
}

func GetOrgMeditationViewTypeMap() map[string]string {
	return map[string]string{"EndedAt": "*int64", "OrgId": "*uuid.UUID", "Properties": "*datastore.RawMessage", "SessionId": "*uuid.UUID", "SessionName": "*string", "SessionType": "*string", "StartedAt": "*int64", "UserId": "*uuid.UUID"}
}

type OrgSleepView struct {
	SessionId   *uuid.UUID            `db:"session_id" json:",omitempty"`
	UserId      *uuid.UUID            `db:"user_id" json:",omitempty"`
	OrgId       *uuid.UUID            `db:"org_id" json:",omitempty"`
	SessionName *string               `db:"session_name" json:",omitempty"`
	SessionType *string               `db:"session_type" json:",omitempty"`
	StartedAt   *int64                `db:"started_at" json:",omitempty"`
	EndedAt     *int64                `db:"ended_at" json:",omitempty"`
	Properties  *datastore.RawMessage `db:"properties" json:",omitempty"`
}

func GetOrgSleepViewFieldMap() map[string]string {
	return map[string]string{"EndedAt": "ended_at", "OrgId": "org_id", "Properties": "properties", "SessionId": "session_id", "SessionName": "session_name", "SessionType": "session_type", "StartedAt": "started_at", "UserId": "user_id"}
}

func GetOrgSleepViewReverseFieldMap() map[string]string {
	return map[string]string{"ended_at": "EndedAt", "org_id": "OrgId", "properties": "Properties", "session_id": "SessionId", "session_name": "SessionName", "session_type": "SessionType", "started_at": "StartedAt", "user_id": "UserId"}
}

func GetOrgSleepViewTypeMap() map[string]string {
	return map[string]string{"EndedAt": "*int64", "OrgId": "*uuid.UUID", "Properties": "*datastore.RawMessage", "SessionId": "*uuid.UUID", "SessionName": "*string", "SessionType": "*string", "StartedAt": "*int64", "UserId": "*uuid.UUID"}
}

type AlertRule struct {
	AlertRuleId   *uuid.UUID `db:"alert_rule_id" json:",omitempty"`
	OrgId         *uuid.UUID `db:"org_id" json:",omitempty"`
	AlertName     *string    `db:"alert_name" json:",omitempty"`
	Key           *string    `db:"key" json:",omitempty"`
	Duration      *int64     `db:"duration" json:",omitempty"`
	Enabled       *bool      `db:"enabled" json:",omitempty"`
	CreatedAt     *int64     `db:"created_at" json:",omitempty"`
	UpdatedAt     *int64     `db:"updated_at" json:",omitempty"`
	UpperLimit    *float64   `db:"upper_limit" json:",omitempty"`
	LowerLimit    *float64   `db:"lower_limit" json:",omitempty"`
	ValidFrom     *string    `db:"valid_from" json:",omitempty"`
	ValidTill     *string    `db:"valid_till" json:",omitempty"`
	SnoozedAt     *int64     `db:"snoozed_at" json:",omitempty"`
	SnoozedFor    *int64     `db:"snoozed_for" json:",omitempty"`
	DefaultSnooze *int64     `db:"default_snooze" json:",omitempty"`
}

func GetAlertRuleFieldMap() map[string]string {
	return map[string]string{"AlertName": "alert_name", "AlertRuleId": "alert_rule_id", "CreatedAt": "created_at", "DefaultSnooze": "default_snooze", "Duration": "duration", "Enabled": "enabled", "Key": "key", "LowerLimit": "lower_limit", "OrgId": "org_id", "SnoozedAt": "snoozed_at", "SnoozedFor": "snoozed_for", "UpdatedAt": "updated_at", "UpperLimit": "upper_limit", "ValidFrom": "valid_from", "ValidTill": "valid_till"}
}

func GetAlertRuleReverseFieldMap() map[string]string {
	return map[string]string{"alert_name": "AlertName", "alert_rule_id": "AlertRuleId", "created_at": "CreatedAt", "default_snooze": "DefaultSnooze", "duration": "Duration", "enabled": "Enabled", "key": "Key", "lower_limit": "LowerLimit", "org_id": "OrgId", "snoozed_at": "SnoozedAt", "snoozed_for": "SnoozedFor", "updated_at": "UpdatedAt", "upper_limit": "UpperLimit", "valid_from": "ValidFrom", "valid_till": "ValidTill"}
}

func GetAlertRuleTypeMap() map[string]string {
	return map[string]string{"AlertName": "*string", "AlertRuleId": "*uuid.UUID", "CreatedAt": "*int64", "DefaultSnooze": "*int64", "Duration": "*int64", "Enabled": "*bool", "Key": "*string", "LowerLimit": "*float64", "OrgId": "*uuid.UUID", "SnoozedAt": "*int64", "SnoozedFor": "*int64", "UpdatedAt": "*int64", "UpperLimit": "*float64", "ValidFrom": "*string", "ValidTill": "*string"}
}
