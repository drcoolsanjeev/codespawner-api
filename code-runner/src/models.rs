use schema::users_code;

#[derive(Queryable)]
pub struct UsersCode {
    pub id: String,
    pub user_id : String,
    pub code_buffer: String,
    pub input_buffer: String,
    pub ts : String,
    pub ts_mode: String,
}

#[derive(Insertable)]
#[table_name="users_code"]
pub struct NewCode<'a> {
    pub user_id: &'a str,
    pub code_buffer: &'a str,
    pub input_buffer: &'a str,
}