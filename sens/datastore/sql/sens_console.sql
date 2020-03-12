CREATE TABLE [orgs] (
  [id] uuid PRIMARY KEY,
  [name] text UNIQUE
)
GO

CREATE TABLE [ops] (
  [id] uuid PRIMARY KEY,
  [org_id] uuid
)
GO

CREATE TABLE [users] (
  [id] uuid PRIMARY KEY,
  [org_id] uuid
)
GO

CREATE TABLE [auths] (
  [id] uuid PRIMARY KEY,
  [email] text UNIQUE NOT NULL,
  [mobile] text UNIQUE NOT NULL,
  [social] text UNIQUE NOT NULL,
  [first_name] text,
  [last_name] text
)
GO

CREATE TABLE [org_auths] (
  [org_id] uuid,
  [auth_id] uuid,
  PRIMARY KEY ([org_id], [auth_id])
)
GO

CREATE TABLE [user_auths] (
  [user_id] uuid,
  [auth_id] uuid
)
GO

CREATE TABLE [api_groups] (
  [id] uuid PRIMARY KEY,
  [name] text UNIQUE
)
GO

CREATE TABLE [apis] (
  [id] uuid PRIMARY KEY,
  [api_group_id] uuid,
  [path] text UNIQUE
)
GO

CREATE TABLE [org_apis] (
  [org_id] uuid,
  [api_id] uuid,
  PRIMARY KEY ([org_id], [api_id])
)
GO

CREATE TABLE [op_apis] (
  [op_id] uuid,
  [api_id] uuid,
  PRIMARY KEY ([op_id], [api_id])
)
GO

CREATE TABLE [user_apis] (
  [user_id] uuid,
  [api_id] uuid,
  PRIMARY KEY ([user_id], [api_id])
)
GO

ALTER TABLE [ops] ADD FOREIGN KEY ([org_id]) REFERENCES [orgs] ([id])
GO

ALTER TABLE [users] ADD FOREIGN KEY ([org_id]) REFERENCES [orgs] ([id])
GO

ALTER TABLE [org_auths] ADD FOREIGN KEY ([org_id]) REFERENCES [orgs] ([id])
GO

ALTER TABLE [org_auths] ADD FOREIGN KEY ([auth_id]) REFERENCES [auths] ([id])
GO

ALTER TABLE [user_auths] ADD FOREIGN KEY ([user_id]) REFERENCES [users] ([id])
GO

ALTER TABLE [user_auths] ADD FOREIGN KEY ([auth_id]) REFERENCES [auths] ([id])
GO

ALTER TABLE [apis] ADD FOREIGN KEY ([api_group_id]) REFERENCES [api_groups] ([id])
GO

ALTER TABLE [org_apis] ADD FOREIGN KEY ([org_id]) REFERENCES [orgs] ([id])
GO

ALTER TABLE [org_apis] ADD FOREIGN KEY ([api_id]) REFERENCES [apis] ([id])
GO

ALTER TABLE [op_apis] ADD FOREIGN KEY ([op_id]) REFERENCES [ops] ([id])
GO

ALTER TABLE [op_apis] ADD FOREIGN KEY ([api_id]) REFERENCES [apis] ([id])
GO

ALTER TABLE [user_apis] ADD FOREIGN KEY ([user_id]) REFERENCES [users] ([id])
GO

ALTER TABLE [user_apis] ADD FOREIGN KEY ([api_id]) REFERENCES [apis] ([id])
GO

