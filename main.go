/**
 * Copyright 2023 Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
**/
package main

import (
    "image"
    "image/color"
    "image/draw"
    "image/png"
    "net/http"
)

func main() {
    http.HandleFunc("/blue", blueHandler)
    http.HandleFunc("/red", redHandler)  // added red endpoint
    http.ListenAndServe(":8080", nil)
}

func blueHandler(w http.ResponseWriter, r *http.Request) {
    img := image.NewRGBA(image.Rect(0, 0, 100, 100))
    draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{0, 0, 255, 255}}, image.ZP, draw.Src)
    w.Header().Set("Content-Type", "image/png")
    png.Encode(w, img)
}

// new redHandler function
func redHandler(w http.ResponseWriter, r *http.Request) {
    img := image.NewRGBA(image.Rect(0, 0, 100, 100))
    draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{255, 0, 0, 255}}, image.ZP, draw.Src)
    w.Header().Set("Content-Type", "image/png")
    png.Encode(w, img)
}
