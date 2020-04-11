-- Insert Sens Auth
INSERT INTO public.auths
(auth_id, email, mobile, social, first_name, last_name, created_at, updated_at, is_sens)
VALUES(gen_random_uuid(), 'emayank@gmail.com', '+917032806003', 'Mayank', 'Joshi', 'NA', now()::INT8, 0, true);


-- create 7 auths
INSERT INTO public.auths
(auth_id, email, mobile, social, first_name, last_name, created_at, updated_at, is_sens)
VALUES('1b620e07-9ff8-40a2-8d79-9a320b463c60', 'e4@gmail.com', '+919876543214', 'S4', 'F4', 'L4', 1586288496, 0, false);
INSERT INTO public.auths
(auth_id, email, mobile, social, first_name, last_name, created_at, updated_at, is_sens)
VALUES('3a7f319b-7cad-4a75-81fd-d9536c63ff75', 'e3@gmail.com', '+919876543213', 'S3', 'F3', 'L3', 1586288496, 0, false);
INSERT INTO public.auths
VALUES('66428ee6-f10d-44de-9618-5f605f50403b', 'e1@gmail.com', '+919876543211', 'S1', 'F1', 'L1', 1586288496, 0, false);
INSERT INTO public.auths
(auth_id, email, mobile, social, first_name, last_name, created_at, updated_at, is_sens)
VALUES('8df74df4-f718-41f2-8548-56a6edd3facc', 'e2@gmail.com', '+919876543212', 'S2', 'F2', 'L2', 1586288496, 0, false);
INSERT INTO public.auths
(auth_id, email, mobile, social, first_name, last_name, created_at, updated_at, is_sens)
VALUES('aa40bfe3-d734-46c8-99ab-b26bb1632277', 'e5@gmail.com', '+919876543215', 'S5', 'F5', 'L5', 1586288496, 0, false);
INSERT INTO public.auths
(auth_id, email, mobile, social, first_name, last_name, created_at, updated_at, is_sens)
VALUES('bab09cb2-6605-45a7-a581-5876d92116ec', 'e7@gmail.com', '+919876543217', 'S7', 'F7', 'L7', 1586288496, 0, false);
INSERT INTO public.auths
(auth_id, email, mobile, social, first_name, last_name, created_at, updated_at, is_sens)
VALUES('ddf0529e-d0f1-416b-9e96-6522813f0113', 'e6@gmail.com', '+919876543216', 'S6', 'F6', 'L6', 1586288496, 0, false);


-- create 1 org
INSERT INTO public.orgs
(org_id, auth_id, org_name, created_at, updated_at)
VALUES('83632c67-5208-4949-9216-0b809a4ed1cb', '1b620e07-9ff8-40a2-8d79-9a320b463c60', 'Dozee', 1586288623, 0);


-- create 1 op
INSERT INTO public.ops
(op_id, auth_id, org_id, created_at, updated_at, "status")
VALUES('10f5e046-bf33-42c9-b42d-5dca2226c2c6', '3a7f319b-7cad-4a75-81fd-d9536c63ff75', '83632c67-5208-4949-9216-0b809a4ed1cb', 1586288750, 0, 'APPROVED');


-- create 5 users
INSERT INTO public."users"
(user_id, auth_id, org_id, access_group, created_at, updated_at, age, "status")
VALUES('208cf5c8-0fa9-47d6-9ed8-ef5cdd2cb5d6', '66428ee6-f10d-44de-9618-5f605f50403b', '83632c67-5208-4949-9216-0b809a4ed1cb', 'DEFAULT'::STRING, now()::INT8, 0, 0, 'APPROVED');

INSERT INTO public."users"
(user_id, auth_id, org_id, access_group, created_at, updated_at, age, "status")
VALUES('d99e07c1-bcb7-4e6e-819d-f66b22bb0282', '8df74df4-f718-41f2-8548-56a6edd3facc', '83632c67-5208-4949-9216-0b809a4ed1cb', 'DEFAULT'::STRING, now()::INT8, 0, 0, 'APPROVED');

INSERT INTO public."users"
(user_id, auth_id, org_id, access_group, created_at, updated_at, age, "status")
VALUES('92575c99-a421-44f2-b1c5-89e2b4ea3e44', 'aa40bfe3-d734-46c8-99ab-b26bb1632277', '83632c67-5208-4949-9216-0b809a4ed1cb', 'DEFAULT'::STRING, now()::INT8, 0, 0, 'APPROVED');

