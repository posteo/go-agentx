{ pkgs ? import <nixpkgs> {} }:
pkgs.mkShell {
    name = "dev-environment";
    buildInputs = [
        pkgs.net-snmp
    ];
}
