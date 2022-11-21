# Nix

## Introduction

<https://nixos.org/manual/nix/stable/introduction.html>

Nixは関数型のパッケージマネージャー
すなわち、関数型言語の値のようにパッケージを取り扱う。副作用は無いし、一度ビルドされると変化しない。

NixはパッケージをNixストア内に保持する。ストアは/nix/store配下にある。下記に確認できるようにサブディレクトリはユニークである。

`/nix/store/b6gvzjyb2pg0kjfwrjmg1vfhh54ad73z-firefox-33.1/`

b6gvzjyb2pg0…はすべての依存関係を記録するそのパッケージのためのユニークなIDである。
（パッケージのビルドの依存関係グラフの暗号的ハッシュ値）

### 複数バージョン

インストール済みの一パッケージの複数バージョンやvariantを同時に持つことが出来ます。
同じパッケージに別バージョンに依存する異なるアプリケーションがある場合に特に重要。
ストア配下のパス名の設計より自明。"DLL地獄"を防ぐ。

### Complete dependencies

Nix helps you make sure that package dependency specifications are complete. In general, when you’re making a package for a package management system like RPM, you have to specify for each package what its dependencies are, but there are no guarantees that this specification is complete. If you forget a dependency, then the package will build and work correctly on your machine if you have the dependency installed, but not on the end user's machine if it's not there.

Since Nix on the other hand doesn’t install packages in “global” locations like /usr/bin but in package-specific directories, the risk of incomplete dependencies is greatly reduced. This is because tools such as compilers don’t search in per-packages directories such as /nix/store/5lbfaxb722zp…-openssl-0.9.8d/include, so if a package builds correctly on your system, this is because you specified the dependency explicitly. This takes care of the build-time dependencies.

Once a package is built, runtime dependencies are found by scanning binaries for the hash parts of Nix store paths (such as r8vvq9kq…). This sounds risky, but it works extremely well.

### Mulit-user support

権限の無いユーザでも安全にインストールができる。
それぞれのユーザが異なるプロファイル（ユーザPATHに現れるNixストアのパッケージリスト）を持つ。
すでにインストール済みのパッケージは二度とビルド、ダウンロードされない。ユーザによるパッケージへのトロイの木馬の注入もできない。

### Atomic upgrades and rollbacks

パッケージは上書きされない。アップグレード後も残る。すなわち古いバージョンに上書きできる。

```shell
$ nix-env --upgrade -A nixpkgs.some-package
$ nix-env --rollback
```

### Garbage collection

アンインストールしても、

```shell
$ nix-env --uninstall firefox
```

パッケージは消えない（ロールバックしたくなるかもだし、別ユーザのプロファイルで参照しているかもしれないし。）。
代わりにgarbage collectorを動かして安全に削除することができる。

```shell
$ nix-collect-garbage
```

動作中のプログラムがない、いずれのユーザプロファイルからの参照が無い場合に実行できる。

### Functional package language

Packages are built from Nix expressions, which is a simple functional language. A Nix expression describes everything that goes into a package build action (a “derivation”): other packages, sources, the build script, environment variables for the build script, etc. Nix tries very hard to ensure that Nix expressions are deterministic: building a Nix expression twice should yield the same result.

Because it’s a functional language, it’s easy to support building variants of a package: turn the Nix expression into a function and call it any number of times with the appropriate arguments. Due to the hashing scheme, variants don’t conflict with each other in the Nix store.

### Transparent source/binary deployment

Nix expressions generally describe how to build a package from source, so an installation action like


$ nix-env --install -A nixpkgs.firefox
could cause quite a bit of build activity, as not only Firefox but also all its dependencies (all the way up to the C library and the compiler) would have to be built, at least if they are not already in the Nix store. This is a source deployment model. For most users, building from source is not very pleasant as it takes far too long. However, Nix can automatically skip building from source and instead use a binary cache, a web server that provides pre-built binaries. For instance, when asked to build /nix/store/b6gvzjyb2pg0…-firefox-33.1 from source, Nix would first check if the file https://cache.nixos.org/b6gvzjyb2pg0….narinfo exists, and if so, fetch the pre-built binary referenced from there; otherwise, it would fall back to building from source.

### Nix Packages collection

We provide a large set of Nix expressions containing hundreds of existing Unix packages, the Nix Packages collection (Nixpkgs).

### Managing build environments

Nix is extremely useful for developers as it makes it easy to automatically set up the build environment for a package. Given a Nix expression that describes the dependencies of your package, the command nix-shell will build or download those dependencies if they’re not already in your Nix store, and then start a Bash shell in which all necessary environment variables (such as compiler search paths) are set.

For example, the following command gets all dependencies of the Pan newsreader, as described by its Nix expression:


$ nix-shell '<nixpkgs>' -A pan
You’re then dropped into a shell where you can edit, build and test the package:


[nix-shell]$ unpackPhase
[nix-shell]$ cd pan-*
[nix-shell]$ configurePhase
[nix-shell]$ buildPhase
[nix-shell]$ ./pan/gui/pan
Portability
Nix runs on Linux and macOS.

### NixOS

NixOS is a Linux distribution based on Nix. It uses Nix not just for package management but also to manage the system configuration (e.g., to build configuration files in /etc). This means, among other things, that it is easy to roll back the entire configuration of the system to an earlier state. Also, users can install software without root privileges. For more information and downloads, see the NixOS homepage.

### License

Nix is released under the terms of the GNU LGPLv2.1 or (at your option) any later version.


## Installation

* Binary distribution
  * Single-user
  * Multiple-user
* From source
