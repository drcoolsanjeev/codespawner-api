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

