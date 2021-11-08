use crate::diesel::{QueryDsl, RunQueryDsl};
use crate::models::NewSessionParam;
use crate::schema::user::dsl::*;
use diesel::r2d2::{ConnectionManager, Pool};
use diesel::PgConnection;
use rocket::http::Status;
use rocket::serde::json::Json;
use rocket::State;

#[get("/")]
pub async fn index(db: &State<Pool<ConnectionManager<PgConnection>>>) -> Result<String, Status> {
    let conn = db.get().map_err(|_| Status::InternalServerError)?;
    let count = user
        .count()
        .get_result::<i64>(&conn)
        .map_err(|_| Status::InternalServerError)?;

    Ok(format!("user count: {}", count))
}

#[post("/v1/sessions")]
pub async fn newSession(
    db: &State<Pool<ConnectionManager<PgConnection>>>,
    param: Json<NewSessionParam>,
) -> Result<String, Status> {
    let conn = db.get().map_err(|_| Status::InternalServerError)?;
    let count = user
        .count()
        .get_result::<i64>(&conn)
        .map_err(|_| Status::InternalServerError)?;

    Ok(format!("user count: {}", count))
}