INSERT INTO public."users"
(user_id, auth_id, org_id, access_group, created_at, updated_at, age, "status")
VALUES('98a798a0-4e38-46ef-9b12-95c01c8e96ed', 'bab09cb2-6605-45a7-a581-5876d92116ec', '83632c67-5208-4949-9216-0b809a4ed1cb', 'DEFAULT'::STRING, now()::INT8, 0, 0, 'APPROVED');

INSERT INTO public."users"
(user_id, auth_id, org_id, access_group, created_at, updated_at, age, "status")
VALUES('dce42705-73cf-4c28-8940-f7228ec506eb', 'ddf0529e-d0f1-416b-9e96-6522813f0113', '83632c67-5208-4949-9216-0b809a4ed1cb', 'DEFAULT'::STRING, now()::INT8, 0, 0, 'APPROVED');

-- create 8 devices

INSERT INTO public.devices
(device_id, device_name, org_id, user_id, created_at, "status")
VALUES('dc641f85-529d-4cf5-8f0e-b80dc43309ed', 'Device1', '83632c67-5208-4949-9216-0b809a4ed1cb', '208cf5c8-0fa9-47d6-9ed8-ef5cdd2cb5d6', 1586341177, 'PAIRED');
INSERT INTO public.devices
(device_id, device_name, org_id, user_id, created_at, "status")
VALUES('8aefe53f-bec8-411d-9495-b1277bbff0e2', 'Device2', '83632c67-5208-4949-9216-0b809a4ed1cb', 'd99e07c1-bcb7-4e6e-819d-f66b22bb0282', 1586341220, 'PAIRED');
INSERT INTO public.devices
(device_id, device_name, org_id, user_id, created_at, "status")
VALUES('caec6326-13e0-44e5-be16-679de2f719c7', 'Device3', '83632c67-5208-4949-9216-0b809a4ed1cb', '92575c99-a421-44f2-b1c5-89e2b4ea3e44', 1586341177, 'PAIRED');
INSERT INTO public.devices
(device_id, device_name, org_id, user_id, created_at, "status")
VALUES('711c0ca3-1a25-415c-bb4e-39d52e7ad8a9', 'Device4', '83632c67-5208-4949-9216-0b809a4ed1cb', '98a798a0-4e38-46ef-9b12-95c01c8e96ed', 1586341220, 'PAIRED');
INSERT INTO public.devices
(device_id, device_name, org_id, user_id, created_at, "status")
VALUES('c3f4c2f7-c5cd-48aa-b4c5-a92b520a66ee', 'Device5', '83632c67-5208-4949-9216-0b809a4ed1cb', 'dce42705-73cf-4c28-8940-f7228ec506eb', 1586341220, 'PAIRED');
INSERT INTO public.devices
(device_id, device_name, org_id, user_id, created_at, "status")
VALUES('ad7f7218-5d34-4510-8382-46ec8237be54', 'Device6', '83632c67-5208-4949-9216-0b809a4ed1cb', NULL, 1586341220, 'REGISTERED');
INSERT INTO public.devices
(device_id, device_name, org_id, user_id, created_at, "status")
VALUES('e0fb8e1d-7cf7-495f-ae68-31b0ea1af920', 'Device7', '83632c67-5208-4949-9216-0b809a4ed1cb', NULL, 1586341220, 'REGISTERED');
INSERT INTO public.devices
(device_id, device_name, org_id, user_id, created_at, "status")
VALUES('0b2746b7-44ff-43a8-8676-9173fd3acfaf', 'Device8', '83632c67-5208-4949-9216-0b809a4ed1cb', NULL, 1586341220, 'REGISTERED');



-- for each user create 2 sleeps and 2 meditations

INSERT INTO public."sessions"
(session_id, user_id, session_name, session_type, started_at, ended_at)
VALUES('001c0042-f0f3-4e69-98f7-109d7265822f', '208cf5c8-0fa9-47d6-9ed8-ef5cdd2cb5d6', 'SN', 'Sleep', '1586072569', '1586292569');

INSERT INTO public."sessions"
(session_id, user_id, session_name, session_type, started_at, ended_at)
VALUES('02744d0f-c257-4c16-9db9-750917493918', '92575c99-a421-44f2-b1c5-89e2b4ea3e44', 'SN', 'Sleep', '1586076569', '1586292569');

INSERT INTO public."sessions"
(session_id, user_id, session_name, session_type, started_at, ended_at)
VALUES('57ccb971-22f6-4b02-8fb4-cac4563602f8', '98a798a0-4e38-46ef-9b12-95c01c8e96ed', 'SN', 'Sleep', '1586078569', '1586292569');

