use dirs::home_dir;
use std::path::PathBuf;

pub fn dotfiles_dir() -> PathBuf {
    home_dir().unwrap().join(".dotfiles")
}

pub fn expand_user_path(path: &str) -> PathBuf {
    if let Some(stripped) = path.strip_prefix('~') {
        home_dir().unwrap().join(stripped.trim_start_matches('/'))
    } else {
        PathBuf::from(path)
    }
}
