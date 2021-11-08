use rocket::serde::Deserialize;

#[derive(Debug, Deserialize)]
pub struct NewSessionParam {
    pub email: String,
    pub password: String,
}