INSERT INTO public."sessions"
(session_id, user_id, session_name, session_type, started_at, ended_at)
VALUES('72894b1b-f5f5-46e6-a7c7-886aa6a0d5ff', 'd99e07c1-bcb7-4e6e-819d-f66b22bb0282', 'SN', 'Sleep', '1586075569', '1586292569');

INSERT INTO public."sessions"
(session_id, user_id, session_name, session_type, started_at, ended_at)
VALUES('77312aee-192a-4a10-927a-c529c941ccff', 'dce42705-73cf-4c28-8940-f7228ec506eb', 'SN', 'Sleep', '1586074569', '1586292569');

INSERT INTO public."sessions"
(session_id, user_id, session_name, session_type, started_at, ended_at)
VALUES('8043ba29-7624-4640-a829-47fcb1d0eda7', '208cf5c8-0fa9-47d6-9ed8-ef5cdd2cb5d6', 'SN', 'Sleep', '1586206169', '1586225969');

INSERT INTO public."sessions"
(session_id, user_id, session_name, session_type, started_at, ended_at)
VALUES('851b992f-1364-49ae-8953-31b80a256667', '92575c99-a421-44f2-b1c5-89e2b4ea3e44', 'SN', 'Sleep', '1586202169', '1586223969');

INSERT INTO public."sessions"
(session_id, user_id, session_name, session_type, started_at, ended_at)
VALUES('ce58342e-4973-4c3c-960a-fcabbcc16e11', '98a798a0-4e38-46ef-9b12-95c01c8e96ed', 'SN', 'Sleep', '1586208169', '1586225369');

INSERT INTO public."sessions"
(session_id, user_id, session_name, session_type, started_at, ended_at)
VALUES('ddbca8b9-6ed4-4c71-81b1-ced144b7548c', 'd99e07c1-bcb7-4e6e-819d-f66b22bb0282', 'SN', 'Sleep', '1586206169', '1586221969');

INSERT INTO public."sessions"
(session_id, user_id, session_name, session_type, started_at, ended_at)
VALUES('ef4d7791-2aea-4819-a1c9-75ac05f1b5e2', 'dce42705-73cf-4c28-8940-f7228ec506eb', 'SN', 'Sleep', '1586206169', '1586228969');

-- Meditations
INSERT INTO public."sessions"
(session_id, user_id, session_name, session_type, started_at, ended_at)
VALUES('0a4a1e5a-a7c1-4cc3-b1e4-e67945368967', '208cf5c8-0fa9-47d6-9ed8-ef5cdd2cb5d6', 'SN', 'Meditation', '1586115769', '1586114769');

INSERT INTO public."sessions"
(session_id, user_id, session_name, session_type, started_at, ended_at)
VALUES('3964667b-06db-46bd-9cf4-3e6767415a5a', '92575c99-a421-44f2-b1c5-89e2b4ea3e44', 'SN', 'Meditation', '1586115769', '1586114069');

INSERT INTO public."sessions"
(session_id, user_id, session_name, session_type, started_at, ended_at)
VALUES('3a6e43f1-cf2c-4abc-998e-3388ea9df2d4', '98a798a0-4e38-46ef-9b12-95c01c8e96ed', 'SN', 'Meditation', '1586115769', '1586113769');

INSERT INTO public."sessions"
(session_id, user_id, session_name, session_type, started_at, ended_at)
VALUES('3bb4a251-7e18-4733-a20a-7f769756a5b1', 'd99e07c1-bcb7-4e6e-819d-f66b22bb0282', 'SN', 'Meditation', '1586115769', '1586114369');

INSERT INTO public."sessions"
(session_id, user_id, session_name, session_type, started_at, ended_at)
VALUES('45134944-211c-4e9b-8886-677f2b6faaf2', 'dce42705-73cf-4c28-8940-f7228ec506eb', 'SN', 'Meditation', '1586115769', '1586114769');

INSERT INTO public."sessions"
(session_id, user_id, session_name, session_type, started_at, ended_at)
VALUES('48e01b74-6f88-4f01-a3f3-420310b50be6', '208cf5c8-0fa9-47d6-9ed8-ef5cdd2cb5d6', 'SN', 'Meditation', '1586115769', '1586113769');

INSERT INTO public."sessions"
(session_id, user_id, session_name, session_type, started_at, ended_at)
VALUES('51a49fa4-b4e2-4510-8b2a-a8fdc7be839f', '92575c99-a421-44f2-b1c5-89e2b4ea3e44', 'SN', 'Meditation', '1586115769', '1586112700');

INSERT INTO public."sessions"
(session_id, user_id, session_name, session_type, started_at, ended_at)
VALUES('a2da13e1-2dd1-4c6a-8494-4b4d0f01ff75', '98a798a0-4e38-46ef-9b12-95c01c8e96ed', 'SN', 'Meditation', '1586115769', '1586114769');

