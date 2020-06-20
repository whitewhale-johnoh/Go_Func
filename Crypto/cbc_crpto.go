package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "fmt"
    "io"
)

func encrypt(b cipher.Block, plaintext []byte) []byte {
    if mod := len(plaintext) % aes.BlockSize; mod != 0 {
        padding := make([]byte, aes.BlockSize-mod)
        plaintext = append(plaintext, padding...)
    }
    
    ciphertext := make([]byte, aes.BlockSize+len(plaintext))
    
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        fmt.Println(err)
        return nil
    }
    
    mode := cipher.NewCBCEncrypter(b, iv)
    mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)
    return ciphertext
}

func decrypt(b cipher.Block, ciphertext []byte) []byte {
    if len(ciphertext)%aes.BlockSize != 0 {
        fmt.Println("암호화된 데이터의 길이는 블록 크기의 배수가 되어야 합니다.")
        return nil
    }
    
    iv := ciphertext[:aes.BlockSize]
    ciphertext = ciphertext[aes.BlockSize:]
    plaintext := make([]byte, len(ciphertext))
    mode := cipher.NewCBCDecrypter(b, iv)
    mode.CryptBlocks(plaintext,ciphertext)
    return plaintext
}

func main() {
    key := "Hello Key 123456"
    s := `동해 물과 백두산이 마르고 닳도록 하느님이 보우하사 우리나라 만세. 무궁화 삼천리 화려강산 대한사람, 대한으로 길이 보전하세.`
    
    block, err := aes.NewCipher([]byte(key))
    if err != nil {
        fmt.Println(err)
        return
    }
    
    ciphertext := encrypt(block, []byte(s))
    fmt.Printf("%x\n", ciphertext)
    plaintext := decrypt(block, ciphertext)
    fmt.Println(string(plaintext))
}