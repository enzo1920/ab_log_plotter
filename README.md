# ab_log_inserter

##Create db and table

```
 CREATE DATABASE ab_log_db;

CREATE TABLE light (
    light_id serial not null primary key,
    light_val float  NOT NULL,
    light_date  timestamp default NULL
);
```
