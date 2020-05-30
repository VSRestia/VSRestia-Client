package encryptor

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

//The length of key must be 16 24 or 32
//The complexity of these three mode is incremental
//Complexity increases burden of your device
//
//
//Examples:
//
//
//func main() {
//	origData := []byte("Hello World") // Data to be encrypted
//	key := []byte("ABCDEFGHIJKLMNOP") // Encrypt key
//	log.Println("Original:", string(origData))
//
//	log.Println("------------------ CBC --------------------")
//	encrypted := AesEncryptCBC(origData, key)
//	log.Println("cipher(hex):", hex.EncodeToString(encrypted))
//	log.Println("cipher(base64):", base64.StdEncoding.EncodeToString(encrypted))
//	decrypted := AesDecryptCBC(encrypted, key)
//	log.Println("Decryption results:", string(decrypted))
//
//	log.Println("------------------ ECB --------------------")
//	encrypted = AesEncryptECB(origData, key)
//	log.Println("cipher(hex):", hex.EncodeToString(encrypted))
//	log.Println("cipher(base64):", base64.StdEncoding.EncodeToString(encrypted))
//	decrypted = AesDecryptECB(encrypted, key)
//	log.Println("Decryption results:", string(decrypted))
//
//	log.Println("------------------ CFB --------------------")
//	encrypted = AesEncryptCFB(origData, key)
//	log.Println("cipher(hex):", hex.EncodeToString(encrypted))
//	log.Println("cipher(base64):", base64.StdEncoding.EncodeToString(encrypted))
//	decrypted = AesDecryptCFB(encrypted, key)
//	log.Println("Decryption results:", string(decrypted))
//}

// =================== CBC ======================
func AesEncryptCBC(origData []byte, key []byte) (encrypted []byte) {
	// Newcipher this function limits the length of the input k to 16, 24, or 32
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()                              // Get the length of the key block
	origData = pkcs5Padding(origData, blockSize)                // Complement code
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) // Encryption mode
	encrypted = make([]byte, len(origData))                     // Create array
	blockMode.CryptBlocks(encrypted, origData)                  // Start encrypt
	return encrypted
}
func AesDecryptCBC(encrypted []byte, key []byte) (decrypted []byte) {
	block, _ := aes.NewCipher(key)                              // Group the key
	blockSize := block.BlockSize()                              // Get the length of the key block
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) // Encryption mode
	decrypted = make([]byte, len(encrypted))                    // Create array
	blockMode.CryptBlocks(decrypted, encrypted)                 // Start decrypt
	decrypted = pkcs5UnPadding(decrypted)                       // Remove complement code
	return decrypted
}
func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// =================== ECB ======================
func AesEncryptECB(origData []byte, key []byte) (encrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}
func AesDecryptECB(encrypted []byte, key []byte) (decrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	decrypted = make([]byte, len(encrypted))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim]
}
func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

// =================== CFB ======================
func AesEncryptCFB(origData []byte, key []byte) (encrypted []byte) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	encrypted = make([]byte, aes.BlockSize+len(origData))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], origData)
	return encrypted
}
func AesDecryptCFB(encrypted []byte, key []byte) (decrypted []byte) {
	block, _ := aes.NewCipher(key)
	if len(encrypted) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encrypted, encrypted)
	return encrypted
}
