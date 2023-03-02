-- default procedures
-- Update trigger to compare only specific values like (new.email != old.email ...)
create or replace function modified_at_refresh() returns trigger as
$$
begin
    if (new != old) then
        new.modified_at = now();
        return new;
    else
        return old;
    end if;
end;
$$ language plpgsql;

-- main tables
-- Effectively logins should be in a separate table, as user could login using different devices in many cases
create table users
(
    id              uuid primary key,
    username        text unique not null,
    hashed_password text not null,
    email           text unique,
    is_active       boolean default true not null,
    created_at      timestamp not null default current_timestamp,
    modified_at     timestamp,
    last_login_at   timestamp,
    session_id      uuid
);

create trigger users_modified_at
    before update
    on users
    for each row
execute procedure modified_at_refresh();

create table user_pages
(
    user_id    uuid not null,
    url        text not null,
    created_at timestamp not null default current_timestamp,
    foreign key (user_id) references users (id) on delete cascade
);

create unique index user_id_url_idx on user_pages (user_id, url);

-- Seed DB with data for the demo
insert into users (id, username, hashed_password)
values
       ('d77308d3-a0f7-4ce3-993c-00502754789c', 'alex', '0ccb57fb62af497bbdcdcf9d9b12f99a15b3802dcb9f59cc6718e5894d820f9f'),
       ('a3e5b3ab-4c80-4b9d-8ae5-7a1732497778', 'mark', 'bc99405d7f53a6d5c6942cf7bea8edd4fe80d508562a01003b20f1dfd062ca64'),
       ('26c8b55a-c95f-4aea-9c1d-6384d687c0ea', 'james', 'f221e2496c5e7209979375c038af30250adafd0bf4ecddc028777a2da012a0a0');
