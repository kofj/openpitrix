ALTER TABLE app
	ADD COLUMN active BOOL NOT NULL DEFAULT false;
ALTER TABLE app
	DROP PRIMARY KEY,
	ADD PRIMARY KEY (app_id, active);

ALTER TABLE app_version
	ADD COLUMN active BOOL NOT NULL DEFAULT false;
ALTER TABLE app_version
	DROP PRIMARY KEY,
	ADD PRIMARY KEY (version_id, active);
