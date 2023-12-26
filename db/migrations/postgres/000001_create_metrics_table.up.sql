CREATE TABLE metrics (
    id              VARCHAR     PRIMARY KEY,
    user_id         VARCHAR     NOT NULL,
    event_id        INTEGER     NOT NULL,
    event_name      VARCHAR     NOT NULL,
    layout_id       INTEGER     NOT NULL,
    layout_name     VARCHAR     NOT NULL,
    created_date    timestamp   NOT NULL
);

CREATE INDEX newest_metrics on metrics (created_date);