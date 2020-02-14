#![allow(proc_macro_derive_resolution_fallback)]
use super::schema::*;
use uuid::Uuid;
use serde::{Serialize, Deserialize};
use chrono::NaiveDateTime;

#[derive(Debug, Serialize, Deserialize, Queryable)]
pub struct UsersCode {
    pub id: Uuid,
    pub user_id : Option<Uuid>,
    pub code_buffer: String,
    pub input_buffer: Option<String>,
    pub ts : Option<NaiveDateTime>,
    pub ts_mod: Option<NaiveDateTime>,
}
