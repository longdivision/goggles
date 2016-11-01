PACKAGE=github.com/longdivision/goggles
RELEASE_DIR=release
TMP_DIR=tmp

function build {
  export GOOS=$1
  export GOARCH=$2
  BUILD_VARIANT=$GOOS-$GOARCH

  echo == Building $PACKAGE for $BUILD_VARIANT

  mkdir $TMP_DIR && pushd $TMP_DIR
  go build -v $PACKAGE
  zip -r ../$BUILD_VARIANT.zip *
  popd
  rm -rf $TMP_DIR
}

rm -rf $RELEASE_DIR
mkdir $RELEASE_DIR && pushd $RELEASE_DIR

build darwin amd64
build linux 386
build linux amd64
build windows 386
build windows amd64
