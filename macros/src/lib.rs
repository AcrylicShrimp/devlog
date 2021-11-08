use proc_macro::TokenStream;
use proc_macro_error::*;
use quote::quote;
use syn::parse::{Parse, ParseStream};
use syn::{parse_macro_input, Block, Expr, LitStr, Result as SynResult, Token};

#[proc_macro]
#[proc_macro_error]
pub fn use_db(_: TokenStream) -> TokenStream {
    let expanded = quote! {
        db.get().map_err(|err| {
            tracing::event!(
                tracing::Level::ERROR,
                msg = "failed to get database connection",
                ?err
            );
            return rocket::http::Status::InternalServerError;
        })?
    };

    TokenStream::from(expanded)
}

struct TaggedExpr {
    tag: LitStr,
    expr: Expr,
}

impl Parse for TaggedExpr {
    fn parse(input: ParseStream) -> SynResult<Self> {
        let tag = input.parse()?;
        input.parse::<Token![,]>()?;
        let expr = input.parse()?;

        Ok(Self { tag, expr })
    }
}

#[proc_macro]
#[proc_macro_error]
pub fn return_500_if_err(input: TokenStream) -> TokenStream {
    let TaggedExpr { tag, expr } = parse_macro_input!(input as TaggedExpr);
    let msg = format!("an error occurred while {}", tag.value());

    let expanded = quote! {
        match #expr {
            Ok(result) => result,
            Err(err) => {
                tracing::event!(
                    tracing::Level::ERROR,
                    msg = #msg,
                    ?err,
                    expr = stringify!(#expr),
                );
                return Err(rocket::http::Status::InternalServerError);
            }
        }
    };

    TokenStream::from(expanded)
}

struct Query {
    query: Expr,
}

impl Parse for Query {
    fn parse(input: ParseStream) -> SynResult<Self> {
        let query = input.parse()?;

        Ok(Self { query })
    }
}

#[proc_macro]
#[proc_macro_error]
pub fn query(input: TokenStream) -> TokenStream {
    let Query { query } = parse_macro_input!(input as Query);

    let expanded = quote! {
        match #query {
            Ok(result) => result,
            Err(err) => {
                tracing::event!(
                    tracing::Level::ERROR,
                    msg = "failed to run query",
                    ?err,
                    query = stringify!(#query),
                );
                return Err(rocket::http::Status::InternalServerError);
            }
        }
    };

    TokenStream::from(expanded)
}

struct QueryIfNotFound {
    query: Expr,
    if_not_found: Block,
}

impl Parse for QueryIfNotFound {
    fn parse(input: ParseStream) -> SynResult<Self> {
        let query = input.parse()?;
        input.parse::<Token![=>]>()?;
        let if_not_found = input.parse()?;

        Ok(Self {
            query,
            if_not_found,
        })
    }
}

#[proc_macro]
#[proc_macro_error]
pub fn query_if_not_found(input: TokenStream) -> TokenStream {
    let QueryIfNotFound {
        query,
        if_not_found,
    } = parse_macro_input!(input as QueryIfNotFound);

    let expanded = quote! {
        macros::query!({
            let result = #query;

            if let Err(err) = &result {
                if let diesel::result::Error::NotFound = err {
                    #if_not_found
                }
            }

            result
        })
    };

    TokenStream::from(expanded)
}
