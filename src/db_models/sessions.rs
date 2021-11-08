use crate::schema::session;
use chrono::NaiveDateTime;

#[derive(Queryable)]
pub struct AuthenticatingUser {
    pub id: i32,
    pub password: String,
}

#[derive(Insertable)]
#[table_name = "session"]
pub struct NewSession<'s> {
    pub token: &'s str,
    pub expires_at: Option<NaiveDateTime>,
    pub user_id: i32,
}
