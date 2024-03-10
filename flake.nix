{
  description = "Yet another neofetch clone";

  inputs = {
    nixpkgs.url = "nixpkgs/nixos-unstable";

    utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, utils }:
    utils.lib.eachSystem [
      "x86_64-linux"
      "aarch64-linux"
    ] (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in {
        packages.default = pkgs.buildGoModule {
          pname = "fetch";
          version = "latest";
          src = ./.;

          vendorHash =
            "sha256-uTRLJlrUPGdPeIX7vDuv86IqMRvutVwJCuUS/KLdpXs=";

          CGO_ENABLED = 0;

          # removes debug info, making the binary smaller
          ldflags = [ "-s" "-w" ];
        };

        apps.default = utils.lib.mkApp { drv = self.packages.${system}.default; };

        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [ go ];
        };

      });

}
