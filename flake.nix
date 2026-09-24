{
  description = "go";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-26.05-darwin";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = {
    nixpkgs,
    flake-utils,
    ...
  }:
    flake-utils.lib.eachDefaultSystem (
      system: let
        pkgs = nixpkgs.legacyPackages.${system};

        delve = pkgs.delve.overrideAttrs (old: rec {
          version = "1.27.1";
          hash = "sha256-H91QnLyqywgoc3zdTaclzzUxVPagNnxLzKub2gnL25w=";
          vendorHash = null;
          src = pkgs.fetchFromGitHub {
            owner = "go-delve";
            repo = "delve";
            rev = "v${version}";
            hash = hash;
          };
        });

        golangci-lint =
          (pkgs.golangci-lint.override {
            buildGo126Module = pkgs.buildGo127Module;
          }).overrideAttrs (old: rec {
            version = "2.14.0";

            src = pkgs.fetchFromGitHub {
              owner = "golangci";
              repo = "golangci-lint";
              rev = "v${version}";
              hash = "sha256-HATA7JKHwEouM+8jYZbQrkX7p4gut4IpyTvcBexu/4o=";
            };

            vendorHash = "sha256-ekP/zDhYMpMG+tYAyYHNfLOCt/JkxrXUw1jHUEfsM8k=";
          });
      in {
        devShells.default = pkgs.mkShell {
          packages = [
            delve
            golangci-lint

            pkgs.go
            pkgs.gofumpt
            pkgs.golines
            pkgs.gotestsum
            pkgs.just
            pkgs.sqlc
          ];
        };
      }
    );
}
