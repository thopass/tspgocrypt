package tspgocrypt

const VER_MAJOR = 0
const VER_MINOR = 1

func GetVersion() (uint,uint) {
  return VER_MAJOR,VER_MINOR
}

type TSPCrypter interface {
  Encrypt(data []byte, key []byte, counter uint) int
  Decrypt(data []byte, key []byte, counter uint) int

  GetName() string
}

func GetXORCrypter() TSPCrypter {

  return new(XORCrypter)
}
