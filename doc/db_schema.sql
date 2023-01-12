CREATE TABLE `daily_price` (
    `code`   CHAR(8)        NOT NULL,
    `date`   DATE           NOT NULL,
    `open`   DECIMAL(20, 3) NOT NULL,
    `high`   DECIMAL(20, 3) NOT NULL,
    `low`    DECIMAL(20, 3) NOT NULL,
    `close`  DECIMAL(20, 3) NOT NULL,
    `volume` BIGINT(20) NOT NULL,
    PRIMARY KEY (`code`, `date`)
);

CREATE TABLE `amount_by_participants` (
    `code`       CHAR(8) NOT NULL,
    `date`       DATE    NOT NULL,
    `institution`   DOUBLE  NOT NULL,
    `foreigner`  DOUBLE  NOT NULL,
    `individual` DOUBLE  NOT NULL,
    PRIMARY KEY (`code`, `date`)
);