INSERT INTO public."sessions"
(session_id, user_id, session_name, session_type, started_at, ended_at)
VALUES('aa060b70-25aa-4c0d-9985-91540dfec35e', 'd99e07c1-bcb7-4e6e-819d-f66b22bb0282', 'SN', 'Meditation', '1586115769', '1586113769');

INSERT INTO public."sessions"
(session_id, user_id, session_name, session_type, started_at, ended_at)
VALUES('aeb82f5c-6f63-46e8-a4d9-4c97e1d6ecee', 'dce42705-73cf-4c28-8940-f7228ec506eb', 'SN', 'Meditation', '1586115769', '1586114769');

-- add properties for each
INSERT INTO public.session_properties (session_id, "key", value) VALUES
('001c0042-f0f3-4e69-98f7-109d7265822f', 'HeartRate', '90'), ('001c0042-f0f3-4e69-98f7-109d7265822f', 'BreathRate', '100'), ('001c0042-f0f3-4e69-98f7-109d7265822f', 'Score', '100'), ('001c0042-f0f3-4e69-98f7-109d7265822f', 'LastSyncedAt', '100'),
('02744d0f-c257-4c16-9db9-750917493918', 'HeartRate', '100'), ('02744d0f-c257-4c16-9db9-750917493918', 'BreathRate', '100'), ('02744d0f-c257-4c16-9db9-750917493918', 'Score', '100'), ('02744d0f-c257-4c16-9db9-750917493918', 'LastSyncedAt', '100'),
('57ccb971-22f6-4b02-8fb4-cac4563602f8', 'HeartRate', '66'), ('57ccb971-22f6-4b02-8fb4-cac4563602f8', 'BreathRate', '65'), ('57ccb971-22f6-4b02-8fb4-cac4563602f8', 'Score', '100'), ('57ccb971-22f6-4b02-8fb4-cac4563602f8', 'LastSyncedAt', '100'),
('72894b1b-f5f5-46e6-a7c7-886aa6a0d5ff', 'HeartRate', '100'), ('72894b1b-f5f5-46e6-a7c7-886aa6a0d5ff', 'BreathRate', '66'), ('72894b1b-f5f5-46e6-a7c7-886aa6a0d5ff', 'Score', '77'), ('72894b1b-f5f5-46e6-a7c7-886aa6a0d5ff', 'LastSyncedAt', '100'),
('77312aee-192a-4a10-927a-c529c941ccff', 'HeartRate', '65'), ('77312aee-192a-4a10-927a-c529c941ccff', 'BreathRate', '66'), ('77312aee-192a-4a10-927a-c529c941ccff', 'Score', '100'), ('77312aee-192a-4a10-927a-c529c941ccff', 'LastSyncedAt', '100'),
('8043ba29-7624-4640-a829-47fcb1d0eda7', 'HeartRate', '100'), ('8043ba29-7624-4640-a829-47fcb1d0eda7', 'BreathRate', '100'), ('8043ba29-7624-4640-a829-47fcb1d0eda7', 'Score', '67'), ('8043ba29-7624-4640-a829-47fcb1d0eda7', 'LastSyncedAt', '100'),
('851b992f-1364-49ae-8953-31b80a256667', 'HeartRate', '100'), ('851b992f-1364-49ae-8953-31b80a256667', 'BreathRate', '100'), ('851b992f-1364-49ae-8953-31b80a256667', 'Score', '77'), ('851b992f-1364-49ae-8953-31b80a256667', 'LastSyncedAt', '100'),
('ce58342e-4973-4c3c-960a-fcabbcc16e11', 'HeartRate', '77'), ('ce58342e-4973-4c3c-960a-fcabbcc16e11', 'BreathRate', '76'), ('ce58342e-4973-4c3c-960a-fcabbcc16e11', 'Score', '100'), ('ce58342e-4973-4c3c-960a-fcabbcc16e11', 'LastSyncedAt', '100'),
('ddbca8b9-6ed4-4c71-81b1-ced144b7548c', 'HeartRate', '100'), ('ddbca8b9-6ed4-4c71-81b1-ced144b7548c', 'BreathRate', '100'), ('ddbca8b9-6ed4-4c71-81b1-ced144b7548c', 'Score', '100'), ('ddbca8b9-6ed4-4c71-81b1-ced144b7548c', 'LastSyncedAt', '100'),
('ef4d7791-2aea-4819-a1c9-75ac05f1b5e2', 'HeartRate', '100'), ('ef4d7791-2aea-4819-a1c9-75ac05f1b5e2', 'BreathRate', '100'), ('ef4d7791-2aea-4819-a1c9-75ac05f1b5e2', 'Score', '100'), ('ef4d7791-2aea-4819-a1c9-75ac05f1b5e2', 'LastSyncedAt', '100');

