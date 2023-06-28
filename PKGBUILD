pkgname=neogo
pkgver=1.0.1
pkgrel=1
pkgdesc="neofetch but in golang"
arch=('x86_64')
url="https://github.com/CheemsBread505/neogo"
license=('GPL v3.0')
depends=('go')

source=("$pkgname-$pkgver.tar.gz::https://github.com/CheemsBread505/neogo/archive/refs/tags/v$pkgver.tar.gz")
sha256sums=('f4afc52798037376a8c2015cecc1c1d1aaa0cf87be3f7b5e11a9df3194e4f116')

build() {
    cd "$srcdir/$pkgname-$pkgver"
    go build -o "$pkgname"
}

package() {
    install -Dm755 "$srcdir/$pkgname-$pkgver/$pkgname" "$pkgdir/usr/bin/$pkgname"
}
