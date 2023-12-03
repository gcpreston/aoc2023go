package main

func Check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    RunDay03()
}
