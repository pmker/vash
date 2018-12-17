-- +migrate Up
CREATE TABLE IF NOT EXISTS data_changes_archive (
    seq int(20) primary key,
    node_id varchar(255),
    type enum('create', 'delete', 'path', 'content'),
    source text,
    target text
);

-- +migrate Down
DROP TABLE data_changes_archive;
