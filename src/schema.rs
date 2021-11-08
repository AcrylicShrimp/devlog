table! {
    session (token) {
        token -> Bpchar,
        expires_at -> Nullable<Timestamp>,
        created_at -> Timestamp,
        user_id -> Int4,
    }
}

table! {
    user (id) {
        id -> Int4,
        email -> Varchar,
        username -> Varchar,
        password -> Varchar,
        joined_at -> Timestamp,
    }
}

joinable!(session -> user (user_id));

allow_tables_to_appear_in_same_query!(
    session,
    user,
);
