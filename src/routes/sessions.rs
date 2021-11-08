use crate::db_models::*;
use crate::models::{CreatedSession, NewSessionParam};
use crate::schema::session::table as db_session_table;
use crate::schema::user::dsl as db_user;
use crate::utils::gen_token_256;
use argon2::verify_encoded;
use diesel::prelude::*;
use diesel::r2d2::{ConnectionManager, Pool};
use diesel::PgConnection;
use rocket::http::Status;
use rocket::serde::json::Json;
use rocket::State;
use tracing::{event, instrument, Level};

#[instrument(skip(db))]
#[post("/sessions", data = "<param>")]
pub async fn new_session(
    db: &State<Pool<ConnectionManager<PgConnection>>>,
    param: Json<NewSessionParam>,
) -> Result<Json<CreatedSession>, Status> {
    let conn = use_db!();

    let user = query_if_not_found!(
        db_user::user
        .filter(db_user::username.eq(&param.username))
        .select((db_user::id, db_user::password))
        .first::<AuthenticatingUser>(&conn) => {
        return Err(Status::Unauthorized);
    });

    if !return_500_if_err!(
        "verifying password",
        verify_encoded(&user.password, param.password.as_bytes())
    ) {
        return Err(Status::Unauthorized);
    }

    let token = gen_token_256();
    let new_session = NewSession {
        token: &token,
        expires_at: None,
        user_id: user.id,
    };

    query!(diesel::insert_into(db_session_table)
        .values(&new_session)
        .execute(&conn));

    event!(
        Level::INFO,
        msg = "a new session has been created",
        username = param.username.as_str()
    );

    Ok(Json(CreatedSession {
        token,
        expiry: None,
    }))
}
