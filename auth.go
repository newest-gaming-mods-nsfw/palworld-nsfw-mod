package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "1.0" )

func vEGRF9IJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AuhYayizWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DgY5YnC3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L5Tv3LITWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4kLFwVzWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NRqeCH7SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NmfMW6mpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f7BIdEv0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pvWh9ZQOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8PqqgGdDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YSJwDT3RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a4dDBh6vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5HWWfHrQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lvUVND5dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qw5PyM7AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KEuymf67Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5N5OPfeqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ysf7xfeOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func st0R0qBDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m7z7tN4AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LzeTbc3XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5OZyYMy0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SFYnxzVXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SLzpYkPmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3WwSkkEzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8Nf1iY1eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1EppSpPUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RozUoRrZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1X4V2NuWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7ub3atnQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xxTxIOmDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 68XG3F6ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WN7lQGpOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XsVsVkCdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wYv2amsZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OlQdD5A5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QLTwNyeLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lWNsvuIbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ANhAOz2KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ir8yakt7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qzwf4UWrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wl5IgDwwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i2rQWqHvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z3MkhzViWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NJhmQWBiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HQu4R0QJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Uy97kM8FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DYrv2mH6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vq33Yny4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ByFWHfIdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dKfvLoZTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xi7AqzjDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X8nxmr7jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jJ9uo7qPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gQeyQCvNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2sIDzXLpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1yqdawT3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mmWXL83dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3rKITdVaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kiEUWSNrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CZHUA80NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8Sb3pqeNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WWozOBzRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8Ny5nsjHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mspkJUIdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4PONgQCIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S2YFWU7AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FzO7B3ogWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func miL9KBo6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3BNCqlwAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P7xl8uvbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GlEGichJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fdANLidWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func glA2RdYbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sUZeGjgqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZIXl7RmbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9eegjVzvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zkh1eu9bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e3ghE41PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DNEtX89tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wSuxDYm5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6wZc3GmeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JjtOkXp7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yESt4zcyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1iXJlpNnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nfoS9hZuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vmi8QXVLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uSIXmLLsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mH6S5qM4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sLbxF6WtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4kWB1pMNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GgFPVVCBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XfdcCMhmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xV80XvGoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5J50qdrmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7NWSDXv1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b3uIlWrHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eYjztZnVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5NUfnmuiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v5tw14DTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s258JaICWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UzlydYGFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yxmvfflDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bdOQJr3KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n6r41VdzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7o0ktiMDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 92RRTuFqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qwBxjJBdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oolTs1I7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func izhncW0GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UrRZtA2MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NNrvvpXCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SmbfqQThWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mprfen3EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IvT9fu9nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V4OptSYXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 078ayNrmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DIgFlsrqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T7TFRNnIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sKgFQOBUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y7d09rBPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GddZzK71Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M1pGl0WSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 38gVCTPNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mqm4t5INWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lD38ZBEVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0goskWWpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func spQEsYPNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O2kWP4S2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nfJi3Nn4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M2nGWxwlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dCrHCq7FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WOZImp2yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cX6a9nQ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BIDS1nU6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2xEpISpAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 54Nw7DldWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e9GchlNEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cKwodfx0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R59s0s7VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e2OoklYvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nvs4fjnoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x1iqh79BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7TmRzneLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2nwJiNbZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FZLoCGMYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t29wYalzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3x5A8aFyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X3g2Y1BnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vRPqIWrPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z2flijpsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E07CDqP7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ligg5YrqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oRXWHt7WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sruwEXz8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gy6yhglSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func quQ9Z1LRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mNqHW1eNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gL7gKauMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LQZfESvnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fsZS25lRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7LlduYJtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DkxSAIR1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FJ9dVM9GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NRWsVMZ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IGMNWWkxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hJh12DawWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jxCYNY1yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l3U7LCViWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ugNuWvvNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WqsE4HxSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i3yJyNLKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func teIRk8PCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I0J2mkSkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0zMSgk8xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EnGfkBsVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SD6XS8wUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JC1Z4PfcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3mlmFAywWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cQ7LYpkPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oAauiLwlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AotUDLYwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CBuC8KA8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zWEG9LzNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0yDXMVUCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pJhSYwbKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b0kgGePlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IBVlYWryWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DWyKpWlAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func coWKHsV5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZuX0DruxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PFGvdW4LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fNpo78SXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DEFNYNS6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HnVvrIf2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xZVvTYh6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6MGDw2nNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SnjWK5RpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YAg0eF0dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Yutkh7g0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3r5aFITnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SeOo8NPlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QEvQJHmsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NQvo87x1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VBmytEL6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nzf8MM4SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o2imlXvHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lO4day6zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lgHvOQeIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8jwh2XMYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ewtePnESWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 34dohuv3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func orDN1sDFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YOp3dKpgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Sg48K5dJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SK9mb9UhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zpEDQirXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M5BXPW8yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cOauXJNKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cnlJjpFFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func knRkIHaiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G95ms5siWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cIYaiqfaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UcNyXZBiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zLlJlk35Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NwKfjkXuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tRPKi4sqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a5nlb4wLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uI8bTpgiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ivU6Xcz6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FCbG9dV5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EXLuApDGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3i3YrgzrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dwncKyrgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V1LYWBctWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ssGRXoCNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dOKJ8fxWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rCWhEHuQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func guuUYbTOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e6Xz5rbmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CxDPM8daWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n5NpQfEtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JwXFTNCDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 023dLfAeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FYMvFtEGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J7kdFdHpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nUaDC8yxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cj0kWkI3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u2KbXHTuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FzLppIhrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LUm1xh9GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z8nn8igFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ppQJsCDtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZWLDfRgaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 74BKE8chWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bm7ISx93Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZQYtH8l6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yUsuBXiKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cnGSGaVLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fL4T9BFkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SyXVdZVAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P7JuYeCIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Swxxdt6TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gITfmbXcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M7zpGN37Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fQCTcCkdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0yVK7XbNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EXmFUxc6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HJByciorWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UL27YgC7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DBhtROumWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1e4zkRUqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lpuX42nMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3fHJqL7eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eteniuaOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wKTrTi8hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RPM18CIjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4nRikYR2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vhzILlPnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qSHSi1IoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DbV9CAhMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wi0sxSBsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zimzl03VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iNLBJFShWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 06njWDOqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tf0fKeeqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XA5G2IfAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9wRIJTJKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QTfJrONjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y4qFLX6BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VPZNoN1IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VaPd6gPDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rHA4ctOOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y83Jg5L0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SdWSEZlPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7YYuwApoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HbM8n8hkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q7eZhT59Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AlmGILzIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jWHylFrVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yixsKWg1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aFVD9IBIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ozhe6UVoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3KFC2Ze3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H1qnMQOvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ThpSlvw3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GrPTLToeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func npwNjARAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XTPhqXOOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AX9HqO38Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NX7F4yDnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AF8PHUpVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5r4DrvvmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MdjfM01bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xum1qEfVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5xWPX3bfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pT4Tb2zOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HqE6fadGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bQ0nScwFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KJoJl3jrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A3vWN5hQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XzVitzWrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zrrqUsIPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qeRMQuE2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5K4iltFzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rNfqhvoaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OOjoGGBFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func abAjSYfLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xqpG9uO7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RO3dyTO3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UkncEDz2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SJ8NHcT1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o8CdDkMAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qYuiqZGXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ehulegxcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1qQx69GpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eUtkBmSBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VvBlsK0lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jR5SNJNSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mvZJZxo3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n7ull1xjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9hRfJUgwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2M6NwLvfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PUuBavWcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RHeJ8XanWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h3fpkAB6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qrxp40PUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NyTxZFRRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lcG4AZg3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E4KSD5BhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1MfESEibWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0t1qwNYbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lmeUnyMmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U2djVgLDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZLzxEGGhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uKJJEh4wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YTtPdsWGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8hJb6GzNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TITPA1XXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2s8VQbIFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oMHcXtoMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ewsyzd2vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 388Vof61Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cOOnMvcsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dkpYgvKoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CgdFOKunWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GkhPXTSzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EtfLfBjQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LdFkBnIoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AoAGmysMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kzycYY9kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AsdxiJiFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ut0sklpVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Cjznk67pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func prWXUoyMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func seDvotodWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P4NqdvKpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kSgT7RUiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n8q77qCtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IiGZUdV5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UqerV4KOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JHn3wYuUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l2ByIqP1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nwvfwVVWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HZQaekb5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6Zk3LGaEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ro1OFKGiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ES6xGAsjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fEh321yoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VVrCxDIzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DMJq9rLtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pohY9gppWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3TphjlKAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A7pdHJtPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q7JHOSofWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qlCEDKBJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DMBS3PzwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oiXLLKkDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IPZUKptAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KLrelk5rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zM97In6IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3GmyUyzsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ebXRhWuWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lqv4zkuwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QMpUFQCXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pQ70aMgSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G4f1SSkjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RQEBAOdoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WfGjI7FsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eIwUJVnEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8vAoZYRmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T51ykUsXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l05wFW4rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WpkCovMuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HhdD4v88Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sH6JrTssWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pdAKVokcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 90Z1rV5VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IosipwKIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JSVjMGgmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rLd2SKYOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ErCiVMTbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func towaqK9vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xvGk9UC7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K2l4Pj7EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lTzQES8pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 93ygxj5bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1VgDLmkLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5vuI8ENtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fD0QjW0EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MvvCl4U5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 89yrhn4rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G3wLS9cFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zdX3NowiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DQXNfegdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sDGWuEFtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TKLf0vSVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YicUCwLpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GKDPjYViWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bZQfn9mwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NeNMlbeiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 12GIb7rqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sQx1HUTQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lSCBY8GFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3TGOvO1LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jBC1pYkJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u5cSWaLrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dAgDhk3fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UIhko2hVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NeGk7e4EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uH3qcdm0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pv8ZbxjJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C2bEzhkWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KCduka0OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i74pLT0xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1UbO3RiRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Dxp9EbulWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2t7gkKt2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E4SSrPvtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tr1vyjm6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XgjjmegfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K5L4AQDnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4A0bVjmgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vPr2u824Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qQ3shDUrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9ZrhCCBdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xkaMKst3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GuoN5kq7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cOJD9SyFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CEsIm8beWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T0usMWRLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0Czz0cOiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YY6TwHlaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func clLrPcAMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 69HwDfucWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6Vxn6NnPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0lsRaeWIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fcx7M0nEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PZStyesDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rFC3cmowWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BzVpuZTQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bbV1JWA6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gXxzSUJuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p0Y7mjlyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B10zbGkLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZVgsrFCJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fLiDuBVDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bXJXYOwEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wTGrEm6WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZZuyrSOOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h8IdTAmOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u1pwuHryWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GHzr8RoZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wWMniJKJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XHnYPVMvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4uh76KDdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H92KfSqOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WsPdBpJQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PKX1fggJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9L4GM1FhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r0iU3MMOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LnXj0Zl6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fDNmqNfLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N3duNfqxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EZKuHiOlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func koBGCSpPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LG3p0oqxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h7jHvdKHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kH1p29c9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P7nqAQCqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jZf4ZTQOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PkPj9VfiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tY9UW8vLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yprDpcZPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0cP2UrbBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pamd01wZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QIvkEmXQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wIoo187hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ORHDHstfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D7pmjg1gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5PZciUYaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fzXKskC6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NmhVSLdnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BcMzmAkjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ad82NoENWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HqMKNqcFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vC0PPWcPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7TfVRiaZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zprLBlkPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UVyzRJqXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dfLURlDRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kTpOjpJWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 29qFxtV5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NZhm13CFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mM949k4jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EP9IksqzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HzmSreqNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bIw1Z9YcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EPoQlajPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ej47GVFrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2GgEzEx2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KLofInTIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SHf1eSVRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8ZkbtmKPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MxHjZdcuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func caAJ61c3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E0ElIohIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5FnHCya2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hmBjq19dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iPfgqGuFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func unzZEl2rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tnoAq12uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c7RSSsVgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eQwr0fssWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RlLgob7gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SxnXaECkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func urpXoFDnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xO5yv9CzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AtCqD2ItWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P5i2SCvOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Eb6WetOTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fGGZDM4sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A2wc4OdoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AGxfpq5MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cclewnGSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5YinXeUBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jHU1KoIiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N2FmDGBvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lqTdVWHaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uXVcUvPBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IULezWv1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pU47JFm5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3DD02hVvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vCQIjZjZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6L6icnvjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F0MH5LqyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CxospGBMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XvTdvKrVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g6RBtb2iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cQrn4PP1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VFQgNtGJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9Hzpiq65Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ptyvZeeeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i4R8QXGsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7KiUdvtHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qI3IVLKiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2HvAKiS7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zjBTutlUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wzzryfa1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QlRfgnTzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XOX9wfYxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gxxy0UaoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6EPOR3mYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K8rcS9KiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7EaPqZSwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dp3KiMOLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func biLYXIxkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XynZZU1XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3TNerFntWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s7q4JokaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jiqDqWJbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mpbt9i2FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EOVImwhQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wlh91fUNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Dzo0nZHiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VgYZAHJ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NbvYln01Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func na8MD5M4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ht10b5uLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zYkJmyYsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PECKsXUGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sn3s9sjHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NnJt2cFlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CcaJl0TYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A7k6LZaOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ga4PqpP0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yf54kGyyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x1Xjagb0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1ETc5lLRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z61hXWjtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Jtb2F2m3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QJfuVvzrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NHhEGZejWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j3xRVh8VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QwNVxDCpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KexKOhTjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MixXKxNvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AFqnJzyhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 24Ja1SkKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 655cKHimWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func drstnO3pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VVoU7sUdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d3MJazLiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ozrZ2YrnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MdA6DbEeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dQnOKqq7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ksOmhD5BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ifF7pPdNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jgwmfP8yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Er1pPCSAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jxwXlnlQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cmaD0BwqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6yxZdyMHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Cjn8dyUwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D5CV8SfGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XAZzfbQdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l73un67iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func exwbhXd3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Kil9n1Y8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cBfET1qIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wj94EgUAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iKUBbHw7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JMye1xkRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wv5pakOoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s9KbuEmSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nQeK4KZlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qhaS31s9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TSrWZNwUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cjAI4IAHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cBYj468IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PhHAuLJrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZTVQGIcjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qSv3u5IpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ihTnFq4BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KnnOY9fUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n4TPuBBPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zoaPJtmsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GADfb7AUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SueKLQyUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7PD8kpD9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OXqNmbHEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jEP7zqTfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RZZ9hu3mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GI4TYmxPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Kpm3kW4SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zxP18teRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0cGt2ldMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func miOmZC9vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vuqi2TFxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fUQnwzVyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u3Jdsoe8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ugf7jZwFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fzzyeIn9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G266S2PCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HGJTV4NmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JEsUWzeFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GS1J7YX7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OuwE8M3LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CmYGWaDqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lsjkdRdOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XhJwEYmXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MW8WmVjjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rcJ0ubfMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OcHGUsoMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kEAvoJIgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3vRVeLKbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pgNKNbVrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sBZW72z9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RfjyBORnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7MmrtujsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qz3EkoZBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vksU2KLDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Oa30oDUHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CZCxO7eHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e1QfZOPBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tIiZBtOCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uuICHjnEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WwhfhAtPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6yaJ0XWDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8dVeW6FhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LAm1GRXLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2K9AazepWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A6KJbmgmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aRdH2lKRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HkyPqHtFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iWDazbDTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qn35BP5lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RNNFs0z0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SABW4UwiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3C8zzA2EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M6Cj4fyAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EqO3r51dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lvC6TtV6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B8SgOLvnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PESzrw4zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AYwbktmVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R7bGoQKgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func joS5doq9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UvhPkQR1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hm8dKVE9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z9Vur1ONWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y31D05esWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aHRELWDkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Sna0FRNsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VTIxCDTMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SC1sySjjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func josmVJDkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eUSwqFvbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7uFEIRdcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QelL3v3ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 75luHBxUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eA0uk7hTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x6b3Eh31Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y8EBnGLdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xNcwQaOnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6EIJNmTaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jCzy6eQgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Dae5SQtkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G6Hqwka3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CrSwYdXiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3QaN2hhvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MllHTN9wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3OnzRVPDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J7bqS1D7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DwsfExDYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func klRme0otWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IcmsnxghWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YNBwUBsCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lQF1oiOAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oXB8cVHrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hckTUkIuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uAGRWcQkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ODNQ8K5zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kmL4MtVBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iXcVy1JyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EBl9AWPgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fTAooYhWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DloVESAlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SMN1iLKdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8HQldRDWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J2A6DnbAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ckJdnPROWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1Ok9iBKzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u1DydfhdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HABH0KP8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bA6xwOZ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UFFhbM67Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ByTZDyM6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NvvuSOBbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZWoStYrOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JBhQBfJeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f6UIraqmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KyCV7gadWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ssBPXDxbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dDqx8h5VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JOtzzGDHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0rEBnmunWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dsw8t2taWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S1jLp3T9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NZB1rTNuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WpumK2WEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FvsZpKtWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ChvKGadhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LbqeZGJTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S2pjfHVhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M9lUn5EXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8ATRnhHOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u4jnabeOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a0mM6ViDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kJgvX6VCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3e2TpFEZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GXy4FKNsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JFVoirAkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AEvp36tQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bLqySAZQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NJr2zq32Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PVl59IgoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NH783kQlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qpRw4MP5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4d0kZkG4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xtBd9XORWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9a4fvLpYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YP7XsaNfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xPfeFcgTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8fn6Ksd9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gncbxwcUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nSIDbMXxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YqPYK51fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lGyIiJDGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s4jmli3CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3UvqONW9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IQRSrot5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mqROdLM3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iuE9o9KlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HtTGmFEKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3oAWM4PUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AFG4VG7fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aqIxrQxgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CMi2iG80Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9I7kpboCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qoJEvR1MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EMzaDmOjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r6hAqFrJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2sEyirTKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QNs6yGy2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iM6nldgYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3ZsGqhYnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mdEZ9l3QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 09doghL3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YiylwqArWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n3PclW3JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Sz5A1lq8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zwypeEAHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mpytotvXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lYPKN2MEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func softKO4TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fx3gNx0KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CYG5EJk8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PqGfNTxYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XBURyoV3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QBtNhuaQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NgyIEkPXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ev3XCXRyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cOtLXMZ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WW5EO0q1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gLRLOfNVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ocAqSAV3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func axPsYLuVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h2UsWRgPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DC7QGQakWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k7AdLNeeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5gfhINRJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HZCxxWeOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0Zh7zE1hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YtMabiWbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 12vAS2DmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CJSZVV9UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jEG9UpevWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uAbfEIxeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JYFPOaQ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DLe7LktXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yC4OSAf5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WmueuX6XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gqejrEPbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dD2JkrqfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ohhMC8V6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E8bgdNIHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7iUfo0ojWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QM5vkzLCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SZIkZIorWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D3UrpRTbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cMagRLFBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SLQhIBA2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YRTyzDiBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DzS61Y7lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3K6hqzcKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XDW42seiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WBTbQ61oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2OHhHe4WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aJudZ9Q1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kU6d2UqtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fVT1vazhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1itHfvsuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n8FP8eA6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3fEpIjWMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wS0NFGfcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wQARhCxVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AEkpl4IUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i5Sb9iI4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BKSrcPGnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 75cQ60w5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fUAjhktcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YiZwQwoEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RsBjpw5pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nFrw2KztWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xHdaaYpxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 83aIGWCgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w5bXk1XhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JRRMXA2lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dDVLeIlAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cUrI7l2uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wow2sf8XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GdzVJTYTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3uFdGfmDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5iVji6niWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BFoT2wuoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3mTEodFBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tHjyP6CNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CjNgWjRaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GKXXpGN0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PFqZHrLdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 55MQ6uRkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KB951OLgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JqqFTlRNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iBt0kTC5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jgsALiZRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vEOPeKyhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y09vmQNdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JANaAJeOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C9BPGrHTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func obSlIO9pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FYSxvXiSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lnUqfSagWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1Q5Pg8pLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lm5xlTbAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DzrHFzRQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YpsEiwnyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hWuRm9nBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1VUvdJn2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HblVt85iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iAH2jfzMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5ZMMqpZuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ey8h49oQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YQSikrpUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AjGg7oh4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cL8B7d4cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Md7SuZ9pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ToEVFNZ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SGGshgLMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cBu3OKxUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Yd32k1wnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wvm41IvRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fnDo9WuTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DxrRsPSFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p83916QNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UnwtjetwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fxOCKSK2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZMWUdNVxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WOhWD6eLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lfd5lmH1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KwbzjQrXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t2ghxh83Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ysy1FCslWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vjG7J3O4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kL2fM9T8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hAY7RCNEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ivYSBuAAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func na9BIxNEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cgjqWmglWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XOKaR7MYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func esDWhMxNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vra7kT9KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4akTmauwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h1Y4YT7eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BJ3YcPE8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZyMgr6dIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D9Ev75c8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zUJzpjdAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jfVLsHwsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qyala0rJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UafmdQkLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9uzOIesuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YbHhg32HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VQkWaAonWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5C7ySbaSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uvRKvsmnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8TNZSQyvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ohVjI3nSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1PIwY3wWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rGjJAFHrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7Zhq9BTSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hPodiwJHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zE8sWDofWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RlaW8VZSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N1bSwZy5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gC06gvh3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tk5t4EUfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iVaIp9c5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hOeeDCJyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2WIjdDUzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WM9n1PB6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N7ibWdw9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DbPjsIhZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YHIaSoLOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PHJ0BFu3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DYIkmVliWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pw9Rryz0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iy9GMUTwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DfufHObAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cT1rUx7EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1BvOCvTLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2DTYuydxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xBmo8J74Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AXBf2G8nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3tciR4X4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RNKFTINFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func alXAkfDFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1PyQT7olWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2Y4m1a7jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NkREVLq7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pVhVf0ufWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RFojnpP3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WMeDa4SwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 52Grw36AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n73xHOHQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O8bJJQffWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TSYdbaubWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KjAtSkhGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D5oXKvMNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fqw8YQWnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3SIEDUOOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DEo6MSp6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4ED3Vdn2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oQXKfzkhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5dvIu8veWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wYuwm1DWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E9JtMFleWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 16HpluhhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hFno0q1mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N3hgg16SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3AbvYaFMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6TQDvblJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SvPxSyZiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jd9hwB3xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rNHGPliTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A3M4pvx6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 977HIFW0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CL4SHhzKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u5ndwTETWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eMCGUBRlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v90uewnWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gIfiMng0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2AR2mLB5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UXa980wXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tr38HwUvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JAJ6ja6hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func krGOHZuKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ukASRoUBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F64xwBRkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BcBRnCxhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0cgWH3tIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g375bZ9qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ntt9pCwmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nba7Xov7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bK5E9nJjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6Jm33K39Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fMRgSn2AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h4RwLlSZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HEuenCCJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gcbAWj7qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MpsjU0WuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Dgh7NBbyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GdbPnGGdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1oLy50LHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UnyoKoiRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func npuMuTYaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5gUZgTKiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mBWAK2g9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z8mj1vIpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zaHmQaTrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z1rANOxZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JEVwCzztWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gg4PyqhfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QFw5KfONWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XBugcQq8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tj3B2eudWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JSF3ZxONWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y20XwgMsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wJkt7JSrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f9Wgs6eQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jc6B7mM7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Sy1Y8H1lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5pnuF9KyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vyeZNYZ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DNpAJEocWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QhPGyAqzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mfYOtZ9dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BH3OsOt2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9lN73Q82Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cJQneVBMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5Qwg3f1nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ewO3BtMcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PLdspHOQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LzP63YOFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xYf2jQWfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ebOFgJhOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ozg3ts39Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1sdpUjXeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zqWBUen3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dVvWlrRNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p6RDmtSoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w2ZTyKVtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wEhn61jaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Zc3wqI8SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KLHYmJqEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b1bWLmxmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ogrSdQ3ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jnXsIjGiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zR1WmhguWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ipYZhxrdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XZPyz9uKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ru3goHg6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HKoeLwVJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nxDKuZC0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LBxxf83cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vt8l3cTRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uUcO9zPsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 27vKb5IMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qy7qBkWTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0oyWdN60Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XrrxkQSLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dw34qLfwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hocC5ugNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2bNu6BDaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5ctvUytLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mwi08tuuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BlyvBnmaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fByZ5MR0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func La5ERRb5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qmGexKIIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pQBwXH1hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ugCOViQTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i7AsZqW6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3VZX5eNMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sMJRomd0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mj9KtyQxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K76ZG8IrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wsp7FmuWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8W7Gi4MRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UOXQJKsnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5DHAeKooWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IqgayQkDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DN4rXeK8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i3jMxLP5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IB7dry27Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J29uF11wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8Atsw7VbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tHFozZ0HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5ecXImQ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PJpUIRhlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1j3jgqgFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yusOdqFbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RJHXOv10Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S5Ztc1yUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j29TxK4bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CmkpR0h8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CiV4wzEBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z9QhqPqRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zzXkLYBKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Nvu0CLhYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dr18VMYkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cW4Dpv5VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xvUBVsq8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HKIW87vEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GF5HCPJuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mIqmftU7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZvfXXDQ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q1t6IC52Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dY0aQMRpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zlJaldFvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4LQFR0lpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BBzq4ayeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HJFCecU1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WkSoXdRzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DJ5GGeWeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6w00U1yLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K9W8DzFcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1CIEgANPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lUV1EhkZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dGTsnaUfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uBS35sXIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4bGZ03SEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1Jo4VgwLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WJiX8NblWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QIV9ENG0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R1Y6Kc2BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lRvxeuKyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HvLxn8tSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tQbq2RjKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LFCYr1JGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pBgpB4KVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lV5tTZuOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i01vDjGFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p2BKpb2AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h0VoeMz7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GNE3WcyNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XFqpaBroWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FVUoCRemWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t2D92SOlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hasPkFDGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MzC57YroWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func omPyz0eHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SUFQlUiIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gc5FqowaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PvZTJzsHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hUWWWBBgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZyKZHDuXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qltHahV5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GKLf3zjqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PtqUAoIXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PHQIQ5wbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YZUMzXT6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g1HWOwzAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3tBJD7QWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func skSMG7TNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ukky4If2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uQ5Jo0YKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y8lVeVk1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mhVU7sB9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LQXOMlsaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bKXFTO5SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zqdNKAVmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qb47rtc9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 39i64Do5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func acg4tnm3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9PV38iWdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ra1Xn5CYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OD7mSRUZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XmV3JuWPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9p1AR3LZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n8tTCIpJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6P1WJOhaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OsiNJPIjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SrRjlcSJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8WDkcwlMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NEnwpz8MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DLfvh7HmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jZXE2y7gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8DzoVIpwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2Vibt7N4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PtQDLFvGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kBgTDdotWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sVs3stnMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6nXUN0hXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eeS7qGvdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uDcFCfaoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func baIlGVjUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UBQ4Dsd1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func owwgJiQPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WasXLurxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c83aH8yaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 143sBMZDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WEoy0tVrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n97F4rK0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OmVXMmmJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dUzGKAERWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vWQp4jIEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9Aw3YudwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UZ72HPv8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3CJxKcrwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AxsSZEawWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gymDBufzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ns3MUbJCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 76UFrEmoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r18C8duDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FG3GAf0SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iIHTQLMUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FF0RDhJLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YFVD7zMLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gT0yD2iqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PCngNncLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7x8OoUwBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bjEGNoMUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mYA3anhCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l7gnBWwkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wWSSJ36TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ffFPc0v2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i49oxCTqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vwp4QrFrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 35D7JETqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P0HURnznWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KLB7VkzUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fo7c72HfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IErU6jNoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yKubCk2jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mImZ23BFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b9oze5aDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9nqGEAyuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mnDYeyPWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wBKoIxr2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0sy0lQI6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i0SJtVe1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yU2oSi0cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vIluqAkAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CrEtw559Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nY9IWr1cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lhEUjLYxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zX1JtFxSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cEuVmxJHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZhwL54ItWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1uMpgFaFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ed3ewTijWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IJiLHbOkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IZ7mFVunWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1DgDyJBTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RNtxRY6xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2dBwmdenWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gm4Y3cxoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BKfPxu9UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nkD9VOF1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qI5P3WgSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jUp0ljdbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W7GKgLZQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mnx2M0E5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E2KrIkelWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3BYqZEOVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f2PGiVjpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JxHz4a0pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J6H4QZuyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bapPTVx9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J1A0laN3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QypNLfw3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bbIp9bCRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cryPc1tgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oWLd3FJpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dCMBwp8QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aVgC2WbUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N6dsBZSgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WoWIJIvhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T6jeW7yAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JyelK8J8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iLiuIEynWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QrqUBJLtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s39X2v9eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BWhZOqpbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q2ZiNvDxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iMQItjhQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eafDhpzjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S1cgNUsEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BH0L4p9jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fEwElDcqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5g5dg0drWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nzwiEFuvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O1ylpkIAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 85AtO2mrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yrjevKXfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func byUZLLHnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l5vym59FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tzhIKOq9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oke6ta0IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bQtI56srWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 980bAHSwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FreKthYwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fFnITPhGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hbe4tDLGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 00sxyibeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VJlZrr8HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hTpdvLyKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dqn37UREWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wHKgy20MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xgi8uF2QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3WGCiWeMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ew0xxMfuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2yOtWNROWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Cp3gh0mCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AM3FbZORWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4saxIj3iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ejXaWupdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 65fvwYSrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JpI2GsrFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fXnB6u8YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func imbnIWnnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wt6taN1JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IQayEkTDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ha7YDqpsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XaSkMCzCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D476w6bfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qiYDWnKeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pESMUnePWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aEoUg7nmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zoTD8c5UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Yp1DVDfBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w4IvybfNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vXh78lSDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9ahdCCzkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rI2n90UOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QuXdLUOaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lbTW3w0EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YDrDnDLcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IGazte6vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s0O5DDbVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i39RAotoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5bEfidrnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y9hI2tCFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eObmuAtSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c7KDDharWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DSA9mEqmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xEUgU5MkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TxqUgzMjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QQEvZm6uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MALWz6OaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1aQTbYPwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7ftGkkK8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k62u0wj4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wfNF0sutWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xCjmCOM8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z9tvySevWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TIYUdHiAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PKecGiN2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e1FIIntjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tfZ5E8AdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 97WOVsYHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func myPiBeVoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wqeUt7aTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bYp3CMVrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func COvkTin0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZZGHSPVuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Kd1Iq3gYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8lMlNpBpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zDQ3E4e4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xsJ2zyBlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qMxQTOaDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9ykCviidWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YD2RpQstWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5vfQsY6KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 81rMNVaKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RU7PMEDbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oeg7QA49Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IFOPgk5bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oaFcgoPTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U6iaUAavWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JqoeVCm7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LnrzzIOMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W4z4OmebWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BZYYP4DzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w87eD4vLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wjwq6QkMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qk9JKqnXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jEpSCDX4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k0zEZzQcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mc8c5SaRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zcx6NgcxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SAstxUnOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4Zy83g5VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 97HydnX7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nCN7Td91Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tQnhq7v4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func frbsOPARWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a3uj0ytTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gxYWm2ytWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZbdqukuDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ynHZTQSIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gpAmXGFhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fA5KlqywWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vpHFllUGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SgL3EiTAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b7z8Bf8rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Brw9OZ4mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QeyYxzF5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QaoKpZDcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0XJL5BGVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KvMv7mgxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5T2xC3SCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 675JIwFlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NScYOt2rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9Iz78tvLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wkJFXSs3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V8eI0mCZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zYtEwWITWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uopUa38mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func opF9J1cCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nhlffqIKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vBAQ2QwAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kslJbCqcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NxaVDN9rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ro3G6zUYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8nGAYGVaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LbaMQd06Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fo93aolBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EcnvF1tUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iU5FmQPsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1VsVJLrVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Iv6V772PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M7NmHjMUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z3HGhLhyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O6qKEtBGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tJY2Y8O8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b5J429k6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VFXqtRXtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7TLR9knmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ilz7shSfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2UlTspjZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7rTqe9EfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mn1VtvZ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X0GyNNvSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U4cJF49iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qyGM13JXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bEfHUB9aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BtcscFBaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MTaUQ5wLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func enbg1XG0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9AJmiFrfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HsaOlkPuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JvgSdj5OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AtEwErphWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SCtGot9mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CibHkBzVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GWpsKZvgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gBxxJOeoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mgNMee5yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hcdOVVnNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7Wk6Y3MiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ECwTbnw3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QSUAWUl8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mZFNvEgZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cTwSwVhfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2uhTArZhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xpN4o2LZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gI0sbb9XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GoK3g5cKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yb88miNlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ihxVxNIKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w5h2NXgxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aw96b81XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NqOVgdBwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fdL2sQzmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SSFYWo92Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FaxpOeHFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func orFiAtUZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J6rFBDRoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qSz6L45JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TR18O8dPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Bd2xe2bTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dCDpOLPVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SElIDD5DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bkgjC3JFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ckU0fXB5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5OAUkc6TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9WjpbN1BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pC48dRu1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5pkoHezKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NxTAlvjfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FMhkj2yWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tAxJ3y0qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V1YAIN4RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2PwQfzXyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NDahpEblWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GPeAQrzdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hb6vbk93Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S2XNzsm1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V3FNx1TlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ol8piaoPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2bNiVQUfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RRavx8O4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vut0gTSpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7ObmdVecWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ia5fe3NqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YXyDAvUgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tnwiN9wzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func exSsHkbeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T5Byw33rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZVhpgCdqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jLc734aBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AY6CgX03Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N31eePKcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tdqoRdDuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cu5Z3EVdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KRdc7fIbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k5mVVyxRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5vfoOIydWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N05rWmd9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eVhMBYLnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TTOCJwTLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EHi5t4isWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J6QFwbE0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3vuIKj2cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1lETKhCEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5w2pmscgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MwwEJvUvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9uImGG6IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NqkvxHbYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HFubv023Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TA8dHbxtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hu5ThsncWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fr6EdM9LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func th6HK49CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yZPp8KRTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MLg3cxS2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B9y4V5DdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func okBvv8KhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aZEF1ZDEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j8lscn4IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ztRn5cBqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RSqhaWUWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A7LuwrmmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4IAgX6MeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WB5002Q6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mPqdlVD1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gYOZCznSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9eAR9H2JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XDZ1VEPoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yOrpJILIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rObow1PhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mgeaSErsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func am2YJgXXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TgKGBHM3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jZfBEn6DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dB4rhZwsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WMlMjiozWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P3sgtwV6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 91Bh8U68Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3puoOfo5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JmmNiETuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bGDP5GqWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QPrrl5KBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qKWt17WLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lgJZZa5IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eu0S0Ql7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DdpP1WBFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yHqZeEPHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QmHaNk9WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1BnDNJBMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cHmPodQHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M86agdlBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func enALwF8sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aHZLUrFWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lTPqaKH9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JSD4EZWbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QcSgRpLGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BsEjmnGKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EUUNsYUFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func awOcMkagWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DB47G2LiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yzwKET9vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7YjuJbG9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B55ujkruWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9AaZDr8xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gHOV0T3vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gh1sGaqyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Kz8SEfIWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W5M0YxRqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vXz8RI7yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wje6x9dAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 38P8FOwGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0LmHIRLKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BpvNP3njWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DSVKBtd4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BFJvmNGwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AlXO9Lb9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5FAK4RMuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q0I4osh1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h4PC3ZMzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cWJN3L49Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rRua7NQTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S5IxtokLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xZCFAMX4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Is6etYQBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cUQLHyWaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6Tdjql86Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oIW6vrmvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EeQh8zq3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S0gmmcmGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9zX2XM2hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NIpn4XTwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ahi7Dgx4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TEZRx4HmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ejIP1fzXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HbIdEyIeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q03fxmleWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MIAy9Qq2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 356Edj6uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2c5CVyZ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D7zZnl9UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XDkRaoC8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GTV1yjXmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PmSQL9tmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r7lFzz8XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8x1ExD18Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TCKYFlaNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ESmH1AhqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 99yO3vGKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dRuXIEbMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xeftJCTtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pRpNy7UsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tbZlwWVnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GJB8AdRKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WSPmMUNBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bbZffzfsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7N9xh833Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jD82aOzJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mOx0N3vQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TXR3exGJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fjw8D5RsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Co2mHEbpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CyjkxV31Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HMyAY9tcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GVPLzzgGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lPuzdhJKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gnrp9v5XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PNqz6JI5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xzGpZdqeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pac0gG8BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4dn86CT2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aj8bV3nnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ncwW77idWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B1QZkIxLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6AeMZJkcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AqN91uShWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bdqxmf7ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6r3gw6iWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Zzd7XigZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AChivHV4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ACAO0lYZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3ufmDr6zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tVZqRAH7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cHrAWexlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 73IqaxygWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lvszuCRXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YAWr1FBGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oq7hoP0EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5uLz6oKxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OtA9wPrTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5OZHy1jtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rNznMV8ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uAm8FocKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aAFx7RFlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QECDQqLrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lJBXDdlTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sIIxZu4ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bWDpFFi6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3GALJSCfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L4gpEEIxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R5z0YpHvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zq1MzuqGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m8l7yhMMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PNYxK96XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F32pwIoMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C6bLufABWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1VSlYX7qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 40AEJR9cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M6ELevPHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cfJzbvrtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sNmqnbWgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 80FtKFGRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 14tMV34bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Kj5RHCMfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q9V6avzjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lIVnwdXLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OA5ZpN2rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 85xveWgBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AkiLsaknWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xy67Hmd4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VSDE5nCeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lsxfMMMpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T0MEJCI5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dqOE2IBAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UqJ56J60Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9GKRAeqSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KxrK3443Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tgdQVu8ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z4LyVSMGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UFcxE9nYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D9AuOnKxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FzDgvThQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UJ0oLEymWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oW9QsRBzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fRCncVWiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N07jR0VNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bQXYLQ0YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 94utH8GsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QmmWKWXnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5isWATA6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func giy3xEemWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C39DnMEYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HVhLyCy6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0ZYQE2OdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MG5XBc3MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Cpa4i26lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I3iyLN55Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Em8CVsZtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DivxK4O0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UDNDPpkNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OCBNTG1tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R1QPjb50Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CU4HKRP3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9ZemqgKLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iz7UD9BDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hq6KlxzZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JjsNvsW7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c8UEpvM6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 030mxe5ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YhN2TDiGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func irGZZX74Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gHuJeW6sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TdByQgtCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YwYP7TigWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rVvrXqfqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vMDXpkeaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7K3Jad6mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aQwGti4WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HlGajrPQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LRupdNHkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qfsJIAgLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aWE0TVRvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OIPM0byxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gsd23LkKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1pCkR03JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wXQ0rp4WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c5r7coAIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QKSLMekzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3esITJH0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9yIDsW2qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Rad0T1WeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7HbbrNrMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YWLachy7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Lw3k7j77Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8Xzz3Y9dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1Rd8yGWrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Nct3APjeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wjw9WPO4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SAenQNM2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pfbIwY60Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NRXvZVQXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v005wFN6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ncrugSNEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DBpAaWCMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F8oB3573Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yOLaddV2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pvAMpksqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o6p7fWxGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Osfolq9GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GKRZmXOIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tF8R0FKOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YWmaHXGPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9D2oIXHWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m6aczZifWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ndusePUiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7sYQV8yhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RoA8dUgRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sNkaNocpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S4VId2RrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B2lwdqfIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HnG29Qn8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jUqwPxrZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MxrTklPqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bB8UPkBwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 20CIWwrsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fBRRWIGLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WSnFABqcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YC6Ip4i7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3loc6R15Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s6DmXko1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SFIzHdW7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q84xCQkaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UroTdrIlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2nmSs4kxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oGNaBWARWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IOJfBIdaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v4Zp1jHqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GguhHbrcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2pWxDBUOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G2OjEMKoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9bB1QMthWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uaSQVafmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DsuIHfrZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0q2Y5qTlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hKaMS3IDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DiCjKnzGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vOzno78sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0Pr9pedoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yuR4WKOaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wzhDrhwcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J6kqSZxPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TL0a4o2NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dJoU1AarWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fRW53VMOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sdiAtFYdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dn7brEzyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oGjOFxjcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A9lQVPWrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i1lf72VhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ASDEW6WkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XkcTzZ0lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LAfQFjaDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R48JIJ7wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kjaSg6p0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6xSyq5iEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i64r1JG6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eiYPFrZeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b4vcMXjxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TXV2QgMLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aJ0L2pMVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bL9dRRMbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hlXEuOrnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7CBEI8itWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9MIMLiMMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func awXjklXsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cStKYtc2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func exg66LtfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LQ9Ats3WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xtggSVqOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g6c21pU5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m5toVkFhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z02RkD8xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6RH24SfTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W0hhodVMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t6PrTn7RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E89AtdOJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eGUXmMc4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 87becvEyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 27d25EJwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q5cV3hzoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YjZwTyi7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Cj3Fb3osWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n912BakRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e885PN8DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4v557TnUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SQILK7AIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sO5xpQ0KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pJi8TmAqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BSvGwaUXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C59jYw7RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A3NkhImaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6DxHXX7HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func whO5ZzkuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WU1w6o8xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gxNljf56Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JmTd8wWAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OI0Ryaj5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kSqtu4FBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 344HCMY0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ezY6TcZgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func grwI8Ng3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B7JphcRcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T9xzUVHSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1nxO8vqOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U97ZHAZ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7ZwHi73DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4b96yGk9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ptvy0D5qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4SzOsgWqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P7mXl4zUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LQdRpyLAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 37D8R7a5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XUwD8f96Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DhvzAuEhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YOY1eLyeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kPzGt1c0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KPV2fXScWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mSRvnvY2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1jM1Fez7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WXkAE7elWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HCdpmap6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pbDCx8T3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F1OmmyN6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CNLpgIsmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uopVn3VpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gruvl81GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TKIWsGtHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0as9k2i3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lEUuQ3dtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dOuzM06NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sQppwAjeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PEITbyisWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tQQ9MJluWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mFms9eNIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kdbZmS2XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3H5oA0CpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JdZWceC8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pccBvA32Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BFGPGSwYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g8mOkO9LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qLjarf30Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N7Nln3yxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iZY4UcFaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YwD38CtOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZciezpqiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func divsSND0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wl6ijnULWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dTJZ0h9NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t4ZUoGJuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eEui1LlPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uDZN1u41Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nLrHaW9YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Bq3CriLmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FA69CLDeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W8EIDvmxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gLJMsuUvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bZCsosOWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R0YJNNrsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rE9yBpNvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Yu7015J2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6Q9vewsgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z96CeGlMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ias7Ne3TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z7gl7WxKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WzbK3moIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gsWCQyj2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r7VMFh8eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8ce7ojspWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func liwh3BMqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CC9YagzMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func raDnC8m2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Remm9Q8CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TdnrZCmJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H6n5wizpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LL6pN4T6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AGejPNpkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PaMASEZBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func adbkEWl7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kmgxAwOzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MRzdKTBpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TRtzH70GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IzN1cOvdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d2RAtPxpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 71Eg9VIMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y58XlZvwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n4nGvcxaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GQmWsJSMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FSgYATt0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bjQtxj1hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 01Nj8CVgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JTfqfxW9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U75TqGCkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SMHEPTYjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hsMnpWQSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7MPj8MjYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6y1PI968Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bQxbUfD3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qi5t5qZsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2YfBmoPyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B6R0rq4uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cQvjfoTzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vxvdMzkpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pgHt6mwHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jbwYcJEOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AS22s1nEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7P9St9f4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FymZPV4XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hLLib5MLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EFhSSsInWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m9TGc0RaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JiceANhTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i4D1NrXFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D776GkusWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AiJbFB4KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vxLm5xBFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QSad99OCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yyk6SMPNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h8d1yHABWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KxjM44DXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8uG0WvYQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F24tc3Q3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NpngxAqHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G5YsFQiUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 73jjmulDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YLuNAAOsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vZRWuEIRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CbkVZqDrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func trXJq25sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ueNmENQIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IervCWASWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func olQlwxCbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T2OZP7SaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iRqaIW2mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7W8tNGwlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TDgPBTPAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fau8z5nxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9IbvJ218Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JckQ3LNIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2TZjxdoqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PHP8ecXdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QmciPYFGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5ADSYjTqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DQ0WMrzMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zHysJ16eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9v4tne2RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7us3i1KJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oeODNcfgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BTWYiQwXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mYvD1uw3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func avUN520PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I7GmIGgCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gYtXSUjCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r5usbDp5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WZDwLNfKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gRFjnteIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cPX9KK0EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uTGYzeeeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j0KzN3xUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9FvmdCvXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t4dozHxoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 14prZphDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 85XkTkYVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7fTjHoEYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5OuBRqakWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m0nscprSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4xHCKwjQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aoeigL6BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r5sDvw3oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u7HHNGheWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xLcRAu3EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4LDBk4wVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xN1XKvyKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ojRM8NqdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hqxTefEcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kdHGy8mCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pwQUHgj5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zfWHngj1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9UzlCY5tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G4Vka4E8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WBwEktEZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NuC6Co74Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F2RtbP2fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sftDcr7xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z3GrjRzCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rliMR3qMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EoRks2DiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VaENAdSpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4wPnCQ0WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 40zt9LbNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jlKCMtXSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6J6HQl4KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zpeQQRwmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2Ofl00CsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aEF6ucyhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n6cYRj5oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lG6mZ91iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TPlSZuMIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CcdG6zFUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tko0QOLuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jf00W0zAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tt0hw4zlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cm7FPzkdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7qfqkbI5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WXvnqtAgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FhBPhz3XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xG7OKEQLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t67v76OsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N0UXcyXqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zHrVf2LEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 83VHIyYWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D4vrEs4SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j1mO1LCNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XjWiTGCbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Exy0LECyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7OQjPkvCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 47pAJG4EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gXtQTdKmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fuup3IC7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func po92GbABWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1zNDED4AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oGcf4agWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VHGz4QNFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F6lWbVswWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N3z0o4HlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gLHqW6iYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XHzH8U8iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vv3T5CDbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VjUyvThRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N4f5JtKJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7Iykry1ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tKFjrw1wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8NnumW9HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KJgYMShYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J0FGsBmmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oLCHh0ohWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H4dIskANWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SnJxBVR5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tu1a2Zf4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w5BryRqxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3xCmDMzrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IT43Zz9FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 28SxTZgEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d4zAEWc4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KgkdWzifWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bYJRkg9oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7aKa14D2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1AZw2dAEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TX9A0rwdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nL3ypRz2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AZSKWrYQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cUKKfaXvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4w9HnRpzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bx04Jra6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Dj7xo4C5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8L85Cnr5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DDc8lavsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RUCdVNUvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gm8oQr2NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r16UJTLhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LH3Y5NEpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fqoip6mcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ru2hj9yBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9qU3h2odWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oESKhDQ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JW9yPfkGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1t4fCnZlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iwP41hsoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hEbSZnreWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QKm5NOFFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qdJsoB6NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5fJYt94FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zhwBH9g7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2l07OGkcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oF3uuaJnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func whAMLCAEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9OuGquswWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8HxghL6NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tZfb5uAIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UaKIxidgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9XZ4taBVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CzH3qVR1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jdS0RoRlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UfUqcLMoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zXXDqJ5vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jdxOmzirWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ImV8aLTFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2RHZfIXPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gtm5G3i6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZYRSYrabWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kOKy4xWKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k6O0qoZ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func anH1ZZcdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vZB23pEaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VvEaQOnYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1pEtKaHXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HWg4jILlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func axsetf7FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o92PK8peWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D9DLDTJ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2Bn9NNJ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lGaICj4nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wGh6jJRoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VWPC5ULGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LX5lHRdSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P8OKUw6iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w4Lm8IRxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zsnmSyGgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2bqXQZmpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qsgyyRxuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3l7WLdybWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 81muNpHFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i0iHXFYoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func svRudRL3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kaC3OplqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WTJTI6C1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5o9TY04fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y9JTkPhTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dpGrYzdsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func POJI9whUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0DFc7U9cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KENAeHzpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TDwAULDxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5qM5O0ofWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0HwRmGIqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ruaNA7hyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DNxa2OjoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BFyKAVY0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yk8zGTSYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qtCVM78tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1tdL79mBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wcNEvwnnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rNIviZYAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IQe66S0oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sqBdCW0ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func spjyCqq7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a0dbJdgDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3N0YnITGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YwAvk1lsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YDkd4gzbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aDBaRDOcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BDoadNE5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r0Tl6UAuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CA25palpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o49MaxbLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wlEbck3nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Bp4PAQDYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H8R2Q1I8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ks0Zd5ELWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pEQyZrTlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 12kA2eeFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I5YpO0keWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jP4clZfhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4Lkahz1mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9H7VhdtWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JmOljDYVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SPvZ75GmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SR48RUlVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w34cV3IGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XARV4fpMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iLDW9MxoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jSsTSw2YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GR5C6XdTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IlmhiQ9GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yHqcfd3EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jWC0HamdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DNZqSD5NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m2GQ7Fb3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lVUdAhN5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WVr1N8dOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6gMkNuSPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func co8p2keeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JcHfdwikWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BvpKIp3nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p7ZL94AUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4ogxBwFZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GYWV697DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jZr6VwxUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jRvNZMAgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pVAXnmEAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wFGjlTXtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7i1PORuLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j0jcT2x5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t9vKh2QNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AIQm2EW1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bsIbkfOeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OJewWo83Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4osWIpSeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4O1QzlWVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FvDZ3AaAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q5AOzinUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3aeyZgHeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S5Skh14eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NfYhkQaQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ffLX62KBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qcFyYNR2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O8h8UBbkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ha63B9OyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UV2X2jtDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 06TLqzTiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8KmplcVvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ny2QeadiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5o5vENsnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QpK298N0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yMPF2lyqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eFE7FmE2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k1twv8AVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LIYlF42CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gFxmBjxGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4DvjmiJiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J2zt0QYNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wyWKdPhXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cv5Y7wi0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ou2tx80SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 03HfwPaTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R3aZDEQnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fWfvwgmOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e2epI2EUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nhG532EHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EngSVIxiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YaQZmW64Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kGdiHiYVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZUQvg5m5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OrUuz7l9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PxIpgB0OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6IO0vJKFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jMTMFw3kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aNW7Z74JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iD9jRItWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qx4kEqJzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rEPWCX0kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M3PefR99Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 44Lu5ZpfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g4Ltgb7HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dI6wmKdTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XZ9Yfr4bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hK8XHh9tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H5GAcIlAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sYhOxzKcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q0LQPygGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nJYXKv4rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VcmPKNx4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TEE9gRtvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UF0fv7txWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W6Af4QZdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C8UCEL3zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wUaQjZfWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MmuuvKh1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WvugmrwYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hNSVFOQIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AfzrM9FJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mEwHlhstWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O3VWmuwuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1YeagKXiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UC3IA46bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oxmwJ3zfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NJvKgYIPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fz3LTht7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WGb5cRXaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oF13O88SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KfKW8rxSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TDJCm9TJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8SgR7uASWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uKXTaAHtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JWRJ3wzBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VlQjwURKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qL9QAwxxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 88t7x65oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sfdckl7CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OIPowy7vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SlEWmKefWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F9yn06mrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aBfzlcOnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pNkHnXvUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WIkML4ErWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UhTBUcbXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wGLc2hjGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hdIskqB8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PhEOzC8EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LYRVptFIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GxhtFlbvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4XYD2DnyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SG7JZRk0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qJuovnnpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func afwJMfN0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WKdcF6nZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v7ZhOTggWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HS38TyePWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LKkkBSP2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GVConyQGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func acTg6kzVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oWnT6pa1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xu69BP1CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Un65CHRdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yxC6VbonWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rgauj98yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YuTAqmQyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A7cEsZeKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oFRVfriTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jCHB1NAJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XdKXUudGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LX7eEWo5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UUXamTagWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gQutrVPuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TqbKPmVkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WQr4BzH7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H1ZqYqESWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sIz8jNlbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9yf3pPQXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ckAt59aXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uPAVu10KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PR4RgzFVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TzFnQF2HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nXoNrlsKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fRNogqOqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z1dFIxA7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1IEwDYv7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DS6BihFmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TBHtpzMDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lOpYujZdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gs2gHFlZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GZWqLxDTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1IfstzY8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ABcdVUNCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func miOGx3VLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mr2GVLUBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iSMh7gGqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SEO0qE8hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5TLPCYFDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Kb5cTIOPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uIeemQ5LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lssYydhgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WTbcbRjoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Sa26Qcj2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LxLRcop8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pcsFjV3XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 17OJpQeXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G2kzmlHWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g8nvbP8XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gi7zVpgoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pWiIfRweWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mF9GH4jLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JKSCQSV1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4i0K7ZydWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wwYrJj9nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func frVgZyEZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yW8aXbW7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8TCK4W49Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K37D1efwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HsSS5NOwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qlB9JLfGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QodU2mE6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L21h4sr0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W7JRBnZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func srIm94CxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IXnqRCjvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7JX4YYHTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RKhAf2HgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zjTnk0ZpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DmgkzBtjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m8VAuT20Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HnT2R7MUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UsCaf2M5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hruKuIHxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yqeFVe9aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CvRkXJStWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AiENNosAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IyPQBak3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TTy2bbawWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tZXUR7uRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OJwcIWYkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y3AKLZQoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NpwSXocGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nsiPpXqzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GCi5aKlkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QX6MnnH3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N5YIZbwDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ti0vmwMRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8HfN2YLqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9nYZxDvsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SuRGd90OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ebyFTjaIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nj2O4FaTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func USpu03gmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HZmNR9jgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fdneScxeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AQzagEQBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lS9t07BQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2XdnWlLqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 94SfKh4BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func two8onaWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UUkiMb5hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mk2pFicPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1JsTMSCNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kdwjF4fkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yBAI3Wl8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ouOzIPDgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XzaMhoBrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tbWRsJh9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VgWA2RHaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6ee4GNQtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RRlFy9YtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JLSesrJDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OzveJVlmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JVJWcrstWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func msDQGQ35Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IRB4DqpkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l37ymlKEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7nSD4aRoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4YHfHuBlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b6F1mn5uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i1HqvRP7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9rOyYQGKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yIWGOr12Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JFox4rNKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4S4kOPOVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UtDEu127Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fLwBTTqTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3u9YMviYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7DlzzRIjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bGOEOslWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iGytVYRQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v4p0FUKuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5J9uXY9XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func maNjck24Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NAq18OSIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZsR7fAgMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func udcWZcWuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func axG3LCafWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LqQUh2OPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 56sQTG9fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y6T6aJHyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sbyzQw4aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MTBFv1qrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FuEfUMrPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0j1JjuDPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SAjUrUqBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SpWHEnJ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DKytBx1VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cGWKP80oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4ivQtuChWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3sTjzCiuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6hj9kO6FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bmimR8kgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e1cgntofWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o4GTfbQMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QuOq3zBuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S036yG6PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FbFTk6amWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u2qAegDrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oyRMDLCsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3xmv3vWOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0uNNRONcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0XS5EU80Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UbuMmOEFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2XuDBrEzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GKN5F8jYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V24Tnc6JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ywIZihSyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vq24iqFAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KrYxCOOkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JYLRczYZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iQPa5BYsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GFzxGxh1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BeuQtOWQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JLZ4eTr4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b02yZ0gMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VXKzqTOQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qyvTwe7fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FZaSMkeVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bF7GZJ4eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gc5fdkSCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aCOqF2NdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XZKTFRGSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kiiwSHxnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 103EdkbPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aIb97lZeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rhJ5p0bOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uxVlWcdRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PA9Vd65MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9kHYckq3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d6cWztX8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OafUi6CmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CfOEGkxHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uzC1OrCrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QbyFkjXkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yuURpFmYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gJXRCsbOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func inYwHkyaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ClYxlBfcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iqV612pUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8rHHxKTCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YfbNcqB2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mbGEIBPQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AeksS5KhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func agUkYGIjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D8Vned2OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ht2St2CtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SkxskG0FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7Xud0ihQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WizerpinWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IHte2J0DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H88aUp3PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a0nBmlHVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func keOYrjfaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EglOluZ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L1YaCBPbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mxG0BO9zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oJkGbvPTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HHGNIbPIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EeYi9tGkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wg562AqZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yb3WgarFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nBiefOIgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yvp9cdTNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iGbFDhxJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9wJQs1HLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zkCj6wvAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xKzNg7SnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AsfI7mB5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ntaBqTALWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tr16DV5VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g4SLpTkiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7KyeDvdKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1J3yCJOoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SazAE39qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zQ9CBa4RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Lkrc3EEQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GHeeCBOYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CqmZrZHlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WmaI5KPhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1Z9eQR8zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q2TdHyMjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j3dP3jlwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZyyUg4FVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pl7AJuQSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oz0bL4kGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Bl2Ng91HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uDgxjgv4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JwjObo4lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kQL8CHJTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y2ucOtDqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VCfwXpT4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UHoMaxe3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GiViQKdBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GIiEllnXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YZRC59vXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Nlip3spQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EigGiGJTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CGlX3ULRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HPzxjrBMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NDYnnxaSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4Y6z3VnuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ANVEHwcdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WXVkaiV9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JIlnBsN1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TIp37E4kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LozwvJIJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 84U0VZcjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ObrEYDMLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cxZ5hEViWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GjLX3FGrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AGJqeMSRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pldJpFUfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FIwaRExPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K1l56i4FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7a8d2GEuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iRJNBSCVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ft0UwJIIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4cD8Aov4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4D8T3NGIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nAiCp08mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UUzWFtNpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6xp8YAQqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aMKyYPdFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f0zW5nILWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ijQ43rC5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nvHoAPyaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3GDFKD1lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BDn8dXZmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IAOwSZPbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xWlRbq9pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VjIcLiZQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JVeSKa7cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6bhzoWwKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G6Sks11OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bRFgyqDmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func McPqKUg2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func imVvQHbhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CMCagBw5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VCWqIgXmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1siWlZknWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qqz71KjjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rUvKuxmBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O4pAMDYoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5fOQ7tZQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J8kWpLojWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ocgOx1v4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vrgCehQ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ViNxyTu5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JrK8CkqoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 84MGro16Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4I17OnIcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JMoAZsjbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xlaDJARiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IL6QVqrwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sBgEfgkbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e5A7YZWgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T3rzAvHeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IXmzgxZxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XQPVvGTMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RtvMLz4iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tRU0LBgOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8SBf4bgvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 35Dw0KETWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bgrjox9bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i9mcHbBfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n4vYtxMZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AdFqPPYyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ywNsParxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func udcZve99Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S9Vq8FTRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 51FZFemnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h8QxiO46Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ua5Hf0V8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y9BoFTo9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QLcSaTMpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SSyjdzuOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VvI9dhOXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NH5fLZh6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pK0Abx68Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ef2hQNq9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8H1nZNkoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1YeoJrcdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Lems5ZtsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5dP1drCVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SZkEWue2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QegH8cpKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wD8oA25lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eL4hUiroWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D5NXfSnVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func voMBAZEBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P8XyvOJdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WL12mRTYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y8GYuAEWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2onO3hvmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 81dGy1U5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func voSqu4GmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cLOMZpzbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AVRS4ngBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zkuTQyqsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oIaQEMkEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tvoHdxyLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HXsT7sGZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u34HGjMKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m4IKmxQfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KGvM1B8fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KeM80LNtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lto6aqLnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4rdHeoLSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SCz8jEtPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ObS1dCA6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k9lvAyHTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eYWtaPROWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mgc24SHbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0T5MrGyBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jH1ahUpKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AOMc82noWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ft8CO3rTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fr50yx4eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QDmEdRU5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D3oWB5AwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DLlNeT4MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PIay2TNiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2GJ6CrxRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Aq9rtu4pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YUdKTFPsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G6c8mpqfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OgyBV1eJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Sgf0ELLKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 37mbUgEsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NQLitRTQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ejp43wLbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PQEcE8cvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DkW2pHxFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UEjmkPnLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func flqoNtbgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func upvx2QbMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k20QZUcbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hLDazwZUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UmOUaIJwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gC48pAA5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mzmBi2ZCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8vxIlonNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KFzBZnbJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rxoWVOIEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CAhjAXYPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YglN7CfQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 67l5Jy9nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mMP62hqkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P8xx7s6yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dpCCRQ6cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ayh3eCnOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tg7Ftk0FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p4Svqu4cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gWCFnICXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UeJgHw1QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5jDS7CqrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3odusUVHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p5wrCz22Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zGVEAqbjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 035BIrEbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kIwR4GVQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zbSQuXSKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hiibespjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pAUyBH6PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UHMJE2sDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ud8F1nbVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5JQ1n1xxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xEBWvyiQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func On3R5swzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mrc9XtXQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N8qnAsfGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1SCwKBILWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jOMwxqyJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SjNFRxdcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3mARAPlfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zf2lAi6LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XIdfXWKxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4uzbV4aoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0eqX2UuCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i4ifH3UxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gD2zruWyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UipJoBo9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SouABO1kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FrKFWViVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5G7nlUHVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qNqbQJE9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ohWETAPFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func byPm8pbHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nTgVnCFrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b4enTVOgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EJoOSGkPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7BTmas2fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J8k8BMe8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t9qlJsTEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z271h8jlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lT8wvCRtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bwivljpsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8IFzBlhbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PnJ3n6jmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func erIlaRd9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UGgwiTkZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sJOzJgR1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EdPDH9YlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4yHWB0ooWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fvvpOwWbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rnfOiuBQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PyffU9EPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Sw4g6Wi2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PNOzKSASWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PbXuDzGAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iXTUyZK6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GYbmsOSIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q2l4f1W3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qbGAelZAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ewi5HlTcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iL2NWxkuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ntrr3NzQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u66IwxXBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WH0syWCFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BvjRT5TuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TmE55Q5kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jPfxVKrxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qNxPqjACWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ahYgOe09Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rK1lBNjHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e3Km0kjmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PcG6zm5jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QywLomJUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wxKV1ub7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SRJVKY8CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X2rOOeXxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yXgPWi7oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5yfxPCb6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3qoM9XnKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qhzJVwp7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0hMKcFTmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y23JrScxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YaSOPPM3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e6O3VWrRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AHHhPWFgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rIyBYJjfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func InfOzlv6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AvOGsYMIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fNMsPmq2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IdUfR1vRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H7D7Gh4PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fyhCI3FNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SOliFOfrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fsy0Jjg5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5R5XzzbkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ipLHwLaeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r1g1pciPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MjhZzqglWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QQesp3PGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hmOt0LzjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RQVZGoZXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7NeqxzhrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VIgLsyfpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HzwJT6XIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 79T9UchbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bsYPtPQ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Le3oufavWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Nim8BNCsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ch85OxpsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xy233sfkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lVzJ2CR2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J4qzaaxKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xZIaKKzlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func peJCjJJQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vvvMBf8yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gePi7prKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qZotn4fHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mlXFJAwwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MnV3uCtFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hgC9wgMBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HGBJIA38Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KEQu0bc1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func POOoyDM2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func luaREcMnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8JBkWSZTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VOgAY20zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7f8zcZIaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PyNzi3oHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IzWZ1d7bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I5olAjabWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jYftL714Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MiaTk6mtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1kTTgeWpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V6c1rNI3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d0OPgGwLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EDEHqJk6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ptoee2ytWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YcwIyTqBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4VLMnaXKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

