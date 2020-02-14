table! {
    users (id) {
        id -> Uuid,
        #[sql_name = "type"]
        type_ -> Int4,
        name -> Text,
        email -> Text,
        ts -> Nullable<Timestamptz>,
        password -> Nullable<Text>,
        ts_mod -> Nullable<Timestamptz>,
    }
}

table! {
    users_code (id) {
        id -> Uuid,
        user_id -> Nullable<Uuid>,
        code_buffer -> Text,
        input_buffer -> Nullable<Text>,
        ts -> Nullable<Timestamptz>,
        ts_mod -> Nullable<Timestamptz>,
    }
}

joinable!(users_code -> users (user_id));

allow_tables_to_appear_in_same_query!(
    users,
    users_code,
);
