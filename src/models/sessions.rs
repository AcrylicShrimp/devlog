use rocket::serde::Deserialize;
use serde::Serialize;

#[derive(Debug, Deserialize)]
pub struct NewSessionParam {
    pub username: String,
    pub password: String,
}

#[derive(Debug, Serialize)]
pub struct CreatedSession {
    pub token: String,
    pub expiry: Option<u64>,
}
