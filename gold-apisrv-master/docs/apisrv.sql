-- =============================================================================
-- Diagram Name: apisrv
-- Created on: 6/8/2022 12:02:23 PM
-- Diagram Version: 
-- =============================================================================

CREATE TABLE "statuses" (
	"statusId" SERIAL NOT NULL,
	"title" varchar(255) NOT NULL,
	"alias" varchar(64) NOT NULL,
	CONSTRAINT "statuses_pkey" PRIMARY KEY("statusId"),
	CONSTRAINT "statuses_alias_key" UNIQUE("alias")
);

CREATE TABLE "users" (
	"userId" SERIAL NOT NULL,
	"login" varchar(64) NOT NULL,
	"password" varchar(64) NOT NULL,
	"authKey" varchar(32),
	"createdAt" timestamp with time zone NOT NULL DEFAULT now(),
	"lastActivityAt" timestamp with time zone,
	"statusId" int4 NOT NULL,
	CONSTRAINT "users_pkey" PRIMARY KEY("userId")
);

CREATE INDEX "IX_FK_users_statusId_users" ON "users" USING BTREE (
	"statusId"
);


CREATE TABLE "vfsFiles" (
	"fileId" SERIAL NOT NULL,
	"folderId" int4 NOT NULL,
	"title" varchar(255) NOT NULL,
	"path" varchar(255) NOT NULL,
	"params" text,
	"isFavorite" bool DEFAULT false,
	"mimeType" varchar(255) NOT NULL,
	"fileSize" int4 DEFAULT 0,
	"fileExists" bool NOT NULL DEFAULT true,
	"createdAt" timestamp NOT NULL DEFAULT now(),
	"statusId" int4 NOT NULL,
	CONSTRAINT "vfsFiles_pkey" PRIMARY KEY("fileId")
);

CREATE INDEX "IX_FK_vfsFiles_folderId_vfsFiles" ON "vfsFiles" USING BTREE (
	"folderId"
);


CREATE INDEX "IX_FK_vfsFiles_statusId_vfsFiles" ON "vfsFiles" USING BTREE (
	"statusId"
);


CREATE TABLE "vfsFolders" (
	"folderId" SERIAL NOT NULL,
	"parentFolderId" int4,
	"title" varchar(255) NOT NULL,
	"isFavorite" bool DEFAULT false,
	"createdAt" timestamp NOT NULL DEFAULT now(),
	"statusId" int4 NOT NULL,
	CONSTRAINT "vfsFolders_pkey" PRIMARY KEY("folderId")
);

CREATE INDEX "IX_FK_vfsFolders_folderId_vfsFolders" ON "vfsFolders" USING BTREE (
	"parentFolderId"
);


CREATE INDEX "IX_FK_vfsFolders_statusId_vfsFolders" ON "vfsFolders" USING BTREE (
	"statusId"
);


CREATE TABLE "vfsHashes" (
	"hash" varchar(40) NOT NULL,
	"namespace" varchar(32) NOT NULL,
	"extension" varchar(4) NOT NULL,
	"fileSize" int4 NOT NULL DEFAULT 0,
	"width" int4 NOT NULL DEFAULT 0,
	"height" int4 NOT NULL DEFAULT 0,
	"blurhash" text,
	"error" text,
	"createdAt" timestamp with time zone NOT NULL DEFAULT now(),
	"indexedAt" timestamp with time zone,
	CONSTRAINT "vfsHashes_pkey" PRIMARY KEY("hash","namespace")
);

CREATE INDEX "IX_vfsHashes_indexedAt" ON "vfsHashes" USING BTREE (
	"indexedAt"
);

--=============================================================================
--NewsPortal
-- =============================================================================


CREATE TABLE "tags" (
    "tagId" SERIAL NOT NULL,
    "title" varchar(256) NOT NULL,
    "statusId" int4 NOT NULL,
    PRIMARY KEY("tagId")
);

CREATE TABLE "news" (
    "newsId" SERIAL NOT NULL,
    "title" varchar(256) NOT NULL,
    "alias" varchar(32) NOT NULL,
    "content" text,
    "categoryId" int4 NOT NULL,
    "createdAt" timestamp with time zone NOT NULL DEFAULT now(),
    "updatedAt" timestamp with time zone,
    "publicationDate" timestamp with time zone NOT NULL,
    "tagIds" int4[],
    "statusId" int4 NOT NULL,
    PRIMARY KEY("newsId")
);

CREATE TABLE "categories" (
      "categoryId" SERIAL NOT NULL,
      "title" varchar(256) NOT NULL,
      "orderNumber" int4 NOT NULL,
      "statusId" int4 NOT NULL,
      PRIMARY KEY("categoryId")
);


ALTER TABLE "tags" ADD CONSTRAINT "Ref_tags_to_statuses" FOREIGN KEY ("statusId")
    REFERENCES "statuses"("statusId")
        MATCH SIMPLE
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
    NOT DEFERRABLE;

ALTER TABLE "news" ADD CONSTRAINT "Ref_news_to_statuses" FOREIGN KEY ("statusId")
    REFERENCES "statuses"("statusId")
        MATCH SIMPLE
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
    NOT DEFERRABLE;

ALTER TABLE "news" ADD CONSTRAINT "Ref_news_to_categories" FOREIGN KEY ("categoryId")
    REFERENCES "categories"("categoryId")
        MATCH SIMPLE
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
    NOT DEFERRABLE;

ALTER TABLE "categories" ADD CONSTRAINT "Ref_categories_to_statuses" FOREIGN KEY ("statusId")
    REFERENCES "statuses"("statusId")
        MATCH SIMPLE
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
    NOT DEFERRABLE;




ALTER TABLE "users" ADD CONSTRAINT "FK_users_statusId" FOREIGN KEY ("statusId")
	REFERENCES "statuses"("statusId")
	MATCH SIMPLE
	ON DELETE RESTRICT
	ON UPDATE RESTRICT
	NOT DEFERRABLE;

ALTER TABLE "vfsFiles" ADD CONSTRAINT "vfsFiles_folderId_fkey" FOREIGN KEY ("folderId")
	REFERENCES "vfsFolders"("folderId")
	MATCH SIMPLE
	ON DELETE RESTRICT
	ON UPDATE RESTRICT
	NOT DEFERRABLE;

ALTER TABLE "vfsFiles" ADD CONSTRAINT "vfsFiles_statusId_fkey" FOREIGN KEY ("statusId")
	REFERENCES "statuses"("statusId")
	MATCH SIMPLE
	ON DELETE RESTRICT
	ON UPDATE RESTRICT
	NOT DEFERRABLE;

ALTER TABLE "vfsFolders" ADD CONSTRAINT "vfsFolders_parentFolderId_fkey" FOREIGN KEY ("parentFolderId")
	REFERENCES "vfsFolders"("folderId")
	MATCH SIMPLE
	ON DELETE RESTRICT
	ON UPDATE RESTRICT
	NOT DEFERRABLE;

ALTER TABLE "vfsFolders" ADD CONSTRAINT "vfsFolders_statusId_fkey" FOREIGN KEY ("statusId")
	REFERENCES "statuses"("statusId")
	MATCH SIMPLE
	ON DELETE RESTRICT
	ON UPDATE RESTRICT
	NOT DEFERRABLE;


