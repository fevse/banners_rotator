-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS banner (
    banner_id SERIAL PRIMARY KEY,
    descript TEXT
);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS slot (
    slot_id SERIAL PRIMARY KEY,
    descript TEXT
);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS sgroup (
    group_id SERIAL PRIMARY KEY,
    descript TEXT
);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS storage (
    id SERIAL PRIMARY KEY,
    slot_id INT NOT NULL,
    banner_id INT NOT NULL,
    group_id INT,
    view INT DEFAULT 0,
    click INT DEFAULT 0,
    reward NUMERIC DEFAULT 1,
    FOREIGN KEY (slot_id)  REFERENCES slot (slot_id),
    FOREIGN KEY (banner_id) REFERENCES banner (banner_id),
    FOREIGN KEY (group_id) REFERENCES sgroup (group_id)
);
-- +goose StatementEnd
-- +goose StatementBegin
INSERT INTO banner (descript)
VALUES ('banner1'), ('banner2');
-- +goose StatementEnd
-- +goose StatementBegin
INSERT INTO sgroup (descript)
VALUES ('group1'), ('group2');
-- +goose StatementEnd
-- +goose StatementBegin
INSERT INTO slot (descript)
VALUES ('slot1'), ('slot2');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS storage;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS banner;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS sgroup;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS slot;
-- +goose StatementEnd


