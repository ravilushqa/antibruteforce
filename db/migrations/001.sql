create table blacklist
(
    id     serial not null,
    subnet cidr   not null
);

create unique index blacklist_id_uindex
    on blacklist (id);

create unique index blacklist_subnet_uindex
    on blacklist (subnet);


create table whitelist
(
    id     serial  not null,
    subnet cidr not null
);

create unique index whitelist_id_uindex
    on whitelist (id);

create unique index whitelist_subnet_uindex
    on whitelist (subnet);