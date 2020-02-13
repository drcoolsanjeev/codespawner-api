#![allow(proc_macro_derive_resolution_fallback)]

use super::schema::*;
use uuid::Uuid;
use serde::{Serialize, Deserialize};
use chrono::NaiveDateTime;

#[derive(Debug, Serialize, Deserialize, Queryable)]
pub struct UsersCode {
    pub id: Uuid,
    pub user_id : Uuid,
    pub code_buffer: String,
    pub input_buffer: String,
    pub ts : NaiveDateTime,
    pub ts_mod: NaiveDateTime,
}

#[derive(Debug, Insertable)]
#[table_name = "users_code"]
pub struct NewCode<'a> {
    pub user_id: &'a str,
    pub code_buffer: &'a str,
    pub input_buffer: &'a str,
}
