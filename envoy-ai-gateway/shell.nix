{
  pkgs ? import <nixpkgs> {}
}:

pkgs.mkShell {
  buildInputs = [
    pkgs.kind
    pkgs.kubectl
    pkgs.kubernetes-helm
  ];
}
