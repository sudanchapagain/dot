mod cli;
mod config;
mod fsops;
mod state;
mod util;
use crate::{cli::Cli, config::parse_mappings, fsops::*, state::*};

use std::process;

use clap::Parser;

fn main() -> anyhow::Result<()> {
    let cli = Cli::parse();
    let mappings = parse_mappings()?;

    match cli.command {
        Some(cli::Command::Link) => link_files(&mappings, cli.force, cli.dry_run)?,
        Some(cli::Command::Remove) => {
            let state = load_state();
            remove_links(&state)?;
        }
        Some(cli::Command::Status) => status(&mappings)?,
        None => {
            eprintln!("usage: dotfiles [options] [command]");
            process::exit(1);
        }
    }

    Ok(())
}

