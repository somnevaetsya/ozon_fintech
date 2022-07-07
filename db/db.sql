create
unlogged table if not exists links
(
    original_link varchar not null primary key,
    short_link    char(10)             not null unique,
    timestamp timestamp NOT NULL DEFAULT NOW(),
);

create or replace function delete_old_links()
    returns trigger as
$$
begin
    delete from links where timestamp < NOW() - interval '10 minute';
return new;
end;
$$ language plpgsql;

create trigger delete_old_links_trigger
    after insert
    on links
    execute procedure delete_old_links();

create index if not exists short_link_index on links using hash (short_link);
