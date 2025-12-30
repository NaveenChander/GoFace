-- Drop existing tables to perform a clean recreate
DROP TABLE IF EXISTS public.facttransactioninstrumentdetail;

DROP TABLE IF EXISTS public.facttransactioninstrumentdenomination;

DROP TABLE IF EXISTS public.facttransactioninstrument;

DROP TABLE IF EXISTS public.facttransactionsession;

DROP TABLE IF EXISTS public.facttransaction;

DROP TABLE IF EXISTS public.dimtransactiontype;

DROP TABLE IF EXISTS public.diminstrumenttype;

DROP TABLE IF EXISTS public.dimlocation;

DROP TABLE IF EXISTS public.dimworkarea;

DROP TABLE IF EXISTS public.dimproperty;

DROP TABLE IF EXISTS public.dimdate;

DROP TABLE IF EXISTS public.Patron;

DROP EXTENSION IF EXISTS vector;

-- Enable extensions
CREATE EXTENSION IF NOT EXISTS vector;

-- 1. Patron
CREATE TABLE public.Patron (
    PatronID BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    SubjectNumber VARCHAR(10) NOT NULL,
    FirstName VARCHAR(100) NOT NULL,
    LastName VARCHAR(100) NOT NULL,
    InsertedDatetime TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX idx_unique_subject_number ON public.Patron (SubjectNumber);

-- 2. DIMInstrumentType
CREATE TABLE public.diminstrumenttype (
    instrumenttypesk INTEGER GENERATED ALWAYS AS IDENTITY (
        START
        WITH
            1 INCREMENT BY 1
    ) NOT NULL,
    instrumenttypename VARCHAR(128) NOT NULL,
    CONSTRAINT pkc_diminstrumenttype PRIMARY KEY (instrumenttypesk)
);

ALTER TABLE public.diminstrumenttype
ADD CONSTRAINT uqc_instrumenttypename UNIQUE (instrumenttypename);

-- 3. DimProperty
CREATE TABLE public.dimproperty (
    propertysk SMALLINT GENERATED ALWAYS AS IDENTITY (
        START
        WITH
            1 INCREMENT BY 1
    ) NOT NULL,
    propertyname VARCHAR(256) NOT NULL,
    inserteddatetime TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT pkc_dimproperty PRIMARY KEY (propertysk)
);

ALTER TABLE public.dimproperty
ADD CONSTRAINT uqc_propertyname UNIQUE (propertyname);

-- 4. DimWorkArea
CREATE TABLE public.dimworkarea (
    workareask INTEGER GENERATED ALWAYS AS IDENTITY (
        START
        WITH
            1 INCREMENT BY 1
    ) NOT NULL,
    workareaname VARCHAR(256) NOT NULL,
    propertysk SMALLINT NULL,
    inserteddatetime TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT pkc_dimworkarea PRIMARY KEY (workareask),
    CONSTRAINT fkc_dimworkarea_property FOREIGN KEY (propertysk) REFERENCES public.dimproperty (propertysk)
);

ALTER TABLE public.dimworkarea
ADD CONSTRAINT uqc_workareaname UNIQUE (workareaname);

-- 5. DimLocation
CREATE TABLE public.dimlocation (
    locationsk BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    locationname VARCHAR(256) NOT NULL,
    workareaname VARCHAR(256) NOT NULL,
    propertyname VARCHAR(256) NOT NULL
);

ALTER TABLE public.dimlocation
ADD CONSTRAINT uqc_locationname UNIQUE (locationname);

-- 6. DimDate
CREATE TABLE public.dimdate (
    datesk INTEGER NOT NULL,
    fulldate DATE NOT NULL,
    dayofweek SMALLINT NOT NULL,
    month SMALLINT NOT NULL,
    quarter SMALLINT NOT NULL,
    year SMALLINT NOT NULL,
    CONSTRAINT pkc_dimdate PRIMARY KEY (datesk)
);

ALTER TABLE public.dimdate
ADD CONSTRAINT uqc_fulldate UNIQUE (fulldate);

ALTER TABLE public.dimdate
ADD CONSTRAINT chk_dimdate_dayofweek CHECK (dayofweek BETWEEN 1 AND 7),
ADD CONSTRAINT chk_dimdate_month CHECK (month BETWEEN 1 AND 12),
ADD CONSTRAINT chk_dimdate_quarter CHECK (quarter BETWEEN 1 AND 4);

-- 7. DimTransactionType
CREATE TABLE public.dimtransactiontype (
    transactiontypesk BIGINT GENERATED ALWAYS AS IDENTITY (
        START
        WITH
            1 INCREMENT BY 1
    ) NOT NULL,
    transactiontypecode VARCHAR(32) NOT NULL,
    name VARCHAR(512) NOT NULL,
    CONSTRAINT pkc_dimtransactiontype PRIMARY KEY (transactiontypesk)
);

ALTER TABLE public.dimtransactiontype
ADD CONSTRAINT uqc_transactiontypecode UNIQUE (transactiontypecode);

-- 8. FactTransaction
CREATE TABLE public.facttransaction (
    transactionsk BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    sourcetransactionnumber VARCHAR(256) NOT NULL,
    patronid BIGINT NULL,
    locationsk INTEGER NULL,
    transactionutcdatetime TIMESTAMPTZ NOT NULL,
    gamingdatesk INTEGER NOT NULL,
    transactiondatesk INTEGER NOT NULL,
    transactiontypesk BIGINT NOT NULL,
    iscarded BOOLEAN NOT NULL DEFAULT FALSE,
    issenttoincident BOOLEAN NOT NULL DEFAULT FALSE,
    inserteddatetime TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT fkc_facttransaction_patron FOREIGN KEY (patronid) REFERENCES public.Patron (PatronID),
    CONSTRAINT fkc_facttransaction_location FOREIGN KEY (locationsk) REFERENCES public.dimlocation (locationsk),
    CONSTRAINT fkc_facttransaction_transactiontype FOREIGN KEY (transactiontypesk) REFERENCES public.dimtransactiontype (transactiontypesk)
);

-- 9. FactTransactionInstrument
CREATE TABLE public.facttransactioninstrument (
    transactioninstrumentsk BIGINT GENERATED ALWAYS AS IDENTITY (
        START
        WITH
            1 INCREMENT BY 1
    ) NOT NULL,
    transactionsk BIGINT NOT NULL,
    instrumenttypesk INTEGER NOT NULL,
    amountin NUMERIC(18, 2) NULL,
    amountout NUMERIC(18, 2) NULL,
    inserteddatetime TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT pkc_facttransactioninstrument PRIMARY KEY (transactioninstrumentsk)
);

-- 10. FactTransactionInstrumentDenomination
CREATE TABLE public.facttransactioninstrumentdenomination (
    transactioninstrumentdenominationsk BIGINT GENERATED ALWAYS AS IDENTITY (
        START
        WITH
            1 INCREMENT BY 1
    ) NOT NULL,
    transactioninstrumentsk BIGINT NOT NULL,
    denominationvalue SMALLINT NOT NULL,
    denominationcount INTEGER NOT NULL,
    inserteddatetime TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT pkc_facttransactioninstrumentdenomination PRIMARY KEY (
        transactioninstrumentdenominationsk
    )
);

-- 11. FactTransactionInstrumentDetail
CREATE TABLE public.facttransactioninstrumentdetail (
    transactioninstrumentdetailsk BIGINT GENERATED ALWAYS AS IDENTITY (
        START
        WITH
            1 INCREMENT BY 1
    ) NOT NULL,
    transactioninstrumentsk BIGINT NOT NULL,
    detailsk SMALLINT NOT NULL,
    value VARCHAR(512) NOT NULL,
    inserteddatetime TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT pkc_facttransactioninstrumentdetail PRIMARY KEY (transactioninstrumentdetailsk),
    CONSTRAINT fkc_facttransactioninstrumentdetail_transactioninstrument FOREIGN KEY (transactioninstrumentsk) REFERENCES public.facttransactioninstrument (transactioninstrumentsk)
);

-- 12. FactTransactionSession
CREATE TABLE public.facttransactionsession (
    transactionsessionsk BIGINT GENERATED ALWAYS AS IDENTITY (
        START
        WITH
            1 INCREMENT BY 1
    ) NOT NULL,
    subjectsk BIGINT NULL,
    propertysk SMALLINT NOT NULL,
    workareask INTEGER NULL,
    locationsk INTEGER NULL,
    gamingdatesk INTEGER NOT NULL,
    sessionstartutcdatetime TIMESTAMPTZ NULL,
    sessionendutcdatetime TIMESTAMPTZ NULL,
    cashin NUMERIC(18, 2) NULL,
    cashout NUMERIC(18, 2) NULL,
    noncashin NUMERIC(18, 2) NULL,
    noncashout NUMERIC(18, 2) NULL,
    cashincount SMALLINT NOT NULL,
    totalcount SMALLINT NOT NULL,
    avggapseconds INTEGER NULL,
    mingapseconds INTEGER NULL,
    turnover NUMERIC(18, 2) NOT NULL,
    winamount NUMERIC(18, 2) NOT NULL,
    cashbox NUMERIC(18, 2) NOT NULL,
    spincount INTEGER NOT NULL,
    inserteddatetime TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    username VARCHAR(64) NOT NULL,
    clientname VARCHAR(64) NOT NULL,
    CONSTRAINT pkc_facttransactionsession PRIMARY KEY (transactionsessionsk)
);

CREATE INDEX idx_ftsession_gamingdate ON public.facttransactionsession (gamingdatesk);

CREATE INDEX idx_ftsession_subject ON public.facttransactionsession (subjectsk);

CREATE INDEX idx_ftsession_property ON public.facttransactionsession (propertysk);