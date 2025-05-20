use clap::{Parser, Subcommand};

#[derive(Parser)]
#[command(version, about = "dotfiles - Manage dotfiles with symlinks")]
pub struct Cli {
    #[arg(short, long, global = true)]
    pub force: bool,

    #[arg(short = 'd', long, global = true)]
    pub dry_run: bool,

    #[command(subcommand)]
    pub command: Option<Command>,
}

#[derive(Subcommand)]
pub enum Command {
    Link,
    Remove,
    Status,
}

