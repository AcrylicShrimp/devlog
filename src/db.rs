use diesel::prelude::*;
use diesel::r2d2::{ConnectionManager, Pool, PoolError};
use std::env::{var as env_var, VarError};

#[derive(Debug)]
pub enum DBError {
    EnvVarError(String, VarError),
    PoolError(PoolError),
}

impl From<PoolError> for DBError {
    fn from(err: PoolError) -> Self {
        Self::PoolError(err)
    }
}

pub fn db_connection_pool() -> Result<Pool<ConnectionManager<PgConnection>>, DBError> {
    let database_url = env_var("DATABASE_URL")
        .map_err(|err| DBError::EnvVarError("DATABASE_URL".to_string(), err))?;
    let manager = ConnectionManager::<PgConnection>::new(database_url);
    Ok(Pool::builder()
        .min_idle(Some(4))
        .max_size(16)
        .build(manager)?)
}
