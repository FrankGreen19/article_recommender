create table articles
(
    id      serial primary key,
    title   text not null,
    content text
);

CREATE INDEX idx_articles_title_gin ON articles USING gin (to_tsvector('russian', title));