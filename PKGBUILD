# Maintainer: Pramananda <pramanandasarkar02@gmail.com>
pkgname=leetcode-local-cli
pkgver=0.1.0
pkgrel=1
pkgdesc="A command-line tool for interacting with LeetCode locally"
url="https://github.com/pramanandasarkar02/leetcode-local-cli"
arch=('x86_64' 'aarch64')
license=('MIT') # Adjust based on your LICENSE file
makedepends=('go')
source=("$pkgname-$pkgver.tar.gz::https://github.com/pramanandasarkar02/leetcode-local-cli/archive/v$pkgver.tar.gz")
sha256sums=('SKIP') # Replace with actual checksum

build() {
    cd "$srcdir/$pkgname-$pkgver"
    export CGO_ENABLED=0
    go build -o $pkgname -ldflags "-s -w" .
}

package() {
    cd "$srcdir/$pkgname-$pkgver"
    install -Dm755 $pkgname "$pkgdir/usr/bin/$pkgname"
    install -Dm644 LICENSE "$pkgdir/usr/share/licenses/$pkgname/LICENSE"
    # Optionally include README or other docs
    install -Dm644 README.md "$pkgdir/usr/share/doc/$pkgname/README.md"
}