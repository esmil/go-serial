# Contributor: Esmil <esmil@mailme.dk>

pkgname=go-serial
pkgver=$(date +%Y%m%d)
pkgrel=1
pkgdesc='Go library to use serial connections'
arch=('i686' 'x86_64')
url='http://github.com/esmil/go-serial'
license=('GPL')
depends=('go>=2010_08_25' 'go-termios')
options=(!strip)

build() {
  cd ${srcdir}
  ln -s .. $pkgname
  cd $pkgname

  make
}

package() {
  cd ${srcdir}/$pkgname

  make DESTDIR=${pkgdir} install
}
