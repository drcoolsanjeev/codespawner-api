use crate::schema::users_code;
use diesel::pg::types::sql_types::Uuid;
use diesel::sql_types::Timestamptz;
use diesel::sql_types::Nullable;
#[derive(Queryable)]
pub struct UsersCode {
    pub id: Uuid,
    pub user_id : Nullable<Uuid>,
    pub code_buffer: String,
    pub input_buffer: Nullable<String>,
    pub ts : Nullable<Timestamptz>,
    pub ts_mod: Nullable<Timestamptz>,
}

#[derive(Insertable)]
#[table_name = "users_code"]
pub struct NewCode<'a> {
    pub user_id: Nullable<Uuid>,
    pub code_buffer: &'a str,
    pub input_buffer: Nullable<&'a str>,
}
