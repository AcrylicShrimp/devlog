mod db;
mod db_models;
mod models;
mod routes;
mod schema;
mod utils;

#[macro_use]
extern crate diesel;
#[macro_use]
extern crate macros;
#[macro_use]
extern crate rocket;

use db::*;
use rocket::Error as RocketError;
use tracing::{event, instrument, Level};
use tracing_subscriber::{self, EnvFilter};

#[cfg(debug_assertions)]
use dotenv::{dotenv, Error as DotEnvError};

#[derive(Debug)]
enum Error {
    #[cfg(debug_assertions)]
    DotEnvError(DotEnvError),
    DBError(DBError),
    RocketError(RocketError),
}

#[cfg(debug_assertions)]
impl From<DotEnvError> for Error {
    fn from(err: DotEnvError) -> Self {
        Self::DotEnvError(err)
    }
}

impl From<DBError> for Error {
    fn from(err: DBError) -> Self {
        Self::DBError(err)
    }
}

impl From<RocketError> for Error {
    fn from(err: RocketError) -> Self {
        Self::RocketError(err)
    }
}

#[instrument]
#[rocket::main]
async fn main() -> Result<(), Error> {
    #[cfg(debug_assertions)]
    dotenv()?;

    tracing_subscriber::fmt()
        .with_env_filter(EnvFilter::from_env("LOG_LEVEL"))
        .init();

    rocket::build()
        .manage(db_connection_pool().map_err(|err| {
            event!(Level::ERROR, msg = "failed to connect to database", ?err);
            err
        })?)
        .mount("/v1/", routes![routes::new_session])
        .ignite()
        .await?
        .launch()
        .await
        .map_err(|err| {
            event!(
                Level::ERROR,
                msg = "error occurred while running the server",
                ?err
            );
            err
        })?;

    Ok(())
}
