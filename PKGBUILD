pkgname=neogo
pkgver=1.0.1
pkgrel=1
pkgdesc="neofetch but in golang"
arch=('x86_64')
url="https://github.com/CheemsBread505/neogo"
license=('GPL v3.0')
depends=('go')

source=("$pkgname-$pkgver.tar.gz::https://github.com/CheemsBread505/neogo/archive/refs/tags/v$pkgver.tar.gz")
sha256sums=('v1.0.1')

build() {
    cd "$srcdir/$pkgname-$pkgver"
    go build -o "$pkgname"
}

package() {
    install -Dm755 "$srcdir/$pkgname-$pkgver/$pkgname" "$pkgdir/usr/bin/$pkgname"
}
