{
  description = "go";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
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
      in {
        devShells.default = pkgs.mkShell {
          packages = [
            pkgs.delve
            pkgs.go
            pkgs.gofumpt
            pkgs.golangci-lint
            pkgs.golines
            pkgs.gotestsum
            pkgs.just
            pkgs.sqlc
          ];
        };
      }
    );
}
