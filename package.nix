{buildGoApplication, ...}:

buildGoApplication {
  pname = "nekofetch";
  version = "latest";
  src = ./.;
  modules = ./gomod2nix.toml;

  CGO_ENABLED = 0;

  # removes debug info, making the binary smaller
  ldflags = ["-s" "-w"];
}
