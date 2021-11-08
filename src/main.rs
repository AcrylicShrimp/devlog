mod db;
mod models;
mod routes;
mod schema;

#[macro_use]
extern crate diesel;
#[macro_use]
extern crate rocket;

use db::*;
use rocket::Error as RocketError;

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

#[rocket::main]
async fn main() -> Result<(), Error> {
    #[cfg(debug_assertions)]
    dotenv()?;

    rocket::build()
        .manage(db_connection_pool()?)
        .mount("/", routes![routes::index])
        .ignite()
        .await?
        .launch()
        .await?;

    Ok(())
}
