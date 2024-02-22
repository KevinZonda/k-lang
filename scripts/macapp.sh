cd src/out

rm -fr IDLE.app
mkdir -p IDLE.app/Contents/MacOS
cp ../../scripts/macPack/Info.plist IDLE.app/Contents
cp idle IDLE.app/Contents/MacOS