INSERT INTO public.session_properties (session_id, "key", value) VALUES
('0a4a1e5a-a7c1-4cc3-b1e4-e67945368967', 'HeartRate', '100'), ('0a4a1e5a-a7c1-4cc3-b1e4-e67945368967', 'BreathRate', '100'), ('0a4a1e5a-a7c1-4cc3-b1e4-e67945368967', 'Score', '100'), ('0a4a1e5a-a7c1-4cc3-b1e4-e67945368967', 'LastSyncedAt', '100'),
('3964667b-06db-46bd-9cf4-3e6767415a5a', 'HeartRate', '100'), ('3964667b-06db-46bd-9cf4-3e6767415a5a', 'BreathRate', '100'), ('3964667b-06db-46bd-9cf4-3e6767415a5a', 'Score', '100'), ('3964667b-06db-46bd-9cf4-3e6767415a5a', 'LastSyncedAt', '100'),
('3a6e43f1-cf2c-4abc-998e-3388ea9df2d4', 'HeartRate', '100'), ('3a6e43f1-cf2c-4abc-998e-3388ea9df2d4', 'BreathRate', '100'), ('3a6e43f1-cf2c-4abc-998e-3388ea9df2d4', 'Score', '100'), ('3a6e43f1-cf2c-4abc-998e-3388ea9df2d4', 'LastSyncedAt', '100'),
('3bb4a251-7e18-4733-a20a-7f769756a5b1', 'HeartRate', '100'), ('3bb4a251-7e18-4733-a20a-7f769756a5b1', 'BreathRate', '100'), ('3bb4a251-7e18-4733-a20a-7f769756a5b1', 'Score', '100'), ('3bb4a251-7e18-4733-a20a-7f769756a5b1', 'LastSyncedAt', '100'),
('45134944-211c-4e9b-8886-677f2b6faaf2', 'HeartRate', '100'), ('45134944-211c-4e9b-8886-677f2b6faaf2', 'BreathRate', '100'), ('45134944-211c-4e9b-8886-677f2b6faaf2', 'Score', '100'), ('45134944-211c-4e9b-8886-677f2b6faaf2', 'LastSyncedAt', '100'),
('48e01b74-6f88-4f01-a3f3-420310b50be6', 'HeartRate', '100'), ('48e01b74-6f88-4f01-a3f3-420310b50be6', 'BreathRate', '100'), ('48e01b74-6f88-4f01-a3f3-420310b50be6', 'Score', '100'), ('48e01b74-6f88-4f01-a3f3-420310b50be6', 'LastSyncedAt', '100'),
('51a49fa4-b4e2-4510-8b2a-a8fdc7be839f', 'HeartRate', '100'), ('51a49fa4-b4e2-4510-8b2a-a8fdc7be839f', 'BreathRate', '100'), ('51a49fa4-b4e2-4510-8b2a-a8fdc7be839f', 'Score', '100'), ('51a49fa4-b4e2-4510-8b2a-a8fdc7be839f', 'LastSyncedAt', '100'),
('a2da13e1-2dd1-4c6a-8494-4b4d0f01ff75', 'HeartRate', '100'), ('a2da13e1-2dd1-4c6a-8494-4b4d0f01ff75', 'BreathRate', '100'), ('a2da13e1-2dd1-4c6a-8494-4b4d0f01ff75', 'Score', '100'), ('a2da13e1-2dd1-4c6a-8494-4b4d0f01ff75', 'LastSyncedAt', '100'),
('aa060b70-25aa-4c0d-9985-91540dfec35e', 'HeartRate', '100'), ('aa060b70-25aa-4c0d-9985-91540dfec35e', 'BreathRate', '100'), ('aa060b70-25aa-4c0d-9985-91540dfec35e', 'Score', '100'), ('aa060b70-25aa-4c0d-9985-91540dfec35e', 'LastSyncedAt', '100'),
('aeb82f5c-6f63-46e8-a4d9-4c97e1d6ecee', 'HeartRate', '100'), ('aeb82f5c-6f63-46e8-a4d9-4c97e1d6ecee', 'BreathRate', '100'), ('aeb82f5c-6f63-46e8-a4d9-4c97e1d6ecee', 'Score', '100'), ('aeb82f5c-6f63-46e8-a4d9-4c97e1d6ecee', 'LastSyncedAt', '100');
