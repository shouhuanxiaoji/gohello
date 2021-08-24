package main

import (
    "errors"
    "fmt" 
) 

// Hi returns a friendly greeting in language lang
func Hi(name, lang string) (string, error) {
    switch lang {
    case "zh_CN":
        return fmt.Sprintf("你好, %s!", name), nil
    case "en":
        return fmt.Sprintf("Oi, %s!", name), nil
    default:
        return "", errors.New("unknown language")
    }
}