CREATE USER user_schema_content WITH LOGIN PASSWORD 'pass_user_schema_content';

CREATE USER user_schema_action  WITH LOGIN PASSWORD 'pass_user_schema_action';

CREATE USER user_schema_user    WITH LOGIN PASSWORD 'pass_user_schema_user';

GRANT USAGE ON SCHEMA user_schema TO user_schema_user;

GRANT USAGE ON SCHEMA user_action_schema TO user_schema_action;

GRANT USAGE ON SCHEMA content_schema TO user_schema_content;

GRANT SELECT ON 
    content_schema.roles, 
    content_schema.persons, 
    content_schema.content, 
    content_schema.films, 
    content_schema.selections, 
    content_schema.content_selections,
    content_schema.content_roles_persons,
    content_schema.countries,
    content_schema.genres,
    content_schema.content_countries,
    content_schema.content_genres,
    content_schema.series,
    content_schema.episodes
    TO user_schema_content;

GRANT SELECT, UPDATE, INSERT ON 
    user_schema.users
    TO user_schema_user;

GRANT SELECT, UPDATE, INSERT, DELETE ON 
    user_action_schema.users_content_favorites,
    user_action_schema.ratings,
    user_action_schema.history_views
    TO user_schema_action;
