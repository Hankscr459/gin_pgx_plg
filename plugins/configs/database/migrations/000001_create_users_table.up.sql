create table public.user
(
    id            serial      primary key,
    name    varchar(150) not null,
    date_of_birth date,  
    created_at    date,
    description   varchar(150)
);