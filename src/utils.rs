use rand::distributions::Alphanumeric;
use rand::prelude::*;

pub fn gen_token_256() -> String {
    thread_rng()
        .sample_iter(&Alphanumeric)
        .take(256)
        .map(char::from)
        .collect()
}
