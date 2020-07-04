----------------------------------------------------------------- admins -------------------------------------------------
DROP TABLE IF EXISTS admins;
DROP TABLE IF EXISTS admins_operations;
CREATE EXTENSION IF NOT EXISTS CITEXT;


CREATE TABLE IF NOT EXISTS admins
(
    admin_id SERIAL PRIMARY KEY,
    username CITEXT COLLATE "C" NOT NULL UNIQUE,
    password BYTEA              NOT NULL
);
-- added full index for faster scanning --
CREATE INDEX admins_username_search ON admins USING hash (username, password, admin_id);
CLUSTER admins USING admins_username_search;

-- this table was created for loggining all admins actions --
-- operation operator is user who provides such operation --
-- operation_columns is json array of affected columns --
-- operation_old_data is old data of columns ordered by columns in operation_columns --
-- operation_new_data is new data of columns ordered by columns in operation_columns --

-- if we are operating with images like ApplePass.Icon just log column which was updated without any old//new _data --
-- i guess we won't have any problems with loggining jsonb objects --
CREATE TABLE IF NOT EXISTS admins_operations
(
    operation_id       BIGSERIAL PRIMARY KEY,
    operation_operator CITEXT COLLATE "C" NOT NULL REFERENCES admins (username),
    operation_time     TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    operation_table    CITEXT COLLATE "C" NOT NULL,
    operation_columns  JSONB                    DEFAULT NULL,
    operation_old_data JSONB                    DEFAULT NULL,
    operation_new_data JSONB                    DEFAULT NULL
);

CREATE INDEX admins_opeartion_operator_table_search_with_time_data ON admins_operations USING HASH (operation_operator, operation_table, operation_time);
CREATE INDEX admins_opeartion_table_operator_search_with_time_data ON admins_operations USING HASH (operation_table, operation_operator, operation_time);
CREATE INDEX admins_all_data ON admins_operations (operation_operator, operation_time, operation_table,
                                                   operation_columns, operation_old_data, operation_new_data);
CLUSTER admins_operations USING admins_opeartion_operator_table_search_with_time_data;