-- create 7 auths
INSERT INTO public.auths
(auth_id, email, mobile, social, first_name, last_name, created_at, updated_at, is_sens)
VALUES('1b620e07-9ff8-40a2-8d79-9a320b463c60', 'e4@gmail.com', '+919876543214', 'S4', 'F4', 'L4', 1586288496, 0, false);
INSERT INTO public.auths
(auth_id, email, mobile, social, first_name, last_name, created_at, updated_at, is_sens)
VALUES('3a7f319b-7cad-4a75-81fd-d9536c63ff75', 'e3@gmail.com', '+919876543213', 'S3', 'F3', 'L3', 1586288496, 0, false);
INSERT INTO public.auths
(auth_id, email, mobile, social, first_name, last_name, created_at, updated_at, is_sens)
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
VALUES(gen_random_uuid(), '66428ee6-f10d-44de-9618-5f605f50403b', '83632c67-5208-4949-9216-0b809a4ed1cb', 'DEFAULT'::STRING, now()::INT8, 0, 0, 'APPROVED');

INSERT INTO public."users"
(user_id, auth_id, org_id, access_group, created_at, updated_at, age, "status")
VALUES(gen_random_uuid(), '8df74df4-f718-41f2-8548-56a6edd3facc', '83632c67-5208-4949-9216-0b809a4ed1cb', 'DEFAULT'::STRING, now()::INT8, 0, 0, 'APPROVED');

INSERT INTO public."users"
(user_id, auth_id, org_id, access_group, created_at, updated_at, age, "status")
VALUES(gen_random_uuid(), 'aa40bfe3-d734-46c8-99ab-b26bb1632277', '83632c67-5208-4949-9216-0b809a4ed1cb', 'DEFAULT'::STRING, now()::INT8, 0, 0, 'APPROVED');

INSERT INTO public."users"
(user_id, auth_id, org_id, access_group, created_at, updated_at, age, "status")
VALUES(gen_random_uuid(), 'bab09cb2-6605-45a7-a581-5876d92116ec', '83632c67-5208-4949-9216-0b809a4ed1cb', 'DEFAULT'::STRING, now()::INT8, 0, 0, 'APPROVED');

INSERT INTO public."users"
(user_id, auth_id, org_id, access_group, created_at, updated_at, age, "status")
VALUES(gen_random_uuid(), 'ddf0529e-d0f1-416b-9e96-6522813f0113', '83632c67-5208-4949-9216-0b809a4ed1cb', 'DEFAULT'::STRING, now()::INT8, 0, 0, 'APPROVED');

-- create 8 devices

INSERT INTO public.devices
(row_id, device_id, device_name, org_id, user_id, created_at, "status")
VALUES(gen_random_uuid(), gen_random_uuid(), 'Device1', '83632c67-5208-4949-9216-0b809a4ed1cb', '208cf5c8-0fa9-47d6-9ed8-ef5cdd2cb5d6', now()::INT8, 'PAIRED');

INSERT INTO public.devices
(row_id, device_id, device_name, org_id, user_id, created_at, "status")
VALUES(gen_random_uuid(), gen_random_uuid(), 'Device2', '83632c67-5208-4949-9216-0b809a4ed1cb', '92575c99-a421-44f2-b1c5-89e2b4ea3e44', now()::INT8, 'PAIRED');

INSERT INTO public.devices
(row_id, device_id, device_name, org_id, user_id, created_at, "status")
VALUES(gen_random_uuid(), gen_random_uuid(), 'Device3', '83632c67-5208-4949-9216-0b809a4ed1cb', '98a798a0-4e38-46ef-9b12-95c01c8e96ed', now()::INT8, 'PAIRED');

INSERT INTO public.devices
(row_id, device_id, device_name, org_id, user_id, created_at, "status")
VALUES(gen_random_uuid(), gen_random_uuid(), 'Device4', '83632c67-5208-4949-9216-0b809a4ed1cb', 'd99e07c1-bcb7-4e6e-819d-f66b22bb0282', now()::INT8, 'PAIRED');

INSERT INTO public.devices
(row_id, device_id, device_name, org_id, user_id, created_at, "status")
VALUES(gen_random_uuid(), gen_random_uuid(), 'Device5', '83632c67-5208-4949-9216-0b809a4ed1cb', 'dce42705-73cf-4c28-8940-f7228ec506eb', now()::INT8, 'PAIRED');

-- for each user create 2 sleeps and 2 meditations


-- add properties for each